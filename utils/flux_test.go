package utils

import (
	"fmt"
	"testing"
)

func TestFlux(t *testing.T) {
	result, err := Flux("f53276ad225e")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}
