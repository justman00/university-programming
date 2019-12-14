#include <stdio.h>
#include <stdlib.h>
#include <math.h>

void calc_func(int x, int y, double z)
{
    double a = (2*cos(x - M_PI/6))/(1/2 + pow(sin(y), 2.0));
    double b = 1 + pow(z, 2.0)/(3 + pow(z, 2.0) / 5);

    printf("The answer for a is %lf\n", a);
    printf("The answer for b is %lf\n", b);
}

int main(void)
{
    int x = 1426;
    int y = -1220;
    double z = 3.5;

    calc_func(x, y, z);

    return 0;
}