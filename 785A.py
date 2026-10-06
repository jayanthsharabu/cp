if __name__ == "__main__":
    n = int(input())
    res = 0
    for i in range(n):
        shape = input()
        if shape == "Tetrahedron":
            res += 4
        elif shape == "Cube":
            res += 6
        elif shape == "Octahedron":
            res += 8
        elif shape == "Dodecahedron":
            res += 12
        else:
            res += 20

    print(res)
