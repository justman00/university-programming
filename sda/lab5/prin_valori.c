// VARIANTA 4
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// C always passes arguments by value, but the string, like other arrays, 
// is converted to a pointer to its first element, and then that pointer is passed. 
// By value.
void afiseaza_string(char str[])
{
    strcpy(str, "Turcan Vladimir");
    printf("Valoarea string-ului din functie %s\n", str);
}

int main()
{
    int sum = 0;
    char numele_meu[] = "Vladimir Turcan";

    printf("Valoarea string-ului %s\n", numele_meu);
    afiseaza_string(numele_meu);
    printf("Valoarea string-ului care s-a schimbat %s\n", numele_meu);

    return 0;
}