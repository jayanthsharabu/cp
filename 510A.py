if __name__ == "__main__":
    n,m = [int(x) for x in input().split()]
    flag = 1
    for i in range(n):
        if flag and i % 2 != 0:
            print("."*(m-1) + "#")
            flag = 0
        elif flag == 0 and i%2 != 0:
            print("#" + "."*(m-1))
            flag = 1
        else:
            print("#"*m)
