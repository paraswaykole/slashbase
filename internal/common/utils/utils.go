package utils

import (
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
	"unsafe"
)

func ContainsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func ContainsInt(a []int, integer int) bool {
	for _, v := range a {
		if v == integer {
			return true
		}
	}
	return false
}

func UnixNanoToTime(nanoInt int64) time.Time {
	msInt := nanoInt / 1000000000
	remainder := nanoInt % 1000000000
	value := time.Unix(msInt, remainder*int64(time.Nanosecond))
	return value
}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func InterfaceArrayToStringArray(arr []interface{}) []string {
	result := []string{}
	for _, str := range arr {
		result = append(result, str.(string))
	}
	return result
}

// fast & unsafe pointer function
func RandString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

func FileExtensionFromPath(path string) string {
	if strs := strings.Split(path, "."); true {
		slen := len(strs)
		if slen > 1 {
			return strs[slen-1]
		}
	}
	return ""
}

func ExtractDomainFromHost(host string) string {
	parts := strings.Split(host, ":")
	if len(parts) == 1 {
		return host
	}
	return parts[0]
}
