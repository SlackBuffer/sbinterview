## C/C++
C++ 的语法向下兼容 C，C++ 有一些好用的特性和功能，因此可以在 C++ 中写 C（后缀存成 `.cpp`）：主体上使用 C 的语法，混用部分 C++ 语法。

C 的输入输出语句 `scanf` 和 `printf` 比 C++ 的 `cin` 和 `cout` 快很多，但繁琐。  
同一个程序中不要混用两种写法。

头文件的写法 `stdio.h` 和 `cstdio` 等价，推荐后者。

变量声明：

```cpp
int i;
int j = 1;

long long big_num = 123456789012345LL;
```

常用类型：

类型 | 范围 | 输入 | 输出
--|--|--|--
int                 | 32bit, $-2×10^9～2×10^9$          | `scanf("%d", &n)`         | `printf("%d", n)`
long long           | 64bit, $-9×10^{18}～9×10^{18}$    | `scanf("%lld", &n)`       | `printf("%lld", n)`
float               | 32bit，6～7位有效精度               | `scanf("%f", &f)`         | `printf("%f", f)`
double              | 64bit，15～16位效精度               | `scanf("%lf", &d)`        | *`printf("%f", d)`*
char                | 8bit, -128～127                   | `scanf("%c", &c)`         | `printf("%c", c)`
bool                | 0, 1                              | `scanf("%d", &b)`         | `printf("%d", b)`
char str[length]    | -                                 | ***`scanf("%s", str)`***  | `printf("%s", str)`           | 
|

除 `%c` 外，`scanf` 对其他格式符（如 `%d`，`%s`）的输入是以空白符（空格，字符，换行）为结束标志。

`%c` 会把空白符当成字符读入。

输出格式：
- `%md` 使**不足** m 位的 int 型变量以 m 位进行**右对齐**输出，高位用空格补齐，大于等于 m 位的正常输出。
- `%0md` 和 `%md` 的区别是用 0 而不是空格来补位。
- `%.mf` 让浮点数保留 m 位小数输出，规则是四舍六入五成双。四舍五入用 `round` 函数（使用数学函数前需加上 `math.h` 头文件）。

整型类型都可以在前面加 `unsigned`。  
`long long` 型赋大于 $2^{31}-1$ 的初值时需在初值后面加上 `LL`，否则会编译错误。

字符常量用单引号括起来。  
小写字母的 ASCII 码值比大写字母的大 32。  
`'\0'` 代表空字符 NULL，其 ASCII 码为 0，不是空格。

字符串常量用双引号括起来。  
C 的基本数据类型中没有 `string` 类型，只能用字符数组；C++ 有 `string` 类型。  
数组名称本身就代表了这个数组第一个元素的地址，所以 `scanf` 中不需要加取地址运算符。

布尔型在 C++ 中可以直接使用，在 C 中必须添加 `stdbool.h` 头文件才可以使用。  
整型常量在赋值给布尔型变量时会自动转换为 `true`（非零）或者 `false`（零）。

常量：
1. 符号常量

    ```cpp
    // 没有分号
    #define pi 3.14
    ```

    - 本质是替换，也叫宏定义。
2. `const` 常量

    ```cpp
    const double pi = 3.14;
    ```

    - 推荐 `const` 的写法。

`define` 还可以定义任何语句或片段。

```cpp
#include <stdio.h> 
#define CAL1(x) ((x) * 2 + 1)
#define CAL2(x) (x * 2 + 1)
int main() {
    int a = 1; 
    printf("%d\n", CAL1(a + 1)); // 5 == (1 + 1) * 2 + 1
    printf("%d\n", CAL2(a + 1)); // 4 == 1 + 1 * 2 + 1 
    return 0; 
}
```

可以类型转换的类型在赋值时没有显式强制类型转换时编译器会自动进行转换。

可以用连续赋值为多个变量同时赋同一个值：`n = m = 4`。

三目运算符（条件运算符）：`? : `。

`typedef` 可以给复杂的数据类型起别名。

```cpp
typedef long long LL;
LL a = 1234;
```

`switch` 语句默认会 fallthrough。

C 中不允许在 `for` 语句的表达式 1 里定义变量（`for(int i = 1; ...)`），C++ 中可以。

数组大小必须是整数常量。

数组初始化：

```cpp
int a[10] = {5, 3, 2, 6, 8, 4};

// 以下两种写法可以将所有元素初始化为 0
int a[10] = {0}; 
int a[10] = {}; 

int a[5][6] = {{3, 1, 2}, {8, 4}, {}, {1, 2, 3, 4, 5}};

char str[4] = {'a', 'b', 'c', 'd'}
// 仅限初始化时可以这样赋值
char str[4] = "abcd"
```

未被赋初值的元素将会因不同编译器内部实现的不同而被赋以不同的初值（可能是很大的随机数），一般情况默认初值为 0。

