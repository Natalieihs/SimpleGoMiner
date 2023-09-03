package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
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
func mineBlock(data string, difficulty int, start, step int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	prefix := strings.Repeat("0", difficulty)
	for nonce := start; ; nonce += step {
		hash := calculateHash(data, nonce)
		if strings.HasPrefix(hash, prefix) {
			ch <- nonce
			return
		}
	}

}

func main() {

	start := time.Now()
	data := "hello, block chain"
	defficulty := 7
	numGoroutines := 7
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go mineBlock(data, defficulty, i, numGoroutines, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	nonce := <-ch
	hash := calculateHash(data, nonce)
	end := time.Now()
	fmt.Printf("耗时：%v\n", end.Sub(start))
	fmt.Printf("挖矿成功，随机数为：%d, hash为：%s\n", nonce, hash)
}
