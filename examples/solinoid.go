package main

import (
	"fmt"

	"github.com/1lann/fabricator"
)

func main() {
	set := fabricator.ResultSet{
		{0.58, 0.67},
		{1, 1.68},
		{1.6, 3.69},
		{2.12, 6.03},
		{2.55, 8.71},
		{2.7, 9.38},
	}

	fb := set.Fabricate()
	fb.SetErrorRange(0.1)

	newSet := fabricator.ResultSet{
		{0.7, 0},
		{1.1, 0},
		{1.4, 0},
		{1.8, 0},
		{2.2, 0},
		{2.5, 0},
	}

	for i := 0; i < 3; i++ {
		trial := newSet.Copy()
		fb.Apply(false, trial)
		fmt.Println("-- trial", i, "--")
		for _, res := range trial {
			fmt.Println(res)
		}
	}
}
