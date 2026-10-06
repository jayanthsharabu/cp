if __name__ == "__main__":
    n = int(input())
    dum = n
    q = 10
    m = 1
    rem = 0
    while dum:
        dig = dum % 10
        dum //= q
        if dum == 0 and dig == 9:
            rem += dig * m
        elif dig >= 5:
            dig = 9 - dig
            rem += dig*m
        else:
            rem += dig*m

        m *= 10
    print(rem)
