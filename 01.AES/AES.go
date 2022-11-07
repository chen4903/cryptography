package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//定义我们的输入
	fmt.Println("(1)当输入为：32 43 f6 a8 88 5a 30 8d 31 31 98 a2 e0 37 07 34:")
	var arr01 = [16]byte{0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31,
		0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34}
	AES_94(arr01)
	fmt.Println()
	fmt.Println()

	fmt.Println("(2)当输入为：19 3d e3 be a0 f4 e2 2b 9a c6 8d 2a e9 f8 48 08:")
	var arr02 = [16]byte{0x19, 0x3d, 0xe3, 0xbe, 0xa0, 0xf4, 0xe2, 0x2b, 0x9a, 0xc6, 0x8d, 0x2a, 0xe9, 0xf8, 0x48, 0x08}
	AES_94(arr02)
	fmt.Println()
	fmt.Println()

	fmt.Println("(3)当输入为：20 21 13 10 94 00 00 00 00 00 00 00 00 00 00 0B:")
	var arr03 = [16]byte{0x20, 0x21, 0x13, 0x10, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0B}
	AES_94(arr03)

}

// 定义S盒
var Sbox = [16 * 16]byte{
	/*      0     1      2     3     4    5     6     7      8    9     a      b     c     d     e    f */
	/*0*/ 0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76,
	/*1*/ 0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0,
	/*2*/ 0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15,
	/*3*/ 0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75,
	/*4*/ 0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84,
	/*5*/ 0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf,
	/*6*/ 0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8,
	/*7*/ 0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2,
	/*8*/ 0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73,
	/*9*/ 0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb,
	/*a*/ 0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79,
	/*b*/ 0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08,
	/*c*/ 0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a,
	/*d*/ 0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e,
	/*e*/ 0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf,
	/*f*/ 0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16}

// 字节代替
func subBytes_94(arr [16]byte) (newArr [16]byte) {
	for index, value := range arr {
		low := value & 0x0F         //取低4位
		high := (value >> 4) & 0x0f //取高4位
		newArr[index] = Sbox[16*high+low]
	}
	return newArr
}

// 行移位
func ShiftRows_94(arr [16]byte) (newArr [16]byte) {
	for index, _ := range arr {
		newArr[index] = arr[index]
	}
	//第一行不移位
	//第二行移位
	newArr[4] = arr[5]
	newArr[5] = arr[6]
	newArr[6] = arr[7]
	newArr[7] = arr[4]
	//第三行移位
	newArr[8] = arr[10]
	newArr[9] = arr[11]
	newArr[10] = arr[8]
	newArr[11] = arr[9]
	//第四行移位
	newArr[12] = arr[15]
	newArr[13] = arr[12]
	newArr[14] = arr[13]
	newArr[15] = arr[14]
	return newArr
}

// 列混合
func columnMix_94(arr [16]byte) [16]byte {
	//列混合之前先将数组重新排序（因为列混合是一列一列来的）
	var arrByte [16]byte = [16]byte{}
	arrByte[0] = arr[0]
	arrByte[1] = arr[4]
	arrByte[2] = arr[8]
	arrByte[3] = arr[12]
	arrByte[4] = arr[1]
	arrByte[5] = arr[5]
	arrByte[6] = arr[9]
	arrByte[7] = arr[13]
	arrByte[8] = arr[2]
	arrByte[9] = arr[6]
	arrByte[10] = arr[10]
	arrByte[11] = arr[14]
	arrByte[12] = arr[3]
	arrByte[13] = arr[7]
	arrByte[14] = arr[11]
	arrByte[15] = arr[15]

	result_ := [16]byte{} //用来装输出矩阵
	result_[0] = columnMixResult01_94(arrByte[0], arrByte[1], arrByte[2], arrByte[3])
	result_[1] = columnMixResult01_94(arrByte[4], arrByte[5], arrByte[6], arrByte[7])
	result_[2] = columnMixResult01_94(arrByte[8], arrByte[9], arrByte[10], arrByte[11])
	result_[3] = columnMixResult01_94(arrByte[12], arrByte[13], arrByte[14], arrByte[15])

	result_[4] = columnMixResult02_94(arrByte[0], arrByte[1], arrByte[2], arrByte[3])
	result_[5] = columnMixResult02_94(arrByte[4], arrByte[5], arrByte[6], arrByte[7])
	result_[6] = columnMixResult02_94(arrByte[8], arrByte[9], arrByte[10], arrByte[11])
	result_[7] = columnMixResult02_94(arrByte[12], arrByte[13], arrByte[14], arrByte[15])

	result_[8] = columnMixResult03_94(arrByte[0], arrByte[1], arrByte[2], arrByte[3])
	result_[9] = columnMixResult03_94(arrByte[4], arrByte[5], arrByte[6], arrByte[7])
	result_[10] = columnMixResult03_94(arrByte[8], arrByte[9], arrByte[10], arrByte[11])
	result_[11] = columnMixResult03_94(arrByte[12], arrByte[13], arrByte[14], arrByte[15])

	result_[12] = columnMixResult04_94(arrByte[0], arrByte[1], arrByte[2], arrByte[3])
	result_[13] = columnMixResult04_94(arrByte[4], arrByte[5], arrByte[6], arrByte[7])
	result_[14] = columnMixResult04_94(arrByte[8], arrByte[9], arrByte[10], arrByte[11])
	result_[15] = columnMixResult04_94(arrByte[12], arrByte[13], arrByte[14], arrByte[15])
	return result_
} //列混合主函数

