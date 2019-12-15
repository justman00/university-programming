#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int main(void)
{
    int row = 2;
    int col = 3;
    int c, d, matrix[row][col];

    printf("Enter the elements of first matrix\n");

    for (c = 0; c < row; c++)
        for (d = 0; d < col; d++)
            scanf("%d", &matrix[c][d]);

    printf("The matrix:\n");
    for (c = 0; c < row; c++)
    {
        for (d = 0; d < col; d++)
        {
            printf("%d ", matrix[c][d]);
            if (d == row)
            {
                printf("\n");
            }
        }
    }

    int x_axis[] = {};
    int y_axis[] = {};
    int i = 0;
    int j = 0;
    int sum = 0;
    int sum_cols[] = {};

    for (c = 0; c < col; c++)
    {
        if (sum != 0)
        {
            printf("size of column %d is %d\n", j, sum);
            sum_cols[j] = sum;
            sum = 0;
            j++;
        }

        for (d = 0; d < row; d++)
        {
            int current_number = matrix[d][c];

            if (current_number > 0)
            {
                sum = sum + current_number;
                x_axis[i] = c;
                y_axis[i] = d;

                int x = x_axis[i];
                int y = y_axis[i];
                int num = matrix[y][x];
                printf("the number is %d and the x is %d and the y is %d  and the i is %d\n", num, x, y, i);
                i++;
            }
        }
    }

    if (sum != 0)
    {
        printf("size of column %d is %d\n", j, sum);
        sum_cols[j] = sum;
        sum = 0;
        j++;
    }

    return 0;
}