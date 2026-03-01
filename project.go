🔎 What Is This Project?

You are building a text editing tool in Go.

The program:

Takes a text file as input

Modifies the text based on specific rules

Saves the modified result into another file

Think of it as a mini auto-correct + formatter program.

🧱 Step 1: Understand the Goal

Your tool should automatically:

Convert numbers written in special formats

Change letter casing (upper/lower/capitalize)

Fix punctuation spacing

Fix quotation marks

Correct “a” → “an” when needed

💻 Step 2: How the Program Should Work

Your program will be run like this:

go run . input.txt output.txt

input.txt → contains messy or tagged text

output.txt → should contain the corrected text

🧮 Step 3: Number Conversions
1️⃣ Hexadecimal → Decimal

If text says:

1E (hex)

You:

Convert 1E from hex to decimal

Replace the whole thing

Result:

30
2️⃣ Binary → Decimal

If text says:

10 (bin)

Convert from binary to decimal.

Result:

2
🔤 Step 4: Text Case Changes

These modify words before them.

(up)

Make previous word uppercase:

go (up)
→ GO
(low)

Make previous word lowercase:

HELLO (low)
→ hello
(cap)

Capitalize previous word:

bridge (cap)
→ Bridge
With a number:

If written like:

(up, 2)

It affects the previous 2 words.

Example:

so exciting (up, 2)
→ SO EXCITING
✏️ Step 5: Fix Punctuation Spacing

Your program must:

✅ Remove space before punctuation

Wrong:

hello , world !

Correct:

hello, world!
✅ Add space after punctuation (if needed)

Wrong:

hello,world

Correct:

hello, world
✅ Keep grouped punctuation together

Wrong:

thinking ...

Correct:

thinking...
🗣 Step 6: Fix Single Quotes

Quotes ' always come in pairs.

Wrong:

' awesome '

Correct:

'awesome'

If multiple words inside:

' I am great '
→ 'I am great'

No spaces just inside the quotes.

🅰️ Step 7: Fix "a" → "an"

If the next word starts with:

a vowel (a, e, i, o, u)

or h

Change:

a apple
→ an apple
A house
→ An house

(You must also respect uppercase/lowercase.)

🧪 Step 8: Testing

You are encouraged to:

Write unit tests

Create your own example files

Test edge cases

Also:

Other students will audit your project

You will audit other students

So your code must be:

Clean

Structured

Easy to understand

🏗 What You Are Really Building

You are building:

A command-line text processor that reads text, applies smart formatting rules, and writes a corrected version.

It’s like a simplified Grammarly + formatter written in Go.

🎯 Final Summary

Your program must:

Read input file

Apply all transformation rules

Output clean, corrected text

Follow Go best practices

Be testable and well-structured

If you want, I can now break this into a clear coding plan (how to structure your Go project step-by-step).

break into a clear coding plan(how to structure go project step by step)

Perfect — let’s turn this into a clear, practical Go coding plan you can follow step by step.

🏗 Step-by-Step Go Project Structure Plan
✅ STEP 1 — Create Basic Project Structure

Create a new folder:

texttool/

Inside it:

texttool/
│── go.mod
│── main.go
│── internal/
│     ├── processor.go
│     ├── numbers.go
│     ├── case.go
│     ├── punctuation.go
│     ├── quotes.go
│     └── article.go
│── testdata/
│     ├── input.txt
│     └── expected.txt

Why?

main.go → handles arguments + file reading/writing

internal/ → contains your logic

separate files = clean, organized code

✅ STEP 2 — Initialize Go Module

Inside project folder:

go mod init texttool
✅ STEP 3 — Write the Entry Point (main.go)

Responsibilities of main.go:

Check arguments count

Read input file

Send content to processor

Write result to output file

Pseudo-structure:

func main() {
    // 1. Validate arguments
    // 2. Read input file
    // 3. Process text
    // 4. Write output file
}

Keep main.go small. No logic inside it.

✅ STEP 4 — Create a Central Processor

In processor.go:

func ProcessText(text string) string {
    text = HandleHex(text)
    text = HandleBin(text)
    text = HandleCase(text)
    text = FixPunctuation(text)
    text = FixQuotes(text)
    text = FixArticles(text)
    return text
}

This makes everything modular and easy to debug.

✅ STEP 5 — Implement Features One by One

Do NOT code everything at once.

Follow this order:

🔢 1. Handle Hex and Binary

File: numbers.go

Functions:

func HandleHex(text string) string
func HandleBin(text string) string

Logic:

Find patterns like WORD (hex)

Convert WORD

Replace whole match

Tip:
Use:

strings.Fields() OR

regexp (cleaner solution)

🔤 2. Handle Case Conversions

File: case.go

Functions:

func HandleCase(text string) string

Inside it:

Detect (up)

Detect (low)

Detect (cap)

Detect (up, 2) style patterns

Strategy:

Split text into words

Loop through slice

When you see (up) → modify previous word

Remove the tag from slice

