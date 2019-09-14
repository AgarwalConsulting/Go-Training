#ifndef __BYTETEST_H__
#define __BYTETEST_H__

typedef unsigned char UBYTE;

extern void ArrayReadFunc(UBYTE *arrayout);
extern void ArrayWriteFunc(UBYTE *arrayin);

#endif
#include <stdio.h>
#include <string.h>

void ArrayReadFunc(UBYTE *arrayout)
{
     UBYTE array[20] = {1, 2, 3, 4,5, 6, 7, 8, 9, 10,
                        11, 12, 13, 14, 15, 16, 17,
                        18, 19, 20};
     memcpy(arrayout, array, 20);
}

void ArrayWriteFunc(UBYTE *arrayin)
{
     UBYTE array[20];

     memcpy(array, arrayin, 20);

     printf("Byte slice array received from Go:\n");
     for(int i = 0; i < 20; i ++){
             printf("%d ", array[i]);
     }

     printf("\n");
}
