
if __name__ == "__main__":
    n = int(input())
    arr = [int(x) for x in input().split()]
    mxi,mni = 0,0
    mxv, mnv = min(arr),max(arr)
    for i in range(n):
        ar = arr[i]
        if ar <= mnv:
            mni = i
            mnv = ar
        if ar > mxv:
            mxi = i
            mxv = ar
    res = mxi + (n-1 - (mni))
    if mxi > mni:
        res -= 1
    print(res)
