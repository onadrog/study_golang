package main

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestPromptGameRecursive(t *testing.T) {
	inputs := [][]string{
		{"7", "4", "1"},
		{"8", "az", "oi"},
	}

	for _, a := range inputs {
		for j := range a {
			main()
			_, err := io.WriteString(os.Stdout, a[j])
			if err != nil {
				log.Fatal(err)
			}
			t.Log("Ya win son!")

		}
	}
}
