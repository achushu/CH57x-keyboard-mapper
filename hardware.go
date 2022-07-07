package main

import (
	"fmt"
	"time"

	"github.com/achushu/hid"
)

// manu: wch.cn, product: CH57x, vendorID 4489, productID: 34960

const (
	VENDOR_ID  = 4489
	PRODUCT_ID = 34960
	INTERFACE  = 1 // the programmable interface
)

type MacroType uint8

const (
	MACRONONE  MacroType = 0x00
	MACROKEYS            = 0x01
	MACROPLAY            = 0x02
	MACROMOUSE           = 0x03
)

type Layer uint8

const (
	LAYER1 Layer = 0x10
	LAYER2       = 0x20
	LAYER3       = 0x30
)

type Macro struct {
	Type  MacroType
	Layer Layer
	Key   Key
	Combo []Sequence
}

func NewMacro(key Key) *Macro {
	return &Macro{
		Layer: LAYER1,
		Key:   key,
		Combo: make([]Sequence, 0),
	}
}

func NewMacroSequence(key Key, seq Sequence) *Macro {
	m := &Macro{
		Layer: LAYER1,
		Key:   key,
		Combo: []Sequence{seq},
	}
	switch seq.Key.(type) {
	case Keycode:
		m.Type = MACROKEYS
	case Mousecode:
		m.Type = MACROMOUSE
	}
	return m
}

func (m *Macro) Add(mod Modifier, key Code) error {
	switch key.(type) {
	case Keycode:
		if m.Type == MACRONONE {
			m.Type = MACROKEYS
		} else if m.Type != MACROKEYS {
			return ErrTypeMixing
		}
	case Mousecode:
		if m.Type == MACRONONE {
			m.Type = MACROMOUSE
		} else if m.Type != MACROMOUSE {
			return ErrTypeMixing
		}
	}

	m.Combo = append(m.Combo, Sequence{mod, key})
	return nil
}

func (m *Macro) AddKey(key Code) error {
	return m.Add(NOMOD, key)
}

func (m *Macro) Len() int {
	return len(m.Combo)
}

type Sequence struct {
	Mod Modifier
	Key Code
}

var EmptySequence = Sequence{NOMOD, NOKEY}

type Key uint8

const (
	KEY1 Key = iota + 1
	KEY2
	KEY3
	KEY4
	KEY5
	KEY6
	KEY7
	KEY8
	KEY9
	KEY10
	KEY11
	KEY12
	ROT1CCW
	ROT1
	ROT1CW
	ROT2CCW
	ROT2
	ROT2CW
)

type Keyboard struct {
	dev *hid.Device
}

func NewKeyboard(info hid.DeviceInfo) (*Keyboard, error) {
	d, err := info.Open()
	if err != nil {
		return nil, err
	}
	return &Keyboard{d}, nil
}

func (k *Keyboard) Close() {
	k.dev.Close()
}

func (k *Keyboard) Send(data []byte) error {
	_, err := k.dev.Write(append([]byte{3}, data...))
	time.Sleep(15 * time.Millisecond)
	return err
}

func (k *Keyboard) SendHello() error {
	req := make([]byte, 64)
	return k.Send(req)
}

func (k *Keyboard) sendKeybindStart() error {
	req := make([]byte, 64)
	req[0] = 0xa1
	req[1] = 0x01
	return k.Send(req)
}

func (k *Keyboard) sendKeybindEnd() error {
	req := make([]byte, 64)
	req[0] = 0xaa
	req[1] = 0xaa
	return k.Send(req)
}

func (k *Keyboard) BindKeyMacro(macro *Macro) error {
	var err error
	err = k.sendKeybindStart()
	if err != nil {
		return err
	}

	// header
	req := make([]byte, 64)
	req[0] = byte(macro.Key)                      // key ID
	req[1] = byte(macro.Layer) + byte(macro.Type) // layer and macro type
	req[2] = byte(macro.Len())                    // length

	var combo []Sequence
	if macro.Type == MACROKEYS {
		// start key sequences with a blank (for some reason... bug?)
		combo = append([]Sequence{EmptySequence}, macro.Combo...)
	} else {
		combo = macro.Combo
	}
	// bind the key sequence
	for i, seq := range combo {
		req[3] = byte(i)
		req[4] = byte(seq.Mod)
		req[5] = byte(seq.Key.Code())
		err = k.Send(req)
		if err != nil {
			return err
		}
	}

	return k.sendKeybindEnd()
}

func (k *Keyboard) BindMapping(mapping []*Macro) {
	for _, m := range mapping {
		err := k.BindKeyMacro(m)
		if err != nil {
			fmt.Println("error binding key", m.Key, err)
		} else {
			fmt.Println("bound key", m.Key)
		}
	}
}
func MapKeys(seqs []Sequence) []*Macro {
	mapping := make([]*Macro, len(seqs))
	for i, s := range seqs {
		mapping[i] = NewMacroSequence(Key(i+1), s)
	}
	return mapping
}
