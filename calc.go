package calc_lib

type (
	Addition       struct{}
	Subtraction    struct{}
	Multiplication struct{}
	Division       struct{}
)

func (this *Addition) Calculate(a, b int) int       { return a + b }
func (this *Subtraction) Calculate(a, b int) int    { return a - b }
func (this *Multiplication) Calculate(a, b int) int { return a * b }
func (this *Division) Calculate(a, b int) int       { return a / b }
