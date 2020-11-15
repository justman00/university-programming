// VARIANTA 4
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char str_global[] = "Vladimir Turcan";

// C always passes arguments by value, but the string, like other arrays,
// is converted to a pointer to its first element, and then that pointer is passed.
// By value.
void afiseaza_string()
{
    strcpy(str_global, "Turcan Vladimir");
    printf("Valoarea string-ului din functie %s\n", str_global);
}

int main()
{
    int sum = 0;

    printf("Valoarea string-ului %s\n", str_global);
    afiseaza_string();
    printf("Valoarea string-ului care s-a schimbat %s\n", str_global);

    return 0;
}