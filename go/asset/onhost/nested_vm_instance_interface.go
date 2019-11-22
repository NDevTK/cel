// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package onhost

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	hostpb "chromium.googlesource.com/enterprise/cel/go/schema/host"
	"cloud.google.com/go/compute/metadata"
	"github.com/pkg/errors"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const timeOutError = "ssh timeout"

// nestedVMInstanceInterface is the interface that represents an instance running as a nested VM.
type nestedVMInstanceInterface interface {
	instanceInterface

	// Runs a command on the host instance.
	RunLocalCommand(name string, arg ...string) (string, error)

	// Returns the path to a supporting file on the host instance.
	GetLocalSupportingFilePath(filename string) string

	GetNestedVM() *hostpb.NestedVM
	SetInternalIP(ip string)
	GetInternalIP() string
	SetExternalIP(ip string)
}

// sshConnect connects to the instance and returns the ssh client.
func sshConnect(instance nestedVMInstanceInterface) (*ssh.Client, error) {
	nestedVM := instance.GetNestedVM()
	sshConfig := &ssh.ClientConfig{
		User: nestedVM.UserName,
		Auth: []ssh.AuthMethod{
			ssh.Password(nestedVM.Password),

			// ChromeOS test image uses interactive auth method
			ssh.KeyboardInteractive(
				func(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
					answers = make([]string, len(questions))
					for n := range questions {
						answers[n] = nestedVM.Password
					}

					return answers, nil
				}),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	return ssh.Dial("tcp", instance.GetInternalIP()+":22", sshConfig)
}

func sshRunCommand(instance nestedVMInstanceInterface, name string, arg ...string) (string, error) {
	cmd := sshGetCommandString(name, arg...)

	instance.Logf("Run command on nested VM: %s", cmd)
	conn, err := sshConnect(instance)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return "", err
	}

	output, err := session.CombinedOutput(cmd)
	instance.Logf("output from nested VM '%s': [%s], err: %v", cmd, output, err)

	return string(output), err
}

// waitUntilSshIsAlive waits until ssh connection can be made.
func waitUntilSshIsAlive(instance nestedVMInstanceInterface, timeOut time.Duration) error {
	instance.Logf("Try ssh connecting to nested VM")
	for start := time.Now(); time.Since(start) < timeOut; {
		conn, err := sshConnect(instance)
		if err == nil {
			instance.Logf("Connected")
			conn.Close()
			return nil
		}

		instance.Logf("Error: %s. Retry", err)
		time.Sleep(1 * time.Minute)
	}

	return errors.New(timeOutError)
}

func ensureWorkingDirOnNestedVmExists(instance nestedVMInstanceInterface) error {
	directory := workingDirectoryOnNestedVM + `\supporting_files`
	cmd := fmt.Sprintf("if not exist %s mkdir %s", directory, directory)

	instance.Logf("Run command on nested VM: %s", cmd)
	conn, err := sshConnect(instance)
	if err != nil {
		return err
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return err
	}

	output, err := session.CombinedOutput(cmd)
	instance.Logf("output from ensureWorkingDirOnNestedVmExists: [%s], err: %v", output, err)

	return err
}

func uploadSupportingFilesToNestedVM(instance nestedVMInstanceInterface) error {
	err := ensureWorkingDirOnNestedVmExists(instance)
	if err != nil {
		return errors.WithStack(err)
	}

	// save supporting files on disk
	destDirectory := filepath.Join(workingDirectoryOnNestedVM, "supporting_files")
	for filename, file := range _escData {
		if !file.isDir {
			supportingFile := path.Join(instance.GetCurrentDirectory(), filename)
			err := UploadFileToNestedVM(instance, supportingFile, destDirectory)
			if err != nil {
				return errors.WithStack(err)
			}
		}
	}

	// upload cel_ui_agent
	err = UploadFileToNestedVM(
		instance,
		path.Join(instance.GetCurrentDirectory(), "cel_ui_agent.exe"),
		workingDirectoryOnNestedVM)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func isTimeOut(err error) bool {
	return err.Error() == timeOutError
}

func waitForNestedVMRebootComplete(instance nestedVMInstanceInterface) error {
	// Wait a while to give shutdown enough time to finish.
	time.Sleep(1 * time.Minute)

	err := waitUntilSshIsAlive(instance, 5*time.Minute)

	// Sometimes, for unknown reason, a Win10 VM could freeze after reboot:
	// we see the message "Getting Windows ready. Don't turn off your computer."
	// on the screen, but the animation is totally frozen. Since the VM is unresponse,
	// there is nothing we can do to figure out what is going. When this happens,
	// the only thing we can do is to kill and then restart kvm and hope that the
	// VM can boot up successfully this time.
	if err != nil && isTimeOut(err) {
		return killAndRestartKvm(instance)
	}

	return err
}

func killAndRestartKvm(instance nestedVMInstanceInterface) error {
	// find the current qemu/kvm process
	output, err := instance.RunLocalCommand("pgrep", "-a", "qemu")
	if err != nil {
		return err
	}

	// The output has the format:
	//   PID command args ...
	m := strings.Split(strings.TrimSpace(output), " ")

	// kill it. m[0] is the PID of the qemu process.
	_, err = instance.RunLocalCommand("sudo", "kill", "-9", m[0])
	if err != nil {
		return err
	}

	// wait a little bit for qemu to exit
	time.Sleep(5 * time.Second)

	// restart qemu/kvm
	err = runLocalCommandWithoutWait(instance, "sudo", m[1:]...)
	if err != nil {
		return err
	}

	return waitUntilSshIsAlive(instance, 5*time.Minute)
}

func nestedVmInitialSetup(instance nestedVMInstanceInterface, kvmArgs []string) bool {
	if err := setupNestedVM(instance, kvmArgs); err != nil {
		instance.Logf("Error setting up the nested VM: %s.", err)

		// if the status is already 'ready', then we shouldn't change it.
		// The reason is that if it is changed to error, then the nested VM cannot
		// work after host reboot since setupNestedVM() will never be called.
		if instance.GetStatus() != statusReady {
			instance.SetStatus(statusError)
		}

		// should stop
		return true
	}

	// shoud not stop, i.e. continue
	return false
}

func runLocalCommandWithoutWait(instance nestedVMInstanceInterface, name string, arg ...string) error {
	instance.Logf("Run command: %s, args: %s", name, arg)
	command := exec.Command(name, arg...)

	// Log the output of this command asynchronously.
	stdout, err := command.StdoutPipe()
	command.Stderr = command.Stdout
	go func() {
		scanner := bufio.NewScanner(stdout)

		for scanner.Scan() {
			// Text() usually returns a single line of output, but can also
			// return a partial line if it's over 64 * 1024 bytes.
			text := scanner.Text()
			instance.Logf("%s: %s", name, text)
		}
	}()

	err = command.Start()
	return errors.Wrap(err, "run command")
}

// Setup nested VM, mainly the virtual network setup. This step is always needed.
// kvmArgs are arguments passed to kvm that are specific to this nested VM. The common kvm
// arguments are constructed in this method.
func setupNestedVM(instance nestedVMInstanceInterface, kvmArgs []string) error {
	imageFile := path.Join(instance.GetCurrentDirectory(), path.Base(instance.GetNestedVM().Image))

	// download the image file if it does not exist
	if _, err := os.Stat(imageFile); os.IsNotExist(err) {
		retries := 0
		for {
			if _, err := instance.RunLocalCommand(
				"gsutil", "cp", instance.GetNestedVM().Image,
				instance.GetCurrentDirectory()); err == nil {
				break
			}

			retries++
			if retries >= 3 {
				return errors.Wrapf(err, "error downloading image %s", instance.GetNestedVM().Image)
			}

			instance.Logf("Failed image download (will retry): %s", err)
		}

		// if the image file is a zip file, unzip it.
		if strings.HasSuffix(imageFile, "tar.xz") {
			if output, err := instance.RunLocalCommand(
				"tar", "xvf", imageFile,
				"--directory", instance.GetCurrentDirectory()); err != nil {
				return err
			} else {
				// output of the tar command is the file name of the uncompressed
				// file. Update imageFile.
				imageFile = path.Join(instance.GetCurrentDirectory(), strings.TrimSpace(output))
			}
		}
	}

	fileToRun := instance.GetLocalSupportingFilePath("setup_vm_host.sh")
	if _, err := instance.RunLocalCommand("bash", fileToRun); err != nil {
		return err
	}

	// start the VM
	cmd := []string{"kvm", "-m", "4096",
		"-net", "tap,ifname=tap0,script=no", "-usbdevice", "tablet",
		"-cpu", "host",
		// the monitor so that we can tell kvm to cleanly shutdown the VM.
		"-qmp", "tcp:127.0.0.1:25555,server,nowait",
		"-vnc", ":20100", imageFile}
	cmd = append(cmd, kvmArgs...)

	err := runLocalCommandWithoutWait(instance, "sudo", cmd...)
	if err != nil {
		return err
	}

	internalIP, err := waitForVMToStart(instance)
	if err != nil {
		return err
	}

	instance.SetInternalIP(internalIP)

	aliasIP, err := metadata.Get("instance/network-interfaces/0/ip-aliases/0")
	if err != nil {
		return errors.Errorf("Error getting ip alias: %s", err)
	}

	// aliasIP is CIDR such as 10.128.0.3/32 . We need to get the IP address.
	externalIP := strings.Split(aliasIP, "/")[0]
	instance.SetExternalIP(externalIP)
	_, err = instance.RunLocalCommand("bash",
		instance.GetLocalSupportingFilePath("setup_iptables.sh"),
		externalIP, internalIP)
	return err
}

// The output will look like this:
// Expiry Time          MAC address        Protocol  IP address                Hostname        Client ID or DUID
// -------------------------------------------------------------------------------------------------------------------
// 2018-07-27 19:26:29  52:54:00:12:34:56  ipv4      192.168.122.89/24         win7            01:52:54:00:12:34:56
// returns internalIp
func waitForVMToStart(instance nestedVMInstanceInterface) (string, error) {
	// wait until the VM gets its IP address
	for i := 0; i < 10; i++ {
		output, err := instance.RunLocalCommand("sudo", "virsh", "net-dhcp-leases", "default")
		if err != nil {
			return "", err
		}

		lines := strings.Split(output, "\n")
		if len(lines) == 5 {
			fields := strings.Fields(lines[2])

			// the format of field 4 is 192.168.122.89/24, so we need another split
			// to get the IP address.
			return strings.Split(fields[4], "/")[0], nil
		}

		time.Sleep(1 * time.Minute)
	}

	return "", errors.New("Time out")
}

func UploadFileToNestedVM(instance nestedVMInstanceInterface, srcFilePath string, destDirectory string) error {
	fileName := filepath.Base(srcFilePath)
	destFilePath := filepath.Join(destDirectory, fileName)
	instance.Logf("Upload file %s -> %s to nested VM", srcFilePath, destFilePath)

	srcFile, err := os.Open(srcFilePath)
	if err != nil {
		return errors.WithStack(err)
	}

	conn, err := sshConnect(instance)
	if err != nil {
		return errors.WithStack(err)
	}
	defer conn.Close()

	client, err := sftp.NewClient(conn)
	if err != nil {
		return errors.WithStack(err)
	}

	destFile, err := client.Create(destFilePath)
	if err != nil {
		return errors.WithStack(err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
