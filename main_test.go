package main

import (
	"fmt"
	"testing"
)

func TestNeighbors(t *testing.T) {
	world := NewWorld(10, 10)
	world.Cells = [][]bool{
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, true, true, true, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false}}
	fmt.Println(world.Neighbors(4, 4))
}
