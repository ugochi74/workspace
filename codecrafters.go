package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ─────────────────────────────────────────────
// GLOBAL STATE
// ─────────────────────────────────────────────
var (
	lastResult    float64
	hasLastResult bool
	history       []string
	calcHistory   []string
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to C&C Console. Type 'help' for commands.")

	for {
		fmt.Print("\nC&C> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		// ─────────────────────────────────────────────
		// STEP 8 — PIPE HANDLING
		// ─────────────────────────────────────────────
		if strings.Contains(input, "|") {
			handlePipe(input)
			continue
		}

		parts := strings.Fields(input)

		// ─────────────────────────────────────────────
		// STEP 3 — ROUTER
		// ─────────────────────────────────────────────
		switch strings.ToLower(parts[0]) {
		case "calc":
			handleCalc(parts[1:])
		case "base":
			handleBase(parts[1:])
		case "str":
			handleStr(parts[1:])
		case "help":
			printHelp()
		case "history":
			printHistory()
		case "exit":
			fmt.Println("Shutting down. Goodbye.")
			return
		default:
			fmt.Printf("     ✗ Unknown command: %q. Type 'help' for commands.\n", parts[0])
		}
	}
}

// ─────────────────────────────────────────────
// STEP 4 — CALCULATOR
// ─────────────────────────────────────────────
func handleCalc(parts []string) {
	if len(parts) == 0 {
		fmt.Println("     ✗ Error: missing command. Usage: calc <operation> <a> <b>")
		return
	}

	operation := strings.ToLower(parts[0])

	if operation == "last" {
		if !hasLastResult {
			fmt.Println("     ✗ No previous result in this session.")
			return
		}
		fmt.Printf("     ✦ Last Result: %g\n", lastResult)
		return
	}

	if operation == "history" {
		if len(calcHistory) == 0 {
			fmt.Println("     ✗ No calc history yet.")
			return
		}
		fmt.Println("     ✦ Calc History:")
		start := len(calcHistory) - 5
		if start < 0 {
			start = 0
		}
		for i, entry := range calcHistory[start:] {
			fmt.Printf("       %d. %s\n", i+1, entry)
		}
		return
	}

	if len(parts) < 3 {
		fmt.Println("     ✗ Error: missing arguments. Usage: calc <operation> <a> <b>")
		return
	}

	a, ok := resolveNumber(parts[1])
	if !ok {
		return
	}
	b, ok := resolveNumber(parts[2])
	if !ok {
		return
	}

	var result float64
	switch operation {
	case "add":
		result = a + b
	case "sub":
		result = a - b
	case "mul":
		result = a * b
	case "div":
		if b == 0 {
			fmt.Println("     ✗ Error: cannot divide by zero.")
			return
		}
		result = a / b
	case "mod":
		if b == 0 {
			fmt.Println("     ✗ Error: cannot mod by zero.")
			return
		}
		result = math.Mod(a, b)
	case "pow":
		result = math.Pow(a, b)
	default:
		fmt.Printf("     ✗ Error: unknown operation %q. Use: add sub mul div mod pow\n", operation)
		return
	}

	fmt.Printf("     ✦ Result: %g\n", result)
	setLastResult(result)
	entry := fmt.Sprintf("calc %s %g %g → %g", operation, a, b, result)
	calcHistory = append(calcHistory, entry)
	addHistory(entry)
}

// ─────────────────────────────────────────────
// STEP 5 — BASE CONVERTER
// ─────────────────────────────────────────────
func handleBase(parts []string) {
	if len(parts) == 0 {
		fmt.Println("     ✗ Error: missing command. Usage: base <dec|hex|bin> <number>")
		return
	}

	command := strings.ToLower(parts[0])

	if len(parts) < 2 {
		fmt.Printf("     ✗ Error: missing number. Usage: base %s <number>\n", command)
		return
	}

	input := parts[1]

	switch command {
	case "dec":
		n, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Printf("     ✗ Error: %q is not a valid decimal number.\n", input)
			return
		}
		if n < 0 {
			fmt.Println("     ✗ Error: negative numbers are not supported for base conversion.")
			return
		}
		fmt.Printf("     ✦ Binary : %b\n", n)
		fmt.Printf("     ✦ Hex    : %X\n", n)
		setLastResult(float64(n))
		addHistory(fmt.Sprintf("base dec %s → %b / %X", input, n, n))

	case "hex":
		if !isValidHex(input) {
			fmt.Printf("     ✗ Error: %q is not valid hex.\n", input)
			return
		}
		n, err := strconv.ParseInt(input, 16, 64)
		if err != nil {
			fmt.Printf("     ✗ Error: %q is not valid hex.\n", input)
			return
		}
		fmt.Printf("     ✦ Decimal: %d\n", n)
		setLastResult(float64(n))
		addHistory(fmt.Sprintf("base hex %s → %d", input, n))

	case "bin":
		if !isValidBinary(input) {
			fmt.Printf("     ✗ Error: %q is not valid binary.\n", input)
			return
		}
		n, err := strconv.ParseInt(input, 2, 64)
		if err != nil {
			fmt.Printf("     ✗ Error: %q is not valid binary.\n", input)
			return
		}
		fmt.Printf("     ✦ Decimal: %d\n", n)
		setLastResult(float64(n))
		addHistory(fmt.Sprintf("base bin %s → %d", input, n))

	default:
		fmt.Printf("     ✗ Error: unknown command %q. Use: dec hex bin\n", command)
	}
}

