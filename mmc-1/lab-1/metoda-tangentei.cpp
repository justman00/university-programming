#include <iostream>
#include <math.h>

#define erroare 0.000001
#define maxIteratii 100

using namespace std;

double func1(double x)
{
    return 2 * pow(x - 1, 2) - pow(M_E, x);
}

double func2(double x)
{
    return pow(x, 3) + 20 * x - 41;
}

double func_derivata_1(double x)
{
    return 4 * x - 4 - pow(M_E, x);
}

double func_derivata_2(double x)
{
    return 3 * pow(x, 2) + 20;
}

double metoda_tangentei(double a, double b, double (*f)(double), double (*f1)(double))
{
    int i = 0;
    double x, y1, y;
    x = a;
    y = f(x);
    y1 = f1(x);
    while ((i <= maxIteratii) && ((y < -erroare) || (y > erroare)))
    {
        x = x - y / y1;
        y = f(x);
        y1 = f1(x);
        cout << "\n\nf(" << x << ")=" << y << " la iteratia " << (int)i;
        i++;
    }
    if (i > maxIteratii)
    {
        cout << "problema nu se poate rezolva in nr.maxim de iteratii";
        return 0;
    }
    else
        return x;
}
int main()
{
    double x, y1, y, a, b;
    std::cout << "a=";
    cin >> a;
    std::cout << "b=";
    cin >> b;
    x = metoda_tangentei(a, b, func2, func_derivata_2);
    if (x != 0)
        std::cout << "\nSolutia = " << x << endl;
}