✏️ 3. Fix Punctuation

File: punctuation.go

Functions:

func FixPunctuation(text string) string

Tasks:

Remove space before punctuation

Add space after punctuation

Keep grouped punctuation like ..., !!, ?!

Best approach:

Use regex replacements

Or iterate character by character

🗣 4. Fix Quotes

File: quotes.go

Function:

func FixQuotes(text string) string

Logic:

Find '

They always come in pairs

Remove space after first '

Remove space before second '

You can:

Split by '

Fix odd indexes

Rejoin

🅰️ 5. Fix Articles (a → an)

File: article.go

Function:

func FixArticles(text string) string

Logic:

Split into words

If word == "a" or "A"

Check next word

If starts with vowel or h → change to "an" or "An"

✅ STEP 6 — Testing Strategy

Create test files inside testdata/.

Example:

input.txt

1E (hex) files were added

expected.txt

30 files were added

You can write Go test files:

processor_test.go

Example:

func TestHex(t *testing.T) {
    input := "1E (hex)"
    expected := "30"
    result := HandleHex(input)

    if result != expected {
        t.Errorf("Expected %s but got %s", expected, result)
    }
}

Run tests:

go test ./...
✅ STEP 7 — Order of Execution Matters

Correct order in ProcessText:

Numbers (hex/bin)

Case transformations

Quotes

Punctuation

Articles (last)

Why?

Because some transformations affect word positions and spacing.

✅ STEP 8 — Handle Edge Cases

Test things like:

Multiple (up) in a row

(up, 10) bigger than word count

Multiple punctuation marks

Uppercase "A"

Quotes with multiple words

✅ STEP 9 — Code Quality Requirements

Follow good Go practices:

Small functions

Clear names

No duplicated code

Proper error handling

No global variables

🎯 Final Architecture Overview

Flow of execution:

main.go
   ↓
Read file
   ↓
ProcessText()
   ↓
   ├── HandleHex()
   ├── HandleBin()
   ├── HandleCase()
   ├── FixQuotes()
   ├── FixPunctuation()
   └── FixArticles()
   ↓
Write output
🚀 What This Project Really Tests

String manipulation

Regex usage

File handling in Go

Clean architecture

Edge case handling

Writing testable code

If you want, I can now give you:

A clean starter template code

OR a recommended algorithm for parsing words safely

OR common mistakes students make on this project

a clean starter template code

Perfect — here is a clean, minimal, well-structured starter template for your project.

This gives you:

Proper project structure

Clean separation of concerns

Function stubs ready to implement

Error handling

Test-ready architecture

You just need to fill in the logic.

📁 Project Structure
texttool/
│── go.mod
│── main.go
│── internal/
│     ├── processor.go
│     ├── numbers.go
│     ├── case.go
│     ├── punctuation.go
│     ├── quotes.go
│     └── article.go
✅ go.mod

Initialize with:

go mod init texttool
✅ main.go
package main

import (
	"fmt"
	"os"
	"texttool/internal"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	result := internal.ProcessText(string(data))

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
		return
	}
}
✅ internal/processor.go
package internal

func ProcessText(text string) string {
	text = HandleHex(text)
	text = HandleBin(text)
	text = HandleCase(text)
	text = FixQuotes(text)
	text = FixPunctuation(text)
	text = FixArticles(text)
	return text
}
✅ internal/numbers.go
package internal

func HandleHex(text string) string {
	// TODO: implement hex conversion
	return text
}

func HandleBin(text string) string {
	// TODO: implement binary conversion
	return text
}
✅ internal/case.go
package internal

func HandleCase(text string) string {
	// TODO: implement (up), (low), (cap), (up, n)
	return text
}
✅ internal/punctuation.go
package internal

func FixPunctuation(text string) string {
	// TODO: fix punctuation spacing
	return text
}
✅ internal/quotes.go
package internal

func FixQuotes(text string) string {
	// TODO: fix single quote spacing
	return text
}
✅ internal/article.go
package internal

func FixArticles(text string) string {
	// TODO: change "a" to "an" when needed
	return text
}
🎯 Why This Is Clean

main.go only handles I/O

All logic is in internal/

Each feature is isolated

Easy to test each function independently

Easy for auditors to read

Easy to debug

🧪 Example Test Skeleton (Optional but Recommended)

Create:

internal/numbers_test.go
package internal

import "testing"

func TestHandleHex(t *testing.T) {
	input := "1E (hex)"
	expected := "30"

	result := HandleHex(input)

	if result != expected {
		t.Errorf("expected %s but got %s", expected, result)
	}
}

Run tests:

go test ./...
🚀 What You Should Do Next

Implement features in this order:

HandleHex

HandleBin

HandleCase

FixQuotes

FixPunctuation

FixArticles

Test each one before moving to the next.

If you'd like, I can now give you:

A clean implementation example for HandleHex

Or a safe algorithm for parsing words without breaking punctuation

Or a full working reference solution (for understanding)
