Alright little buddy 🧸, let’s imagine we’re playing with words like they’re toy trains 🚂. Each word is made of train cars (letters), and we want to change how the train looks!

---

## 🐫 What is camelCase?

CamelCase is like a train where each new word starts with a **big letter** (capital letter), but the first word might start with a **small letter**.

Examples:
- `"helloWorld"` → small camel 🐪
- `"HelloWorld"` → big camel 🐫

---

## 🐍 What is snake_case?

Snake_case is like a train where we put **underscores** between words instead of big letters.

Examples:
- `"hello_world"`
- `"hello_big_snake"`

---

## 🎯 What does the code do?

It changes camelCase words into snake_case words. Like turning a camel 🐫 into a snake 🐍!

---

## 🧩 Step-by-Step Like a Game

Let’s say we have `"camelToSnakeCase"`:

1. **Look at each letter** in the word.
2. If the letter is a **big letter** (capital), that means a **new word** is starting!
3. So we put an **underscore `_`** before it.
4. Then we keep going until the end.
5. If the word is **not camelCase**, we just return it as it is.

---

## 🧑‍💻 The Code (explained simply)

```go
func CamelToSnakeCase(s string) string {
	if s == "" {
		return "" // If the word is empty, return empty
	}

	// Check if it's a valid camelCase word
	if !isCamelCase(s) {
		return s // If it's not camelCase, return it as it is
	}

	result := "" // Start with an empty result

	for i, ch := range s {
		if ch >= 'A' && ch <= 'Z' { // If the letter is big
			if i != 0 { // Don't add underscore at the start
				result += "_"
			}
		}
		result += string(ch) // Add the letter to the result
	}

	return result
}
```

---

## 🕵️ Helper: isCamelCase

This part checks if the word is really camelCase:

- No numbers
- No punctuation
- No ending with a big letter
- No two big letters together

If it breaks any rule, we say “Nope!” and return the word unchanged.

---

## 🧪 Examples

| Input                  | Output               | Why? |
|------------------------|----------------------|------|
| `"HelloWorld"`         | `"Hello_World"`      | Big letters → new words |
| `"helloWorld"`         | `"hello_World"`      | Small start, big next |
| `"camelCase"`          | `"camel_Case"`       | One big letter |
| `"CAMELtoSnackCASE"`   | `"CAMELtoSnackCASE"` | Not camelCase (bad ending) |
| `"camelToSnakeCase"`   | `"camel_To_Snake_Case"` | All good! |
| `"hey2"`               | `"hey2"`             | Has a number → not camelCase |

---

So it’s like playing with word trains and putting little snake tails between the cars 🐍🚋🚋🚋

Want to build your own word snake next?
