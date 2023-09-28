package validator

import (
	"fmt"
	"main/storage"
	"net/mail"
	"strings"
	"time"
)

const (
	replaceChars      = `!@$&*`
	sepChars          = `_-., `
	otherSpecialChars = `"#%'()+/:;<=>?[\]^{|}~`
	lowerChars        = `abcdefghijklmnopqrstuvwxyz`
	upperChars        = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	digitsChars       = `0123456789`
	maxCharacters     = 32
	minCharacters     = 8
)

func ValidateNewUser(user storage.User) error {
	err := validateEmail(user.Email)
	if err != nil {
		return err
	}

	err = validatePassword(user.Password)
	if err != nil {
		return err
	}

	err = validateNickname(user.Nickname)
	if err != nil {
		return err
	}

	err = validateDate(user.BirthDate)
	if err != nil {
		return err
	}

	return nil
}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func validateDate(dateString string) error {
	fmt.Println(dateString)
	_, err := time.Parse(time.DateOnly, dateString)
	return err
}

func validateNickname(nickname string) error {
	if nickname == "" {
		return fmt.Errorf("nickname should not be empty")
	}
	return nil
}

func validatePassword(password string) error {
	var (
		hasReplace      bool
		hasSep          bool
		hasOtherSpecial bool
		hasLower        bool
		hasUpper        bool
		hasDigits       bool
	)

	letters := len([]rune(password))
	if letters > maxCharacters {
		return fmt.Errorf("too long password")
	} else if letters < minCharacters {
		return fmt.Errorf("insecure password, try using a longer password")
	}

	for _, c := range password {
		switch {
		case strings.ContainsRune(replaceChars, c):
			hasReplace = true
		case strings.ContainsRune(sepChars, c):
			hasSep = true
		case strings.ContainsRune(otherSpecialChars, c):
			hasOtherSpecial = true
		case strings.ContainsRune(lowerChars, c):
			hasLower = true
		case strings.ContainsRune(upperChars, c):
			hasUpper = true
		case strings.ContainsRune(digitsChars, c):
			hasDigits = true
		}
	}

	allMessages := []string{}

	if !hasOtherSpecial || !hasSep || !hasReplace {
		allMessages = append(allMessages, "including more special characters")
	}
	if !hasLower {
		allMessages = append(allMessages, "using lowercase letters")
	}
	if !hasUpper {
		allMessages = append(allMessages, "using uppercase letters")
	}
	if !hasDigits {
		allMessages = append(allMessages, "using numbers")
	}

	if len(allMessages) > 0 {
		return fmt.Errorf("insecure password, try %v", strings.Join(allMessages, ", "))
	}

	return nil
}
