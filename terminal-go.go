package terminal_go

import "fmt"

const (
	// ReverseIndex performs the reverse operation of \n, moves cursor up one line, maintains horizontal position, scrolls buffer if necessary
	ReverseIndex = "\033M"
	// SaveCursorPointerInMemory saves the cursor position in memory
	SaveCursorPointerInMemory = "\0337"
	// RestoreCursorPointerFromMemory restores the cursor position from memory
	RestoreCursorPointerFromMemory = "\0338"

	// CursorBlinking enables cursor blinking
	CursorBlinking = "\033[?12h"
	// CursorBlinkingDisable disables cursor blinking
	CursorBlinkingDisable = "\033[?12l"
	// ShowCursor shows the cursor
	ShowCursor = "\033[?25h"
	// HideCursor hides the cursor
	HideCursor = "\033[?25l"

	// EnterAltScreen switches to the alternate screen buffer
	EnterAltScreen = "\033[?1049h"
	// ExitAltScreen switches back to the main screen buffer
	ExitAltScreen = "\033[?1049l"

	// EnableLineWrap enables line wrapping
	EnableLineWrap = "\033[?7h"
	// DisableLineWrap disables line wrapping
	DisableLineWrap = "\033[?7l"

	// EraseInDisplay clears the screen and moves cursor to home position
	EraseInDisplay = "\033[2J"
	// EraseInLine clears from cursor to the end of line
	EraseInLine = "\033[K"

	// ScrollUp scrolls display up one line
	ScrollUp = "\033[S"
	// ScrollDown scrolls display down one line
	ScrollDown = "\033[T"

	// SaveCursorPosition saves current cursor position
	SaveCursorPosition = "\033[s"
	// RestoreCursorPosition restores cursor to last saved position
	RestoreCursorPosition = "\033[u"

	// EnableVirtualTerminalProcessing enables VT processing
	EnableVirtualTerminalProcessing = "\033[?1h"

	// ResetAllAttributes resets all character attributes
	ResetAllAttributes = "\033[0m"

	// BoldBright sets bold/bright mode
	BoldBright = "\033[1m"
	// NormalIntensity sets normal intensity
	NormalIntensity = "\033[22m"

	// Underline enables underline
	Underline = "\033[4m"
	// UnderlineDisable disables underline
	UnderlineDisable = "\033[24m"

	// Negative sets negative image
	Negative = "\033[7m"
	// Positive sets positive image
	Positive = "\033[27m"

	// TabSet sets a tab stop at current position
	TabSet = "\033H"
	// TabClear clears tab stop at current position
	TabClear = "\033[0g"
	// TabClearAll clears all tab stops
	TabClearAll = "\033[3g"

	// DoubleHeightTop makes current line the top half of double height characters
	DoubleHeightTop = "\033#3"
	// DoubleHeightBottom makes current line the bottom half of double height characters
	DoubleHeightBottom = "\033#4"
	// SingleWidthLine makes current line single-width single-height
	SingleWidthLine = "\033#5"
	// DoubleWidthLine makes current line double-width single-height
	DoubleWidthLine = "\033#6"

	// DeviceStatusReport requests cursor position report
	DeviceStatusReport = "\033[6n"

	// ApplicationKeypad enables application keypad mode
	ApplicationKeypad = "\033="
	// NormalKeypad enables numeric keypad mode
	NormalKeypad = "\033>"

	// AutoWrap enables auto-wrap mode
	AutoWrap = "\033[?7h"
	// AutoWrapOff disables auto-wrap mode
	AutoWrapOff = "\033[?7l"

	// ClearAndResetScrollback clears screen and scrollback buffer
	ClearAndResetScrollback = "\033[3J"
)

// CursorPosition sets the cursor position where subsequent text will begin
func CursorPosition(row, column int) string {
	return fmt.Sprintf("\033[%d;%dH", row, column)
}

// CursorForward moves cursor forward by specified number of columns
func CursorForward(columns int) string {
	return fmt.Sprintf("\033[%dC", columns)
}

// CursorBackward moves cursor backward by specified number of columns
func CursorBackward(columns int) string {
	return fmt.Sprintf("\033[%dD", columns)
}

// CursorDown moves cursor down by specified number of lines
func CursorDown(lines int) string {
	return fmt.Sprintf("\033[%dB", lines)
}

// CursorUp moves cursor up by specified number of lines
func CursorUp(lines int) string {
	return fmt.Sprintf("\033[%dA", lines)
}

// CursorNextLine moves cursor to beginning of line n lines down
func CursorNextLine(lines int) string {
	return fmt.Sprintf("\033[%dE", lines)
}

// CursorPreviousLine moves cursor to beginning of line n lines up
func CursorPreviousLine(lines int) string {
	return fmt.Sprintf("\033[%dF", lines)
}

// CursorHorizontalAbsolute moves cursor to specified column
func CursorHorizontalAbsolute(column int) string {
	return fmt.Sprintf("\033[%dG", column)
}

// SetTextColor sets the foreground color
func SetTextColor(color int) string {
	return fmt.Sprintf("\033[38;5;%dm", color)
}

// SetBackgroundColor sets the background color
func SetBackgroundColor(color int) string {
	return fmt.Sprintf("\033[48;5;%dm", color)
}

