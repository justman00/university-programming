// SIRURI DE CARACTERE
// Să se elaboreze un program, care va introduce 4 şiruri care reprezintă valori întregi, transformă şirurile în numere întregi, sumează şi afişează suma celor patru valori.
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

/**
 * In cadrul acestei lucrari de laborator, am creat un program care citeste siruri de caractere si le transforma
 * dintr-un tip de date in altul(char -> integer). Pentru a face acest lucru am folosit metoda "atoi".
 * Iar citirea datelor de la tastatura a fost posibila datorita functiei "scanf"
*/