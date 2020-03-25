#include <cstdio>

const int maxn = 11;

// 数组 P 用来存放当前排列
// hashTable[x] 当整数 x 在 P 中时为 true
int n, P[maxn], hashTable[maxn] = {false};

/*
按顺序往数组 P 的第 1 位到第 n 位中填入数字。假设当前已经填好了 P[1］～ P[index-1]，准备填 P[index]。

显然需要枚举 1 ～ n，如果当前枚举的数字 x 还没有在 P[1］～ P[index-1］中（即 hashTable[x］== false），那么就把它填入 P[index]，同时将 hashTable[x］置为 true, 然后去处理 P 的第 index + 1 位（即进行递归〉；而当递归完成时，再将 hashTable[x］还原为 false, 以便让 P[index］填下一个数字。

当 index 达到 n + 1 时，说明 P 的第 1 ～ n 位都已经填好了，此时可以把数组 P 输出，表示生成了一个排列，然后直接 return 即可。
*/

// 当前处理排列的第 index 号位
void generateP(int index) {
    if (index == n + 1) { // 递归边界，已经处理完排列的 1~n 位
        for (int i = 1; i <= n; i++) {
            printf("%d", P[i]);
        }
        printf("\n");
        return;
    }

    for (int x = 1; x <= n; x++) {      // 枚举 1~n，试图将 x 填入 P[index]
        if (hashTable[x] == false) {    // x 不在 P[0]~P[index-1] 中
            P[index] = x;               // 令 P 的第 index 位为 x，即把 x 加入当前排列
            hashTable[x] = true;        // 记录 x 已在 P 中
            generateP(index + 1);       // 处理排列的第 index+1 号位
            hashTable[x] = false;       // 已处理完 P[index] 为 x 的子问题，还原状态
        }
    }
}

int main() {
    n = 3;  // 输出 1~4 的全排列
    generateP(1); // 从 P[1] 开始填
    return 0;
}