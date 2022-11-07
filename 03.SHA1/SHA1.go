package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	SHA1_94("abc")
	SHA1_94("2021131094")
	SHA1_94("abcdbcdecdefdefgefghfghighijhijkijkljklmklmnlmnomnopnopq")
}

func SHA1_94(str1 string) {
	//将参数转换为二进制
	str2 := stringToBinary(str1)
	//将二进制划分为若干个512bit块，存放到一个切片中
	str3 := div512BitBlock(str2) //len(str3)是512bit块的数量
	//将每一个512bit扩展成80个W，用切片存放每一个80w
	str__ := "0110011101000101001000110000000111101111110011011010101110001001100110001011101011011100111111100001000000110010010101000111011011000011110100101110000111110000"
	fmt.Println("第1个512bit块的80轮转换：")
	x := _80translate_94(str__[0:32],
		str__[32:64],
		str__[64:96],
		str__[96:128],
		str__[128:160], preHandle_94(str3[0]), 1)
	for i := 1; i < len(str3); i++ {
		fmt.Printf("第%v个512bit块：\n", i+1)
		x = _80translate_94(x[0:32], //1-32
			x[32:64],   //33-64
			x[64:96],   //65-96
			x[96:128],  //97-128
			x[128:160], //129-160
			preHandle_94(str3[i]), 0)
		fmt.Println(binaryToHex_(x))
	}
	fmt.Printf("\n%v的SHA1加密结果为:\n%v", str1, binaryToHex_(x))
} //SHA1算法封装

func div512BitBlock(str string) []string { //传入的字符串会被按照512bit来分成多个组，每个组512bit，用一维数组来装
	strLength := len(str)
	blockGroupNum := int(strLength / 512)       //512组数目：blockGroupNum
	residue := strLength % 512                  //不足512bit的部分长度
	fullFill512Bit_0 := str[512*blockGroupNum:] //取出不足512bit的部分

	//下面开始对这一部分进行补齐512bit
	//计算需要补0的个数
	full0Num := 0
	notEnough65 := 0
	if 512-residue < 65 { //如果剩余位置不足65位
		full0Num = 960 - (residue + 1) //需要0的个数：full0Num个
		notEnough65 = 1
	} else { //如果剩余位置够65位
		full0Num = 448 - (residue + 1) //需要0的个数：full0Num个
	}
	//对不足512bit的部分补1
	fullFill512Bit_1 := fullFill512Bit_0 + "1"
	//对不足512bit的部分补0
	fullFill512Bit_2 := fullFill512Bit_1
	for i := 0; i < full0Num; i++ {
		fullFill512Bit_2 += "0"
	}
	//添加64bit的数值：内容长度
	residueToBinary := strconv.FormatInt(int64(residue), 2) //将内容长度从十进制转为二进制
	fullFill512Bit_3 := residueToBinary                     //声明等待完善的64bit部分
	//前面要补0
	if len(residueToBinary) != 64 { //如果这个内容长度不足64bit
		for j := 0; j < 64-len(residueToBinary); j++ { //看看我们缺64-len(residueToBinary)个bit位
			fullFill512Bit_3 = "0" + fullFill512Bit_3
		}
	}
	fullFill512Bit_4 := fullFill512Bit_2 + fullFill512Bit_3 //不足512bit的部分补足完毕
	//下面我们来将输入转换成的若干个512bit块输出，输出到一个一维数组,这个一维数组的每一个值就是512bit块中的二进制串
	arr := make([]string, 0)
	for i := 0; i < blockGroupNum+1+notEnough65; i++ { //每一个512块（不包括补的）
		arr = append(arr, fullFill512Bit_4[i*512:(i+1)*512])
	}
	//arr = append(arr, fullFill512Bit_4) //加上补的那块
	return arr
} //传入的字符串会被按照512bit来分成多个组，每个组512bit，用一维数组来装

func stringToBinary(str string) string { //字符串转二进制如："abc"==>011000010110001001100011
	result := ""
	sliceTep := make([]string, len(str))
	for i := 0; i < len(str); i++ {
		b := str[i]                  //a,char
		x := int64(b)                //a的ASC码值,int64
		y := strconv.FormatInt(x, 2) //a的ASC码值的二进制,string
		if len(y) != 8 {
			temp := string(y)
			for j := 0; j < 8-len(y); j++ { //看看我们缺8-len(y)个bit位
				sliceTep[i] = "0" + temp
			}
			result += sliceTep[i]
		}
	}
	return result
} //字符串转二进制如："abc"==>011000010110001001100011

