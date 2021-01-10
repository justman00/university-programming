// Modalități de transmitere a valorilor parametrilor funcției prin: valoare, (adresa), variabila globală
// VARIANTA 4
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// C always passes arguments by value, but the string, like other arrays,
// is converted to a pointer to its first element, and then that pointer is passed.
void afiseaza_string(char *str)
{
    strcpy(str, "Turcan Vladimir");
    printf("Valoarea string-ului din functie %s\n", str);
}

int main()
{
    int sum = 0;
    char numele_meu[] = "Vladimir Turcan";

    printf("Valoarea string-ului %s\n", numele_meu);
    afiseaza_string(&numele_meu[0]);
    printf("Valoarea string-ului care s-a schimbat %s\n", numele_meu);

    return 0;
}

/**
 * In cadrul efectuarii acestei lucrari de laborator am observat si analizat diferitele modalitati de transmitere
 * a argumentelor in alte functii. Aceste modalitati fiind: prin valoare, prin referinta(adresa) si prin intermediul unei
 * variabile globale. Cu toate acestea, am mai invatat ca in limbajul C argumentele intotdeauna sunt oferite ca valoare,
 * iar atunci cand sunt array-uri sau string-uri, el le transforma intr-un pointer pe primul element
 * si apoi transmite acel pointer.
*/
