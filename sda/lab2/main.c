// VARIANTA 4
#include <stdio.h>
#include <stdlib.h>

int main()
{
    int sum = 0;

    for (int i = 0; i < 4; i++)
    {
        char sir[64];
        printf("Introduceti un sir de numere fara spatiu: ");
        scanf("%s", sir);

        int aux = atoi(sir);

        sum += aux;
    }

    printf("Aceasta este suma celor 4 siruri de numere %d\n", sum);

    return 0;
}