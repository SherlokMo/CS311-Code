##PseudoCode


Devide & conquer: (top-down approach)
Algorithm factorial(n): 
    if n <= 1:
        return 1
    else:
        k <- n - 1
        result <- n * factorial(k)
        return result

Dyanmic programming: (bottom-top approach)
Algorithm factorial(n):
    if n <= 1:
        return 1
    else:
        fact <- 1
        for i <- 2 to n:
            fact <- fact * i

        return fact