package util

import (
	"database/sql"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxy"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
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
	return RandomString(4)
}
func RandomHouse() string {
	return RandomString(5)
}

func RandomNums() sql.NullInt64 {
	return sql.NullInt64{
		Int64: RandomInt(10, 2000),
		Valid: true,
	}
}