func preHandle_94(str string) [80]string { //预处理部分：将被一个512bit块扩展80个W，每个W用二进制表示

	//0 <= t <= 15块：照搬
	result := [80]string{}
	for i := 0; i < 16; i++ {
		result[i] = str[i*32 : (i+1)*32]
	}
	//16 <= t <= 79：根据公式计算
	for i := 16; i < 80; i++ {
		temp := make([]string, 0)
		result_ := [32]string{}
		for j := 0; j < 32; j++ {
			i_3, _ := strconv.Atoi(string(result[i-3][j]))
			i_8, _ := strconv.Atoi(string(result[i-8][j]))
			i_14, _ := strconv.Atoi(string(result[i-14][j]))
			i_16, _ := strconv.Atoi(string(result[i-16][j]))
			temp = append(temp, strconv.Itoa((i_3 ^ i_8 ^ i_14 ^ i_16)))
		}
		for k := 0; k < len(temp); k++ {
			result_[k] = temp[k]
		}
		//循环左移一位
		RESULT := ""
		for i := 0; i < 32; i++ {
			RESULT += result_[i]
		}
		result[i] = moveNBit(1, RESULT)
	}
	//实验作业要求输出
	fmt.Println("===============扩展的80个W块开始输出===============")
	for i := 0; i < 80; i++ {
		if i == 0 || i == 1 || i == 14 || i == 15 || i == 16 || i == 79 {
			fmt.Printf("W[%v]=%v\n", i, binaryToHex(result[i]))
		}

	}
	fmt.Println("===============扩展的80个W块输出结束===============")
	return result
} //预处理部分：将被一个512bit块分成80个W，每个W用二进制表示

func _80translate_94(a string, b string, c string, d string, e string, str [80]string, time int) string {
	fmt.Println("===============80轮转换输出如下===============")
	fmt.Printf(binaryToHex(a) + ",")
	fmt.Printf(binaryToHex(b) + ",")
	fmt.Printf(binaryToHex(c) + ",")
	fmt.Printf(binaryToHex(d) + ",")
	fmt.Printf(binaryToHex(e) + ",")
	fmt.Println()

	//记录80轮转换之前的结果
	a_old := a
	b_old := b
	c_old := c
	d_old := d
	e_old := e

	for i := 0; i < 80; i++ {
		T := T_94(a, b, c, d, e, str[i], i)
		e = d
		d = c
		c = moveNBit(30, b)
		b = a
		a = T
		fmt.Printf(binaryToHex(a) + ",")
		fmt.Printf(binaryToHex(b) + ",")
		fmt.Printf(binaryToHex(c) + ",")
		fmt.Printf(binaryToHex(d) + ",")
		fmt.Printf(binaryToHex(e) + ",")
		fmt.Printf("[%v]\n", i)
	}
	if time == 1 {
		return addModW(a, b, c, d, e, 0b01100111010001010010001100000001, 0b11101111110011011010101110001001,
			0b10011000101110101101110011111110, 0b00010000001100100101010001110110, 0b11000011110100101110000111110000)
	} else {
		return addModW(a, b, c, d, e, stringToInt(a_old), stringToInt(b_old), stringToInt(c_old), stringToInt(d_old), stringToInt(e_old))
	}

} //80轮转换

func T_94(a string, b string, c string, d string, e string, Wt string, t int) string {
	a_ := moveNBit(5, a)
	f := f_94(b, c, d, t)
	Kt := Kt_94(t)
	//return a_ + f + e + Kt + Wt
	a1 := add(stringToInt(a_), stringToInt(f))  //a_ + f
	a2 := add(stringToInt(a1), stringToInt(e))  //a_ + f + e
	a3 := add(stringToInt(a2), stringToInt(Kt)) //a_ + f + e + Kt
	a4 := add(stringToInt(a3), stringToInt(Wt)) //a_ + f + e
	return a4
} //T函数

