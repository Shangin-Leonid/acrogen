package cio /* Console Input Output */

import (
	"acrogen/fio"
	"errors"
	"fmt"
	"strconv"
)

// Binary choice constants (Yes == true, No == !Yes).
const Yes, No = true, !Yes

// #
// Prints a string inviting user to make a decision: yes or no.
// Returns 'Yes'(==true) or 'No'(== !Yes) and error, if user input is incorrect.
// TODO implement several tries for input (amount of tries as parameter)
// TODO maybe add 2 callbacks as params: 1 for yes and 1 for no. Then implify usage of the function.
// #
func GiveUserYesOrNoChoice(invitingMes, invalidInpMes string) (bool, error) {

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
	returnIfYesOrNoInput := func(inp string) bool { return (inp == "y" || inp == "n") }

	err, _ := ProcessUserInputUntil(
		invitingMes,
		"Print [y/n]",
		invalidInpMes,
		nil,
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
func GiveUserNumberChoice(invitingMes, invalidInpMes string) (userNum int, err error) {

	isInpValid := func(inp string) (bool, error) {
		userNum, err = strconv.Atoi(inp)
		if err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
	returnNeedBreak := func(s string) bool { return true }

	err, _ = ProcessUserInputUntil(
		invitingMes,
		"Print a number",
		invalidInpMes,
		nil,
		isInpValid,
		nil,
		returnNeedBreak)

	return userNum, err
}

// #
// TODO docs
// #
func GiveUserChoiceOfFilename(invitingMes string) (filename string, err error) {
	fmt.Printf(invitingMes)

	_, err = fmt.Scanf("%s", &filename)
	if err != nil {
		return "", err
	}

	if !fio.IsTextFileNameValid(filename) {
		return "", errors.New("incorrect text file name")
	}

	return filename, nil
}

// #
// No words about format, just look inside...
// TODO documentation
// #
func ProcessUserInputUntilExitCommand(
	exitCommand string,
	invitingMes string,
	userGuideMes string,
	invalidInpMes string,
	checkIfInpValid func(string) (bool, error),
	processInp func(string) error) (err error, nProcessed int) {

	returnIfExitCommand := func(s string) bool { return s == exitCommand }

	return ProcessUserInputUntil(
		invitingMes,
		userGuideMes,
		invalidInpMes,
		returnIfExitCommand,
		checkIfInpValid,
		processInp,
		nil)
}

// #
// No words about format, just look inside...
// TODO documentation
// #
func ProcessUserInputUntil(
	invitingMes string,
	userGuideMes string,
	invalidInpMes string,
	checkIfNeedBreakBeforeValidation func(string) bool,
	checkIfInpValid func(string) (bool, error),
	processInp func(string) error,
	checkIfNeedBreakAfterProcess func(string) bool) (err error, nProcessed int) {

	var userInp string

	if invitingMes != "" {
		fmt.Printf("%s\n", invitingMes)
	}

	for {
		fmt.Printf("%s\n", userGuideMes)
		_, err = fmt.Scanf("%s", &userInp)
		if err != nil {
			return err, nProcessed
		}

		if checkIfNeedBreakBeforeValidation != nil && checkIfNeedBreakBeforeValidation(userInp) {
			break
		}

		isInpValid, err := checkIfInpValid(userInp)
		if err != nil {
			return err, nProcessed
		}

		if isInpValid && processInp != nil {
			err = processInp(userInp)
			nProcessed++
			if err != nil {
				return err, nProcessed
			}
		} else {
			if invalidInpMes != "" {
				fmt.Printf("%s\n", invalidInpMes)
			}
		}

		if checkIfNeedBreakAfterProcess != nil && checkIfNeedBreakAfterProcess(userInp) {
			break
		}
	}

	return nil, nProcessed
}
