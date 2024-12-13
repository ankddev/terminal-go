package terminal_go

import (
	"fmt"
	"testing"
)

// TestCursorPosition verifies that CursorPosition function returns correct ANSI escape sequences
// for positioning the cursor at specified coordinates
func TestCursorPosition(t *testing.T) {
	tests := []struct {
		row, col int
		want     string
	}{
		{1, 1, "\033[1;1H"},
		{5, 10, "\033[5;10H"},
		{0, 0, "\033[0;0H"},
	}

	for _, tt := range tests {
		got := CursorPosition(tt.row, tt.col)
		if got != tt.want {
			t.Errorf("CursorPosition(%d, %d) = %q, want %q", tt.row, tt.col, got, tt.want)
		}
	}
}

// ExampleCursorPosition demonstrates moving cursor to specific location
func ExampleCursorPosition() {
	fmt.Print(CursorPosition(5, 10)) // Moves cursor to row 5, column 10
	fmt.Println("Text at position 5,10")
}

// TestCursorMovement verifies that cursor movement functions return correct ANSI escape sequences
// for moving cursor in different directions
func TestCursorMovement(t *testing.T) {
	tests := []struct {
		fn   func(int) string
		n    int
		want string
		name string
	}{
		{CursorForward, 5, "\033[5C", "CursorForward"},
		{CursorBackward, 3, "\033[3D", "CursorBackward"},
		{CursorDown, 2, "\033[2B", "CursorDown"},
		{CursorUp, 4, "\033[4A", "CursorUp"},
		{CursorNextLine, 1, "\033[1E", "CursorNextLine"},
		{CursorPreviousLine, 2, "\033[2F", "CursorPreviousLine"},
		{CursorHorizontalAbsolute, 10, "\033[10G", "CursorHorizontalAbsolute"},
	}

	for _, tt := range tests {
		got := tt.fn(tt.n)
		if got != tt.want {
			t.Errorf("%s(%d) = %q, want %q", tt.name, tt.n, got, tt.want)
		}
	}
}

// ExampleCursorForward demonstrates moving cursor forward
func ExampleCursorForward() {
	fmt.Print(CursorForward(5)) // Moves cursor 5 positions right
	fmt.Println("Text after 5 spaces")
}

// ExampleCursorBackward demonstrates moving cursor backward
func ExampleCursorBackward() {
	fmt.Print("12345")
	fmt.Print(CursorBackward(2)) // Moves cursor 2 positions left
	fmt.Println("X")             // Overwrites character
}

// ExampleCursorUp demonstrates moving cursor up
func ExampleCursorUp() {
	fmt.Println("Line 1")
	fmt.Println("Line 2")
	fmt.Print(CursorUp(1))  // Moves cursor up one line
	fmt.Println("Override") // Overwrites Line 2
}

// TestColorFunctions verifies that color-related functions return correct ANSI escape sequences
// for setting text and background colors
func TestColorFunctions(t *testing.T) {
	tests := []struct {
		name string
		fn   interface{}
		args []int
		want string
	}{
		{"SetTextColor", SetTextColor, []int{1}, "\033[38;5;1m"},
		{"SetBackgroundColor", SetBackgroundColor, []int{2}, "\033[48;5;2m"},
		{"SetRGBTextColor", SetRGBTextColor, []int{255, 128, 0}, "\033[38;2;255;128;0m"},
		{"SetRGBBackgroundColor", SetRGBBackgroundColor, []int{0, 255, 0}, "\033[48;2;0;255;0m"},
	}

	for _, tt := range tests {
		var got string
		switch f := tt.fn.(type) {
		case func(int) string:
			got = f(tt.args[0])
		case func(int, int, int) string:
			got = f(tt.args[0], tt.args[1], tt.args[2])
		}
		if got != tt.want {
			t.Errorf("%s%v = %q, want %q", tt.name, tt.args, got, tt.want)
		}
	}
}

// ExampleSetTextColor demonstrates setting text color
func ExampleSetTextColor() {
	fmt.Print(SetTextColor(1)) // Sets text color to red
	fmt.Println("Red text")
	fmt.Print(ResetAllAttributes)
}

