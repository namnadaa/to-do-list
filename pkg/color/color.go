package color

// ANSI color codes used for colored console output.
var (
	blue    = "\033[34m"       // Menu headers and system messages
	green   = "\033[32m"       // Successful actions and confirmations
	yellow  = "\033[38;5;214m" // Warnings and user prompts
	red     = "\033[31m"       // Cancellations and denials
	magenta = "\033[35m"       // Errors and invalid actions
	reset   = "\033[0m"        // Reset to default color
)

// Blue returns the input text wrapped in ANSI blue color codes.
func Blue(text string) string {
	return blue + text + reset
}

// Green returns the input text wrapped in ANSI green color codes.
func Green(text string) string {
	return green + text + reset
}

// Yellow returns the input text wrapped in ANSI yellow color codes.
func Yellow(text string) string {
	return yellow + text + reset
}

// Red returns the input text wrapped in ANSI red color codes.
func Red(text string) string {
	return red + text + reset
}

// Magenta returns the input text wrapped in ANSI magenta color codes.
func Magenta(text string) string {
	return magenta + text + reset
}
