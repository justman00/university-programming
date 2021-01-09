import numpy as np 

def get_cube_root(num):
    return num ** (1. / 3)

def func(x, y):
    return 0.1 * (get_cube_root(y) + np.log(x + y) -1)
  
# Finds value of y for a given x using step size h 
# and initial value y0 at x0. 
def rungeKutta(x0, y0, x, h): 
    # Count number of iterations using step size or 
    # step height h 
    n = (int)((x - x0)/h)  
    # Iterate for number of iterations 
    y = y0 
    for i in range(1, n + 1): 
        "Apply Runge Kutta Formulas to find next value of y"
        k1 = h * func(x0, y) 
        k2 = h * func(x0 + 0.5 * h, y + 0.5 * k1) 
        k3 = h * func(x0 + 0.5 * h, y + 0.5 * k2) 
        k4 = h * func(x0 + h, y + k3) 
  
        # Update next value of y 
        y = y + (1.0 / 6.0)*(k1 + 2 * k2 + 2 * k3 + k4) 
  
        # Update next value of x 
        x0 = x0 + h 
    print("Codul a facut urmatorul numar de iteratii: ", i)
    return y 
  
# Driver method 
x0 = -1
y = 2
x = 0
h = 0.05
print("The value of y at x is:", rungeKutta(x0, y, x, h))