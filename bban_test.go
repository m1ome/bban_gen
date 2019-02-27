package bban_test

import (
	"fmt"
	"strconv"
	"testing"

	bban "github.com/m1ome/bban_gen"
	"github.com/stretchr/testify/assert"
)

func Example() {
	account := bban.Random("040577", "13439317554524", bban.DoubleMod)
	fmt.Printf("Generated account: %s\n", account)

	next := bban.Next("040577", account, "13439317554524", bban.DoubleMod)
	fmt.Printf("Next account: %s\n", next)

	fmt.Printf("Validity passing: %v\n", bban.Validate("040577", next, "13439317554524", bban.DoubleMod))

	// Output:
	// Generated account: 40954426
	// Next account: 40954431
	// Validity passing: true
}

func TestRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		account := bban.Random("040577", "13439317554524", bban.DoubleMod)
		assert.True(t, len(account) > 0)
	}
}

func TestNext(t *testing.T) {
	start := strconv.Itoa(1000)
	numbers := []string{}
	for i := 0; i < 10; i++ {
		account := bban.Next("040577", start, "13439317554524", bban.DoubleMod)
		numbers = append(numbers, account)
		start = account
	}

	assert.Equal(t, []string{"00001008", "00001013", "00001039", "00001044", "00001051", "00001065", "00001070", "00001077", "00001082", "00001096"}, numbers)
}

func TestValidate(t *testing.T) {
	tests := []struct {
		sortCode string
		account  string
		weight   string
		vtype    int
		result   bool
	}{
		{
			"040577",
			"77777777",
			"13439317554524",
			bban.DoubleMod,
			true,
		},
		{
			"089999",
			"66374958",
			"00000071371371",
			bban.Mod10,
			true,
		},
		{
			"107999",
			"88837491",
			"00000087654321",
			bban.Mod11,
			true,
		},
		{
			"202959",
			"63748472",
			"00000007654321",
			bban.Mod11,
			true,
		},
		{
			"202959",
			"63748472",
			"21212121212121",
			bban.DoubleMod,
			true,
		},
		{
			"1234",
			"567890",
			"123456789",
			bban.DoubleMod,
			false,
		},
		{
			"202959",
			"63748472",
			"21212121212121",
			-1,
			false,
		},
	}

	for _, test := range tests {
		v := bban.Validate(test.sortCode, test.account, test.weight, test.vtype)
		assert.Equal(
			t, test.result, v,
			"expected %s %s / %s [%d] to have validation result '%v', but have '%v'",
			test.sortCode, test.account, test.weight, test.vtype, test.result, v,
		)
	}

}
