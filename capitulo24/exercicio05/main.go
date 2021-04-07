package main

import (
	"testing"
)

func TestVal(t *testing.T) {
	testes := 10

	if testes != 10 {
		t.Errorf("Esperava um int, recebeu outro tipo.")
	}

}
