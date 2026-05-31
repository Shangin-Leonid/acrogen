package ui /* User Interface */

import (
	"fmt"
	"strconv"

	"github.com/Shangin-Leonid/acrogen/utils"
)

// Binary choice constants (Yes == true, No == !Yes).
const Yes, No = true, !Yes

// # giveUserYesOrNoChoice prints a string inviting user to make a decision: yes or no and processes his choice.
//
// # Params:
//
//   - inviting message
//   - message for invalid input
//
// # Returns:
//
//   - user choice as bool: yes or no
//   - error for invalid input
//
// # TODOs:
//
//   - implement several tries for input (amount of tries as parameter).
//   - maybe add 2 callbacks as params: 1 for yes and 1 for no. Then implify usage of the function.
func giveUserYesOrNoChoice(invitingMes, invalidInpMes string) (bool, error) {

	isInpValid := func(inp string) (bool, error) {
		return inp == "y" || inp == "n", nil
	}
	var YesOrNo bool
	isYes := func(inp string) error {
		YesOrNo = utils.TerOp(inp == "y", Yes, No)
		return nil
	}
	returnIfYesOrNoInput := func(inp string) bool { return (inp == "y" || inp == "n") }

	err, _ := processUserInputUntil(
		invitingMes,
		"Print [y/n]:",
		invalidInpMes,
		nil,
		isInpValid,
		isYes,
		returnIfYesOrNoInput)
	return YesOrNo, err
}

// # giveUserNumberChoice prints a string inviting user to make a decision about number and processes his choice.
//
// # Params:
//
//   - inviting message
//   - message for invalid input
//
// # Returns:
//
//   - user choice: number (0 if error)
//   - error for invalid input
//
// # TODOs:
//
//   - implement several tries for input (amount of tries as parameter).
func giveUserNumberChoice(invitingMes, invalidInpMes string) (userNum int, err error) {

	isInpValid := func(inp string) (bool, error) {
		userNum, err = strconv.Atoi(inp)
		if err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
	returnNeedBreak := func(s string) bool { return true }

	err, _ = processUserInputUntil(
		invitingMes,
		"Enter a number:",
		invalidInpMes,
		nil,
		isInpValid,
		nil,
		returnNeedBreak)

	return userNum, err
}

// # giveUserChoiceOfFilename prints a string inviting user to make a decision about filename and processes his choice.
//
// # Params:
//
//   - inviting message
//   - message for invalid input
//
// # Returns:
//
//   - user choice: filename as string
//   - error for invalid input
func giveUserChoiceOfFilename(invitingMes string) (filename string, err error) {
	MenuColor.Printf("%s %s\n", MessagePrefix, invitingMes)

	_, err = fmt.Scanf("%s", &filename)
	if err != nil {
		return "", utils.NewSTError(err.Error())
	}

	if !utils.IsTextFileNameValid(filename) {
		return "", utils.NewSTError("incorrect text file name")
	}

	return filename, nil
}

// # processUserInputUntilExitCommand
//
// # Description:
//
// No words about format, just look inside...
//
// # TODOs:
//
//   - docs
//   - simplify if possible
//   - maybe change the signature
func processUserInputUntilExitCommand(
	exitCommand string,
	invitingMes string,
	userGuideMes string,
	invalidInpMes string,
	checkIfInpValid func(string) (bool, error),
	processInp func(string) error) (err error, nProcessed int) {

	returnIfExitCommand := func(s string) bool { return s == exitCommand }

	return processUserInputUntil(
		invitingMes,
		userGuideMes,
		invalidInpMes,
		returnIfExitCommand,
		checkIfInpValid,
		processInp,
		nil)
}

// # processUserInputUntil
//
// # Description:
//
// No words about format, just look inside...
//
// # TODOs:
//
//   - docs
//   - simplify if possible
//   - maybe change the signature
func processUserInputUntil(
	invitingMes string,
	userGuideMes string,
	invalidInpMes string,
	checkIfNeedBreakBeforeValidation func(string) bool,
	checkIfInpValid func(string) (bool, error),
	processInp func(string) error,
	checkIfNeedBreakAfterProcess func(string) bool) (err error, nProcessed int) {

	var userInp string

	if invitingMes != "" {
		MenuColor.Printf("%s ", invitingMes)
	}

	for {
		MenuColor.Printf("%s\n", userGuideMes)
		_, err = fmt.Scanf("%s", &userInp)
		if err != nil {
			return utils.NewSTError(err.Error()), nProcessed
		}

		if checkIfNeedBreakBeforeValidation != nil && checkIfNeedBreakBeforeValidation(userInp) {
			break
		}

		isInpValid, err := checkIfInpValid(userInp)
		if err != nil {
			return err, nProcessed
		}

		if isInpValid {
			if processInp != nil {
				err = processInp(userInp)
				nProcessed++
				if err != nil {
					return err, nProcessed
				}
			}
		} else {
			if invalidInpMes != "" {
				WarningColor.Printf("%s\n", invalidInpMes)
			}
		}

		if checkIfNeedBreakAfterProcess != nil && checkIfNeedBreakAfterProcess(userInp) {
			break
		}
	}

	return nil, nProcessed
}
