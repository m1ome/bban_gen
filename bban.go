package bban

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// Mods can be used during validation & generation
const (
	DoubleMod int = iota
	Mod10
	Mod11
)

// Random generates random bban number based on sort code, weights and mod.
func Random(sortCode, weights string, mod int) string {
	exp := len(weights) - len(sortCode)

	for {
		rand := strconv.Itoa(rand.Intn(int(math.Pow10(exp)) - 1))
		if Validate(sortCode, rand, weights, mod) {
			return rand
		}
	}
}

// Next try to find next account in a row, using sort code, previous account number, weights and mod
func Next(sortCode, account, weights string, mod int) string {
	exp := len(weights) - len(sortCode)
	acc, _ := strconv.Atoi(account)

	for {
		acc++
		tmp := strconv.Itoa(acc)

		if len(tmp) < exp {
			tmp = strings.Repeat("0", exp-len(tmp)) + tmp
		}

		if Validate(sortCode, tmp, weights, mod) {
			return tmp
		}
	}
}

// Validate validate account number based on sort code and weights with mod
func Validate(sortCode, account, weights string, mod int) bool {
	bban := sortCode + account
	if len(bban) != len(weights) {
		return false
	}

	var (
		mul string
		res int
	)
	for i := 0; i < len(bban); i++ {
		x, _ := strconv.Atoi(string(bban[i]))
		y, _ := strconv.Atoi(string(weights[i]))

		mul += strconv.Itoa(x * y)
		res += x * y
	}

	switch mod {
	case DoubleMod:
		var sum int
		for i := 0; i < len(mul); i++ {
			x, _ := strconv.Atoi(string(mul[i]))
			sum += x
		}

		return sum%10 == 0
	case Mod10:
		return res%10 == 0
	case Mod11:
		return res%11 == 0
	default:
		return false
	}
}
