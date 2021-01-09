import numpy as np

max_iter = 1000

A = np.array([[8.7, -1.2, 0.8, 0.7],
              [-1.2, 9.6, -1.2, 0.8],
              [0.8, -1.2, 8.8, 0.9],
              [0.7, 0.8, 0.9, 11.3]])
B = np.array([-2.7, 8.9, 7.2, 6.4])

x = np.zeros_like(B)
for iteration in range(max_iter):
    x_new = np.zeros_like(x)
    for i in range(A.shape[0]):
        s1 = np.dot(A[i, :i], x[:i])
        s2 = np.dot(A[i, i + 1:], x[i + 1:])
        x_new[i] = (B[i] - s1 - s2) / A[i, i]
    
    if np.allclose(x, x_new, atol=0.01):
        break
    x = x_new
    print("Solutia curenta", x)

print("\nSolutia:")
print(x)
error = np.dot(A, x) - B
print("\nErori:")
print(error)
print("\nIteratii: ", (iteration))
