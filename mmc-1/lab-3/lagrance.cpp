#include <iostream>

using namespace std;
double lagrange(int N, float X[], float Y[], float x1)
{
    float L = 0, b = 1;
    int i, j;

    for (i = 0; i < N; i++)
    {
        b = 1;

        for (j = 0; j < N; j++)
            if (j != i)
                b *= (x1 - X[j]) / (X[i] - X[j]);

        L += Y[i] * b;
    }
    return (L);
}

int main()
{
    float a[10], b[10], x;
    int i, num;

    cout << "Indicati numarul de noduri de interpolare:\n";
    cin >> num;
    cout << "Introduceti elementele tabloului absciselor nodurilor(x=):\n";

    for (i = 0; i < num; i++)
    {
        cout << "x[" << i << "]= ";
        cin >> a[i];
    }

    cout << "Introduceti elementele tabloului valorilor functiei in aceste puncte(y=):\n";

    for (i = 0; i < num; i++)
    {
        cout << "y[" << i << "]= ";
        cin >> b[i];
    }

    cout << "Introduceti punctul in care doriti sa se efectueze interpolarea:\n";
    cin >> x;
    cout << "\n Valoarea functiei in punctul dat este: " << lagrange(num, a, b, x) << endl;
}
