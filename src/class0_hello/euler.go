package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	c := 3 + 4i
	var d = 1i
	fmt.Println(cmplx.Abs(c), cmplx.Pow(math.E, d*math.Pi))
	fmt.Println(cmplx.Exp(d * math.Pi))
	//cmplx.Pow(math.E, d*math.Pi)   e^(iΠ)=-1,出来是浮点数
	const a, b = 3, 4                 //定义常量可以不标注数据类型,而且可以自转换各种动类型
	var e = int(math.Sqrt(a*a + b*b)) //math.Sqrt()接受的参数必须是float64
	fmt.Println(e)
}
