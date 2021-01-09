#include <iostream>
#include <math.h>

#define erroare 0.000001
#define maxIteratii 100

using namespace std;

double func1(double x) {
    return 2 * pow(x - 1, 2) - pow(M_E, x);
}

double func2(double x) {
    return pow(x, 3) + 20 * x - 41;
}

double func(double x)
{
    return func2(x);
}

int main()
{
    int i = 0;
    double x, x0, x1, a, b, y;
    std::cout << "a=";
    cin >> a;
    std::cout << "b=";
    cin >> b;
    x0 = a;
    x1 = b;
    x = x0;
    y = func(x);
    if (func(x0) * func(x1) < 0)
    {
        while ((i <= maxIteratii) && ((y < -erroare) || (y > erroare)))
        {
            x = x0 - func(x0) * (x1 - x0) / (func(x1) - func(x0));
            y = func(x);
            if (func(x0) * y < 0)
            {
                x1 = x;
            }
            else
            {
                x0 = x;
            }
            std::cout << "\n\nf(" << x << ")=" << func(x) << " la iteratia " << (int)i;
            i++;
        }
        if (i > maxIteratii)
        {
            cout << "problema a depasit numarul maxim de iteratii";
        }
    }
    else
    {
        cout << "interval invalid";
    }

    cout << "\n" << endl;
}
