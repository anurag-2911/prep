package calculator

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Calc() {
	// basicCalculation()
}

func basicCalculation() {
	if len(os.Args) < 4 {
		fmt.Println("usage: calc <operation> <operand1> <operand2>")
		fmt.Println("Operations: add, subtract, multiply, divide, power, modulus")
		os.Exit(1)
	}

	operation := os.Args[1]
	operand1, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("invalid first operand ", operand1)
		os.Exit(1)
	}
	operand2, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("invalid second operand ", operand2)
		os.Exit(1)
	}
	calc := &Calculator{}

	switch operation {
	case "add":
		{
			fmt.Println(calc.Add(int(operand1), int(operand2)))
		}
	case "subtract":
		{
			fmt.Println(calc.Subtract(int(operand1), int(operand2)))
		}
	case "pow":
		{
			fmt.Println(calc.Power(operand1, operand2))
		}
	default:
		{
			fmt.Println("not supported operation")
			os.Exit(1)
		}
	}
}

type Calculator struct {
}

func (c *Calculator) Add(x, y int) int {
	return x + y
}
func (c *Calculator) Subtract(x, y int) int {
	return x - y
}
func (c *Calculator) Multiply(x, y int) int {
	return x * y
}
func (c *Calculator) Divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Division by zero is not allowed")
	}

	return x / y, nil
}
func (c *Calculator) Power(x, y float64) float64 {
	return math.Pow(x, y)
}
func (c *Calculator) SquareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("can't take the square root of negative number")
	}
	return math.Sqrt(x), nil
}
func (c *Calculator) Modulus(x, y int) int {
	return x % y
}
