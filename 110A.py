if __name__ == "__main__":
    num = input()
    flag = 0
    for n in range(len(num)):
        if num[n] == "4" or num[n] == "7":
            flag += 1

    if flag == 4 or flag == 7:
       var = "YES"
    else:
       var = "NO"

    print(var)
