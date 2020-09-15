package main

import (
	"fmt"
	"strconv"
)

// 4003600000000014 - VISA
// it doesn't use arrays or struct
// pure mathematic solution

func main() {
	// get card number
	var cardNumber int = -1
	for cardNumber == -1 {
		cardNumber = getCardNumer()
	}

	// get the number of digits from the card
	numDigits := getNumberDigitsCard(cardNumber)

	// AUX VARIABLES ---------------------------
	// aux vars for iterate through the card number
	mod := 10
	divider := 1
	// aux for know the ieration, because it goes from right to left
	var generalOdd bool = (numDigits % 2) != 0
	// aux var for iterate one yes - one no
	odd := generalOdd
	// accumulators
	ac1 := 0 // special - the one that will be multiple by 2
	ac2 := 0 // odds - start from the second digit

	// ITERATE CARD NUMBER
	for cardNumber > divider {

		// get the digit from the right
		var currentVal int = (cardNumber % mod) / divider

		// move one digit to the left
		mod = mod * 10
		divider = divider * 10

		// LOGIC FROM Hans Peter Luhn ALGORITHM
		// make the sum step from digits
		if odd {
			if !generalOdd {
				ac1 = ac1 + specialSum(currentVal)
			} else {
				ac1 = ac1 + currentVal
			}
		} else {
			if generalOdd {
				ac2 = ac2 + specialSum(currentVal)
			} else {
				ac2 = ac2 + currentVal
			}
		}
		odd = !odd
	}
	var result int = ac1 + ac2

	// CHECK THE RESULT
	// compare the info given from the credit cards
	var cardType string = checkResult(result, numDigits, cardNumber)
	fmt.Println(cardType)
}

func getCardNumer() int {
	// get the value
	var str string = ""
	fmt.Print("Enter a card number: ")
	fmt.Scanf("%s", &str)

	// convert value
	cardNumber, err := strconv.Atoi(str)

	// handle error if no number was provided
	if err != nil {
		fmt.Println("Enter only numbers")
		return -1
	}

	return cardNumber
}

func getNumberDigitsCard(card int) (numDigits int) {
	divider := 10
	// because is an int, will delete the decimals from each division
	for card != 0 {
		card = card / divider
		numDigits += 1
	}
	return
}

// SPECIAL SUM
// because the log says if multiple by 2 is equal to a number with two digits
// we should sum the two digits from the number
func specialSum(num int) int {
	if num < 5 {
		return num * 2
	} else {
		num = num * 2
		var n1 int = num % 10
		var n2 int = num / 10
		return n1 + n2
	}
}

// CHECK THE RESULT
// compare the info given from the credit cards
func checkResult(result, numDigits, cardNumber int) string {
	if (result % 10) == 0 {
		var divisor int = 1
		for i := 0; i < numDigits-2; i++ {
			divisor = 10 * divisor
		}
		var firstDigits int = cardNumber / divisor
		var firstDigit = firstDigits / 10
		if firstDigits == 34 || firstDigits == 37 {
			if numDigits == 15 {
				return "AMEX"
			} else {
				return "INVALID"
			}
		} else if firstDigits == 51 || firstDigits == 52 || firstDigits == 53 || firstDigits == 54 || firstDigits == 55 {
			if numDigits == 16 {
				return "MASTERCARD"
			} else {
				return "INVALID"
			}
		} else if firstDigit == 4 {
			if numDigits == 13 || numDigits == 16 {
				return "VISA"
			} else {
				return "INVALID"
			}
		}
	}
	return "INVALID"
}
