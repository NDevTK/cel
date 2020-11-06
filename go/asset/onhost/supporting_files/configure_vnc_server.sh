#! /bin/bash

# Copyright 2020 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

echo '===== Set VNC server password ===='
passwrd="pass@wrd"
sudo apt-get install -y expect

/usr/bin/expect <<EOF

spawn "/usr/bin/vncpasswd"
expect "Password:"
send "$passwrd\r"
expect "Verify:"
send "$passwrd\r"
expect "Would you like to enter a view-only password (y/n)?"
send "n\r"
expect eof
exit
EOF

echo '==== Configure VNC server to use GNOME desktop ===='
vncserver
vncserver -kill :1

cat > ~/.vnc/xstartup <<EOF
export PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/root/bin:$PATH

autocutsel -fork
xrdb $HOME/.Xresources
xsetroot -solid grey
export XKL_XMODMAP_DISABLE=1
export XDG_CURRENT_DESKTOP="GNOME-Flashback:Unity"
export XDG_MENU_PREFIX="gnome-flashback-"
unset DBUS_SESSION_BUS_ADDRESS
gnome-session --session=gnome-flashback-metacity --disable-acceleration-check --debug &
EOF

vncserver -geometry 1920x1080