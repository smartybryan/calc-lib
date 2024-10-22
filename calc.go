package calc_lib

type Addition struct{}

func (this *Addition) Calculate(a, b int) int {
	return a + b
}
