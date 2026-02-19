# Step1 - Do it yourself

# pseudocode
1.write a function "is_palindrome" that takes a string as an input and return a string equal to a reversed string
2.create a variable for user input
3.if the function is true print "string is a palindrome"
4.else print "string is not a palindrome"

# implementation of solution

```py
def is_palindrome(s): #defining the function is_palindrome
    return s == s[::-1]  #the function returns string equal the reverse of the string

s = input("enter word;")  #it provides the user input

if is_palindrome(s):
    print("this is a palindrome")  #this shows that the function is true
else:
    print("this is not a palindrome") #this shows that the function is false
```

# Step2 - Use AI to learn

```py
def is_palindrome(s):
    s = re.sub(r'[^a-z0-9]', '', s.lower())
    return s == s[::-1]

s = input("enter word;")

if is_palindrome(s):
    print("this is a palindrome")
else:
    print("this is not a palindrome")
```
# Step3
# Reflection
* What did you learn from solving it before asking AI?

I had to think logically to get the algorithim myself, i now know how to solve a problem of reverse string and handle a case senstively.

* How is your understanding different now?

By solving it myself i now understand why each step is necessary rather than just getting the solution , i can confidently attempts a question just by being able to think logically and break it down step by step.

* Could you now write similar functions (e.g., reverse a string) without help?

Yes, i can confidently write functions similar to reverse string without help due to my deep understanding of breaking down problems , thinking logically and having a right mindset towards every problem using thesame pattern of approach.




