from math import sqrt
from pprint import pprint


def cholesky(A):
    n = len(A)

    # Creem o matrice de zerouri
    L = [[0.0] * n for i in range(n)]

    for i in range(n):
        for k in range(i+1):
            tmp_sum = sum(L[i][j] * L[k][j] for j in range(k))

            if (i == k):  # Elemente diagonale
                L[i][k] = sqrt(A[i][i] - tmp_sum)
            else:
                L[i][k] = (1.0 / L[k][k] * (A[i][k] - tmp_sum))
    return L


A = [[8.7, -1.2, 0.8, 0.7],
     [-1.2, 9.6, -1.2, 0.8],
     [0.8, -1.2, 8.8, 0.9],
     [0.7, 0.8, 0.9, 11.3]]
B = [-2.7, 8.9, 7.2, 6.4]
S = cholesky(A)

print("A:")
pprint(A)

print("S:")
pprint(S)

y0 = B[0] / S[0][0]
y1 = (B[1] - S[1][0] * y0) / S[1][1]
y2 = (B[2] - S[2][0] * y0 - S[2][1] * y1) / S[2][2]
y3 = (B[3] - S[2][0] * y0 - S[2][1] * y1 - S[3][1] * y2) / S[3][3]

print("\nSolutia(y):\n y0={:>7.3f}\n y1={:>7.3f}\n y2={:>7.3f}\n y3={:>7.3f}\n".format(
    y0, y1, y2, y3))

x4 = y3 / S[3][3]
x3 = (y2 - S[3][2] * x4) / S[2][2]
x2 = (y1 - S[3][1] * x4 - S[2][1] * x3) / S[1][1]
x1 = (y0 - S[3][0] * x4 - S[2][0] * x3 - S[1][0] * x2) / S[0][0]

print("\nSolutia(x):\n y0={:>7.3f}\n y1={:>7.3f}\n y2={:>7.3f}\n y3={:>7.3f}\n".format(
    x1, x2, x3, x4))