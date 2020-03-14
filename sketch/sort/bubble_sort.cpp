#include <cstdio>

const int arr_length = 8;
/* 
冒泡排序
每一轮中，当前元素 curr 依次和下一个元素 next 比较大小，当前元素的下标每次比较后递增。
1. 升序：curr > next 时交换，每轮的最大值冒泡都数组末尾。
2. 降序：curr < next 时交换，每轮的最小值冒泡到数组末尾。
首先排好序的元素在数组末尾。
*/

int main() {
    int a[arr_length] = {3,5,2,8,0,9,4,19};

    // sb's implementation
    // for (int i = 0; i < arr_length-1; i++) {
        // for (int j = 0; j < arr_length-1-i; j++) {

    // ** 外层循环下标从 1 开始
    for (int i = 1; i < arr_length; i++) {
        for (int j = 0; j < arr_length-i; j++) {
            if (a[j] > a[j+1]) {
                int t = a[j];
                a[j] = a[j+1];
                a[j+1] = t;
            }
        }
    }
    for (int i = 0; i < arr_length; i++) {
        printf("%d ", a[i]);
    }
    printf("\n");
}