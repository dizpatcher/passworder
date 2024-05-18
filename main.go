package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// Password requirements:
// 1. содержит минимум 1 заглавную букву
// 2. содержит минимум 1 цифру
// 3. содержит минимум 1 спецсимвол

const passwordLength = 10
const shareOfSpecials = 0.125
const alphabet = "abcdefghijklmnopqrstuvwxyz"
const digits = "0123456789"
const specials = "!#$%&()*+,-./:;<=>?@[\\]^_`{|}~"

var alphabetUpper = strings.ToUpper(alphabet)

func shuffleArray(arr []byte) []byte {
	rand.Seed(time.Now().UnixNano())

	shuffledArr := make([]byte, len(arr))
	copy(shuffledArr, arr)

	for i := len(shuffledArr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)

		shuffledArr[i], shuffledArr[j] = shuffledArr[j], shuffledArr[i]
	}

	return shuffledArr
}

func main() {
	rand.Seed(time.Now().UnixNano()) // инициализация генератора случайных чисел

	var pwd [passwordLength]byte
	// digit block
	pwd[0] = digits[rand.Intn(len(digits))]
	// Upper block
	pwd[1] = alphabetUpper[rand.Intn(len(alphabetUpper))]

	// special character block
	numberOfSpecials := int(math.Round(shareOfSpecials * passwordLength))
	// fmt.Println("Количество спецсимволов:", numberOfSpecials)
	for i := 2; i < 2+numberOfSpecials; i++ {
		pwd[i] = specials[rand.Intn(len(specials))]
	}

	allSymbols := alphabet + digits + alphabetUpper
	for i := 2 + numberOfSpecials; i < passwordLength; i++ {
		pwd[i] = allSymbols[rand.Intn(len(allSymbols))]
	}

	sharr := shuffleArray(pwd[:])
	fmt.Println(string(sharr))
}
