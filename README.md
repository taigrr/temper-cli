# temper

[![Latest Release](https://img.shields.io/github/release/taigrr/temper-cli.svg?style=for-the-badge)](https://github.com/taigrr/temper-cli/releases)
[![Software License](https://img.shields.io/badge/license-0BSD-blue.svg?style=for-the-badge)](/LICENSE)
[![Go ReportCard](https://goreportcard.com/badge/github.com/taigrr/temper-cli?style=for-the-badge)](https://goreportcard.com/report/taigrr/temper-cli)

A simple golang command-line that takes a reading of a TEMPer USB and prints it
to STDOUT. Uses the [temper](https://github.com/taigrr/temper) library.

## Installation

Make sure you have a working Go environment (Go 1.26+ is required).
See the [install instructions](http://golang.org/doc/install.html).

To install temper-cli, run:

    go install github.com/taigrr/temper-cli@latest

## Usage

```
temper-cli          # print temperature in Celsius
temper-cli -f       # print temperature in Fahrenheit
```


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
This cli is tested with the model available [here](https://www.pcsensor.com/product-details?product_id=782&brd=1).
This is not an endorsement of the product.

Make sure your user is part of the `plugdev` group and reload the rules with
`sudo udevadm control --reload-rules`.
Unplug and replug the device.
