package ut /* Unit Tests*/

import (
	"github.com/fatih/color"
)

// Messages
const (
	PassedMes = "PASSED"
	FailedMes = "FAILED"
)

// Color themes
var NeutralColor *color.Color = color.New(color.FgYellow, color.Faint)
var NeutralBoldColor *color.Color = NeutralColor.Add(color.Bold)
var PassedColor *color.Color = color.New(color.FgGreen, color.Bold)
var FailedColor *color.Color = color.New(color.FgRed, color.Bold)
var FailReasonColor *color.Color = color.New(color.FgWhite, color.Faint)
