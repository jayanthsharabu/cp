def helper(num):
    dum = num
    fin = []
    q = 1
    while dum > 0:
        dig = dum % 10
        if dig != 0:
            fin.append(dig * q)
        dum //= 10
        q *= 10

    return fin


if __name__ == "__main__":
    n = int(input())
    inputs = []
    for i in range(n):
        val = int(input())
        inputs.append(val)
    for ele in inputs:
        res = helper(ele)
        print(len(res))
        print(*res)
