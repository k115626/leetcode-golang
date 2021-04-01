package main

import "fmt"

//func reverseBits(num uint32) uint32 {
//	var res uint32 = 0
//	for i := 0; i < 32; i++ {
//		rightBit := num & 1
//		res |= rightBit << (31 - i)
//		num >>= 1
//	}
//	return res
//}

// 分冶
// 反转一个8位数，前4位肯定和后4位换位
// 然后前4位的前2位和后2位换位
// 直到后每一位都进行换位
// 每隔1个交换
// 0xaaaaaaaa = (1010 1010 1010 1010 1010 1010 1010 1010)b
// 0x55555555 = (0101 0101 0101 0101 0101 0101 0101 0101)b

// 每隔2个交换
// 0xcccccccc = (1100 1100 1100 1100 1100 1100 1100 1100)b
// 0x33333333 = (0011 0011 0011 0011 0011 0011 0011 0011)b

// 每隔4个交换
// 0xf0f0f0f0 = (1111 0000 1111 0000 1111 0000 1111 0000)b
// 0x0f0f0f0f = (0000 1111 0000 1111 0000 1111 0000 1111)b

// 每隔8个交换
// 0xff00ff00 = (1111 1111 0000 0000 1111 1111 0000 0000)b
// 0x00ff00ff = (0000 0000 1111 1111 0000 0000 1111 1111)b

func reverseBits(num uint32) uint32 {
	num = num>>16 | num<<16
	num = (num&0xff00ff00)>>8 | (num&0x00ff00ff)<<8
	num = (num&0xf0f0f0f0)>>4 | (num&0x0f0f0f0f)<<4
	num = (num&0xcccccccc)>>2 | (num&0x33333333)<<2
	num = (num&0xaaaaaaaa)>>1 | (num&0x55555555)<<1
	return num
}

func main() {
	var num uint32 = 43261596
	res := reverseBits(num)
	fmt.Println(res)
}
