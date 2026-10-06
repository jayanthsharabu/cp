if __name__ == "__main__":
    n = int(input())
    arr = [int(x) for x in input().split()]
    res = [0]*n
    for i in range(n):
        res[arr[i]-1] = i+1
    print(*res)
