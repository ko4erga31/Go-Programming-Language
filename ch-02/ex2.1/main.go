package main

import (
	"fmt"

	"github.com/ko4erga31/Go-Programming-Language/ch-02/ex2.1/tempconv"
)

func main() {
	fmt.Println(tempconv.CToK(0))       // 273.15K
	fmt.Println(tempconv.FToK(32))      // 273.15K
	fmt.Println(tempconv.AbsoluteZeroK) // 0K
}
