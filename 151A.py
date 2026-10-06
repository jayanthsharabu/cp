if __name__ == "__main__":
    n, k, l, c, d, p, nl, np = [int(x) for x in input().split()]
    liq = k*l
    lem = c*d
    tml = n*nl
    tg = n*np
    res = 0
    while liq >= tml and p >= tg and lem >= n:
        liq -= tml
        p -= tg
        lem -= n
        res += 1

    print(res)
