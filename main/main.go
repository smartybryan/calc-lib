package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/smartybryan/calc-lib"
)

func main() {
	handler := NewHandler(os.Stdout, &calc_lib.Addition{})
	err := handler.Handle(os.Args[1:])
	if err != nil {
		panic(err)
	}
}

type Calculator interface {
	Calculate(a, b int) int
}

type Handler struct {
	stdout     io.Writer
	calculator Calculator
}

func NewHandler(stdout io.Writer, calculator Calculator) *Handler {
	return &Handler{stdout, calculator}
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return errWrongArgCount
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: '%s'", errInvalidArgument, args[0])
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: '%s'", errInvalidArgument, args[1])
	}
	result := this.calculator.Calculate(a, b)
	_, err = fmt.Fprint(this.stdout, result)
	return err
}

var (
	errWrongArgCount   = errors.New("usage: calculator <a> <b>")
	errInvalidArgument = errors.New("invalid syntax")
)
