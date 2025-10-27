Alright, little buddy 🧸, let’s imagine this program is like a magic marker that helps us change letters in a sentence — like coloring over one letter with another!

---

## 🎨 What Does It Do?

This program takes **three things** (we call them “arguments”):

1. A **sentence** — like `"hella there"`
2. A **letter you want to change** — like `"a"`
3. A **new letter to use instead** — like `"o"`

Then it goes through the sentence and **swaps** every `"a"` with `"o"` — like turning all red apples into green apples 🍎→🍏

---

## 🧪 Example

```bash
go run . "hella there" "a" "o"
```

- The sentence is `"hella there"`
- We want to change `"a"` to `"o"`
- So `"hella"` becomes `"hello"`
- Final result: `"hello there"` 🎉

---

## 🚫 What If We Give Too Many or Too Few Things?

If we don’t give **exactly 3 things**, the program says **nothing**. It just stays quiet 🤫

Example:

```bash
go run . "something" "a" "o" "b" "c"
```

Too many things! So the program doesn’t do anything.

---

## 🧼 What If the Letter Isn’t in the Sentence?

If the letter we want to change isn’t there, the program just **prints the sentence** as it is.

Example:

```bash
go run . "abcd" "z" "l"
```

- `"z"` is not in `"abcd"`
- So it just prints: `"abcd"`

---

## 🧑‍💻 What Would the Code Look Like?

Imagine the computer is doing this:

1. “Did you give me 3 things?”
2. “Is the letter you want to change inside the sentence?”
3. “If yes, I’ll swap it!”
4. “If no, I’ll just show the sentence.”

---

So it’s like a magic eraser that replaces one letter with another — but only if you follow the rules! 🪄✨

Want to try changing your name with it next?