// ExampleSetRGBTextColor demonstrates setting RGB text color
func ExampleSetRGBTextColor() {
	fmt.Print(SetRGBTextColor(255, 0, 0)) // Sets text color to bright red
	fmt.Println("Bright red text")
	fmt.Print(ResetAllAttributes)
}

// TestScrollFunctions verifies that scroll functions return correct ANSI escape sequences
// for scrolling screen content
func TestScrollFunctions(t *testing.T) {
	tests := []struct {
		fn   func(int) string
		n    int
		want string
		name string
	}{
		{ScrollUpLines, 3, "\033[3S", "ScrollUpLines"},
		{ScrollDownLines, 2, "\033[2T", "ScrollDownLines"},
	}

	for _, tt := range tests {
		got := tt.fn(tt.n)
		if got != tt.want {
			t.Errorf("%s(%d) = %q, want %q", tt.name, tt.n, got, tt.want)
		}
	}
}

// ExampleScrollUpLines demonstrates scrolling content up
func ExampleScrollUpLines() {
	fmt.Println("Line 1")
	fmt.Println("Line 2")
	fmt.Print(ScrollUpLines(1)) // Scrolls content up by one line
}

// TestEraseFunctions verifies that erase functions return correct ANSI escape sequences
// for clearing screen content
func TestEraseFunctions(t *testing.T) {
	tests := []struct {
		fn   func(int) string
		n    int
		want string
		name string
	}{
		{EraseInDisplayMode, 0, "\033[0J", "EraseInDisplayMode"},
		{EraseInDisplayMode, 1, "\033[1J", "EraseInDisplayMode"},
		{EraseInDisplayMode, 2, "\033[2J", "EraseInDisplayMode"},
		{EraseInLineMode, 0, "\033[0K", "EraseInLineMode"},
		{EraseInLineMode, 1, "\033[1K", "EraseInLineMode"},
		{EraseInLineMode, 2, "\033[2K", "EraseInLineMode"},
	}

	for _, tt := range tests {
		got := tt.fn(tt.n)
		if got != tt.want {
			t.Errorf("%s(%d) = %q, want %q", tt.name, tt.n, got, tt.want)
		}
	}
}

// ExampleEraseInDisplayMode demonstrates clearing screen
func ExampleEraseInDisplayMode() {
	fmt.Print(EraseInDisplayMode(2)) // Clears entire screen
	fmt.Println("Fresh start")
}

// TestModeFunctions verifies that mode setting functions return correct ANSI escape sequences
func TestModeFunctions(t *testing.T) {
	tests := []struct {
		fn   func(int) string
		n    int
		want string
		name string
	}{
		{SetMode, 4, "\033[4h", "SetMode"},
		{ResetMode, 4, "\033[4l", "ResetMode"},
	}

	for _, tt := range tests {
		got := tt.fn(tt.n)
		if got != tt.want {
			t.Errorf("%s(%d) = %q, want %q", tt.name, tt.n, got, tt.want)
		}
	}
}

// ExampleSetMode demonstrates setting terminal mode
func ExampleSetMode() {
	fmt.Print(SetMode(4)) // Enables insert mode
	fmt.Println("Text in insert mode")
}

// TestWindowManipulation verifies that window manipulation functions return correct ANSI escape sequences
// for controlling terminal window
func TestWindowManipulation(t *testing.T) {
	tests := []struct {
		ps   int
		args []int
		want string
	}{
		{1, nil, "\033[1t"},
		{3, []int{100, 200}, "\033[3;100;200t"},
		{4, []int{800, 600}, "\033[4;800;600t"},
		{8, []int{24, 80}, "\033[8;24;80t"},
		{5, nil, "\033[5t"},
	}

	for _, tt := range tests {
		got := WindowManipulation(tt.ps, tt.args...)
		if got != tt.want {
			t.Errorf("WindowManipulation(%d, %v) = %q, want %q", tt.ps, tt.args, got, tt.want)
		}
	}
}

