if __name__ == "__main__":
    n = int(input())
    arr = [int(x) for x in input().split()]
    flag = any(x for x in arr if x == 1)
    if flag:
        print("HARD")
    else:
        print("EASY")
