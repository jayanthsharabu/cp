if __name__ == "__main__":
    arr = [int(x) for x in input().split()]
    arr.sort()
    mean = arr[1]
    dis =  abs(mean-arr[0]) + abs(mean-arr[1]) + abs(mean-arr[2])
    print(dis)
