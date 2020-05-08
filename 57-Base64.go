// Go内置支持base64编解码。

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~" // 编解码这个字符串

	// Go支持标准的和URL兼容的Base64
	// 使用标准编码器进行编码，编码器需要字节数组，所以将字符串转换为字节数组
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 解码会返回错误，可用于检测输入格式是否正确
	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// 使用与URL兼容的Base64编解码方式
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDec, _ := base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

// 对同一个字符串使用标准编码器和URL兼容的编码器进行编码等到的值略有不同(末尾是+或者-)
// 但是都能根据需要解码为原始字符串
