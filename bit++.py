
if __name__ == "__main__":
    n = int(input())
    x = 0
    for i in range(n):
        item = input()
        if "+" in item:
            x += 1
        else:
            x -= 1
    print(x)
