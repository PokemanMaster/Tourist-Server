package midAuto

import (
	"math/rand"
	"time"
)

// GenerateRandomNumber 随机生成7位数字
func GenerateRandomNumber() uint {
	rand.Seed(time.Now().UnixNano())
	return uint(rand.Intn(9000000) + 1000000)
}
