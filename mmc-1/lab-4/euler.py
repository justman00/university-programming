import numpy as np 

def get_cube_root(num):
    return num ** (1. / 3)


def func(x, y):
    return 0.1 * (get_cube_root(y) + np.log(x + y) -1)

# Function for euler formula


def euler(x0, y, h, x):
    iterator = 0
    # Iterating till the point at which we
    # need approximation
    while x0 < x:
        iterator += 1
        y = y + h * func(x0, y)
        x0 = x0 + h

    # Printing approximation
    print("Approximate solution at x = ", x, " is ", "%.6f" % y)
    print("Codul a facut urmatorul numar de iteratii: ", iterator)


# Driver Code
# Initial Values
x0 = -1
y0 = 2
h = 0.05

# Value of x at which we need approximation
x = 0

euler(x0, y0, h, x)
