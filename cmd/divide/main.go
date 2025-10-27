package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/benstrobel-planqc/knowledge-sharing-go/pkg/mycoollib"
)

func main() {
	// Not handling errors for brevity
	dividend, _ := strconv.ParseFloat(os.Args[1], 32)
	divisor, _ := strconv.ParseFloat(os.Args[2], 32)
	result := mycoollib.Divide(dividend, divisor)
	fmt.Println(result)
}
