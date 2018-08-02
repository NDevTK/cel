#! /bin/bash

# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

# Setup the iptabes so that the VM is connected to the outside
# the args are: external_ip internal_ip

set -x

external_ip=$1
internal_ip=$2

sudo iptables -t nat -A PREROUTING -d $external_ip -j DNAT --to-destination $internal_ip
sudo iptables -t nat -A POSTROUTING -d $internal_ip -j MASQUERADE
sudo iptables -A INPUT -p udp -j ACCEPT
sudo iptables -A FORWARD -p tcp -j ACCEPT
sudo iptables -A OUTPUT -p tcp -j ACCEPT
sudo iptables -A OUTPUT -p udp -j ACCEPT