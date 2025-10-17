package main

import (
	"fmt"
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
func giveUserAChoiceYesOrNo(invitingMes, invalidInpMes string) (bool, error) {
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
	returnNoNeedBreak := func(s string) bool { return false }
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
