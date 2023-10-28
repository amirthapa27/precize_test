package helpers

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func CheckError(err error) bool {
	if err != nil {
		fmt.Println("")
		fmt.Printf("***%s\n***", err)
		fmt.Println("")
		return true
	}
	return false
}

func CheckUserDoseNotExists(err error) bool {
	if err == gorm.ErrRecordNotFound {
		fmt.Println("")
		fmt.Println("***User does not exist***")
		fmt.Println("")
		return true
	}
	return false
}

// will continue to loop untill the data in the fields are proper
func ReadField(scanner *bufio.Scanner, fieldName string, isInt bool) interface{} {
	for {
		fmt.Printf("Enter %s: ", fieldName)
		// reads the next line of input
		scanner.Scan()
		value := strings.TrimSpace(scanner.Text())
		if value != "" {
			if isInt {
				intValue, err := strconv.Atoi(value)
				if err != nil {
					fmt.Printf("%s must be a valid integer. Please try again.\n", fieldName)
					continue
				}
				if fieldName == "Score" && CheckScore(intValue) {
					fmt.Println("Score must be between 0-100")
					continue
				}
				return intValue
			}
			return value
		}

		fmt.Printf("%s cannot be empty. Please try again.\n", fieldName)
	}
}

func CheckScore(score int) bool {
	return score > 100 || score < 0
}

func CalculateIfPassed(score int) string {
	if score > 30 {
		return "PASS"
	}
	return "FAIL"
}
