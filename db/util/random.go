package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxy"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomTitle() string {
	return RandomString(3)
}
func RandomHouse() string {
	return RandomString(3)
}

//func RandomEggs() int32 {
//	return RandomInt(5, 10)
//}
//func RandomDirty() int32 {
//	return RandomInt(3, 5)
//}
//func RandomWrongShape() int32 {
//	return RandomInt(3, 5)
//}
//func RandomWeakShell() int32 {
//	return RandomInt(3, 5)
//}
//func RandomDamaged() int32 {
//	return RandomInt(3, 5)
//}
//func RandomHatchingEggs() int32 {
//	return RandomInt(5, 8)
//}
