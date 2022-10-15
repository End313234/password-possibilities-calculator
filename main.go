package main

import "errors"

type Criteria = int

const (
	LETTER Criteria = iota
	NUMBER
	LETTER_NOT_USED
	NUMBER_NOT_USED
	ANY
)

var ALPHABET = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var NUMBERS = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func in[T comparable](items []T, target T) bool {
	for _, item := range items {
		if item == target {
			return true
		}
	}

	return false
}

func filterLetters(arr []string) []string {
	letters := []string{}

	for _, item := range arr {
		if in(ALPHABET, item) {
			letters = append(letters, item)
		}
	}

	return letters
}

func filterNumbers(arr []string) []string {
	numbers := []string{}

	for _, item := range arr {
		if in(NUMBERS, item) {
			numbers = append(numbers, item)
		}
	}

	return numbers
}

func CalculatePasswordPossibilities(possibilities []string, quantityOfChars int, criteria []Criteria) (int, error) {
	quantityOfPossibilities := 0
	possibilitiesThatAreLetters := filterLetters(possibilities)
	possibilitiesThatAreNumbers := filterNumbers(possibilities)
	quantityOfUsedLetters := 0
	quantityOfUsedNumbers := 0

	if quantityOfChars > 0 {
		quantityOfPossibilities = 1
	}

	for i := 0; i < quantityOfChars; i++ {
		possibilityCriteria := ANY
		if len(criteria) > i {
			possibilityCriteria = criteria[i]
		}

		switch possibilityCriteria {
		case LETTER:
			quantityOfPossibilities *= len(possibilitiesThatAreLetters)
			quantityOfUsedLetters += 1
		case LETTER_NOT_USED:
			quantityOfPossibilities *= len(possibilitiesThatAreLetters) - quantityOfUsedLetters
			quantityOfUsedLetters += 1
		case NUMBER:
			quantityOfPossibilities *= len(possibilitiesThatAreNumbers)
			quantityOfUsedNumbers += 1
		case NUMBER_NOT_USED:
			quantityOfPossibilities *= len(possibilitiesThatAreNumbers) - quantityOfUsedNumbers
			quantityOfUsedNumbers += 1
		case ANY:
			quantityOfPossibilities *= len(possibilities)
		default:
			return 0, errors.New("invalid criteria")
		}
	}

	return quantityOfPossibilities, nil
}
