
// SORTAREA DATELOR
// VARIANTA 4
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

int nr_iteratii = 0, nr_comparatii = 0, nr_permutari = 0;

int bubble_sort(int array[], int size)
{
    int counter, counter1, swap_var;
    nr_iteratii = 0, nr_comparatii = 0, nr_permutari = 0;
    float timp_de_executie;

    clock_t begin = clock();

    for (counter = 0; counter < size - 1; counter++)
    {
        for (counter1 = 0; counter1 < size - counter - 1; counter1++)
        {
            nr_comparatii++;
            nr_iteratii++;
            if (array[counter1] > array[counter1 + 1])
            {
                swap_var = array[counter1];
                array[counter1] = array[counter1 + 1];
                array[counter1 + 1] = swap_var;
                nr_permutari++;
            }
        }
    }

    clock_t end = clock();
    timp_de_executie += (double)(end - begin) / CLOCKS_PER_SEC;

    printf("Array-ul sortat:\n");
    for (int i = 0; i < size; i++)
    {
        printf("%d  ", array[i]);
    }

    printf("\nMETODA BULELOR:\n");
    printf("Timp de executie: \t %f secunde \n", timp_de_executie);
    printf("Numar iteratii: \t %d \n", nr_iteratii);
    printf("Numar comparatii: \t %d \n", nr_comparatii);
    printf("Numar permutari: \t %d \n", nr_permutari);

    return 0;
}

void selection_sort(int array[], int size)
{
    int counter1, counter2, minimum, temp_value;
    nr_iteratii = 0, nr_comparatii = 0, nr_permutari = 0;
    float timp_de_executie;

    clock_t begin = clock();

    for (counter1 = 0; counter1 < size - 1; counter1++)
    {
        minimum = counter1;
        for (counter2 = counter1 + 1; counter2 < size; counter2++)
        {
            nr_iteratii++;
            nr_comparatii++;
            if (array[minimum] > array[counter2])
                minimum = counter2;
            nr_permutari++;
        }
        nr_comparatii++;
        if (minimum != counter1)
        {
            temp_value = array[counter1];
            array[counter1] = array[minimum];
            array[minimum] = temp_value;
            nr_permutari++;
        }
    }

    clock_t end = clock();
    timp_de_executie += (double)(end - begin) / CLOCKS_PER_SEC;

    printf("Array-ul sortat:\n");
    for (int i = 0; i < size; i++)
    {
        printf("%d  ", array[i]);
    }

    printf("\nMETODA SELECTION SORT:\n");
    printf("Timp de executie: \t %f secunde \n", timp_de_executie);
    printf("Numar iteratii: \t %d \n", nr_iteratii);
    printf("Numar comparatii: \t %d \n", nr_comparatii);
    printf("Numar permutari: \t %d \n", nr_permutari);
}

void quicksort_method(int element_list[], int low, int high)
{
    int pivot, value1, value2, temp;
    nr_comparatii++;
    if (low < high)
    {
        pivot = low;
        value1 = low;
        value2 = high;
        while (value1 < value2)
        {
            nr_iteratii++;
            while (element_list[value1] <= element_list[pivot] && value1 <= high)
            {
                nr_iteratii++;
                value1++;
            }
            while (element_list[value2] > element_list[pivot] && value2 >= low)
            {
                nr_iteratii++;
                value2--;
            }

            nr_comparatii++;
            if (value1 < value2)
            {
                temp = element_list[value1];
                element_list[value1] = element_list[value2];
                element_list[value2] = temp;
                nr_permutari++;
            }
        }
        temp = element_list[value2];
        element_list[value2] = element_list[pivot];
        element_list[pivot] = temp;
        nr_permutari++;
        quicksort_method(element_list, low, value2 - 1);
        quicksort_method(element_list, value2 + 1, high);
    }
}

void quick_sort_invoke(int array[], int size)
{
    nr_iteratii = 0, nr_comparatii = 0, nr_permutari = 0;
    float timp_de_executie;

    clock_t begin = clock();

    quicksort_method(array, 0, size - 1);

    clock_t end = clock();
    timp_de_executie += (double)(end - begin) / CLOCKS_PER_SEC;

    printf("Array-ul sortat:\n");
    for (int i = 0; i < size; i++)
    {
        printf("%d  ", array[i]);
    }

    printf("\nMETODA QUICKSORT:\n");
    printf("Timp de executie: \t %f secunde \n", timp_de_executie);
    printf("Numar iteratii: \t %d \n", nr_iteratii);
    printf("Numar comparatii: \t %d \n", nr_comparatii);
    printf("Numar permutari: \t %d \n", nr_permutari);
}

void perform_merge(int val[], int counter11, int counter12, int counter22, int counter21)
{
    int temp_val[100];
    int c1, c2, c3;
    c1 = counter11;
    c2 = counter22;
    c3 = 0;
    while (c1 <= counter12 && c2 <= counter21)
    {
        nr_iteratii++;
        nr_comparatii += 1;
        if (val[c1] < val[c2])
        {
            nr_permutari++;
            temp_val[c3++] = val[c1++];
        }
        else
        {

            nr_permutari++;
            temp_val[c3++] = val[c2++];
        }
    }
    while (c1 <= counter12)
    {

        nr_iteratii++;
        temp_val[c3++] = val[c1++];
    }
    while (c2 <= counter21)
    {
        nr_iteratii++;
        temp_val[c3++] = val[c2++];
    }
    for (c1 = counter11, c2 = 0; c1 <= counter21; c1++, c2++)
    {
        nr_iteratii++;
        val[c1] = temp_val[c2];
    }
}

