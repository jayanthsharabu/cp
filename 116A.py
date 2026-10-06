def max_req(arr) -> int:
    res = 0
    rem = 0
    for ar in arr:
        x,y = ar
        rem += (y - x)
        res = max(res, rem)
    return res

if __name__ == "__main__":
    n = int(input())
    arr = []
    for _ in range(n):
        t = list(map(int,input().split()))
        arr.append(t)
    print(max_req(arr))
