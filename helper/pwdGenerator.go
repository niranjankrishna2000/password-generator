package helper

import (
	"math/rand"
)

func GeneratePassword(length int, includeUppercase, includeLowercase, includeNumbers, includeSpecialChars bool) string {
	UpperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerCase := "abcdefghijklmnopqrstuvwxyz"
	Numbers := "0123456789"
	SplChar := "!@#$%^&*()-_=+[]{}|;:'\",.<>?/\\"
	password := make([]byte, length)
	for i := 0; i < length; {
		if includeUppercase {
			password[i] = UpperCase[rand.Intn(len(UpperCase))]
			i++
			if i == length {
				break
			}
		}
		if includeLowercase {
			password[i] = LowerCase[rand.Intn(len(LowerCase))]
			i++
			if i == length {
				break
			}
		}
		if includeNumbers {
			password[i] = Numbers[rand.Intn(len(Numbers))]
			i++
			if i == length {
				break
			}
		}
		if includeSpecialChars {
			password[i] = SplChar[rand.Intn(len(SplChar))]
			i++
			if i == length {
				break
			}
		}
	}
	return string(password)
}