func f_94(x string, y string, z string, t int) string { //输入是二进制串，共32bit
	if 0 <= t && t <= 19 {
		return XOR_94(and_94(x, y), and_94(reverse(x), z))
	} else if 20 <= t && t <= 39 {
		return XOR_94(XOR_94(x, y), z)
	} else if 40 <= t && t <= 59 {
		return XOR_94(XOR_94(and_94(x, y), and_94(x, z)), and_94(y, z))
	} else if 60 <= t && t <= 79 {
		return XOR_94(XOR_94(x, y), z)
	} else {
		return "-1"
	}
} //f函数

func Kt_94(t int) string {
	if 0 <= t && t <= 19 {
		return "5a827999"
	} else if 20 <= t && t <= 39 {
		return "6ed9eba1"
	} else if 40 <= t && t <= 59 {
		return "8f1bbcdc"
	} else if 60 <= t && t <= 79 {
		return "ca62c1d6"
	} else {
		return "-1"
	}
} //Kt函数

func XOR_94(x string, y string) string { //两字符串异或,每个字符串32bit
	temp := make([]string, 0)
	for j := 0; j < 32; j++ {
		if j == 31 {
			a, _ := strconv.Atoi(string(x[31]))
			b, _ := strconv.Atoi(string(y[31]))
			temp = append(temp, strconv.Itoa(a^b))
			continue
		}
		a, _ := strconv.Atoi(x[j : j+1]) //将每个数字取出，string=>int
		b, _ := strconv.Atoi(y[j : j+1])
		temp = append(temp, strconv.Itoa(a^b)) //取出的数字异或，然后添加到切片。int=>string
	}
	var str string = "" //用于输出
	for k := 0; k < len(temp); k++ {
		str += temp[k]
	}
	return str
} //两字符串异或

func reverse(str string) string {
	var result string = ""
	for i := 0; i < 32; i++ {
		a, _ := strconv.Atoi(str[i : i+1])
		if a == 0 {
			result += "1"
		} else {
			result += "0"
		}
	}
	return result
} //将字符串取反

func and_94(x string, y string) string {
	//1010 and 1101 =
	temp := make([]string, 0)
	for j := 0; j < 32; j++ {
		if j == 31 {
			a, _ := strconv.Atoi(string(x[31]))
			b, _ := strconv.Atoi(string(y[31]))
			c := 0
			if a == 1 && b == 1 {
				c = 1
			}
			temp = append(temp, strconv.Itoa(c))
			continue
		}
		a, _ := strconv.Atoi(x[j : j+1]) //将每个数字取出，string=>int
		b, _ := strconv.Atoi(y[j : j+1])
		c := 0
		if a == 1 && b == 1 {
			c = 1
		}
		temp = append(temp, strconv.Itoa(c)) //取出的数字异或，然后添加到切片。int=>string
	}
	var str string = "" //用于输出
	for k := 0; k < len(temp); k++ {
		str += temp[k]
	}
	return str
} //两字符串与运算

func moveNBit(n int, str string) string {
	/*//左移实现
	var result string = ""
	var j int = n
	for i := 0; i < len(str)-n; i++ {
		result += str[j : j+1]
		j++
	}
	for i := 0; i < n; i++ {
		result += "0"
	}
	return result
	*/

	//循环左移n位实现
	//10010=>01010:循环左移两位
	var result string = ""
	var j int = n
	for i := 0; i < len(str)-n; i++ {
		result += str[j : j+1]
		j++
	}
	for i := 0; i < n; i++ {
		result += str[i : i+1]
	}
	return result
} //循环左移n位

func addModW(str01 string, str02 string, str03 string, str04 string, str05 string,
	h0 int, h1 int, h2 int, h3 int, h4 int) string {
	/*
		a := 0b01100111010001010010001100000001
		b := 0b11101111110011011010101110001001
		c := 0b10011000101110101101110011111110
		d := 0b00010000001100100101010001110110
		e := 0b11000011110100101110000111110000
	*/
	//应该将string转成int，然后相加，取模2^32
	//现在我们的问题是：如何将string转为int 。比如"101"，要转为int的101
	H0_ := add(h0, stringToInt(str01))
	H1_ := add(h1, stringToInt(str02))
	H2_ := add(h2, stringToInt(str03))
	H3_ := add(h3, stringToInt(str04))
	H4_ := add(h4, stringToInt(str05))
	//return binaryToHex(H0_) + binaryToHex(H1_) + binaryToHex(H2_) + binaryToHex(H3_) + binaryToHex(H4_)
	return H0_ + H1_ + H2_ + H3_ + H4_
} //mode 2^32加法

