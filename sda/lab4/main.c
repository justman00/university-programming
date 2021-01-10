// STRUCTURE DE DATE FISIERE SI INREGISTRARI
// VARIANTA 4
// Să se creeze fişierul student. Introduceţi în fişiere separate înregistrările referitor la studenţii cu diverse forme de studii.
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
        // parcurgem prin fiecare fiecare student
        struct student std = studenti[n];
        char fileName = std.formaInstruirii;
        char path[64];
        // creem denumirea de file care o sa fie in conformitate cu forma lor de instruire
        snprintf(path, sizeof(path), "./%C.%s", fileName, fileEnding);
        printf("Acest este npp unui student %s\n", std.npp);
        printf("acesta este file-ul in care va fi scris studentul %s\n", path);

        // deschidem acel file pentru a adauga, daca file-ul nu exista se va crea
        fptr = fopen(path, "ab");
        // scriem direct in file
        fwrite(&std, sizeof(struct student), 1, fptr);
    }

    return 0;
}

/**
 * In cadrul acestui program am invatat cum poate fi manipulat sistemul de fisiere al calculatorului
 * cu ajutorul unor functii/comenzi predefinite din librariile standarde ale limbajului C.
 * Cu atat mai mult, am lucrat cu structuri mai complexe de date, in cazul acesta am folosit
 * structuri de date originale, definite prin "struct" si oferindu-le un set de proprietati.
 * Totodata, am scris anumite date din memoria C in fisiere separate, precum si am citit din
 * acel fisier pentru a structura datele separat dupa un anumit criteriu, in cazul problemei mele
 * fiind forma instruirii stundetilor, exprimata prin 0 si 1, astfel obtinand 0.bin si 1.bin.
 * */