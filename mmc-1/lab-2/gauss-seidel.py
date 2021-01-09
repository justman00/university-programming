def seidel(a, x ,b): 
    #Finding length of a(3)        
    n = len(a)                    
    # for loop for 3 times as to calculate x, y , z
    # in cazul nostru de 4 ori deoarece matricea e mai mare
    for j in range(0, n):         
        # temp variable d to store b[j] 
        d = b[j]                   
          
        # to calculate respective xi, yi, zi 
        for i in range(0, n):      
            if(j != i): 
                d-=a[j][i] * x[i] 
        # updating the value of our solution         
        x[j] = d / a[j][j] 
    # returning our updated solution            
    return x     
   
X = [0, 0, 0, 0]                         
A = [[8.7, -1.2, 0.8, 0.7],
     [-1.2, 9.6, -1.2, 0.8],
     [0.8, -1.2, 8.8, 0.9],
     [0.7, 0.8, 0.9, 11.3]]
B = [-2.7, 8.9, 7.2, 6.4]
print(X) 
  
#loop run for m times depending on m the error value 
for i in range(0, 25):             
    X = seidel(A, X, B) 
    #print each time the updated solution 
    print(X) 