void algo_merge_sort(int val[], int counter1, int counter2)
{
    int mid;
    if (counter1 < counter2)
    {
        mid = (counter1 + counter2) / 2;
        algo_merge_sort(val, counter1, mid);
        algo_merge_sort(val, mid + 1, counter2);
        perform_merge(val, counter1, mid, mid + 1, counter2);
    }
}

void invoke_merge_sort(int array[], int size)
{
    nr_iteratii = 0, nr_comparatii = 0, nr_permutari = 0;
    float timp_de_executie;

    clock_t begin = clock();

    algo_merge_sort(array, 0, size - 1);

    clock_t end = clock();
    timp_de_executie += (double)(end - begin) / CLOCKS_PER_SEC;

    printf("Array-ul sortat:\n");
    for (int i = 0; i < size; i++)
    {
        printf("%d  ", array[i]);
    }

    printf("\nMETODA MERGE SORT:\n");
    printf("Timp de executie: \t %f secunde \n", timp_de_executie);
    printf("Numar iteratii: \t %d \n", nr_iteratii);
    printf("Numar comparatii: \t %d \n", nr_comparatii);
    printf("Numar permutari: \t %d \n", nr_permutari);
}

int insertion_sort(int array[], int size)
{
    nr_iteratii = 0, nr_comparatii = 0, nr_permutari = 0;
    float timp_de_executie;

    clock_t begin = clock();
    int counter1, counter2, temp_val;

    for (counter1 = 1; counter1 <= size - 1; counter1++)
    {
        nr_iteratii++;
        temp_val = array[counter1];
        counter2 = counter1 - 1;
        nr_comparatii++;
        while ((temp_val < array[counter2]) && (counter2 >= 0))
        {
            // deoarece conditia din while e o comparatie si e verificata la fiecare iteratie
            nr_comparatii++;
            nr_iteratii++;
            nr_permutari++;

            array[counter2 + 1] = array[counter2];
            counter2 = counter2 - 1;
        }
        nr_permutari++;
        array[counter2 + 1] = temp_val;
    }

    clock_t end = clock();
    timp_de_executie += (double)(end - begin) / CLOCKS_PER_SEC;

    printf("Array-ul sortat:\n");
    for (int i = 0; i < size; i++)
    {
        printf("%d  ", array[i]);
    }

    printf("\nMETODA INSERTION SORT:\n");
    printf("Timp de executie: \t %f secunde \n", timp_de_executie);
    printf("Numar iteratii: \t %d \n", nr_iteratii);
    printf("Numar comparatii: \t %d \n", nr_comparatii);
    printf("Numar permutari: \t %d \n", nr_permutari);

    return 0;
}

int main()
{
    time_t t;
    int ARRAY_SIZE = 100;
    int arr[ARRAY_SIZE];
    int lower = 0, upper = 100;
    srand((unsigned)time(&t)); //current time as seed of random number generators

    // construim lista
    for (int i = 0; i < ARRAY_SIZE; i++)
    {
        int rand_num = (rand() % (upper + 1));
        arr[i] = rand_num;
    }

    printf("\nArray-ul nesortat:\n");
    for (int i = 0; i < ARRAY_SIZE; i++)
    {
        printf("%d  ", arr[i]);
    }

    printf("\n");

    int arr_1_copy[ARRAY_SIZE];
    memcpy(arr_1_copy, arr, sizeof(arr_1_copy));
    bubble_sort(arr_1_copy, ARRAY_SIZE);

    printf("-------------------------\n");

    int arr_2_copy[ARRAY_SIZE];
    memcpy(arr_2_copy, arr, sizeof(arr_2_copy));
    selection_sort(arr_2_copy, ARRAY_SIZE);

    printf("-------------------------\n");

    int arr_3_copy[ARRAY_SIZE];
    memcpy(arr_3_copy, arr, sizeof(arr_3_copy));
    quick_sort_invoke(arr_3_copy, ARRAY_SIZE);

    printf("-------------------------\n");

    int arr_4_copy[ARRAY_SIZE];
    memcpy(arr_4_copy, arr, sizeof(arr_4_copy));
    invoke_merge_sort(arr_4_copy, ARRAY_SIZE);

    printf("-------------------------\n");

    int arr_5_copy[ARRAY_SIZE];
    memcpy(arr_5_copy, arr, sizeof(arr_5_copy));
    insertion_sort(arr_5_copy, ARRAY_SIZE);

    printf("-------------------------\n");

    return 0;
}

/**
 * In cadrul efectuarii acestei lucrari de laborator am urmarit ca scop principal compararea intre ei a mai multor algoritme de
 * sortare, ele fiind: al bulelor, selection sort, quick sort, merge sort si insertion sort. Acestea sunt mai mult sau mai putin
 * cele mai populare si cunoscute algoritme. Compararea timpului de executie, al numarului de iteratii, comparatii si permutari
 * ne-a aratat care este totusi diferenta intre ele si care este cel mai eficient, in cazul dat fiind quicksort, care este lider
 * dupa toate metricele colectate.
*/