#include <stddef.h>
#include <stdlib.h>

static int*   alloc_int_array(size_t size) { return (int*)calloc(sizeof(int), size); }
static char** alloc_str_array(size_t size) { return (char**)calloc(sizeof(char*), size); }
static void   set_int_array(int* arr, int at, int val) { arr[at] = val; }
static void   set_str_array(char** arr, int at, char* val) { arr[at] = val; }

static void free_int_array(int* arr) {
        free(arr);
}

static void free_str_array(char** arr, size_t size) {
        for (int i = 0; i < size; i++) free(arr[i]);
        free(arr);
}