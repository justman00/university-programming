#include <stdio.h>
#include <stdlib.h>
#include <math.h>

void get_sum(int *arr, int size, int *total, int *positives)
{
    int sum = 0;
    int count = 0;
    int i;

    for (i = 0; i < size; i++)
    {
        int current_number = *(arr + i);

        if (current_number >= 0)
        {
            sum = sum + current_number;
            count++;
        }
    }

    *total = sum;
    *positives = count;
};

float get_result(int s1, int s2, int k1, int k2)
{
    float result = (float)(s1 + s2) / (float)(k1 * k2);

    return result;
}

int main(void)
{
    int s1 = 0;
    int k1 = 0;
    int arr1[] = {1, 2, 3, 4, 5, 6, 6, -1, 0, -4, 10};
    int size1 = sizeof(arr1) / sizeof(arr1[0]);
    get_sum(arr1, size1, &s1, &k1);

    printf("sum and count %d, %d\n", s1, k1);

    int s2 = 0;
    int k2 = 0;
    int arr2[] = {20, 4, -6, 3, 90, -12, -8, 9, 1, 0};
    int size2 = sizeof(arr2) / sizeof(arr2[0]);
    get_sum(arr2, size2, &s2, &k2);

    printf("sum and count2 %d, %d\n", s2, k2);

    float result = get_result(s1, s2, k1, k2);

    printf("here is the result %.2f\n", result);

    return 0;
}