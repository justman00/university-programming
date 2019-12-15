#include <stdio.h>
#include <stdlib.h>
#include <math.h>

void calc_func(float a, int b, float t)
{
    float comp = t;

    float solution;

    if (comp < 1)
    {
        goto smaller_than_one;
    }
    else if (comp > 2)
    {
        goto bigger_than_two;
    }
    else
    {
        goto between_one_and_two;
    }

    smaller_than_one:
        solution = 1;
    between_one_and_two:
        solution = a * pow(t, 2.0) * log(t);
    bigger_than_two:
        solution = pow(M_E, a * t) * cos(b * t);


    printf("the solution for a:%f, b:%d and t:%f is => %f\n", a, b, t, solution);
}

int main(void)
{
    float a = -0.5;
    int b = 2;
    // here you can choose t to be from a range 0f 0 to 3
    float t = 3;

    calc_func(a, b, t);

    return 0;
}