package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateAlphabet() []string {
	alphabet := []string{}

	for ch := 'a'; ch <= 'z'; ch++ {
		alphabet = append(alphabet, string(ch))
	}

	return alphabet
}

func provideNumbers() []string {
	numbers := []string{}

	for n := 0; n < 10; n++ {
		numbers = append(numbers, fmt.Sprint(n))
	}

	return numbers
}

func TestCalculatePasswordPossibilitiesWithAlphabetSuccess(t *testing.T) {
	assert := assert.New(t)

	possibilities := generateAlphabet()
	quantityOfChars := 2
	criteria := []Criteria{LETTER, LETTER}

	quantityOfPossibilities, err := CalculatePasswordPossibilities(possibilities, quantityOfChars, criteria)
	assert.NoError(err)
	assert.Equal(676, quantityOfPossibilities)
}

func TestCalculatePasswordPossibilitiesWithAlphanumericSuccess(t *testing.T) {
	assert := assert.New(t)

	alphabet := generateAlphabet()
	possibilities := append(alphabet, provideNumbers()...)
	quantityOfChars := 5
	criteria := []Criteria{LETTER, NUMBER, LETTER_NOT_USED, NUMBER_NOT_USED, LETTER_NOT_USED}

	quantityOfPossibilities, err := CalculatePasswordPossibilities(possibilities, quantityOfChars, criteria)
	assert.NoError(err)
	assert.Equal(1404000, quantityOfPossibilities)
}

func TestCalculatePasswordPossibilitiesWithAlphanumericAndInsufficientCriteriaSuccess(t *testing.T) {
	assert := assert.New(t)

	alphabet := generateAlphabet()
	possibilities := append(alphabet, provideNumbers()...)
	quantityOfChars := 6
	criteria := []Criteria{LETTER, NUMBER}

	quantityOfPossibilities, err := CalculatePasswordPossibilities(possibilities, quantityOfChars, criteria)
	assert.NoError(err)
	assert.Equal(436700160, quantityOfPossibilities)
}

func TestCalculatePasswordPossibilitiesFailure(t *testing.T) {
	assert := assert.New(t)

	possibilities := generateAlphabet()
	quantityOfChars := 2
	criteria := []Criteria{LETTER, 20}

	quantityOfPossibilities, err := CalculatePasswordPossibilities(possibilities, quantityOfChars, criteria)
	assert.Error(err)
	assert.Equal(0, quantityOfPossibilities)
}
