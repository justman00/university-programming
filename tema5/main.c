#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int main(void)
{
    int row = 3;
    int col = 2;
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

    printf("\n");

    int i = 0;
    int j = 0;
    int sum = 0;
    int pos = 0;
    int sum_cols[] = {};
    int pos_nums[] = {};

    for (c = 0; c < col; c++)
    {
        if (sum != 0)
        {
            printf("sum of column %d is %d\n", j, sum);
            printf("number of positives of column %d is %d\n", j, pos);
            sum_cols[j] = sum;
            pos_nums[j] = pos;
            pos = 0;
            sum = 0;
            j++;
        }

        for (d = 0; d < row; d++)
        {
            int current_number = matrix[d][c];

            if (current_number > 0)
            {
                sum = sum + current_number;
                i++;
                pos++;
            }
        }
    }

    if (sum != 0)
    {
        printf("sum of column %d is %d\n", j, sum);
        printf("number of positives of column %d is %d\n", j, pos);
        sum_cols[j] = sum;
        pos_nums[j] = pos;
    }

    return 0;
}