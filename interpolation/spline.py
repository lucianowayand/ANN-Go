import numpy as np
import math

def spline(x,y):
    n = len(x)
    a = {k: v for k, v in enumerate(y)}
    h = {k: x[k+1] - x[k] for k in range(n - 1)}
    A = [[1] + [0] * (n-1)]
    for i in range(1, n-1):
        row = [0] * n
        row[i-1] = h[i-1]
        row[i] = 2*(h[i-1] + h[i])
        row[i+1] = h[i]
        A.append(row)
    A.append([0] * (n-1) + [1])

    B = [0]
    for k in range(1,n-1):
        row = 3 * (a[k+1] - a[k]) / h[k] - 3 * (a[k] - a[k-1]) / h[k-1]
        B.append(row)
    B.append(0)
    
    c = dict(zip(range(n), np.linalg.solve(A,B)))
    
    b = {}
    d = {}
    for k in range(n-1):
        b[k] = (1/h[k]) * (a[k+1] - a[k]) - (h[k]/3) * (2*c[k]+c[k+1])
        d[k] = (c[k+1] - c[k])/(3*h[k])
    
    s = {}
    for k in range(n-1):
        #print(a[k], b[k], c[k], d[k])
        eq = f'{a[k]}{b[k]:+} * (x{-x[k]:+}){c[k]:+}*(x{-x[k]:+})**2{d[k]:+}*(x{-x[k]:+})**3 '
        s[k] = {'eq': eq, 'domain': [x[k], x[k+1]]}

    return s

def f(x):
    return math.cos(x) ** 3 + 2*math.cos(x)**2 + 1

x =  [-1.468, -0.18, 0.8, 1.987, 3.08, 4.521, 5.555, 6.77]
y =  [1.406, 1.119, 1.81, 1.279, 0.394, -0.018, 0.409, 1.522]
valores =  [1.719, 1.935, 5.178]


eqs = spline(x,y)
for i in eqs:
    for va in valores:
        if va>eqs[i]['domain'][0] and eqs[i]['domain'][1]>va:
            x = va
            print(eval(eqs[i]['eq']))
        
#print(eqs)