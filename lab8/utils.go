package main

import (
	"fmt"
	"math/rand"
)

func GenerateId() string {
	r := make([]byte, 16)
	rand.Read(r)
	return fmt.Sprintf("%x", r)
}
