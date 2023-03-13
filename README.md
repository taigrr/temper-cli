# temper

[![Latest Release](https://img.shields.io/github/release/taigrr/temper.svg?style=for-the-badge)](https://github.com/taigrr/temper/releases)
[![Software License](https://img.shields.io/badge/license-0BSD-blue.svg?style=for-the-badge)](/LICENSE)
[![Go ReportCard](https://goreportcard.com/badge/github.com/taigrr/temper?style=for-the-badge)](https://goreportcard.com/report/taigrr/temper)

A simple golang command-line that takes a reading of a TEMPer USB and prints it
to STDOUT. Uses the [temper](https://github.com/taigrr/temper) library.

## Installation

Make sure you have a working Go environment (Go 1.12+ is required).
See the [install instructions](http://golang.org/doc/install.html).

To install temper-cli, run:

    go get github.com/taigrr/temper-cli


On Linux you need to set up some udev rules to be able to access the device as
a non-root/regular user.
Edit `/etc/udev/rules.d/99-temper.rules` and add these lines:

```
SUBSYSTEM=="hidraw", ATTRS{idVendor}=="1a86", ATTRS{idProduct}=="e025", GROUP="plugdev", SYMLINK+="temper%n"
SUBSYSTEM=="hidraw", ATTRS{idVendor}=="0c45", ATTRS{idProduct}=="7401", GROUP="plugdev", SYMLINK+="temper%n"
SUBSYSTEM=="hidraw", ATTRS{idVendor}=="0c45", ATTRS{idProduct}=="7402", GROUP="plugdev", SYMLINK+="temper%n"
SUBSYSTEM=="hidraw", ATTRS{idVendor}=="1130", ATTRS{idProduct}=="660c", GROUP="plugdev", SYMLINK+="temper%n"
```
Note that there are many versions of the TEMPer USB and your
`idVendor` and `idProduct` ATTRs may differ.

Make sure your user is part of the `plugdev` group and reload the rules with
`sudo udevadm control --reload-rules`.
Unplug and replug the device.
