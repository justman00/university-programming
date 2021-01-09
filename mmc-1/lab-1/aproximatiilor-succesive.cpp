#include <iostream>
#include <math.h>

using namespace std;

double func1(double x) {
    return 2 * pow(x - 1, 2) - pow(M_E, x);
}

double func2(double x) {
    return pow(x, 3) + 20 * x - 41;
}

int algoritm(double func(double), double *x)
{
    double erroare = 0.000001;
    int maxIteratii = 100;
    double dx, f;
    int it;
    dx = -func(*x);
    for (it = 1; it <= maxIteratii; it++)
    {
        f = func(*x);
        if (fabs(f) > fabs(dx))
            goto divergent;
        dx = -f;
        *x += dx;
        if (fabs(dx) <= erroare * fabs(*x))
            return 0;
        cout << "\n\nF(" << *x << ")=" << f << " la iteratia " << it;
    }
    cout << "nr.maxim de iteratii depasit\n";
    return 1;
divergent:
    cout << "proces divergent\n";
    return 2;
}
int main()
{
    double x;
    cout << "Introduceti valoare pentru x de pe interval" << endl;
    cout << "x0 = ";
    cin >> x;
    algoritm(func1, &x);
    return 0;
}