// ─────────────────────────────────────────────
// STEP 6 — STRING TRANSFORMER
// ─────────────────────────────────────────────
func handleStr(parts []string) {
	if len(parts) == 0 {
		fmt.Println("     ✗ Error: missing command. Usage: str <command> <text>")
		return
	}

	command := strings.ToLower(parts[0])

	// join remaining parts as the text
	text := strings.Join(parts[1:], " ")
	text = strings.TrimSpace(text)

	if text == "" {
		fmt.Println("     ✗ Error: no text provided.")
		return
	}

	var result string

	switch command {
	case "upper":
		result = strings.ToUpper(text)

	case "lower":
		result = strings.ToLower(text)

	case "cap":
		words := strings.Fields(text)
		for i, w := range words {
			if len(w) > 0 {
				words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
			}
		}
		result = strings.Join(words, " ")

	case "title":
		smallWords := map[string]bool{
			"a": true, "an": true, "the": true,
			"and": true, "or": true, "but": true,
			"of": true, "in": true, "on": true,
			"at": true, "to": true, "with": true,
		}
		words := strings.Fields(text)
		for i, w := range words {
			lower := strings.ToLower(w)
			if i == 0 || !smallWords[lower] {
				words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
			} else {
				words[i] = lower
			}
		}
		result = strings.Join(words, " ")

	case "snake":
		// lowercase, remove punctuation, replace spaces with underscores
		lower := strings.ToLower(text)
		var cleaned strings.Builder
		for _, ch := range lower {
			if ch == ' ' {
				cleaned.WriteRune('_')
			} else if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') || ch == '_' {
				cleaned.WriteRune(ch)
			}
		}
		result = cleaned.String()

	case "reverse":
		words := strings.Fields(text)
		for i, w := range words {
			runes := []rune(w)
			for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
				runes[l], runes[r] = runes[r], runes[l]
			}
			words[i] = string(runes)
		}
		result = strings.Join(words, " ")

	default:
		fmt.Printf("     ✗ Error: unknown command %q. Use: upper lower cap title snake reverse\n", command)
		return
	}

	fmt.Printf("     ✦ %s\n", result)
	addHistory(fmt.Sprintf("str %s %s → %s", command, text, result))
}

// ─────────────────────────────────────────────
// STEP 7 — HISTORY AND HELP
// ─────────────────────────────────────────────
func printHistory() {
	if len(history) == 0 {
		fmt.Println("     ✗ No history yet.")
		return
	}
	fmt.Println("     ✦ Session History:")
	start := len(history) - 10
	if start < 0 {
		start = 0
	}
	for i, entry := range history[start:] {
		fmt.Printf("       %d. %s\n", i+1, entry)
	}
}

