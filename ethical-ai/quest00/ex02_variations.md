# 1. I modified my palindrome function to:
```py
import re

def palindrome_break_position(s):
    s = re.sub(r'[^a-z0-9]', '', s.lower())
    
    for i in range(len(s) // 2):
        if s[i] != s[-(i + 1)]:
            return i   
    
    return -1   
s = input("enter word: ")

position = palindrome_break_position(s)

if position == -1:
    print("this is a palindrome")
else:
    print("this is not a palindrome")
    print("breaks at position:", position)
```

## 2. After my  first attempt, asked AI:
```py
import re
def palindrome_break_position(s):
    s = re.sub(r'[^a-z0-9]', '', s.lower())
    
    left = 0
    right = len(s) - 1

    while left < right:
        if s[left] != s[right]:
            return left   
        left += 1
        right -= 1

    return -1   
s = input("enter word: ")

position = palindrome_break_position(s)

if position == -1:
    print("this is a palindrome")
else:
    print("this is not a palindrome")
    print("breaks at position:", position)
```
# 3. Reflection on what AI added that i didn't consider initially.
> I didn't consider the two pointer while loop which compares character from both side of the string at thesame time, the left starts at the begining and the right starts at the end.
   

