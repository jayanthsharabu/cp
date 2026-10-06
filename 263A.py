def dis(mat):
    x,y = 0,0
    for i in range(5):
        for j in range(5):
            if mat[i][j] == "1":
                x,y = i,j


    #(1,4) -> (2,2)
    return abs(x-2) + abs(y-2)



if __name__ == "__main__":
    mat = []
    for i in range(5):
        arr = input().split(" ")
        mat.append(arr)
    print(dis(mat))
