package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// 计算给定数据和Nonce的SHA-256哈希
func calculateHash(data string, nonce int) string {
	input := fmt.Sprintf("%s%d", data, nonce)
	hash := sha256.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

// 挖矿函数
func mineBlock(data string, difficulty int) (int, string) {
	var nonce = 0
	prefix := strings.Repeat("0", difficulty)
	for {
		hash := calculateHash(data, nonce)
		if strings.HasPrefix(hash, prefix) {
			return nonce, hash
		}
		fmt.Printf("当前nonce值为：%d, hash值为：%s\n", nonce, hash)
		nonce++
	}

}

func main() {

	start := time.Now()
	data := "hello, block chain"
	defficulty := 6
	nonce, hash := mineBlock(data, defficulty)
	end := time.Now()
	fmt.Printf("耗时：%v\n", end.Sub(start))
	fmt.Printf("挖矿成功，nonce值为：%d, hash值为：%s\n", nonce, hash)
}
