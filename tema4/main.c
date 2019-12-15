#include <stdio.h>
#include <stdlib.h>
#include <math.h>

void calc_func(float * arr, int size)
{
    float sum = 0;
    int i;

    for (i = 0; i < size; i++)
    {
        float current_number = *(arr + i);

        if (0 <= current_number && 1 >= current_number) {
            sum = sum + current_number;
        }
    }

    printf("The lenght of the array is %d, and the sum of the numbers between 0 and 1 is %f\n", size, sum);
}

int main(void)
{
    float arr[] = {0, 3, 9, 0.1, 0.4, 2, 0.9, 1};
    int size = sizeof(arr) / sizeof(arr[0]);

    calc_func(arr, size);

    return 0;
}