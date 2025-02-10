package main

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadState(t *testing.T) {
	f, _ := os.ReadFile("input.txt")
	fmt.Println(f)
}