func printHelp() {
	fmt.Println(`
  ════════════════════════════════════════════════
  C&C CONSOLE — HELP
  ════════════════════════════════════════════════

  CALCULATOR
    calc add <a> <b>      → addition
    calc sub <a> <b>      → subtraction
    calc mul <a> <b>      → multiplication
    calc div <a> <b>      → division
    calc mod <a> <b>      → remainder
    calc pow <a> <b>      → a to the power of b
    calc last             → show last result
    calc history          → show last 5 calc results

  BASE CONVERTER
    base dec <number>     → decimal to binary and hex
    base hex <number>     → hex to decimal
    base bin <number>     → binary to decimal

  STRING TRANSFORMER
    str upper  <text>     → ALL UPPERCASE
    str lower  <text>     → all lowercase
    str cap    <text>     → Every Word Capitalised
    str title  <text>     → Title Case With Small Words
    str snake  <text>     → convert_to_snake_case
    str reverse <text>    → esreveR hcaE droW

  GENERAL
    history               → last 10 commands
    help                  → show this menu
    exit                  → quit

  PIPE
    cmd1 | cmd2           → pass output of cmd1 into cmd2
    Example: calc add 200 55 | base dec

  NOTE: Use 'last' as a number to reference the
        most recent result across all modules.
  ════════════════════════════════════════════════`)
}

// ─────────────────────────────────────────────
// STEP 8 — PIPE
// ─────────────────────────────────────────────
func handlePipe(input string) {
	pipeParts := strings.SplitN(input, "|", 2)
	left := strings.TrimSpace(pipeParts[0])
	right := strings.TrimSpace(pipeParts[1])

	// execute left command and capture result
	leftResult, leftNum, ok := executeForPipe(left)
	if !ok {
		return
	}

	// print left result
	fmt.Printf("     ✦ Result : %s\n", leftResult)

	// replace 'last' in right command with the left result number
	if hasLastResult {
		right = strings.ReplaceAll(right, "last", fmt.Sprintf("%g", leftNum))
	}

	// execute right command normally
	rightParts := strings.Fields(right)
	if len(rightParts) == 0 {
		return
	}

	switch strings.ToLower(rightParts[0]) {
	case "calc":
		handleCalc(rightParts[1:])
	case "base":
		handleBase(rightParts[1:])
	case "str":
		handleStr(rightParts[1:])
	default:
		fmt.Printf("     ✗ Unknown command in pipe: %q\n", rightParts[0])
	}
}

func executeForPipe(input string) (string, float64, bool) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return "", 0, false
	}

	switch strings.ToLower(parts[0]) {
	case "calc":
		if len(parts) < 4 {
			fmt.Println("     ✗ Error: invalid calc command in pipe")
			return "", 0, false
		}
		a, ok1 := resolveNumber(parts[2])
		b, ok2 := resolveNumber(parts[3])
		if !ok1 || !ok2 {
			return "", 0, false
		}
		var result float64
		switch strings.ToLower(parts[1]) {
		case "add":
			result = a + b
		case "sub":
			result = a - b
		case "mul":
			result = a * b
		case "div":
			if b == 0 {
				fmt.Println("     ✗ Error: cannot divide by zero.")
				return "", 0, false
			}
			result = a / b
		case "mod":
			result = math.Mod(a, b)
		case "pow":
			result = math.Pow(a, b)
		}
		setLastResult(result)
		return fmt.Sprintf("%g", result), result, true

	case "base":
		if len(parts) < 3 {
			fmt.Println("     ✗ Error: invalid base command in pipe")
			return "", 0, false
		}
		n, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			fmt.Println("     ✗ Error: invalid number in pipe")
			return "", 0, false
		}
		setLastResult(float64(n))
		return fmt.Sprintf("%g", float64(n)), float64(n), true
	}

	return "", 0, false
}

// ─────────────────────────────────────────────
// STEP 9 — HELPERS
// ─────────────────────────────────────────────
func resolveNumber(s string) (float64, bool) {
	if strings.ToLower(s) == "last" {
		if !hasLastResult {
			fmt.Println("     ✗ Error: no previous result to use as 'last'")
			return 0, false
		}
		return lastResult, true
	}
	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("     ✗ Error: %q is not a valid number\n", s)
		return 0, false
	}
	return val, true
}

func setLastResult(n float64) {
	lastResult = n
	hasLastResult = true
}

func addHistory(entry string) {
	history = append(history, entry)
}

func isValidHex(s string) bool {
	for _, ch := range strings.ToUpper(s) {
		if !((ch >= '0' && ch <= '9') || (ch >= 'A' && ch <= 'F')) {
			return false
		}
	}
	return len(s) > 0
}

func isValidBinary(s string) bool {
	for _, ch := range s {
		if ch != '0' && ch != '1' {
			return false
		}
	}
	return len(s) > 0
}
