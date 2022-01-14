## Palindrome Solution

# What is palindrom?

"ROTOR" is Palindrome
"انا" is palindrome
"baba" not palindrome 

# Iterative pseudo code
```
    Algorithm isPalindrome(str, n)
        Input: str the string we want to check, n size of the string
        Output: palindrome boolean check

        low <- 0, high <- n
        while low < high do
            if str[low] != str[high] then
                return false
            low <- low + 1
            high <- high - 1

        return true

    Complexity:
        Space O(1)
        Time  O(N)
```


# Recusrive Pseudo code

```
    Algorithm isPalindrome(str, low, high) 
        Input: str is the string we want to check, low is the lower boundry, high is higher boundry
        Output: palindrome boolean check

        if low = high then
            return true
        
        if str[low] <> str[high] then
            return false

        low <- low + 1
        high <- high - 1
        return isPalindrome(str, low, high)

    Complexity:
        Space: O(N)
        Time:  O(N)
```