func columnMixCount01_94(num byte) byte { //与{01}相乘
	return num
}

func columnMixCount02_94(num byte) byte { //与{02}相乘
	if num > 128 {
		return (num << 1) ^ 0b00011011
	} else {
		return num << 1
	}
}

func columnMixCount03_94(num byte) byte { //与{03}相乘
	return columnMixCount01_94(num) ^ columnMixCount02_94(num)
}

func columnMixResult01_94(num01 byte, num02 byte, num03 byte, num04 byte) byte { //固定矩阵第1行
	var s byte = columnMixCount02_94(num01) ^ columnMixCount03_94(num02) ^ columnMixCount01_94(num03) ^ columnMixCount01_94(num04)
	return s
}

func columnMixResult02_94(num01 byte, num02 byte, num03 byte, num04 byte) byte { //固定矩阵第2行
	var s byte = columnMixCount01_94(num01) ^ columnMixCount02_94(num02) ^ columnMixCount03_94(num03) ^ columnMixCount01_94(num04)
	return s
}

func columnMixResult03_94(num01 byte, num02 byte, num03 byte, num04 byte) byte { //固定矩阵第3行
	var s byte = columnMixCount01_94(num01) ^ columnMixCount01_94(num02) ^ columnMixCount02_94(num03) ^ columnMixCount03_94(num04)
	return s
}

func columnMixResult04_94(num01 byte, num02 byte, num03 byte, num04 byte) byte { //固定矩阵第4行
	var s byte = columnMixCount03_94(num01) ^ columnMixCount01_94(num02) ^ columnMixCount01_94(num03) ^ columnMixCount02_94(num04)
	return s
}

// 轮密钥加
func AddRoundKey_94(arr [16]byte) (newArr [16]byte) {
	//定义子密钥.
	var arrKey [16]byte = [16]byte{0xa0, 0x88, 0x23, 0x2a, 0xfa, 0x54, 0xa3, 0x6c, 0xfe, 0x2c, 0x39,
		0x76, 0x17, 0xb1, 0x39, 0x05}
	for i := 0; i < 16; i++ {
		newArr[i] = arrKey[i] ^ arr[i]
	}
	return newArr
}

func AES_94(arrBegin [16]byte) {
	step1 := subBytes_94(arrBegin)
	fmt.Print("字节代替:")
	for i := 0; i < 16; i++ {
		if i%4 == 0 {
			fmt.Println()
		}
		fmt.Printf("%3v", strings.ToUpper(strconv.FormatInt(int64(step1[i]), 16)))
	}
	fmt.Println()
	step2 := ShiftRows_94(step1)
	fmt.Print("行移位 :")
	for i := 0; i < 16; i++ {
		if i%4 == 0 {
			fmt.Println()
		}
		fmt.Printf("%3v", strings.ToUpper(strconv.FormatInt(int64(step2[i]), 16)))
	}
	fmt.Println()
	step3 := columnMix_94(step2)
	fmt.Print("列混合 :")
	for i := 0; i < 16; i++ {
		if i%4 == 0 {
			fmt.Println()
		}
		fmt.Printf("%3v", strings.ToUpper(strconv.FormatInt(int64(step3[i]), 16)))
	}
	fmt.Println()
	step4 := AddRoundKey_94(step3)
	fmt.Print("轮密钥加:")
	for i := 0; i < 16; i++ {
		if i%4 == 0 {
			fmt.Println()
		}
		fmt.Printf("%3v", strings.ToUpper(strconv.FormatInt(int64(step4[i]), 16)))
	}
}
