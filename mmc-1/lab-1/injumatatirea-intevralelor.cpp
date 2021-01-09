#include <iostream>
#include <math.h>

using namespace std;
using std::endl;

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
    double eroare = 0.01;
    double x, x0, x1, a, b, y;
    cout << "Introduceti intervalul [a, b]" << endl;
    cout << "a=";
    cin >> a;
    cout << "b=";
    cin >> b;
    x0 = a;
    x1 = b;
    x = x0;
    y = func(x);
    if (func(x0) * func(x1) < 0)
    {
        while ((y < -eroare) || (y > eroare))
        {
            x = (x0 + x1) / 2;
            y = func(x);
            if (func(x0) * y < 0)
            {
                x1 = x;
            }
            else
            {
                x0 = x;
            }
            cout << "\n\nF(" << x << ")=" << func(x);
        }
    }
    else
    {
        cout << "Erroare interval";
    }

    cout << "\n" << endl;
}