// SetRGBTextColor sets the foreground color using RGB values
func SetRGBTextColor(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// SetRGBBackgroundColor sets the background color using RGB values
func SetRGBBackgroundColor(r, g, b int) string {
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

// ScrollUpLines scrolls screen up by n lines
func ScrollUpLines(lines int) string {
	return fmt.Sprintf("\033[%dS", lines)
}

// ScrollDownLines scrolls screen down by n lines
func ScrollDownLines(lines int) string {
	return fmt.Sprintf("\033[%dT", lines)
}

// EraseInDisplayMode erases display with specified mode
// n=0: erase from cursor to end of display
// n=1: erase from start of display to cursor
// n=2: erase complete display
// n=3: erase scrollback buffer
func EraseInDisplayMode(n int) string {
	return fmt.Sprintf("\033[%dJ", n)
}

// EraseInLineMode erases line with specified mode
// n=0: erase from cursor to end of line
// n=1: erase from start of line to cursor
// n=2: erase complete line
func EraseInLineMode(n int) string {
	return fmt.Sprintf("\033[%dK", n)
}

// SetMode sets various terminal modes
func SetMode(mode int) string {
	return fmt.Sprintf("\033[%dh", mode)
}

// ResetMode resets various terminal modes
func ResetMode(mode int) string {
	return fmt.Sprintf("\033[%dl", mode)
}

// WindowManipulation performs window manipulation
// ps=1: de-iconify window
// ps=2: iconify window
// ps=3: move window to x,y
// ps=4: resize window to height,width in pixels
// ps=5: raise window to top of stack
// ps=6: lower window to bottom of stack
// ps=7: refresh window
// ps=8: resize window to rows,cols in characters
// ps=9: maximize/restore window
func WindowManipulation(ps int, args ...int) string {
	if len(args) == 0 {
		return fmt.Sprintf("\033[%dt", ps)
	}
	switch ps {
	case 3:
		if len(args) >= 2 {
			return fmt.Sprintf("\033[3;%d;%dt", args[0], args[1])
		}
	case 4:
		if len(args) >= 2 {
			return fmt.Sprintf("\033[4;%d;%dt", args[0], args[1])
		}
	case 8:
		if len(args) >= 2 {
			return fmt.Sprintf("\033[8;%d;%dt", args[0], args[1])
		}
	}
	return fmt.Sprintf("\033[%dt", ps)
}

// SetScrollingRegion sets top and bottom margins
func SetScrollingRegion(top, bottom int) string {
	return fmt.Sprintf("\033[%d;%dr", top, bottom)
}

// DeleteLines deletes n lines at cursor
func DeleteLines(n int) string {
	return fmt.Sprintf("\033[%dM", n)
}

// InsertLines inserts n lines at cursor
func InsertLines(n int) string {
	return fmt.Sprintf("\033[%dL", n)
}

// DeleteCharacters deletes n characters at cursor
func DeleteCharacters(n int) string {
	return fmt.Sprintf("\033[%dP", n)
}

// InsertCharacters inserts n spaces at cursor
func InsertCharacters(n int) string {
	return fmt.Sprintf("\033[%d@", n)
}

// SetGraphicsRendition sets various text attributes
// Multiple parameters can be combined
func SetGraphicsRendition(params ...int) string {
	if len(params) == 0 {
		return "\033[m"
	}
	var args string
	for i, param := range params {
		if i > 0 {
			args += ";"
		}
		args += fmt.Sprintf("%d", param)
	}
	return fmt.Sprintf("\033[%sm", args)
}

// RequestCursorPosition requests cursor position and returns response sequence
func RequestCursorPosition() string {
	return "\033[6n"
}

// ReportCursorPosition formats cursor position report response
func ReportCursorPosition(row, col int) string {
	return fmt.Sprintf("\033[%d;%dR", row, col)
}

// SetCharacterSet selects character set
// g can be 0 or 1 for default/alternate
// charset can be B for US ASCII, 0 for DEC Special Graphics
func SetCharacterSet(g int, charset byte) string {
	if g == 0 {
		return fmt.Sprintf("\033(%c", charset)
	}
	return fmt.Sprintf("\033)%c", charset)
}

// DesignateCharacterSet designates character set
// g can be 0-3 for G0-G3
// charset can be B for US ASCII, 0 for DEC Special Graphics
func DesignateCharacterSet(g int, charset byte) string {
	switch g {
	case 0:
		return fmt.Sprintf("\033(%c", charset)
	case 1:
		return fmt.Sprintf("\033)%c", charset)
	case 2:
		return fmt.Sprintf("\033*%c", charset)
	case 3:
		return fmt.Sprintf("\033+%c", charset)
	default:
		return ""
	}
}

// SoftTerminalReset performs a soft terminal reset
func SoftTerminalReset() string {
	return "\033[!p"
}

// RequestTerminalParameters requests terminal parameters
func RequestTerminalParameters() string {
	return "\033[x"
}

// SetConformanceLevel sets ANSI conformance level
// level can be 1, 2, or 3
func SetConformanceLevel(level int) string {
	return fmt.Sprintf("\033[%d\"p", level)
}
