#! /bin/bash

# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

set -x

sudo apt-get update
sudo apt-get install uml-utilities qemu-kvm bridge-utils virtinst libvirt-bin -y

# net-start will fail if the network is already started.
# we ignore this error here.
sudo virsh net-start default || true

sudo tunctl -t tap0
sudo ifconfig tap0 up
sudo brctl addif virbr0 tap0

echo "Current iptable rules:"
sudo iptables -L --line-numbers
sudo iptables -D FORWARD 4
sudo iptables -D FORWARD 4

# For debugging purpose, output the iptable rules. There should be no REJECT rules
# at this point
echo "There should be no REJECT rules now:"
sudo iptables -L --line-numbers