// ExampleWindowManipulation demonstrates window manipulation
func ExampleWindowManipulation() {
	fmt.Print(WindowManipulation(8, 24, 80)) // Resizes window to 24 rows and 80 columns
	fmt.Println("Window resized")
}

// TestSetScrollingRegion verifies that SetScrollingRegion returns correct ANSI escape sequences
// for setting scrollable region
func TestSetScrollingRegion(t *testing.T) {
	tests := []struct {
		top, bottom int
		want        string
	}{
		{0, 24, "\033[0;24r"},
		{1, 10, "\033[1;10r"},
	}

	for _, tt := range tests {
		got := SetScrollingRegion(tt.top, tt.bottom)
		if got != tt.want {
			t.Errorf("SetScrollingRegion(%d, %d) = %q, want %q", tt.top, tt.bottom, got, tt.want)
		}
	}
}

// ExampleSetScrollingRegion demonstrates setting scroll region
func ExampleSetScrollingRegion() {
	fmt.Print(SetScrollingRegion(1, 10)) // Sets scrolling region from line 1 to 10
	fmt.Println("Scrolling region set")
}

// TestLineManipulation verifies that line manipulation functions return correct ANSI escape sequences
// for inserting and deleting lines and characters
func TestLineManipulation(t *testing.T) {
	tests := []struct {
		fn   func(int) string
		n    int
		want string
		name string
	}{
		{DeleteLines, 2, "\033[2M", "DeleteLines"},
		{InsertLines, 3, "\033[3L", "InsertLines"},
		{DeleteCharacters, 4, "\033[4P", "DeleteCharacters"},
		{InsertCharacters, 5, "\033[5@", "InsertCharacters"},
	}

	for _, tt := range tests {
		got := tt.fn(tt.n)
		if got != tt.want {
			t.Errorf("%s(%d) = %q, want %q", tt.name, tt.n, got, tt.want)
		}
	}
}

// ExampleDeleteLines demonstrates deleting lines
func ExampleDeleteLines() {
	fmt.Println("Line 1")
	fmt.Println("Line 2")
	fmt.Print(CursorUp(1))
	fmt.Print(DeleteLines(1)) // Deletes current line
}

// TestSetGraphicsRendition verifies that SetGraphicsRendition returns correct ANSI escape sequences
// for setting text attributes
func TestSetGraphicsRendition(t *testing.T) {
	tests := []struct {
		params []int
		want   string
	}{
		{nil, "\033[m"},
		{[]int{0}, "\033[0m"},
		{[]int{1, 4, 31}, "\033[1;4;31m"},
	}

	for _, tt := range tests {
		got := SetGraphicsRendition(tt.params...)
		if got != tt.want {
			t.Errorf("SetGraphicsRendition(%v) = %q, want %q", tt.params, got, tt.want)
		}
	}
}

// ExampleSetGraphicsRendition demonstrates setting multiple text attributes
func ExampleSetGraphicsRendition() {
	fmt.Print(SetGraphicsRendition(1, 4)) // Sets bold and underline
	fmt.Println("Bold and underlined text")
	fmt.Print(ResetAllAttributes)
}

// TestCursorPositionReport verifies that cursor position reporting functions return correct ANSI escape sequences
func TestCursorPositionReport(t *testing.T) {
	if got := RequestCursorPosition(); got != "\033[6n" {
		t.Errorf("RequestCursorPosition() = %q, want %q", got, "\033[6n")
	}

	if got := ReportCursorPosition(5, 10); got != "\033[5;10R" {
		t.Errorf("ReportCursorPosition(5, 10) = %q, want %q", got, "\033[5;10R")
	}
}

// ExampleRequestCursorPosition demonstrates requesting cursor position
func ExampleRequestCursorPosition() {
	fmt.Print(RequestCursorPosition())
	// Note: The actual response needs to be read from stdin
}

