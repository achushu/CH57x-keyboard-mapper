package main

import "fmt"

var (
	ErrTypeMixing = fmt.Errorf("error cannot mix keyboard and mouse in a macro")
)