func add(str01 int, str02 int) string {
	s := str01 + str02
	x := strconv.FormatInt(int64(s), 2)
	c := x
	if len(x) == 32 { //如果没进位就不变
		c = x
	} else if len(x) == 33 { //如果进位就取模
		c = x[1:]
	} else { //不足32位要补位，前面补0
		for i := 0; i < 32-len(x); i++ {
			c = "0" + c
		}
	}
	return c
} //相加并取模,w位的数，就取模2^w

func stringToInt(str string) int { // 现在我们的问题是：如何将string转为int 。比如"101"，要转为十进制的5
	var num float64 = 0
	var j int = 0

	//如果传入的是Kt（十六进制）
	if len(str) == 8 {
		for i := len(str) - 1; i >= 0; i-- {
			if string(str[i]) == "1" {
				num += math.Pow(2, float64(j))
			}
			if string(str[i]) == "2" {
				num += 2 * math.Pow(2, float64(j))
			}
			if string(str[i]) == "3" {
				num += 3 * math.Pow(2, float64(j))
			}
			if string(str[i]) == "4" {
				num += 4 * math.Pow(2, float64(j))
			}
			if string(str[i]) == "5" {
				num += 5 * math.Pow(2, float64(j))
			}
			if string(str[i]) == "6" {
				num += 6 * math.Pow(2, float64(j))
			}
			if string(str[i]) == "7" {
				num += 7 * math.Pow(2, float64(j))
			}
			if string(str[i]) == "8" {
				num += 8 * math.Pow(2, float64(j))
			}
			if string(str[i]) == "9" {
				num += 9 * math.Pow(2, float64(j))
			}
			if string(str[i]) == "a" {
				num += 10 * math.Pow(2, float64(j))
			}
			if string(str[i]) == "b" {
				num += 11 * math.Pow(2, float64(j))
			}
			if string(str[i]) == "c" {
				num += 12 * math.Pow(2, float64(j))
			}
			if string(str[i]) == "d" {
				num += 13 * math.Pow(2, float64(j))
			}
			if string(str[i]) == "e" {
				num += 14 * math.Pow(2, float64(j))
			}
			if string(str[i]) == "f" {
				num += 15 * math.Pow(2, float64(j))
			}
			j = j + 4
		}
		return int(num)
	}

	for i := 31; i >= 0; i-- {
		if string(str[i]) == "1" {
			num += math.Pow(2, float64(j))
		}
		j++
	}
	return int(num)
} // 现在我们的问题是：如何将string转为int 。比如"101"，要转为int的0b101

func binaryToHex(str string) string {
	result := ""
	for i := 0; i < 8; i++ {
		switch str[i*4 : i*4+4] {
		case "0000":
			result += "0"
		case "0001":
			result += "1"
		case "0010":
			result += "2"
		case "0011":
			result += "3"
		case "0100":
			result += "4"
		case "0101":
			result += "5"
		case "0110":
			result += "6"
		case "0111":
			result += "7"
		case "1000":
			result += "8"
		case "1001":
			result += "9"
		case "1010":
			result += "A"
		case "1011":
			result += "B"
		case "1100":
			result += "C"
		case "1101":
			result += "D"
		case "1110":
			result += "E"
		case "1111":
			result += "F"
		}
	}
	return result
} //二进制转十六进制

func binaryToHex_(str string) string {
	result := ""
	for i := 0; i < 40; i++ {
		switch str[i*4 : i*4+4] {
		case "0000":
			result += "0"
		case "0001":
			result += "1"
		case "0010":
			result += "2"
		case "0011":
			result += "3"
		case "0100":
			result += "4"
		case "0101":
			result += "5"
		case "0110":
			result += "6"
		case "0111":
			result += "7"
		case "1000":
			result += "8"
		case "1001":
			result += "9"
		case "1010":
			result += "A"
		case "1011":
			result += "B"
		case "1100":
			result += "C"
		case "1101":
			result += "D"
		case "1110":
			result += "E"
		case "1111":
			result += "F"
		}
	}
	return result
} //二进制转十六进制
