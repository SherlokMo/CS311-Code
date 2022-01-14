## Solution

General solution for task 1 in slide 2

## Solution number 1

```
    Algorithm solution(n)
        inputs: n intergers of numbers to read

        declear arr[n] // More space
        for i <- 0 to n-1 do
            arr[i] <- Read("Enter your number")
        
        max <- arr[0], min <- arr[0], avg <- arr[0]

        for i <- 1 to n-1 do
            avg <- avg + arr[i]
            if arr[i] > max then
                max <- arr[i]
            if arr[i] < min then
                min <- arr[i]            

        avg <- avg/n
        Output("Average is", avg)
        Output("Min is", min)
        Output("Max is", max)

    Space O(n)
    Time  O(n)
```


## Solution number 2
```
    Algorithm solution(n)
        inputs: n intergers of numbers to read

        max <- (-inf), min <- (inf), avg = 0

        for i <- 0 to n-1 do
            userInput <- Read("Enter the number")
            avg <- avg + userInput
            if userInput > max then
                max <- userInput
            if userInput < min then
                min <- userInput
                avg <- avg/n
        Output("Average is", avg)
        Output("Min is", min)
        Output("Max is", max)

    Space O(1)
    Time  O(n)
            

        


```