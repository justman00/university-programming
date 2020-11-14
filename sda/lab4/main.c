// Să se elaboreze un program, care va insera 25 de valori întregi aleatoare de la 0 până la 100
// într-o listă ordonată înlănţuită.
// Programul trebuie să calculeze suma elementelor şi media aritmetică,
//  care trebuie să fie număr cu virgulă mobilă.

// VARIANTA 4
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

struct Node
{
    int data;
    struct Node *next;
};

void printList(struct Node *n)
{
    while (n != NULL)
    {
        printf(" %d ", n->data);
        n = n->next;
    }
    printf("\n");
}

int main()
{
    struct Node *head;
    head = (struct Node *)malloc(sizeof(struct Node));
    time_t t;

    int lower = 0, upper = 100, count = 25;
    srand((unsigned)time(&t)); //current time as seed of random number generators

    struct Node *currNode = head;

    // construim lista
    for (int i = 0; i < count; i++)
    {
        int rand_num = (rand() % (upper + 1));

        if (i == 0)
        {
            currNode->data = rand_num;
        }
        else
        {
            struct Node *nextNode;
            nextNode = (struct Node *)malloc(sizeof(struct Node));

            nextNode->data = rand_num;
            currNode->next = nextNode;
            currNode = nextNode;
        }
    }

    printList(head);

    int sum = 0;
    currNode = head;

    while (currNode != NULL)
    {
        sum += currNode->data;
        currNode = currNode->next;
    }

    float media_aritmetica = (float)sum / count;
    printf("Suma numerelor din linked list este %d, iar media lor aritmetica %f\n", sum, media_aritmetica);

    return 0;
}