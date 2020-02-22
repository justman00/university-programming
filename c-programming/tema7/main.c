#include <stdio.h>
#include <stdlib.h>

int main()
{
    char str[20];
    int i = 0;

    printf("Enter String: ");
    fgets(str, sizeof(str), stdin);

    int count = 0;
    while (str[i] != '\0')
    {
        printf("%c\n", str[i]);
        if (str[i] == '}')
            count--;
        if (str[i] == '{')
            count++;
        if (count < 0)
        {
            printf("\nInvalid");
            break;
        }
        i++;
    }
    if (count == 0)
        printf("\nValid\n");
    else
    {
        printf("Invalid\n");
    }
    return 0;
}