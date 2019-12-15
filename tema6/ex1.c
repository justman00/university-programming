#include <stdio.h>
#include <stdlib.h>
#include <math.h>

void find_result(double a, double b, double c) {
    double delta, x1, x2;
    delta = b * b - 4 * a * c;

    // delta mai mare ca 0
    if (delta > 0) {
        x1 = (-b + sqrt(delta)) / (2 * a);
        x2 = (-b - sqrt(delta)) / (2 * a);
        printf("x1 = %.2lf and x2 = %.2lf\n", x1, x2);
    }
    // delta egal cu 0
    else if (delta == 0) {
        x1 = x2 = -b / (2 * a);
        printf("x1 = x2 = %.2lf;\n", x1);
    }
    // delta mai mic ca 0
    else {
        printf("Delta mai mic ca 0\n");
    }
}

int main(void)
{   
    double a, b, c;
    printf("Introduceti coeficientii a, b si c: ");
    scanf("%lf %lf %lf", &a, &b, &c);

    find_result(a, b, c);

    return 0;
}