#! /bin/bash

# Copyright 2020 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
REQUIRED_PACKAGES=("autocutsel"
                   "gnome-core"
                   "gnome-panel"
                   "gnome-shell"
                   "gnome-themes-standard"
                   "tightvncserver"
                   "xtightvncviewer"
                   "ubuntu-gnome-desktop")

# Check to see if all of the required packages are installed.
function join_by { local IFS="$1"; shift; echo "$*"; }

PACKAGE_QUERY_OUTPUT=$(dpkg-query -l $(join_by " " "${REQUIRED_PACKAGES[@]}") 2>&1)
echo "${PACKAGE_QUERY_OUTPUT}"

# The installation should occur only once per setup, especially since the
# installation requires reboot.
NOT_FOUND="dpkg-query: no packages found matching "
if [[ "${PACKAGE_QUERY_OUTPUT}" =~ .*"${NOT_FOUND}".* ]]; then
  echo "==== Upgrading system ===="
  sudo apt-get update
  sudo apt-get upgrade -y

  echo "==== Installing packages ===="
  for PACKAGE in ${REQUIRED_PACKAGES[@]}; do
    sudo apt-get install -y ${PACKAGE}
  done

  echo "==== Setting vncserver startup script to autorun on login ==="
  sudo cp --force /cel/supporting_files/configure_vnc_server.sh /etc/profile.d/
  sudo chmod +x /etc/profile.d/configure_vnc_server.sh

  touch ~/.Xresources
  echo "==== Rebooting ===="
  # First time installation of gnome requires a reboot.
  # Note that cel_agent.exe specifically parses for the
  # "==== Rebooting ====" output string to determine if
  # the machine will reboot.
  sudo reboot
else
  echo "==== All packages installed. ===="
fi