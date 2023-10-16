package main

/*
 * Keycodes based on HID Usage Keyboard/Keypad Page modified for
 * wch keyboard firmware
 */

type Code interface {
	Code() uint8
	Type() uint8 // key, media, or mouse
}

type Keycode uint8

func (c Keycode) Code() uint8 {
	return uint8(c)
}

func (c Keycode) Type() uint8 {
	return 0x01
}

type Mousecode uint8

func (c Mousecode) Code() uint8 {
	return uint8(c)
}

func (c Mousecode) Type() uint8 {
	return 0x03
}

type Wheelcode uint8

func (c Wheelcode) Code() uint8 {
	return uint8(c)
}

func (c Wheelcode) Type() uint8 {
	return 0x03
}

type Mediacode uint8

func (c Mediacode) Code() uint8 {
	return uint8(c)
}

func (c Mediacode) Type() uint8 {
	if c == PLAY {
		return 0x02
	}
	// not sure if this is a bug in the original software...
	return 0x03
}

// Modifier flags are added together when used in combination
type Modifier uint8

/* USB HID Keyboard/Keypad Usage(0x07) */
const NOKEY Keycode = 0x00
const (
	A Keycode = iota + 0x04 /* 0x04 */
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M /* 0x10 */
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
	N1
	N2
	N3 /* 0x20 */
	N4
	N5
	N6
	N7
	N8
	N9
	N0
	ENTER
	ESCAPE
	BSPACE
	TAB
	SPACE
	MINUS
	EQUAL
	LBRACKET
	RBRACKET /* 0x30 */
	BSLASH   /* \ and |*/
	NONUS_HASH
	SCOLON /* ; and : */
	QUOTE  /* ' and " */
	GRAVE  /* ` and ~ */
	COMMA  /* , and < */
	DOT    /* . and > */
	SLASH  /* / and ? */
	CAPSLOCK
	F1
	F2
	F3
	F4
	F5
	F6
	F7 /* 0x40 */
	F8
	F9
	F10
	F11
	F12
	PSCREEN
	SCROLLLOCK
	PAUSE
	INSERT
	HOME
	PGUP
	DELETE
	END
	PGDOWN
	RIGHT
	LEFT /* 0x50 */
	DOWN
	UP
	NUMLOCK /* Numpad keys */
	KP_SLASH
	KP_ASTERISK
	KP_MINUS
	KP_PLUS
	KP_ENTER
	KP_1
	KP_2
	KP_3
	KP_4
	KP_5
	KP_6
	KP_7
	KP_8 /* 0x60 */
	KP_9
	KP_0
	KP_DOT
	NONUS_BSLASH
	APPLICATION
	POWER
	KP_EQUAL
)

/* Modifiers */
// Simultaneous modifier presses are added
const (
	NOMOD  Modifier = 0
	CTRL            = 0x01
	SHIFT           = 0x02
	ALT             = 0x04
	WIN             = 0x08
	RCTRL           = 0x10
	RSHIFT          = 0x20
	RALT            = 0x40
	RWIN            = 0x80
)

/* Media */
// 03 01 12 cd 00 00 00
const (
	PLAY   Mediacode = 0xcd
	PREV   Mediacode = 0xb6
	NEXT   Mediacode = 0xb5
	MUTE   Mediacode = 0xe2
	VOL_UP Mediacode = 0xe9
	VOL_DN Mediacode = 0xea
)

// 03 01 13 04 00 00 00
const (
	MS_LEFT   Mousecode = 0x01
	MS_RIGHT            = 0x02
	MS_CENTER           = 0x04
)

/* Mouse wheel */
// 03    01  13    00  00  00 ff   01
// magic key layer len seq ?? code mod
const (
	MS_WL      Wheelcode = 0x00
	MS_WL_UP   Wheelcode = 0x01
	MS_WL_DOWN Wheelcode = 0xff
)