// TestCharacterSetFunctions verifies that character set functions return correct ANSI escape sequences
// for selecting and designating character sets
func TestCharacterSetFunctions(t *testing.T) {
	tests := []struct {
		fn      func(int, byte) string
		g       int
		charset byte
		want    string
		name    string
	}{
		{SetCharacterSet, 0, 'B', "\033(B", "SetCharacterSet"},
		{SetCharacterSet, 1, '0', "\033)0", "SetCharacterSet"},
		{DesignateCharacterSet, 0, 'B', "\033(B", "DesignateCharacterSet"},
		{DesignateCharacterSet, 1, '0', "\033)0", "DesignateCharacterSet"},
		{DesignateCharacterSet, 2, 'B', "\033*B", "DesignateCharacterSet"},
		{DesignateCharacterSet, 3, '0', "\033+0", "DesignateCharacterSet"},
	}

	for _, tt := range tests {
		got := tt.fn(tt.g, tt.charset)
		if got != tt.want {
			t.Errorf("%s(%d, %c) = %q, want %q", tt.name, tt.g, tt.charset, got, tt.want)
		}
	}
}

// ExampleSetCharacterSet demonstrates setting character set
func ExampleSetCharacterSet() {
	fmt.Print(SetCharacterSet(0, 'B')) // Sets default character set to US ASCII
	fmt.Println("ASCII text")
}

// TestTerminalControl verifies that terminal control functions return correct ANSI escape sequences
// for controlling terminal behavior
func TestTerminalControl(t *testing.T) {
	if got := SoftTerminalReset(); got != "\033[!p" {
		t.Errorf("SoftTerminalReset() = %q, want %q", got, "\033[!p")
	}

	if got := RequestTerminalParameters(); got != "\033[x" {
		t.Errorf("RequestTerminalParameters() = %q, want %q", got, "\033[x")
	}

	if got := SetConformanceLevel(2); got != "\033[2\"p" {
		t.Errorf("SetConformanceLevel(2) = %q, want %q", got, "\033[2\"p")
	}
}

// ExampleSoftTerminalReset demonstrates soft terminal reset
func ExampleSoftTerminalReset() {
	fmt.Print(SoftTerminalReset())
	fmt.Println("Terminal reset")
}

// TestConstants verifies that all constant ANSI escape sequences are non-empty
func TestConstants(t *testing.T) {
	constants := map[string]string{
		"ReverseIndex":                    ReverseIndex,
		"SaveCursorPointerInMemory":       SaveCursorPointerInMemory,
		"RestoreCursorPointerFromMemory":  RestoreCursorPointerFromMemory,
		"CursorBlinking":                  CursorBlinking,
		"CursorBlinkingDisable":           CursorBlinkingDisable,
		"ShowCursor":                      ShowCursor,
		"HideCursor":                      HideCursor,
		"EnterAltScreen":                  EnterAltScreen,
		"ExitAltScreen":                   ExitAltScreen,
		"EnableLineWrap":                  EnableLineWrap,
		"DisableLineWrap":                 DisableLineWrap,
		"EraseInDisplay":                  EraseInDisplay,
		"EraseInLine":                     EraseInLine,
		"ScrollUp":                        ScrollUp,
		"ScrollDown":                      ScrollDown,
		"SaveCursorPosition":              SaveCursorPosition,
		"RestoreCursorPosition":           RestoreCursorPosition,
		"EnableVirtualTerminalProcessing": EnableVirtualTerminalProcessing,
		"ResetAllAttributes":              ResetAllAttributes,
		"BoldBright":                      BoldBright,
		"NormalIntensity":                 NormalIntensity,
		"Underline":                       Underline,
		"UnderlineDisable":                UnderlineDisable,
		"Negative":                        Negative,
		"Positive":                        Positive,
		"TabSet":                          TabSet,
		"TabClear":                        TabClear,
		"TabClearAll":                     TabClearAll,
		"ApplicationKeypad":               ApplicationKeypad,
		"NormalKeypad":                    NormalKeypad,
		"AutoWrap":                        AutoWrap,
		"AutoWrapOff":                     AutoWrapOff,
		"ClearAndResetScrollback":         ClearAndResetScrollback,
	}

	for name, constant := range constants {
		if constant == "" {
			t.Errorf("Constant %s is empty", name)
		}
	}
}
