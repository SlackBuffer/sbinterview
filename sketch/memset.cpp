#include <stdio.h> 
#include <string.h> 

// from https://stackoverflow.com/a/3974138/6902525
void printBits(size_t const size, void const * const ptr)
{
    unsigned char *b = (unsigned char*) ptr;
    unsigned char byte;
    int i, j;

    for (i=size-1;i>=0;i--)
    {
        for (j=7;j>=0;j--)
        {
            byte = (b[i] >> j) & 1;
            printf("%u", byte);
        }
    }
    puts("");
}

int main() {
    /* 
    void *memset(void *s, int c, size_t n);
    The memset() function fills the first n bytes of the memory area pointed to by s with the constant **byte** c. 
    按字节填充的！
    */

    int i = 0, j = -1, k = 1;
    printBits(sizeof(i), &i);   // 00000000000000000000000000000000   
    printBits(sizeof(j), &j);   // 11111111111111111111111111111111
    printBits(sizeof(k), &k);   // 00000000000000000000000000000001

    char trunc_num = char(k);
    printBits(sizeof(trunc_num), &trunc_num);
    // 00000001
    // 00000001000000010000000100000001 = 16843009


    int d[5];
    printf("%lu\n", sizeof(d)); // 20
    memset(d, 1, sizeof(d));
    for (int i = 0; i < 5; i++) {
        printf("%d ", d[i]);    // 16843009
    }
    return 0; 
}