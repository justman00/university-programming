import numpy as np


def linearsolver(A, b):
    n = len(A)
    M = A

    i = 0
    for x in M:
        x.append(b[i])
        i += 1

    for k in range(n):
        for i in range(k, n):
            if abs(M[i][k]) > abs(M[k][k]):
                M[k], M[i] = M[i], M[k]
            else:
                pass

        for j in range(k+1, n):
            q = float(M[j][k]) / M[k][k]
            for m in range(k, n+1):
                M[j][m] -= q * M[k][m]

    x = [0 for i in range(n)]

    x[n-1] = float(M[n-1][n])/M[n-1][n-1]
    for i in range(n-1, -1, -1):
        z = 0
        for j in range(i+1, n):
            z = z + float(M[i][j])*x[j]
        x[i] = float(M[i][n] - z)/M[i][i]
    print(x)


A = [[8.7, -1.2, 0.8, 0.7],
     [-1.2, 9.6, -1.2, 0.8],
     [0.8, -1.2, 8.8, 0.9],
     [0.7, 0.8, 0.9, 11.3]]
B = [-2.7, 8.9, 7.2, 6.4]

linearsolver(A, B)