若数组大小较大（约 $10^6$ 级别），则需要定义在主函数外面，否则会使程序异常退出，原因是函数内部申请的局部变量来自系统栈，允许申请的空间较小：而函数外部申请的全局变量来自静态存储区，允许申请的空间较大。

`memset` 用来对数组中每个元素赋相同的值：`memset(arr_name, value, sizeof(arr_name))`。

只建议初学者使用 `memset` 赋 0 或 -1（理由见 `sketch/memset.cpp`），要赋其它值建议用 `fill` 函数（`memset` 执行速度更快）。  
使用 `memset` 前需在程序开头添加 `string.h` 头文件。

`getchar` 用来输入单个字符，`putchar` 用来输出单个字符。

`gets` 用来输入一行字符串。  
`gets` 将换行符 `\n` 作为输入结束标志。因此 `scanf` 读入一个整数并按下回车后，若要使用 `gets` 读入字符串，需先用 `getchar` 接收换行符。

`puts` 用来输出一行字符串并紧跟一个换行。

字符数组的每一位都是一个 `char` 字符。

一维**字符**数组（或是二维字符数组的第二维）的末尾都有一个空字符 `＼0`（只有 `char` 型数组有），以表示存放的字符串的结尾，因此声明的字符数组长度一定要比实际存储字符串的长度至少多 1。  

`puts` 和 `printf` 通过识别 `\0` 作为字符串的结尾来进行输出。

空字符 `＼0` 在使用 `gets` 或 `scanf` 输入字符串时会自动添加在输入的字符串后面，并**占用一个字符位**。若字符串不是通过以上两个函数输入的，则需手动添加 `\0`，否则 `printf` 和 `puts` 输出时会因无法识别字符串结尾而输出乱码。

常用字符数组函数（需添加 `string.h` 头文件）：
- `strlen` 返回字符数组中第一个 `＼0` 前的字符的个数。
- `strcmp` 返回两个字符数组比较大小的结果，依次比较每个字符。
    - 前小于后返回负整数。
    - 相等返回 0。
    - 前大于后返回正整数。
- `strcpy` 把第二个字符数组复制给第一个字符数组（包括结束符 `\0`）。
- `strcat` 把第二个字符数组拼接到第一个字符数组后面（第一个字符串末尾的 `\0` 会被切掉）。

`sscanf` 可以理解为 `string`（字符数组） + `scanf`，`sprintf` 可以理解为 `string` + `printf`。

```cpp
int n, m = 233;
char str1[100] = "123", str2[100];
// 把字符数组 str1 中的内容以 %d 的格式写到 n 中
sscanf(str1, "%d", &n); 
printf("%d\n", n);      // 123

// 把 m 以 %d 的格式写到字符数组 str2 中
sprintf(str2, "%d", m);
printf("%s\n", str2);   // 233
return 0; 
```

`sscanf` 支持正则表达式。


函数定义的小括号内的参数称为形式参数，实际调用时小括号内的参数称为实际参数。

数组作为参数时，参数中数组的第一维不需要写长度，二维数组的第二维要写长度。

数组不允许作为返回类型。


在某种数据类型后加星号 `*` 表示指针变量。

星号的位置在数据类型之后或是变量名之前都可以。

同一行同时定义多个同种类型的指针变量时，星号只会与第一个变量名结合。

指针是 `unsigned` 型整数。

指针变量存储的地址类型称为基类型，基类型必须和指针变量存储的地址类型相同。

指针变量可以进行加减运算。

对于 `int *p;`，`p+1` 是指 `p` 所指的 `int` 型变量的下一个 `int` 型变量的地址，下一个跨越一整个 `int` 型（4Byte）。

指针减法的结果就是两个地址**偏移的距离**，距离以基类型为单位，即相差几个基类型。

指针变量支持自增和自减操作。

可以在元素前面加取地址运算符 `&` 来获取它的地址：`a[0]` 的地址为 `&a[0]`， 即数组 a 的首地址为 `&a[0]`。  
C 中数组名称也作为数组的首地址使用（`a == &a[0]`）。

定义指针变量 `p` 时，若没有将其初始化，`p` 中存放的地址是随机的，如果 `p` 指向的是系统工作区间，改变 `p` 的值就会出问题。

```cpp
// very bad
int *p;
// ok     
int t;
int *p = &t;
```

指针类型的参数传入时也是值传递，传递的值是 unsigned 类型，是一个变量的地址的副本。函数里通过 `*` 操作符取到改地址存放的值、对值进行操作，才能改变这两个指针副本指向的同一个地址的内容。

C++ 的引用不产生副本，而是给原变量起了一个别名，别名和原名指向同一个变量，使用方法是在函数参数类型后面加 `&`。

```cpp
void change(int &n) {
    n = 1;
}
int main() {
    int x = 2;
    change(x);
    printf("%d\n", x);  // 1
    return 0;
}
```

