package main

import (
	"fmt"
	"strconv"
)

// 用于求私钥d
var x int = 0 //调用exGcd()之后，x的值就是私钥的d的值
var y int = 0 //线性同余方程式y的值，没用，用来占位

func main() {
	RSA_94(7, 131, 17, stringToSlice_94("RSA"))
}

func stringToSlice_94(str string) int { //如果输入是字符串
	var slice_ []int
	for index, _ := range str { //将字符串输入转换成int存储再切片中
		slice_ = append(slice_, int(str[index]))
	}
	string_ := ""
	for index, _ := range str { //将切片中的数值连接
		string_ += strconv.Itoa(slice_[index])
	}
	a, _ := strconv.Atoi(string_) //将string转换成int
	return a % 1000               //因为字符串转换拼接出来的值很大，会超过int范围，因此做取模操作
}

// RSA算法封装
func RSA_94(e int, p int, q int, m int) {
	n := p * q
	Fn := (p - 1) * (q - 1)
	fmt.Printf("用户输入的值为: e=%v, p=%v, q=%v, m=%v\n", e, p, q, m)
	fmt.Printf("(1)n=%v, n的欧拉函数值=%v\n", n, Fn)
	exGcd_94(e, Fn)
	temp := modRepeateQuadratic_94(m, e, n)
	fmt.Printf("(3)加密结果(密文c)= %v\n", temp)
	fmt.Printf("(4)解密结果(明文m)= %v", modRepeateQuadratic_94(temp, x+Fn, n))
}

// 求私钥d
// a：公钥e。n：φ(n)
func privateKey_94(a int, n int) int {
	if n == 0 {
		x = 1
		y = 0
		return a
	} else if a == 0 {
		x = 0
		y = 1
		return n
	} else {
		c := privateKey_94(n, a%n)
		tmp := x
		x = y
		y = tmp - a/n*y
		return c
	}
}

// 求私钥d
// a：公钥e。n：φ(n)
func exGcd_94(a int, n int) {
	privateKey_94(a, n)
	fmt.Printf("(2)私钥d= %v\n", x+n)
}

// 模重复平方法
func modRepeateQuadratic_94(b int, n int, m int) int {
	s := 1                              //用于累加累乘得到结果
	x := strconv.FormatInt(int64(n), 2) //将数n转换为二进制
	//从低位开始遍历（右边）
	fmt.Printf("  【模重复平方法过程】")
	for i := len(x) - 1; i >= 0; i-- {
		if int(x[i])-48 == 1 { //减48：因为ASCII码
			s = (s * b) % m //累乘取模
		}
		b = (b * b) % m //将b指数+1
		fmt.Printf("[i=%v,中间值=%v]", i, s)
	}
	fmt.Println()
	return s
}
