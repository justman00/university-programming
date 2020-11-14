#include <stdio.h>
#include <stdlib.h>

struct examen
{
    char denumireaDisciplinei[10];
    char nota;
};

struct student
{
    char npp[40];
    char numarulSeriei;
    char numarulGrupei;
    char formaInstruirii; // 0 pt frecventa si 1 pentru frecv redusa
    struct examen reusita[10][5];
    struct examen ex;
};

struct student studenti[2] = {
    {.npp = "12345678901234567890",
     .numarulSeriei = '1',
     .numarulGrupei = '3',
     .formaInstruirii = '0', // la frecventa
     .reusita = {},
     .ex = {
         .denumireaDisciplinei = "Disc 1",
         .nota = '8'}},
    {.npp = "98765432101234567890", .numarulSeriei = '3', .numarulGrupei = '1',
     .formaInstruirii = '1', // la frecventa
     .reusita = {},
     .ex = {.denumireaDisciplinei = "Disc 1", .nota = '5'}}};

void populate_file_with_data()
{
    FILE *fptr = fopen("./date.bin", "wb");

    for (int n = 0; n < 2; ++n)
    {
        struct student std = studenti[n];

        fwrite(&std, sizeof(struct student), 1, fptr);
    }
    fclose(fptr);
}

int main()
{
    populate_file_with_data();

    struct student std;
    char fileEnding[4] = "bin";

    FILE *fptr;

    for (int n = 0; n < 2; ++n)
    {
        struct student std = studenti[n];
        char fileName = std.formaInstruirii;
        char path[64];
        snprintf(path, sizeof(path), "./%C.%s", fileName, fileEnding);
        printf("Acest este npp unui student %s\n", std.npp);
        printf("acesta este file-ul in care va fi scris studentul %s\n", path);

        fptr = fopen(path, "ab");

        fwrite(&std, sizeof(struct student), 1, fptr);
    }

    return 0;
}