引用是产生变量的别名，因此常量不可使用引用。


结构体里能定义除了自己本身（避免循环定义）之外的任意数据类型（包括自身类型的指针变量）。

```cpp
struct node {
    node* next;
};

struct student_info {
    int id;
    char gender;
    char name[20];
    char major[20];
} Alice, Bob, stu[100]; // 紧接着直接声明变量
```

访问结构体内元素：

```cpp
struct student_info {
    int id;
    student_info* next;
} stu, *p;

stu.id;
stu.next;
(*p).id;
(*p).next;

stu->id;
stu->next;
p->id;
p->next;
```

一个普通定义的结构体的内部会生成一个默认的构造函数（不可见）。

构造函数不需要写返回值类型，函数名称与结构体名相同。

构造函数的作用是初始化结构体。  

默认的构造函数没有参数，函数体为空。  

由于默认构造函数的存在，才可以直接定义 `student_info` 类型的变量而不进行初始化，因为默认构造函数不需要任何初始化参数。

可以重载构造函数，但之后定义变量的时候就需要按新的构造函数初始化结构体变量。

可以定义任意多个构造函数（只要它们的函数签名不同，如只初始化部分字段的构造函数）以适应不同的初始化场景。

```cpp
struct student_info {
    int id;
    char gender;
    // 默认生成的构造函数
    student_info() {}
}

struct student_info {
    int id;

    student_info() {}

    student_info(int _id) {
        id = _id;
    }

    // 简化成一行的写法
    student_info(int _id, char _gender): id(_id), gender(_gender) {}
}

// 手动初始化
stu.id = 1;
scanf("%d", &stu.id);
```

推荐使用构造函数初始化结构体。


`cin` 和 `cout` 不需要指定输入输出的格式，不需要使用取地址运算符 `&`，可以直接进行输入（`>>`）和输出（`<<`）。

需添加 `#include <iostream>` 和 `using namespace std;`。

```cpp
// 依次读入多个变量
cin >> n >> db >> c >> str;

// 读入一整行
char str[100];
cin.getline(str, 100);

// string 容器
string str;
getline(cin, str);

// 值之间没有空格分隔
cout << n << db << c << str;
// 两种换行方式
cout << n << " " << db << "\n" << c << str << endl;
// 输出 2 位小数
#include <iomanip>
cout << setiosflags(ios::fixed) << setprecision(2) << 123.4567 << endl;
```

计算机中采用有限位数的二进制编码，因此浮点数在计算机中的存储并不总是精确。

若经过容易损失精度的运算后，需要比较浮点数大小，可以引入一个极小数 eps（一般取 10^-8）来对误差进行修正。  
没有经过损精度的运算时则可以直接比较。

```cpp
#define Equal(a, b)         ((fabs((a) - (b))) < (eps))
#define Larger(a, b)        (((a) - (b)) > (eps))
#define Smaller(a, b)       (((a) - (b)) < (-eps))
// b 的范围是 [b-eps, b+eps]，大于 b-eps 的数就可以判定为大于等于 b，即需要满足 a > b-eps
#define LargerEqual(a, b)   (((a) - (b)) > (-eps)
// 小于 b+eps 的数就可以判定为小于等于 b，a < b+eps
#define SmallerEqual(a, b)  (((a) - (b)) < (eps)

const double eps = 1e-8;
```

圆周率值：

```cpp
// const(pi) = -1 => pi = arccos(-1)
const double pi = acos(-1.0);
```

由于精度问题，经过大量运算后一个变量中存储的 0 可能是一个很小的负数，此时对其开根号 `sqrt` 会因为参数不在定义域内而报错。同样的问题还出现在 `asin(x)` 的 x 存放 +1、`acos(x)` 的 x 存放 -1 时的情况。这些情况下要用 eps 保证变量在定义域内。

某些由编译环境下本应为 0.00 的变量在输出时会变成 -0.00，这是编译环境本身的问题，只能把结果存放到字符串中，然后与 -0.00 进行比较，如果比对成功，则加上 eps 将它的值修正为 0.00。

`scanf` 函数的返回值为其成功读入的参数的个数。  
读取文件读到文件末尾导致无内容可读取时会产生读入失败的错误，即 EOF，C 中 EOF 用 -1 表示。

```cpp
while (scanf("%d", &n) != EOF) {}
while (scanf("%s", str) != EOF) {}
// 字符串还有另一种读法可以用
while(gets(str) != NULL) {}

// a 或 b 有一个不为零就继续循环
while (scanf("%d%d", &a, &b), a || b) {}
```


base^exponent

---
- [ ] Range and precision: [IEEE 754-1985](https://en.wikipedia.org/wiki/IEEE_754-1985), https://blog.csdn.net/ucan23/article/details/44965609
- [ ] 四舍六入五成双
---
> ***References***
> - [算法笔记](https://book.douban.com/subject/26827295/)