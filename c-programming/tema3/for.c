#include <stdio.h>
#include <stdlib.h>
#include <math.h>

void calc_func(float a, int b, int range[], float pace)
{
    int start = range[0];
    int end = range[1];

    float t;

    for(t = start; t <= end + pace; t = t+ pace) {
        float solution;

        if (t < 1)
        {
            solution = 1;
        }
        else if (t > 2)
        {
            solution = pow(M_E, a * t) * cos(b * t);
        }
        else
        {
            solution = a * pow(t, 2.0) * log(t);
        }

        printf("the solution for a:%f, b:%d and t:%f is => %f\n", a, b, t, solution);
    }
}

int main(void)
{
    float a = -0.5;
    int b = 2;
    int range[2] = {0, 3};
    float pace = 0.15;

    calc_func(a, b, range, pace);

    return 0;
}