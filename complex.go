package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Complexer interface {
	Add() Complex
	Sub() Complex
	Mul() Complex
	Div() Complex
	Mod() float64
	ToString() string
}
type Complex struct {
	a float64
	b float64
}

// 加减乘除方法
func (com Complex) add(num Complex) Complex {
	numFirst := com.a + num.a
	numSecond := com.b + num.b
	return Complex{a: numFirst, b: numSecond}
}
func (com Complex) sub(num Complex) Complex {
	numFirst := com.a - num.a
	numSecond := com.b - num.b
	return Complex{a: numFirst, b: numSecond}
}
func (com Complex) mul(num Complex) Complex {
	numFirst := com.a*num.a - com.b*num.b
	numSecond := com.a*num.b + com.b*num.a
	return Complex{a: numFirst, b: numSecond}
}
func (com Complex) div(num Complex) Complex {
	numFirst := (com.a*num.a + com.b*num.b) / (math.Pow(num.a, 2) + math.Pow(num.b, 2))
	numSecond := (com.b*num.a - com.a*num.b) / (math.Pow(num.a, 2) + math.Pow(num.b, 2))
	return Complex{a: numFirst, b: numSecond}
}

// 求模长方法
func (com Complex) mod() float64 {
	sumSqu := math.Pow(com.a, 2) + math.Pow(com.b, 2)
	return math.Sqrt(sumSqu)
}

// toString方法
func (com Complex) toString() string {
	switch {
	case com.b > 0:
		return fmt.Sprintf("%.2f+%.2fi", com.a, com.b)
	case com.b == 0:
		return fmt.Sprintf("%.2f", com.a)
	case com.b < 0:
		return fmt.Sprintf("%.2f%.2fi", com.a, com.b)
	default:
		return ""
	}

}

// 单输入求模长初始
func oneInput() {
	var input string
	fmt.Println("输入一个复数")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
		return
	}
	//判断b是否为负
	if strings.Contains(input, "+") {
		var inputSlice = strings.Trim(input, "i")
		var inputSliceTure = strings.Split(inputSlice, "+")
		a1, err := strconv.ParseFloat(inputSliceTure[0], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		b1, err := strconv.ParseFloat(inputSliceTure[1], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		c = Complex{a: a1, b: b1}
	} else if strings.Contains(input, "-") {
		var inputSlice = strings.Trim(input, "i")
		var inputSliceTure = strings.Split(inputSlice, "-")
		a1, err := strconv.ParseFloat(inputSliceTure[0], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		b1, err := strconv.ParseFloat(inputSliceTure[1], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		c = Complex{a: a1, b: -b1}
	}
}

// 单输入toString初始
func oneInputStrange() {
	var num1, num2 float64
	fmt.Println("输入实部a")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("输入虚部b")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println(err)
		return
	}
	c = Complex{a: num1, b: num2}
}

// 双输入加减乘除初始
func twoInput() {
	var input1 string
	var input2 string
	fmt.Println("请输入第一个复数")
	_, err := fmt.Scanln(&input1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("请输入第二个复数")
	_, err = fmt.Scanln(&input2)
	if err != nil {
		fmt.Println(err)
		return
	}
	//判断b是否为负
	if strings.Contains(input1, "+") {
		var inputSlice = strings.Trim(input1, "i")
		var inputSliceTure = strings.Split(inputSlice, "+")
		a1, err := strconv.ParseFloat(inputSliceTure[0], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		b1, err := strconv.ParseFloat(inputSliceTure[1], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		c = Complex{a: a1, b: b1}
	} else if strings.Contains(input1, "-") {
		var inputSlice = strings.Trim(input1, "i")
		var inputSliceTure = strings.Split(inputSlice, "-")
		a1, err := strconv.ParseFloat(inputSliceTure[0], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		b1, err := strconv.ParseFloat(inputSliceTure[1], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		c = Complex{a: a1, b: -b1}
	}
	if strings.Contains(input2, "+") {
		var inputSlice = strings.Trim(input2, "i")
		var inputSliceTure = strings.Split(inputSlice, "+")
		a2, err := strconv.ParseFloat(inputSliceTure[0], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		b2, err := strconv.ParseFloat(inputSliceTure[1], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		d = Complex{a: a2, b: b2}
	} else if strings.Contains(input2, "-") {
		var inputSlice = strings.Trim(input2, "i")
		var inputSliceTure = strings.Split(inputSlice, "-")
		a2, err := strconv.ParseFloat(inputSliceTure[0], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		b2, err := strconv.ParseFloat(inputSliceTure[1], 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		d = Complex{a: a2, b: -b2}
	}
}

// 结果符号判断修正
func judge(real float64, imag float64) {
	if imag > 0 {
		fmt.Printf("结果是 %.2f + %.2fi\n", real, imag)
	} else if imag == 0 {
		fmt.Printf("结果是 %.2f\n", real)
	} else if imag < 0 {
		fmt.Printf("结果是 %.2f%.2fi\n", real, imag)
	}
}

var c, d Complex

func main() {
	hotaru := 0
	for hotaru == 0 {
		fmt.Println("扣1加减乘除,扣2求模,扣3toString,扣4不玩了")
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch choice {
		//选择了加减乘除
		case 1:
			fmt.Println("你选择了1,现在输入1、2、3、4中任一个,它们分别代表加、减、乘、除")
			var yourChoice int
			_, err = fmt.Scanln(&yourChoice)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("现在依次输入两个复数")
			twoInput() //进行双输入初始
			switch yourChoice {
			case 1:
				resultAdd := c.add(d)
				judge(resultAdd.a, resultAdd.b)
			case 2:
				resultSub := c.sub(d)
				judge(resultSub.a, resultSub.b)
			case 3:
				resultMul := c.mul(d)
				judge(resultMul.a, resultMul.b)
			case 4:
				resultDiv := c.div(d)
				judge(resultDiv.a, resultDiv.b)
			default:
				fmt.Println("请输入1、2、3、4中的一个!")
			}
		//选择了求模长
		case 2:
			oneInput() //进行单输入初始
			resultMod := c.mod()
			fmt.Printf("该复数的模是 %.2f\n", resultMod)
		//选择了toString
		case 3:
			oneInputStrange() //进行单输入初始
			resultToString := c.toString()
			fmt.Printf("得到字符串 %s\n", resultToString)
		//选择结束
		case 4:
			fmt.Println("现在退出了")
			hotaru = 1
		default:
			fmt.Println("请输入1、2、3、4中的一个!")
		}
	}
}
