package main

import (
	"fmt"
	"math/rand"
	"time"
)

var S0 = [][][]int{
	{{0, 1}, {0, 0}, {1, 1}, {1, 0}},
	{{1, 1}, {1, 0}, {0, 1}, {0, 0}},
	{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
	{{1, 1}, {0, 1}, {1, 1}, {1, 0}},
}

var S1 = [][][]int{
	{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
	{{1, 0}, {0, 0}, {0, 1}, {1, 1}},
	{{1, 1}, {0, 0}, {0, 1}, {0, 0}},
	{{1, 0}, {0, 1}, {0, 0}, {1, 1}},
}

func main() {
	input := "Buna ziua domnul profesor, as putea sa am nota 10 la examen?"

	// convert the input to bits
	inputInBits := ASCIItoBits(input)

	encryptedInput := make([][]int, len(inputInBits))
	decryptedInput := make([][]int, len(inputInBits))

	// generate the key
	key := randomKey()
	kOne, kTwo := generateKeyOneAndKeyTwo(key)

	// encrypt the input
	for i, bits := range inputInBits {
		encryptedInput[i] = crypt(bits, kOne, kTwo)
	}

	fmt.Println("Encrypted input:", bitsToASCII(encryptedInput))

	// decrypt the input
	for i, bits := range encryptedInput {
		decryptedInput[i] = crypt(bits, kTwo, kOne)
	}

	fmt.Println("Decrypted input:", bitsToASCII(decryptedInput))

	// check if the decrypted input is equal to the original input
	if input == bitsToASCII(decryptedInput) {
		fmt.Println("SUCCESS: Decrypted input is equal to the original input")
	} else {
		fmt.Println("ERROR: Decrypted input is NOT equal to the original input")
	}
}

func generateKeyOneAndKeyTwo(key []int) ([]int, []int) {
	p10Permutation := []int{3, 5, 2, 7, 4, 10, 1, 9, 8, 6}
	// lsOnePermutation := []int{2, 3, 4, 5, 1}
	// lsTwoPermutation := []int{3, 4, 5, 1, 2}
	p8Permutation := []int{6, 3, 7, 4, 8, 5, 10, 9}

	// permutate the key with the p10 permutation
	key = permutate(key, p10Permutation)

	// split the key into two 5 bit blocks and shift left by one each block
	firstLS1 := shiftByX(key[:5], 1)
	secondLS1 := shiftByX(key[5:], 1)

	// permutate the two 5 bit blocks with the p8 permutation
	kOne := permutate(append(firstLS1, secondLS1...), p8Permutation)

	// split the key into two 5 bit blocks and shift left by two each block
	firstLS2 := shiftByX(firstLS1, 2)
	secondLS2 := shiftByX(secondLS1, 2)

	// permutate the two 5 bit blocks with the p8 permutation
	kTwo := permutate(append(firstLS2, secondLS2...), p8Permutation)

	return kOne, kTwo
}

func shiftByX(val []int, x int) []int {
	result := make([]int, 0, len(val))

	result = append(result, val[x:]...)
	result = append(result, val[:x]...)

	return result
}

func permutate(val []int, permutation []int) []int {
	result := make([]int, len(permutation))

	for i, v := range permutation {
		result[i] = val[v-1]
	}

	return result
}

// in order to decrypt, just the order of the keys is reversed
func crypt(payload, firstKey, secondKey []int) []int {
	IPPermutation := []int{2, 6, 3, 1, 4, 8, 5, 7}
	inverseIPPermutation := []int{4, 1, 3, 5, 7, 2, 8, 6}

	// permutate the plain text with the IP permutation
	payload = permutate(payload, IPPermutation)

	// split the plain text into two 4 bit blocks
	left := payload[:4]
	right := payload[4:]

	ep1Res := epFunc(right, firstKey)

	// xor the left block with the ep1 result
	xoredLeft := xor(ep1Res, left)

	// swap the left and right blocks
	left = right
	right = xoredLeft

	ep2Res := epFunc(right, secondKey)

	// xor the left block with the ep2 result
	xoredLeft = xor(ep2Res, left)

	// combine the left and right blocks
	combined := append(xoredLeft, right...)

	// permutate the combined block with the inverse IP permutation
	combined = permutate(combined, inverseIPPermutation)

	return combined
}

func epFunc(right []int, key []int) []int {
	// expand the right block with the EP permutation
	EPPermutation := []int{4, 1, 2, 3, 2, 3, 4, 1}
	expandedRight := permutate(right, EPPermutation)

	// xor the expanded right block with the key one
	xored := xor(expandedRight, key)

	// split the xored block into two 4 bit blocks
	leftXored := xored[:4]
	rightXored := xored[4:]

	// lookup the left xored block in S0
	s0Row := bitsToInt([]int{leftXored[0], leftXored[3]})
	s0Col := bitsToInt([]int{leftXored[1], leftXored[2]})

	s0Result := S0[s0Row][s0Col]

	// lookup the right xored block in S1
	s1Row := bitsToInt([]int{rightXored[0], rightXored[3]})
	s1Col := bitsToInt([]int{rightXored[1], rightXored[2]})
	s1Result := S1[s1Row][s1Col]

	// permutate the s0 and s1 results with the P4 permutation
	p4Permutation := []int{2, 4, 3, 1}
	p4Result := permutate(append(s0Result, s1Result...), p4Permutation)

	return p4Result
}

func randomKey() []int {
	key := make([]int, 10)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10; i++ {
		key[i] = r.Intn(2)
	}

	return key
}

func validateASCIIinput(input string) bool {
	for _, r := range input {
		if r > 127 {
			return false
		}
	}

	return true
}

func ASCIItoBits(input string) [][]int {
	result := make([][]int, len(input))

	for charIdx, r := range input {
		for i := 7; i >= 0; i-- {
			result[charIdx] = append(result[charIdx], int(r>>i)&1)
		}
	}

	return result
}

func bitsToASCII(input [][]int) string {
	result := ""

	for _, bits := range input {
		result += string(bitsToInt(bits))
	}

	return result
}

func xor(val1, val2 []int) []int {
	result := make([]int, len(val1))

	for i, v := range val1 {
		result[i] = v ^ val2[i]
	}

	return result
}

func bitsToInt(val []int) int {
	result := 0

	for _, v := range val {
		result = (result << 1) | v
	}

	return result
}

func areEqual(val1, val2 []int) bool {
	if len(val1) != len(val2) {
		return false
	}

	for i, v := range val1 {
		if v != val2[i] {
			return false
		}
	}

	return true
}
