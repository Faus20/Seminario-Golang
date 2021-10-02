package model

import (
	"errors"
	"strconv"
)

type Result struct {
	Type   string
	Length int
	Value  string
}

const minL = 5

func (c *Result) GetResult(cad string) (*Result, error) {
	if len(cad) >= minL {
		t := cad[0:2]
		l := cad[2:4]
		v := cad[4:]

		lInt, err := strconv.Atoi(l)
		if err == nil && lInt == len(v) {
			if checkType(t) == nil {
				r := NewResult(t, v, lInt)
				return &r, nil
			}
		}
	}

	return nil, errors.New("Cadena INVALIDA: " + cad)

}

func checkType(t string) error {
	if t == "NN" {
		return nil
	}
	if t == "TX" {
		return nil
	}

	return errors.New("cadena no valida")

}

func NewResult(typ string, value string, length int) Result {
	return Result{typ, length, value}
}
