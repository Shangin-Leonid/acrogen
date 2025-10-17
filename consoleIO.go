package main

import (
	"fmt"
	"strconv"
)

// #
// Prints acronym in console in detailed format (decodes each letter).
// #
func printAcronymInDetail(acr Acronym) {
	fmt.Printf("%s%s%d\n", acr.word, TokenSeparator, acr.sumEstimation)
	for i, letter := range []rune(acr.word) {
		fmt.Printf("%c -- %s\n", letter, acr.letterDecodings[i])
	}
}

// Binary choice constants (Yes == true, No == !Yes).
const Yes, No = true, !Yes

// #
// Prints a string inviting user to make a decision: yes or no.
// Returns 'Yes'(==true) or 'No'(== !Yes) and error, if user input is incorrect.
// TODO implement several tries for input (amount of tries as parameter)
// #
func giveUserYesOrNoChoice(invitingMes, invalidInpMes string) (bool, error) {

	returnNoNeedBreak := func(s string) bool { return false }
	isInpValid := func(inp string) (bool, error) {
		return inp == "y" || inp == "n", nil
	}
	var YesOrNo bool
	isYes := func(inp string) error {
		if inp == "y" {
			YesOrNo = Yes
		} else {
			YesOrNo = No
		}
		return nil
	}
	returnIfYesOrNoInput := func(s string) bool { return (s == "y" || s == "n") }

	err, _ := processUserInputUntil(
		invitingMes,
		"Print [y/n]",
		invalidInpMes,
		returnNoNeedBreak,
		isInpValid,
		isYes,
		returnIfYesOrNoInput)
	return YesOrNo, err
}

// #
// Prints a string inviting user to make a decision about number.
// Returns the entered number (0, if err) and error, if user input is incorrect.
// TODO implement several tries for input (amount of tries as parameter)
// #
func giveUserNumberChoice(invitingMes, invalidInpMes string) (userNum int, err error) {

	returnNoNeedBreak := func(s string) bool { return false }
	isInpValid := func(inp string) (bool, error) {
		userNum, err = strconv.Atoi(inp)
		if err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
	doNothing := func(inp string) error {
		return nil
	}
	returnNeedBreak := func(s string) bool { return true }

	err, _ = processUserInputUntil(
		invitingMes,
		"Print a number",
		invalidInpMes,
		returnNoNeedBreak,
		isInpValid,
		doNothing,
		returnNeedBreak)

	return userNum, err
}

// Equals to string that represents user's query of exit.
const ExitCommand = "!q"

// #
// No words about format, just look inside...
// TODO documentation
// #
func processUserInputUntilExitCommand(
	invitingMes string,
	userGuideMes string,
	invalidInpMes string,
	checkIfInpValid func(string) (bool, error),
	processInp func(string) error) (err error, nProcessed int) {

	returnIfExitCommand := func(s string) bool { return s == ExitCommand }
	returnNoNeedBreak := func(s string) bool { return false }

	fmt.Printf("\nTo exit (to stop) enter \"%s\"\n", ExitCommand)

	return processUserInputUntil(
		invitingMes,
		userGuideMes,
		invalidInpMes,
		returnIfExitCommand,
		checkIfInpValid,
		processInp,
		returnNoNeedBreak)
}

// #
// No words about format, just look inside...
// TODO documentation
// #
func processUserInputUntil(
	invitingMes string,
	userGuideMes string,
	invalidInpMes string,
	checkIfNeedBreakBeforeValidationAndProcess func(string) bool,
	checkIfInpValid func(string) (bool, error),
	processInp func(string) error,
	checkIfNeedBreakAfterProcess func(string) bool) (err error, nProcessed int) {

	var userInp string

	if invitingMes != "" {
		fmt.Printf("%s %s\n", invitingMes, userGuideMes)
	} else {
		fmt.Printf("%s\n", userGuideMes)
	}

	for {
		_, err = fmt.Scanf("%s", &userInp)
		if err != nil {
			return err, nProcessed
		}

		if checkIfNeedBreakBeforeValidationAndProcess(userInp) {
			break
		}

		isInpValid, err := checkIfInpValid(userInp)
		if err != nil {
			return err, nProcessed
		}

		if isInpValid {
			err = processInp(userInp)
			nProcessed++
			if err != nil {
				return err, nProcessed
			}
		} else {
			if invalidInpMes != "" {
				fmt.Printf("%s %s\n", invalidInpMes, userGuideMes)
			}
		}

		if checkIfNeedBreakAfterProcess(userInp) {
			break
		}
	}

	return nil, nProcessed
}
