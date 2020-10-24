package uid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"strings"
)

// Generator ...
type Generator struct {
	Bits         int
	RandomSource io.Reader
	Dictionary   []string
}

// DefaultGenerator ...
var DefaultGenerator = NewGenerator(48)

// NewGenerator ...
func NewGenerator(bits int) *Generator {
	return &Generator{
		Bits:         bits,
		RandomSource: rand.Reader,
		Dictionary:   Dictionary,
	}
}

// UID ...
func (gen *Generator) UID() string {
	maxWord := int64(len(gen.Dictionary))
	bitsPerWord := big.NewInt(maxWord).BitLen()
	var wordCount, remainingBits int

	if bitsPerWord > gen.Bits {
		wordCount, remainingBits = 0, gen.Bits
	} else {
		wordCount, remainingBits = gen.Bits/bitsPerWord, gen.Bits%bitsPerWord
	}

	words := make([]string, wordCount)
	for i := range words {
		word, err := rand.Int(gen.RandomSource, big.NewInt(maxWord))
		if err != nil {
			panic(err)
		}
		words[i] = gen.Dictionary[int(word.Uint64())]
	}

	if remainingBits > 0 {
		remainingMax := big.NewInt(0)
		remainingMax.SetBit(remainingMax, remainingBits, 1)
		remaining, err := rand.Int(gen.RandomSource, remainingMax)
		if err != nil {
			panic(err)
		}

		remainingBytes := make([]byte, (remainingBits+7)/8)
		remaining.FillBytes(remainingBytes)

		if len(words) > 0 {
			return fmt.Sprintf("%s-%s",
				strings.Join(words, "-"),
				hex.EncodeToString(remainingBytes))
		}
		return hex.EncodeToString(remainingBytes)
	}

	return strings.Join(words, "-")
}

// UID ...
func UID() string {
	return DefaultGenerator.UID()
}
