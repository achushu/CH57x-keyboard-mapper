package main

import (
	"fmt"
	"os"

	"github.com/achushu/hid"
)

var (
	Custom = []Sequence{
		{NOMOD, N1}, {NOMOD, N2}, {NOMOD, N3}, {NOMOD, N4},
		{NOMOD, N5}, {NOMOD, N6}, {NOMOD, N7}, {NOMOD, N8},
		{NOMOD, N9}, {NOMOD, N0}, {NOMOD, MINUS}, {NOMOD, EQUAL},

		{NOMOD, COMMA}, {NOMOD, SLASH}, {NOMOD, DOT},
		{NOMOD, SCOLON}, {NOMOD, BSLASH}, {NOMOD, QUOTE},
	}
)

func main() {
	var err error

	if !hid.Supported() {
		fmt.Println("this platform / binary does not support HID")
		os.Exit(1)
	}

	dev := SelectInterface()

	if dev.Path == "" {
		fmt.Printf("could not find the device interface")
		os.Exit(2)
	}

	kbd, err := NewKeyboard(dev)
	if err != nil {
		fmt.Printf("error opening device %s\n%s\n", dev.Path, err)
		os.Exit(2)
	}
	defer kbd.Close()
	fmt.Println("connected to keyboard")

	err = kbd.SendHello()
	if err != nil {
		fmt.Println("error writing to device:", err)
	}
	fmt.Println("sent hello")

	kbd.BindMapping(MapKeys(Custom))
	fmt.Println("done!")
}

func SelectInterface() hid.DeviceInfo {
	var info hid.DeviceInfo

	devices := hid.Enumerate(VENDOR_ID, PRODUCT_ID)
	if len(devices) == 0 {
		fmt.Println("no macro keyboard detected")
		os.Exit(1)
	}
	for _, d := range devices {
		if d.Interface == INTERFACE {
			return d
		}
	}
	return info
}
