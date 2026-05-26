# 20260421 go教程

## 一、 Golang执行流程分析

### 执行方式对比

**真实环境下（go build）：**
1. `.go文件` → **编译** → `可执行文件`（.exe或可执行文件）
2. `可执行文件` → **运行** → `结果`

**开发环境下（go run）：**
- 直接输入 `go run hello.go`
- 编译并运行一步完成 → `结果`

### 流程详解

#### go build 方式
- **源文件**（Source Code）：.go文件
- **编译**（Compile）：将Go源代码编译成可执行文件
- **可执行文件**：.exe或可执行文件（取决于操作系统）
- **运行**（Execute）：执行编译后的二进制文件
- **结果**（Result）：程序执行结果

#### go run 方式
- **快速执行**：无需显式编译步骤
- **编译运行一步**：自动编译并立即执行
- 按住左键选择代码区域，按ESC退出
- **结果**（Result）：程序执行结果

---

## 二、 重要说明

### 执行流程的方式区别

**1) 编译后的可执行文件可跨平台运行**
   - 如果我们先编译生成了可执行文件，那么我们可以将该可执行文件拷贝到没有Go开发环境的机器上，**仍然可以运行**

**2) go run 方式需要Go开发环境**
   - 如果我们是直接 `go run` Go源代码，那么如果要在另外一个机器上这么运行，也需要Go开发环境
   - 否则无法执行

**3) 编译时文件包含的库**
   - 在编译时，编译器会将程序运行依赖的库文件包含在可执行文件中
   - 所以，可执行文件变大了很多

---

## 三、 VSCode快捷键 - 基础编辑（Mac键盘）

| 按键 | 功能 |
|------|------|
| `Cmd+X` | 剪切行（空选定）Cut line (empty selection) |
| `Cmd+C` | 复制行（空选定）Copy line (empty selection) |
| `Option+↑` / `Option+↓` | 向上/向下移动行 Move line up/down |
| `Shift+Option+↑` / `Shift+Option+↓` | 向上/向下复制行 Copy line up/down |
| `Cmd+Shift+K` | 删除行 Delete line |
| `Cmd+Enter` | 在下面插入行 Insert line below |
| `Cmd+Shift+Enter` | 在上面插入行 Insert line above |
| `Cmd+Shift+\` | 跳到匹配的括号 Jump to matching bracket |
| `Cmd+]` / `Cmd+[` | 缩进/反缩进行 Indent/outdent line |

---

## 四、 go build 详解

### 编译原理

**1) Go源文件的编译过程**
   - 有了 go 源文件，通过编译器将其编译成机器可以识别的二进制码文件

**2) 编译和指定输出名称**
   - 在该源文件目录下，通过 `go build` 对 `hello.go` 文件进行编译
   - 可以指定生成的可执行文件名，在 Windows 下必须是 `.exe` 后缀
   
   **示例命令：**
   ```
   go build -o myhello.exe hello.go
   ```

**3) 编译成功的表现**
   - 如果程序没有错误，没有任何提示，会在当前目录下会出现一个可执行文件
   - **Windows 下是 `.exe` 文件**
   - **Linux 下是一个可执行文件**，该文件是二进制码文件，也是可以执行的程序

**4) 编译时的错误提示**
   - 如果程序有错误，编译时，会在错误的那行报错
   - 有助于程序员调试错误

   **示例错误：**
   ```
   go build hello.go
   # command-line-arguments
   hello.go:7:2: undefined: fmt.Println
   ```

---

## 五、 Go程序开发注意事项（重点）

**1) Go源文件以 "go" 为扩展名**
   - 所有Go源代码文件都必须使用 `.go` 作为文件扩展名

**2) Go应用程序的执行入口是main()函数**
   - 每个Go程序都需要有一个 `main()` 函数作为程序的入口点

**3) Go语言严格区分大小写**
   - 变量名、函数名等都要注意大小写的区分

**4) Go方法由一条条语句构成，每个语句后不需要分号**
   - **Go语言会在每行语句后自动加分号**，这体现出Golang的简洁性
   - 每一行一条语句，不能把多条语句写在同一个

**5) Go编译器是一行一行进行编译的**
   - 因此我们一行就写一条语句，不能把多条语句写在同一个
   - 否则报错

**6) go语言定义的变量或者import的包如果没有使用到，代码不能编译通过**
   - 这是Go的严格特性，所有定义的变量和导入的包都必须被使用

**7) 大括号都是成对出现的，缺一不可**
   - Go语言要求所有的括号必须配对，否则无法编译

---

## 六、 规范的代码风格

### 正确的注释和注释风格

**1) Go官方推荐使用行注释来注释整个方法和语句**
   - 使用 `//` 进行行注释

**2) 带有Go源码**
   - 注释应该简洁清晰，说明代码的功能

### 正确的缩进和空白

**1) 使用一次tab操作，实现缩进**
   - 认为整体向右边移动，时候用 `Shift+Tab` 整体向左移

**2) 或者使用 gofmt 来进行格式化 【演示】**
   - gofmt 是Go官方提供的代码格式化工具
   - 可以自动规范代码的缩进和空白

**3) 运算符两边习惯性各加一个空格。比如：2 + 4 * 5**
   - 这样可以提高代码的可读性

### 代码风格示例

**正确的代码风格：**
```go
package main
import "fmt"
func main(){
  fmt.Println("hello,world!")
}
```

**错误的代码风格（Go语言不允许这样写，是错误的！）：**
```go
package main
import "fmt"
func main()
{
  fmt.Println("hello,world!")
}
```

**注意：** 在Go语言中，`func main()` 和 `{` 必须在同一行，不能换行，否则会编译错误！

---

## 七、 Go程序的基本结构

### Go程序组成的三个部分

**//1) go语言以包作为管理单位**
   - 每个Go程序都必须属于某个包

**//2) 每个文件必须先声明包**
   - 文件开头要声明该文件属于哪个包

**//3) 程序必须有一个main包（重要）**
   - 每个可执行的Go程序都需要一个 `main` 包
   - 程序的入口函数必须是 `main()` 函数

### 完整的Go程序示例

```go
//1) go语言以包作为管理单位
//2) 每个文件必须先声明包
//3) 程序必须有一个main包（重要）

package main

import "fmt"

//入口函数
func main() { //左括号必须和函数名同行
  //打印
  //"hello go"打印到屏幕，Println()会自动换行
  //调用函数，大都分都需要要导入包
  fmt.Println("hello go")
}
```

### 代码说明

- `package main`：声明当前文件属于 main 包
- `import "fmt"`：导入fmt包（用于打印输出）
- `func main()`：程序的入口函数
- `fmt.Println("hello go")`：调用fmt包的Println函数打印输出，自动换行
- `//`：行注释符号，用于解释代码

---

## 八、 将程序类比于商场

### 程序结构类比

**1) 入口**
   - Go有且只有一个入口函数：`main`
   - liteide：直接图形界面编译，一个文件夹里的文件只能有一个main函数

**2) 干活，调用函数**
   - 通过调用不同的函数来实现程序的功能

### 实际开发流程

**编译和运行示例：**

```bash
# 编译Go代码，生成一个可执行程序
go build xxx.go
# 编译go代码，生成一个可执行程序
# 然后，运行可执行程序 xxx
```

**Windows清屏命令：**
```bash
cls: windows清屏命令
```

**快速运行示例：**
```bash
go run xxx.go  //不生成程序，直接运行
```

### 关键要点总结

- **go build**：生成可执行程序后再运行
- **go run**：直接编译并运行（不生成程序文件）
- **cls**：Windows系统清屏命令

---

## 九、 Go语言变量声明

### 变量的定义和声明

**变量、程序运行期间、可以改变的量**

#### 1) 声明格式：var 变量名 类型
   - 变量声明了，必须要使用
   - 只是声明没有初始化的变量，默认值为0

#### 2) 只是声明没有初始化的变量
   - 默认值为0

#### 3) 同一个 {} 里，声明的变量名是唯一的
   - 相同作用域内不能重复声明

**示例代码：**
```go
var a int
fmt.Println("a =", a)
```

#### 4) 可以同时声明多个变量
   - 支持在一行声明多个变量

**示例代码：**
```go
//var b, c int
```

### 变量的赋值

#### 变量的赋值
   - 先声明，后赋值

**示例代码：**
```go
a = 10  //变量的赋值
fmt.Println("a =", a)
```

### 变量的初始化

#### 变量的初始化、声明变量时，同时赋值

**两种方式：**

**方式1：声明并初始化（一步到位）**
```go
var b int = 10  //初始化，声明变量时，同时赋值（一步到位）
b = 20  //赋值，先声明，后赋值
fmt.Println("b =", b)

var c = 10
```

**方式2：自动推导类型，必须初始化，通过初始化的值确定类型（常用）**
```go
//3、自动推导类型，必须初始化，通过初始化的值确定类型（常用）
c := 30
//％T打印变量所属的类型，必须要用printf
fmt.Printf("c type is %T\n", c)
```

### 变量声明的三种方式总结

| 方式 | 写法 | 说明 |
|------|------|------|
| **声明** | `var a int` | 只声明变量，默认值为0 |
| **初始化** | `var b int = 10` | 声明并赋值 |
| **自动推导** | `c := 30` | 自动推导类型，必须初始化（最常用） |

---

## 十、 变量的交换和多重赋值

### 使用临时变量交换两个变量的值

**传统方法：使用临时变量**

```go
var tmp int
tmp = a
a = b
b = tmp
fmt.Printf("a = %d, b = %d\n", a, b)
```

**说明：** 通过第三个变量 `tmp` 来存储中间值，实现两个变量的交换

### Go语言的简化方式：多重赋值

**Go语言支持多重赋值，可以直接交换：**

```go
i, j := 10, 20
i, j = j, i  //直接交换两个变量
fmt.Printf("i = %d, j = %d\n", i, j)
```

**说明：** Go语言支持在一行代码中同时赋值多个变量，无需临时变量

### 匿名变量

**什么是匿名变量？**
- 去弃数据不处理，匿名变量配合函数返回值使用，才有优势

**匿名变量的特点：**
- 用 `_` 表示匿名变量
- 匿名变量不占用命名空间，不会与其他变量冲突

**示例代码：**

```go
var c, d, e int
c, d, e = test()  //return 1, 2, 3
fmt.Printf("c = %d, d = %d, e = %d\n", c, d, e)

// 使用匿名变量，只接收需要的返回值
_, d, e = test()  //return 1, 2, 3
fmt.Printf("d = %d, e = %d\n", d, e)
```

**说明：**
- 第一次接收函数返回的所有值
- 第二次使用 `_` 作为匿名变量，丢弃第一个返回值，只接收需要的值

---

## 十一、 常量（const）

### 常量的定义

**变量 vs 常量：**
- **变量**：程序运行期间，可以改变的量，变量声明需要用 `var`
- **常量**：程序运行期间，不可以改变的量，常量声明需要用 `const`

### 常量的特点

**1) 常量必须初始化（声明时必须赋值）**
   - 声明常量时，必须同时赋值

**2) 常量一旦赋值，就不能修改**
   - 尝试修改常量会报编译错误

**3) 常量没有使用 `:=` 语法**
   - 常量只能用 `const` 关键字声明
   - 不能使用 `:=` 操作符

### 常量的声明方式

**方式1：声明整型常量**
```go
const a int = 10
//a = 20 //err, 常量不允许修改
fmt.Println("a =", a)
```

**方式2：声明浮点型常量（无需使用 `:=`）**
```go
const b = 11.2  //没有使用:=
fmt.Printf("b type is %T\n", b)
fmt.Println("b =", b)
```

### 常量的特点总结

| 特点 | 说明 |
|------|------|
| **必须初始化** | 声明常量时必须同时赋值 |
| **不可修改** | 常量一旦赋值就不能改变 |
| **无 := 语法** | 常量只能用 `const` 关键字声明 |
| **自动类型推导** | 可以省略类型，由赋值推导 |

---

## 十二、 批量声明变量和常量

### 不同类型变量的声明（定义）

**一个一个声明不同类型的变量：**
```go
// var a int = 1
// var b float64 = 2.0
```

**使用 var() 一次性声明多个不同类型的变量：**
```go
var (
    a int = 1
    b float64 = 2.0
)

a, b = 10, 3.14
fmt.Println("a =", a)
fmt.Println("b =", b)
```

**说明：** 使用 `var()` 语法可以方便地在一个代码块中声明多个不同类型的变量

### 批量声明常量

**一个一个声明常量：**
```go
// const i int = 10
// const j float64 = 3.14
```

**使用 const() 一次性声明多个常量：**
```go
const (
    i int = 10
    j float64 = 3.14
)

fmt.Println("i =", i)
fmt.Println("j =", j)
```

**说明：** 使用 `const()` 语法可以方便地在一个代码块中声明多个常量

### 批量声明的好处

| 方法 | 优点 |
|------|------|
| **单个声明** | 代码清晰但冗长 |
| **批量声明 var()** | 集中管理多个不同类型的变量 |
| **批量声明 const()** | 集中管理多个常量，代码更整洁 |

---

## 十三、 枚举（iota）

### 什么是 iota？

**iota 是 Go 语言中的一个预定义常量**
- 它可以用在常量定义中
- 在每个 const 块中自动递增
- 主要用于生成一组相关的常量值

### iota 的核心特性

**1) iota 只能在 const 中使用**
   - iota 不能在 const 外使用
   - 每次遇到新的 const 块时，iota 重置为 0

**2) iota 在 const 块内自动递增**
   - 第一个常量的 iota 值为 0
   - 后续每一行 iota 值递增 1

**3) 同一行的多个 iota 值相同**
   - 在同一行中使用多个 iota，它们的值都相同

### iota 使用示例

**示例1：基本的 iota 用法**
```go
const (
    a = iota  //0
    b = iota  //1
    c = iota  //2
)
fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
// 输出: a = 0, b = 1, c = 2
```

**示例2：iota 重置**
```go
// 每个新的 const 块，iota 重置为 0
const d = iota  //0
fmt.Printf("d = %d\n", d)
// 输出: d = 0
```

**示例3：隐式赋值（省略后续的 = iota）**
```go
const (
    a1 = iota  //0
    b1         //1，自动继承 = iota
    c1         //2，自动继承 = iota
)
fmt.Printf("a1 = %d, b1 = %d, c1 = %d\n", a1, b1, c1)
// 输出: a1 = 0, b1 = 1, c1 = 2
```

**示例4：同一行的多个 iota**
```go
const (
    i          = iota        //0
    j1, j2, j3 = iota, iota, iota  //1, 1, 1（同行值相同）
    k          = iota        //2
)
fmt.Printf("i = %d, j1 = %d, j2 = %d, j3 = %d, k = %d\n", i, j1, j2, j3, k)
// 输出: i = 0, j1 = 1, j2 = 1, j3 = 1, k = 2
```

### iota 的常见用途

**1) 创建枚举类型**
```go
const (
    Sunday    = iota  //0
    Monday            //1
    Tuesday           //2
    Wednesday         //3
    Thursday          //4
    Friday            //5
    Saturday          //6
)
```

**2) 定义权限级别**
```go
const (
    ReadOnly    = iota  //0
    ReadWrite           //1
    Execute             //2
)
```

### iota 的重要性质总结

| 特性 | 说明 |
|------|------|
| **只在const中使用** | iota 是常量专用的预定义常量 |
| **自动递增** | 每一行自动递增 1 |
| **块级重置** | 新的 const 块重置为 0 |
| **隐式继承** | 省略 = iota 会自动继承表达式 |
| **同行相同** | 同一行的多个 iota 值相同 |

---

## 十四、 Go 语言基础类型详解

### Go 语言内置基础类型概览

Go 语言内置以下这些基础类型：

### 详细类型说明

| 类型 | 名称 | 长度 | 零值 | 说明 |
|------|------|------|------|------|
| **bool** | 布尔类型 | 1 | false | 其值不为真即为假，不可以用数字代表 true 或 false |
| **byte** | 字节类型 | 1 | 0 | uint8 的别名 |
| **rune** | 字符类型 | 4 | 0 | 专用于存储 unicode 编码，等价于 uint32 |
| **int, uint** | 整型 | 4 或 8 | 0 | 32 位或 64 位 |
| **int8, uint8** | 整型 | 1 | 0 | -128 ~ 127, 0 ~ 255 |
| **int16, uint16** | 整型 | 2 | 0 | -32768 ~ 32767, 0 ~ 65535 |
| **int32, uint32** | 整型 | 4 | 0 | -21 亿 ~ 21 亿, 0 ~ 42 亿 |
| **int64, uint64** | 整型 | 8 | 0 | 更大范围 |
| **float32** | 浮点型 | 4 | 0.0 | 小数位精确到 7 位 |
| **float64** | 浮点型 | 8 | 0.0 | 小数位精确到 15 位 |
| **complex64** | 复数类型 | 8 | | 64 位复数 |
| **complex128** | 复数类型 | 16 | | 128 位复数 |
| **uintptr** | 整型 | 4 或 8 | | 足以存储指针的 uint32 或 uint64 整数 |
| **string** | 字符串 | | "" | utf-8 字符串 |

### 类型分类详解

#### 1) 布尔类型（bool）
- 只有两个值：`true` 和 `false`
- 不能用数字代表（不能用 0 和 1）
- 长度为 1 字节

#### 2) 数字类型

**整数类型（Integer Types）：**
- `int, uint`：大小取决于操作系统（32 位或 64 位）
- `int8, int16, int32, int64`：有符号整数
- `uint8, uint16, uint32, uint64`：无符号整数
- `byte`：`uint8` 的别名
- `rune`：用于存储 Unicode 编码，等价于 `int32`

**浮点类型（Floating Point Types）：**
- `float32`：单精度浮点数，精确到 7 位小数
- `float64`：双精度浮点数，精确到 15 位小数（推荐使用）

**复数类型（Complex Types）：**
- `complex64`：由两个 float32 组成
- `complex128`：由两个 float64 组成

#### 3) 字符串类型（String）
- 使用 UTF-8 编码
- 是一种引用类型
- 字符串中的字符是不可变的

#### 4) 指针类型（uintptr）
- 用于存储指针地址
- 大小取决于操作系统（32 位或 64 位）

### 类型使用建议

| 场景 | 推荐类型 | 说明 |
|------|----------|------|
| **整数** | `int` | 除非有特殊需求，一般使用 int |
| **浮点数** | `float64` | 精度更高，推荐使用 |
| **字符** | `rune` | 用于 Unicode 字符 |
| **字节** | `byte` | 用于字节操作 |
| **布尔值** | `bool` | 只用 true/false，不用数字 |

### 零值说明

所有基础类型的零值（未初始化时的默认值）：
- **整数类型**：0
- **浮点类型**：0.0
- **布尔类型**：false
- **字符串类型**：""（空字符串）

---

## 十五、 字符（Character）和字符串（String）

### 字符（Character）

#### 什么是字符？

**字符是指单个的字符**

#### 字符的定义方式

**1) 单引号（''）**
   - 用单引号表示字符
   - 字符，往往都只有一个字符，转义字符除外 `\n`

**示例代码：**
```go
var ch byte
ch = 'a'
fmt.Println("ch =", ch)
```

#### 字符的本质

- **字符在底层就是一个整数**
- 字符可以参与整数运算

**示例代码：**
```go
var ch byte = 'a'
fmt.Println("ch =", ch)  // 输出: ch = 97（'a' 的 ASCII 值）
```

#### 字符类型详解

| 类型 | 说明 | 字节数 |
|------|------|--------|
| **byte** | 无符号整数，等同于 uint8 | 1 |
| **rune** | 用于表示 Unicode 字符，等同于 int32 | 4 |

**选择建议：**
- 普通 ASCII 字符使用 `byte`
- Unicode 字符和更大范围的字符使用 `rune`

### 字符串（String）

#### 什么是字符串？

**字符串是指由若干个字符组成的字符序列**

#### 字符串的定义方式

**1) 双引号（""）**
   - 字符串由 1 个或多个字符组成
   - 字符串都是隐藏了一个结束符 `'\0'`

**示例代码：**
```go
var str string
str = "a"  // 由 'a' 和 '\0' 组成了一个字符串
fmt.Println("str =", str)
```

#### 字符串的操作

**字符串的索引访问（从 0 开始）：**

```go
str = "hello go"
// 只想操作字符串的某个字符，从0开始操作
fmt.Printf("str[0] = %c, str[1] = %c\n", str[0], str[1])
// 输出: str[0] = h, str[1] = e
```

#### 字符串访问细节

- **字符串是不可变的**：不能修改字符串中的单个字符
- **下标访问**：使用 `[]` 访问字符串中的字符
- **下标从 0 开始**：第一个字符的下标是 0
- **格式化输出**：
  - `%s`：打印字符串
  - `%c`：打印字符

### 字符与字符串的区别对比

| 特性 | 字符 | 字符串 |
|------|------|--------|
| **符号** | 单引号 '' | 双引号 "" |
| **组成** | 单个字符 | 1 个或多个字符 |
| **类型** | byte 或 rune | string |
| **本质** | 整数 | 字符序列 |
| **可修改** | 可重新赋值 | 不可修改（不可变） |
| **结束符** | 无 | 隐含 '\0' |

### 实际应用示例

**完整代码示例：**
```go
package main

import "fmt"

func main() {
    // 字符示例
    var ch byte
    ch = 'a'
    fmt.Println("ch =", ch)  // 输出: ch = 97
    
    // 字符串示例
    var str string
    str = "a"  // 由 'a' 和 '\0' 组成
    fmt.Println("str =", str)  // 输出: str = a
    
    // 字符串的字符访问
    str = "hello go"
    fmt.Printf("str[0] = %c, str[1] = %c\n", str[0], str[1])
    // 输出: str[0] = h, str[1] = e
}
```

### 转义字符（Escape Characters）

常用的转义字符：

| 转义字符 | 说明 |
|---------|------|
| `\n` | 换行符 |
| `\t` | 制表符（tab） |
| `\r` | 回车符 |
| `\\` | 反斜杠 |
| `\'` | 单引号 |
| `\"` | 双引号 |

**示例代码：**
```go
str := "hello\nworld"      // 含有换行符
str := "hello\tworld"      // 含有制表符
str := "hello\\world"      // 含有反斜杠
```

---

## 十六、 复数（Complex Number）

### 什么是复数？

**复数是由实部和虚部组成的数字**
- 实部：复数的实数部分
- 虚部：复数的虚数部分，用 `i` 表示（虚数单位）
- 一般形式：`a + bi`，其中 `a` 是实部，`b` 是虚部

### Go 中的复数类型

Go 语言提供两种复数类型：

| 类型 | 字节数 | 说明 |
|------|--------|------|
| **complex64** | 8 | 由两个 float32 组成（实部和虚部各 4 字节） |
| **complex128** | 16 | 由两个 float64 组成（实部和虚部各 8 字节） |

**推荐使用 `complex128`**，精度更高

### 复数的定义和使用

#### 定义复数变量

```go
var t complex128
t = 2.1 + 1.28i
fmt.Println("t===", t)
// 输出: t=== (2.1+1.28i)
```

#### 复数的格式

- **虚数单位**：使用 `i` 表示（不是 `j`）
- **实部和虚部**：直接用 `+` 或 `-` 连接
- **示例**：
  - `2.1 + 1.28i`（正虚部）
  - `3 - 2i`（负虚部）
  - `5i`（纯虚数，实部为 0）

### 复数的操作

#### 提取实部和虚部

Go 提供两个内置函数：

| 函数 | 说明 | 返回类型 |
|------|------|---------|
| **real(c)** | 提取复数的实部 | float（对应 float32 或 float64） |
| **imag(c)** | 提取复数的虚部 | float（对应 float32 或 float64） |

**示例代码：**
```go
var t complex128
t = 2.1 + 1.28i

// 获取实部
fmt.Println("实部=", real(t))  // 输出: 实部= 2.1

// 获取虚部
fmt.Println("虚部=", imag(t))  // 输出: 虚部= 1.28
```

#### 完整示例

```go
package main

import "fmt"

func main() {
    // 定义复数
    var t complex128
    t = 2.1 + 1.28i
    
    // 打印复数
    fmt.Println("复数：", t)
    
    // 提取实部
    fmt.Println("实部=", real(t))
    
    // 提取虚部
    fmt.Println("虚部=", imag(t))
}
```

### 复数的算术运算

复数支持基本的算术运算：

```go
c1 := 1 + 2i
c2 := 3 + 4i

// 加法
c3 := c1 + c2  // (1+3) + (2+4)i = 4 + 6i

// 减法
c4 := c1 - c2  // (1-3) + (2-4)i = -2 - 2i

// 乘法
c5 := c1 * c2  // (1*3 - 2*4) + (1*4 + 2*3)i = -5 + 10i

// 除法
c6 := c1 / c2
```

### 复数的应用场景

- **信号处理**：用于表示和处理复信号
- **傅里叶变换**：频域分析
- **电路分析**：交流电压和电流的计算
- **控制理论**：系统分析和设计
- **数值计算**：复杂的数学计算

---

## 十七、 类型转换（Type Conversion）

### 什么是类型转换？

**类型转换是将一个数据类型的值转换为另一个数据类型的值**
- 在 Go 中，不同类型的变量之间不能直接进行运算
- 必须通过显式的类型转换才能兼容

### Go 的类型转换语法

```go
// 语法格式
newType := TargetType(value)

// 例如
a := 10          // int
b := float64(a)  // 将 int 转换为 float64
```

#### 常见的类型转换

**1) 整数到浮点数**
```go
a := 10
b := float64(a)  // 10 → 10.0
```

**2) 浮点数到整数（会丢失小数部分）**
```go
d := 3.14
a := int(d)      // 3.14 → 3
```

**3) 字符到整数**
```go
c := 'a'
a := int(c)      // 'a' → 97（ASCII 值）
```

**4) 整数到字符**
```go
a := 97
c := byte(a)     // 97 → 'a'
```

### 类型转换的注意事项

- **显式转换**：Go 不支持隐式类型转换，必须显式指定
- **精度损失**：从高精度向低精度转换时可能丢失数据
  - 浮点数转整数会丢失小数部分
  - 大整数转小整数可能溢出
- **运算类型必须一致**：不同类型的数据不能直接运算

### 类型转换示例

```go
package main

import "fmt"

func main() {
    a := 10
    b := "abc"
    c := 'a'
    d := 3.14
    
    // %T 操作变量所属类型
    fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)
    // 输出: int, string, int32, float64
    
    // 类型转换后的输出
    fmt.Printf("a = %d, b = %s, c = %c, d = %f\n", a, b, c, d)
    // 输出: a = 10, b = abc, c = a, d = 3.140000
}
```

---

## 十八、 格式化输出（Formatted Output）

### Printf 的格式化占位符

Go 的 `fmt.Printf()` 支持多种格式化占位符来输出不同类型的数据：

#### 常用的格式化占位符

| 占位符 | 说明 | 示例 | 输出 |
|--------|------|------|------|
| **%T** | 打印变量所属的类型 | `fmt.Printf("%T", 10)` | `int` |
| **%d** | 整型格式 | `fmt.Printf("%d", 10)` | `10` |
| **%s** | 字符串格式 | `fmt.Printf("%s", "abc")` | `abc` |
| **%c** | 字符格式（打印字符对应的字符，不是数值） | `fmt.Printf("%c", 97)` | `a` |
| **%f** | 浮点型格式 | `fmt.Printf("%f", 3.14)` | `3.140000` |
| **%v** | 自动匹配格式输出（通用格式） | `fmt.Printf("%v", value)` | 自动适配 |

### 格式化占位符的详细说明

**1) %T 操作 - 打印类型**
```go
a := 10
fmt.Printf("%T\n", a)  // 输出: int
```

**2) %d - 整型格式**
```go
a := 10
fmt.Printf("%d\n", a)  // 输出: 10
```

**3) %s - 字符串格式**
```go
b := "abc"
fmt.Printf("%s\n", b)  // 输出: abc
```

**4) %c - 字符格式**
```go
c := 'a'
fmt.Printf("%c\n", c)  // 输出: a（打印字符，不是数值）
```

**5) %f - 浮点型格式**
```go
d := 3.14
fmt.Printf("%f\n", d)  // 输出: 3.140000
```

**6) %v - 自动匹配格式（通用）**
```go
a := 10
b := "abc"
c := 'a'
d := 3.14

fmt.Printf("a = %v, b = %v, c = %v, d = %v\n", a, b, c, d)
// 输出: a = 10, b = abc, c = a, d = 3.14
```

### 完整的格式化输出示例

```go
package main

import "fmt"

func main() {
    a := 10
    b := "abc"
    c := 'a'
    d := 3.14
    
    // 方法1：%T 操作变量所属类型
    fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)
    // 输出: int, string, int32, float64
    
    // 方法2：使用具体的格式说明符
    fmt.Printf("a = %d, b = %s, c = %c, d = %f\n", a, b, c, d)
    // 输出: a = 10, b = abc, c = a, d = 3.140000
    
    // 方法3：%v 自动匹配格式输出
    fmt.Printf("a = %v, b = %v, c = %v, d = %v\n", a, b, c, d)
    // 输出: a = 10, b = abc, c = a, d = 3.14
}
```

### Printf vs Println vs Print

| 函数 | 说明 | 示例 |
|------|------|------|
| **Printf** | 格式化输出，不自动换行 | `fmt.Printf("%d\n", 10)` |
| **Println** | 自动换行，不支持格式化 | `fmt.Println("hello")` |
| **Print** | 简单输出，不支持格式化，不自动换行 | `fmt.Print("hello")` |

### 浮点数的精度控制

```go
d := 3.14159

fmt.Printf("%f\n", d)      // 默认精度: 3.141590
fmt.Printf("%.2f\n", d)    // 保留 2 位小数: 3.14
fmt.Printf("%.4f\n", d)    // 保留 4 位小数: 3.1416
```

---

## 十九、 从控制台读取输入（Console Input）

### 什么是控制台输入？

**从键盘读取用户输入的数据，存储到变量中**

### Go 中的输入方法

Go 语言提供两种主要的输入函数：

| 函数 | 说明 | 是否支持格式化 |
|------|------|----------------|
| **fmt.Scan()** | 从标准输入读取数据，以空格或换行符分隔 | 否 |
| **fmt.Scanf()** | 从标准输入按指定格式读取数据 | 是 |

### 1) fmt.Scan() - 简单输入

#### 语法

```go
fmt.Scan(&variable)
```

- **&**：取地址符，获取变量的内存地址
- 数据写入到该地址对应的变量中

#### 使用示例

**单个变量输入：**
```go
var a int
fmt.Printf("请输入a:")
fmt.Scan(&a)
fmt.Println("a===", a)
```

**执行结果：**
```
请输入a:10
a=== 10
```

**多个变量输入：**
```go
var a int
var b string
var c float64

fmt.Println("请依次输入 a, b, c:")
fmt.Scan(&a, &b, &c)

fmt.Println("a=", a, "b=", b, "c=", c)
```

**执行结果：**
```
请依次输入 a, b, c:
10 hello 3.14
a= 10 b= hello c= 3.14
```

### 2) fmt.Scanf() - 格式化输入

#### 语法

```go
fmt.Scanf("format", &variable1, &variable2, ...)
```

- **format**：指定输入的格式
- 格式说明符与 Printf 类似（%d、%s、%f 等）

#### 使用示例

```go
var a int
fmt.Printf("请输入整数:")
fmt.Scanf("%d", &a)
fmt.Println("a===", a)
```

#### Scanf 的格式说明符

| 格式符 | 说明 |
|--------|------|
| **%d** | 整数 |
| **%f** | 浮点数 |
| **%s** | 字符串 |
| **%c** | 字符 |

### 3) 输入的细节问题

#### & 符号的含义

- **&**：取地址操作符
- 用于获取变量在内存中的地址
- 函数需要将数据写入这个地址，所以必须使用 &

**错误示例：**
```go
fmt.Scan(a)  // 错误！缺少 &
```

**正确示例：**
```go
fmt.Scan(&a)  // 正确
```

#### 换行符的影响

- `fmt.Scan()` 会跳过前导的空格和换行符
- 遇到空格或换行符时停止读取当前数据
- 这使得多变量输入很方便

```go
// 可以在同一行输入
fmt.Scan(&a, &b, &c)  // 输入: 10 20 30

// 也可以在不同行输入
fmt.Scan(&a, &b, &c)
// 输入:
// 10
// 20
// 30
```

### 完整的输入示例

```go
package main

import "fmt"

func main() {
    // 示例1：读取整数
    var a int
    fmt.Printf("请输入一个整数: ")
    fmt.Scan(&a)
    fmt.Println("您输入的整数是:", a)
    
    // 示例2：读取字符串
    var name string
    fmt.Printf("请输入您的名字: ")
    fmt.Scan(&name)
    fmt.Println("您的名字是:", name)
    
    // 示例3：读取多个数据
    var x, y int
    fmt.Printf("请输入两个整数 x 和 y: ")
    fmt.Scan(&x, &y)
    fmt.Printf("x = %d, y = %d\n", x, y)
    
    // 示例4：读取浮点数
    var price float64
    fmt.Printf("请输入价格: ")
    fmt.Scan(&price)
    fmt.Printf("价格为: %.2f\n", price)
}
```

### 常见错误总结

| 错误 | 原因 | 解决方案 |
|------|------|---------|
| **忘记 &** | 没有取地址 | 使用 `fmt.Scan(&variable)` |
| **类型不匹配** | 输入数据类型与变量类型不符 | 确保输入数据类型正确 |
| **格式错误** | Scanf 的格式字符串不正确 | 检查格式说明符是否正确 |

### Scan vs Scanf 的选择

| 场景 | 推荐函数 | 说明 |
|------|---------|------|
| **简单输入** | `Scan()` | 不需要指定格式，自动识别 |
| **格式化输入** | `Scanf()` | 需要严格按照格式读取 |
| **多变量** | `Scan()` | 以空格或换行分隔，更灵活 |

---

## 二十、 类型转换兼容性

| 源类型 | 目标类型 | 兼容 | 说明 |
|--------|---------|:----:|------|
| int | float | ✓ | 可转换 |
| float | int | ✓ | 丢失小数部分 |
| byte/rune | int | ✓ | 得到 ASCII 值 |
| int | byte/rune | ✓ | 得到对应字符 |
| bool | 其他 | ✗ | bool 不能与任何类型互转 |
| string | 其他 | ✗ | 需要 strconv 包 |

```go
var flag bool = true
// int(flag)  // 编译错误！bool 不兼容
```

---

## 二十一、 自定义类型（Type Definition）

用 `type` 关键字基于已有类型创建新类型，增强代码语义和类型安全。

```go
type bigint int64   // 单个定义

type (              // 批量定义
    long int64
    char byte
)

var a bigint = 10
var b long = 11
var ch char = 'a'
fmt.Printf("%T\n", a)  // main.bigint
```

**重要：自定义类型与基础类型是不同类型，不能直接互赋值，需显式转换：**

```go
type bigint int64
var a bigint = 10
var b int64 = 20
// a = b  // 编译错误！
a = bigint(b)  // ✓ 显式转换
```

**应用场景：**
```go
type UserId int64      // 业务ID区分
type Temperature float64  // 度量单位区分
// 避免 userId 和 productId 混用
```

| 特性 | 自定义类型 `type A B` | 类型别名 `type A = B` |
|-----|--------------------|--------------------|
| 是否新类型 | 是，类型严格不兼容 | 否，完全兼容 |

---

## 二十二、 选择结构 - switch 语句

### 什么是 switch 语句？

**switch 语句是一种多分支选择结构**
- 根据一个表达式的值，从多个分支中选择一个执行
- 比多个 if-else 更清晰，代码结构更好

### switch 语句的基本语法

```go
switch 表达式 {
case 值1:
    // 当表达式等于值1时执行
case 值2:
    // 当表达式等于值2时执行
case 值3:
    // 当表达式等于值3时执行
default:
    // 都不匹配时执行（可选）
}
```

### switch 语句的执行流程

1. 计算 `switch` 后面表达式的值
2. 依次与各 `case` 的值进行比较
3. 当找到匹配的 `case` 时，执行该 `case` 下的代码
4. 遇到 `break` 或到达分支末尾时，退出 switch
5. 如果没有匹配的 `case`，执行 `default`（如果存在）

### switch 的重要特性

#### 1) 默认有 break 效果

**Go 中的 switch 与 C 不同，每个 case 默认有 break**
- 执行完一个 case 后会自动退出，不会"穿透"到下一个 case
- 无需手动写 break

**示例：**
```go
var num int = 2
switch num {
case 1:
    fmt.Println("一楼")
case 2:
    fmt.Println("二楼")  // 只执行这一行
case 3:
    fmt.Println("三楼")  // 不会执行
}
// 输出: 二楼
```

#### 2) fallthrough - 继续执行下一个 case

**如果需要"穿透"到下一个 case，使用 `fallthrough` 关键字**
- `fallthrough` 会继续执行下一个 case 的代码
- 注意：不会再进行比较，直接执行下一个 case

**示例：**
```go
var num int = 1
switch num {
case 1:
    fmt.Println("一楼")
    fallthrough
case 2:
    fmt.Println("二楼")  // 会执行
    fallthrough
case 3:
    fmt.Println("三楼")  // 也会执行
case 4:
    fmt.Println("四楼")  // 不会执行
}
// 输出:
// 一楼
// 二楼
// 三楼
```

**重要说明：**
- `fallthrough` 必须在 case 的最后
- `fallthrough` 会无条件执行下一个 case（不再进行值比较）
- 可以有多个 `fallthrough`

### 完整的 switch 示例

**电梯楼层例子：**
```go
package main

import "fmt"

func main() {
    // 默认有 break，需要 fallthrough 可以继续执行下面的 case
    var num int
    fmt.Printf("请按下楼层: ")
    fmt.Scan(&num)
    
    switch num {
    case 1:
        fmt.Println("你按下了一楼")
        fallthrough
    case 2:
        fmt.Println("你按下了二楼")
        fallthrough
    case 3:
        fmt.Println("你按下了三楼")
    case 4:
        fmt.Println("你按下了四楼")
    default:
        fmt.Println("你按下了其他楼层")
    }
}
```

**执行结果（按 1 时）：**
```
请按下楼层: 1
你按下了一楼
你按下了二楼
你按下了三楼
```

### switch 的变体

#### 1) 不带表达式的 switch

```go
var score int = 85
switch {
case score >= 90:
    fmt.Println("优秀")
case score >= 80:
    fmt.Println("良好")  // 会执行
case score >= 60:
    fmt.Println("及格")
default:
    fmt.Println("不及格")
}
// 输出: 良好
```

#### 2) switch 多个值

```go
var ch byte = 'a'
switch ch {
case 'a', 'e', 'i', 'o', 'u':
    fmt.Println("元音")
default:
    fmt.Println("辅音")
}
// 输出: 元音
```

### switch vs if-else

| 特性 | switch | if-else |
|------|--------|---------|
| **用途** | 多个确定值的比较 | 条件判断（范围、逻辑） |
| **代码清晰度** | 多分支时清晰 | 复杂逻辑时更灵活 |
| **break** | Go 默认有 break | 不需要 break |
| **默认行为** | 不穿透（需 fallthrough） | 顺序执行 |
| **推荐场景** | 枚举、固定值 | 范围判断、逻辑复杂 |

### switch 的最佳实践

| 实践 | 说明 |
|------|------|
| **谨慎使用 fallthrough** | 虽然功能强大但容易造成混淆 |
| **包含 default** | 除非有特殊原因，通常应该有 default |
| **避免过长** | switch 分支过多时考虑其他方案 |
| **合理分组** | 相关的 case 放在一起，提高可读性 |

---

## 二十四、 循环结构 - for 循环

Go 只有 `for` 循环，没有 `while`，但 for 可以模拟所有循环模式。

### 方法1：传统 for（C 风格）

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}

// 求和
sum := 0
for i := 1; i <= 100; i++ {
    sum += i
}
fmt.Println(sum) // 5050
```

### 方法2：range 循环

```go
str := "abc"

// 获取索引和值
for i, data := range str {
    fmt.Printf("索引:%d, 字符:%c, ASCII:%d\n", i, data, data)
}

// 只用值（忽略索引）
for _, data := range str {
    fmt.Printf("%c ", data) // a b c
}
```

### 方法3：类似 while

```go
i := 1
for i <= 5 {
    fmt.Println(i)
    i++
}
```

### 方法4：无限循环

```go
for {
    fmt.Println("循环")
    break // 用 break 退出
}
```

### 循环控制

```go
// break：退出循环
for i := 1; i <= 10; i++ {
    if i == 5 { break }
    fmt.Println(i) // 1 2 3 4
}

// continue：跳过本次
for i := 1; i <= 5; i++ {
    if i == 3 { continue }
    fmt.Println(i) // 1 2 4 5
}
```

| 方式 | 语法 | 适用场景 |
|-----|-----|---------|
| 传统 for | `for i:=0; i<n; i++` | 固定次数、需要索引 |
| range | `for i, v := range x` | 遍历数组/字符串 |
| 类 while | `for 条件` | 动态条件 |
| 无限循环 | `for { }` | 需要手动控制退出 |

---

## 二十五、 goto 语句和标签

`goto` 是无条件跳转语句，实际中很少使用。**优先使用 break/continue 代替。**

```go
fmt.Println("开始")
goto End          // 跳过中间代码
fmt.Println("跳过")
End:
fmt.Println("结束")
// 输出: 开始 → 结束
```

**主要用途：跳出嵌套循环**

```go
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            goto Exit  // 一次性跳出两层循环
        }
        fmt.Printf("(%d,%d) ", i, j)
    }
}
Exit:
fmt.Println("结束")
// 输出: (0,0) (0,1) (0,2) (1,0) 结束
```

| 规则 | 说明 |
|-----|-----|
| **不能跨函数** | 只能在同一函数内跳转 |
| **break/continue 优先** | 可以用 break/continue 解决的不用 goto |

---

## 二十六、函数传参

### 基本概念

函数传参是将数据传递给函数的机制，Go语言中函数可以接收参数并进行处理。

### 函数定义与调用

```go
package main

import "fmt"

func myFun(a int) {
	fmt.Println("a=", a)
}

func main() {
	myFun(1)
}
```

### 传参方式说明
- 单个参数：`a int` - 参数名和类型 |
- 多个参数：`a int,b int` - 参数名和类型 |
- 参数类型相同 `a,b int` - 相同类型可省略 
- rest参数: 在类型处拓展... func myFun2(args ...int) 
- 将rest参数传递给其他函数的方式
    - myFun3(args...)     //传递全部参数
	- myFun3(args[2:]...) //从args[2]开始传递参数，包含本身
	- myFun3(args[:2]...) // 传递args[0]-args[2]，不包含args[2]

---

## 二十七、函数返回值

Go 函数支持三种返回值写法和多返回值。

```go
// 1. 直接返回
func f1(a int) int { return a + 1 }

// 2. 命名返回值（推荐）
func f2(a int) (result int) {
    result = a + 1
    return  // 自动返回 result
}

// 3. 多返回值（常用于错误处理）
func divide(a, b int) (int, error) {
    if b == 0 { return 0, errors.New("division by zero") }
    return a / b, nil
}

// 多命名返回值
func f3() (a, b, c int) {
    a, b, c = 11, 22, 33
    return
}
```

| 写法 | return 语句 | 适用场景 |
|-----|-----------|---------|
| 直接返回 | `return 表达式` | 简单函数 |
| 命名返回 | `return`（空） | 复杂逻辑，多处赋值 |
| 多返回值 | `return v1, v2` | 返回结果+错误 |

---

## 二十八、回调函数

### 什么是回调函数

**回调函数**：函数有一个参数是函数类型，这个函数就是回调函数。

即：将一个函数作为参数传递给另一个函数，由被调用函数在合适的时机去调用它。

### 自定义函数类型

在 Go 中，可以使用 `type` 关键字定义函数类型：

```go
type FuncType func(int, int) int
```

### 示例：计算器（回调函数实现）

```go
type FuncType func(int, int) int

func Add(a, b int) int {
    return a + b
}

func Minus(a, b int) int {
    return a - b
}

func Mul(a, b int) int {
    return a * b
}

// Calc 是回调函数：fTest 参数是函数类型
func Calc(a, b int, fTest FuncType) (result int) {
    fmt.Println("Calc")
    result = fTest(a, b) // 调用传入的函数
    return
}

func main() {
    a := Calc(1, 1, Mul) // 将 Mul 函数作为参数传入
    fmt.Println("a =", a)
}
```

**输出结果：**
```
Calc
a = 1
```

### 回调函数特点

| 特点 | 说明 |
|-----|-----|
| **函数作为参数** | 将函数像普通变量一样传递 |
| **解耦逻辑** | 调用方与实现方分离 |
| **灵活扩展** | 传入不同函数，实现不同行为 |

---

## 二十九、多态

### 什么是多态

**多态**：多种形态，调用同一个接口，不同的表现，可以实现不同表现（加减乘除）。

Go 语言通过**函数类型**或**接口**实现多态。

### 基于函数类型的多态（回调函数多态）

利用上面的 `Calc` 函数，传入不同的函数实现不同运算：

```go
type FuncType func(int, int) int

func Add(a, b int) int   { return a + b }
func Minus(a, b int) int { return a - b }
func Mul(a, b int) int   { return a * b }

func Calc(a, b int, fTest FuncType) (result int) {
    fmt.Println("Calc")
    result = fTest(a, b)
    return
}

func main() {
    // 同一个接口 Calc，传入不同函数，表现不同 —— 多态
    fmt.Println(Calc(10, 2, Add))   // 加法：12
    fmt.Println(Calc(10, 2, Minus)) // 减法：8
    fmt.Println(Calc(10, 2, Mul))   // 乘法：20
}
```

**输出结果：**
```
Calc
12
Calc
8
Calc
20
```

### 多态核心思想

- **同一接口**：`Calc(a, b int, fTest FuncType)`
- **不同实现**：`Add`、`Minus`、`Mul` 是同一类型 `FuncType` 的不同实现
- **灵活调用**：调用方只需要传入符合类型的函数，无需关心内部实现

### 多态与回调函数的关系

| 概念 | 说明 |
|-----|-----|
| **回调函数** | 技术手段：将函数作为参数传递 |
| **多态** | 设计思想：同一接口，不同表现 |
| **关系** | 回调函数是实现多态的一种方式 |

---

## 三十、闭包

### 什么是闭包

**闭包**：函数的返回值是一个匿名函数（函数类型）。

闭包会捕获外部函数中的变量和常量，**不关心这些捕获的变量和常量是否已经超出了作用域**。只要闭包还在使用它，这些变量就还会存在。

### 示例代码

```go
// 函数的返回值是一个匿名函数，返回一个函数类型
func test02() func() int {
    var x int // 没有初始化，值为 0

    return func() int {
        x++
        return x * x
    }
}

func main() {
    // f 接收返回的匿名函数，f 来调用闭包函数
    // 闭包捕获了 x，只要 f 还在使用，x 就一直存在
    f := test02()
    fmt.Println(f()) // 1   x=1, 1*1
    fmt.Println(f()) // 4   x=2, 2*2
    fmt.Println(f()) // 9   x=3, 3*3
    fmt.Println(f()) // 16  x=4, 4*4
    fmt.Println(f()) // 25  x=5, 5*5
}
```

**输出结果：**
```
1
4
9
16
25
```

### 执行过程分析

| 调用次数 | x 的变化 | 计算 | 结果 |
|---------|---------|-----|-----|
| 第 1 次 | 0 → 1   | 1×1 | 1   |
| 第 2 次 | 1 → 2   | 2×2 | 4   |
| 第 3 次 | 2 → 3   | 3×3 | 9   |
| 第 4 次 | 3 → 4   | 4×4 | 16  |
| 第 5 次 | 4 → 5   | 5×5 | 25  |

### 闭包核心特点

- **捕获外部变量**：匿名函数捕获了外层函数的变量 `x`
- **变量持久存在**：`x` 不会随 `test02()` 执行结束而销毁，只要闭包 `f` 还在使用，`x` 就一直存在
- **状态保留**：每次调用 `f()` 时，`x` 在上次的基础上累加，而不是重新初始化
- **作用域突破**：闭包不关心捕获的变量是否已超出原来的作用域

### 与普通函数对比

```go
// 普通函数：每次调用 x 重新初始化为 0
func test01() int {
    var x int // 函数被调用时，x 才分配空间，初始化为 0
    x++
    return x * x
}

func main() {
    fmt.Println(test01()) // 每次都是 1
    fmt.Println(test01()) // 每次都是 1
}
```

| 对比项 | 普通函数 | 闭包 |
|-------|---------|-----|
| **变量生命周期** | 函数调用结束即销毁 | 与闭包绑定，持续存在 |
| **状态保留** | 每次调用重新初始化 | 保留上次调用后的状态 |
| **适用场景** | 无状态计算 | 需要保留状态的场景 |

---

## 三十一、defer（延迟执行）

### 什么是 defer

`defer` 用于延迟执行一个函数调用，**在当前函数返回前执行**。常用于资源释放、解锁、关闭文件等收尾工作。

### 执行顺序：先进后出（栈）

多个 `defer` 按**压栈顺序**执行，即**后注册的先执行**（LIFO）。

```go
func main() {
    defer fmt.Println("bbb") // 第1个注册
    defer fmt.Println("aaa") // 第2个注册
    defer aa(0)               // 第3个注册
    defer fmt.Println("ccc") // 第4个注册（最后）
}
```

**注册顺序（压栈）：** bbb → aaa → aa(0) → ccc

**执行顺序（出栈）：** ccc → aa(0) → aaa → bbb

### defer 遇到 panic

```go
func aa(a int) {
    result := 100 / a        // a=0 时触发 panic: integer divide by zero
    fmt.Println("result", result)
}

func main() {
    defer fmt.Println("bbb")
    defer fmt.Println("aaa")
    defer aa(0)               // 会触发 panic
    defer fmt.Println("ccc")
}
```

**实际输出：**
```
ccc
aaa
bbb
panic: runtime error: integer divide by zero
```

**关键结论**：`aa(0)` 在执行时触发了 panic，但**其余已注册的 defer（aaa、bbb）仍然继续执行**，最后才抛出 panic 信息。

### defer 执行规则总结

| 规则 | 说明 |
|-----|-----|
| **后进先出** | 多个 defer 按注册的逆序执行（栈结构） |
| **panic 不中断 defer** | 某个 defer 触发 panic，其余已注册的 defer 仍会执行 |
| **函数返回前执行** | defer 在 `return` 执行后、函数真正返回前触发 |
| **常见用途** | 关闭文件、释放锁、捕获 panic（配合 recover）|

---

## 三十二、获取命令行参数

### 使用 os.Args

通过 `os` 包的 `Args` 变量获取命令行参数，类型为 `[]string`。

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    list := os.Args
    n := len(list)
    fmt.Println("The number of args:", n)

    for i := 0; i < n; i++ {
        fmt.Printf("OS Args %d: %s\n", i, list[i])
    }
}
```

### 运行示例

```bash
go build a.go
./a hello world 123
```

**输出：**
```
The number of args: 4
OS Args 0: ./a        # 第0个固定是程序本身的路径
OS Args 1: hello
OS Args 2: world
OS Args 3: 123
```

### 注意事项

| 说明 | 内容 |
|-----|-----|
| **`os.Args[0]`** | 固定是程序本身的路径，不是用户传入的参数 |
| **用户参数从 `[1]` 开始** | `os.Args[1:]` 才是真正传入的参数 |
| **macOS/Linux** | `go build a.go` 生成无后缀可执行文件 `a` |
| **Windows** | `go build a.go` 生成 `a.exe` |

---

## 三十三、导入包的方式

Go 支持多种 `import` 写法，满足不同使用场景。

```go
import (
    // 第一种：点导入，省略包名直接调用
    . "fmt"

    // 第二种：起别名
    o1 "os"

    // 第三种：忽略包（只执行包的 init 函数，不直接使用）
    _ "fmt"
)

func main() {
    Println("hello world") // 点导入后，直接用 Println，无需 fmt.Println
    o1.Exit(1)             // 别名调用，用 o1 代替 os
}
```

### 三种导入方式对比

| 方式 | 语法 | 调用方式 | 说明 |
|-----|-----|---------|-----|
| **普通导入** | `"fmt"` | `fmt.Println()` | 最常用，包名作前缀 |
| **点导入** | `. "fmt"` | `Println()` | 省略包名，直接调用，不推荐（易冲突） |
| **别名导入** | `o1 "os"` | `o1.Exit()` | 自定义包名，用于包名冲突或简化长包名 |
| **忽略导入** | `_ "pkg"` | 无法调用 | 只触发包的 `init()` 函数，常用于驱动注册 |

### _ 操作补充说明

有时需要导入一个包，但不需要引用这个包的任何标识符。可以用空白标识符 `_` 来重命名这个导入：

```go
import (
    _ "fmt"
)
```

**`_` 操作的本质**：引入该包，但不直接使用包里的函数，而是**调用了该包里面的 `init` 函数**。

常见场景：数据库驱动注册

```go
import (
    _ "github.com/go-sql-driver/mysql"  // 只需触发 mysql 驱动的 init() 完成注册
)
```

---

## 三十四、工程管理（分文件编程）

### 规则

1. 分文件编程（多个源文件），必须放在 `src` 目录下
2. 设置 `GOPATH` 环境变量指向项目根目录
3. **同一个目录，包名必须一样**
4. 用 `go env` 查看 Go 相关的环境路径
5. **同一个目录，调用别的文件的函数，直接调用即可，无需包名引用**

### 目录结构示例

```
src/
├── main.go
└── test.go
```

### 代码示例

**main.go**
```go
package main  // 必须是 main 包

func main() {
    test()  // 直接调用同目录 test.go 中的函数，无需包名
}
```

**test.go**
```go
package main  // 与 main.go 同目录，包名必须一样

import "fmt"

func test() {
    fmt.Println("this is a test func")
}
```

### 运行方式

```bash
# 运行同目录下所有 .go 文件
go run .

# 或指定多个文件
go run main.go test.go
```

### 关键规则总结

| 规则 | 说明 |
|-----|-----|
| **src 目录** | 多文件项目源码放在 src 目录 |
| **GOPATH** | 设置为项目根目录，`go env` 可查看当前值 |
| **同目录包名相同** | 同一文件夹下所有 `.go` 文件的 `package` 名必须一致 |
| **跨文件调用** | 同包内直接调用函数，不需要 `包名.函数名` |

---

## 三十五、不同目录包的调用

### 规则

1. **不同目录，包名不一样**（子目录用自己的包名，如 `package calc`）
2. **调用不同包的函数**，格式：`包名.函数名()`
3. **函数首字母必须大写才能被外部包调用**；首字母小写为包内私有，外部无法访问

### 目录结构

```
src/
├── main.go          → package main
└── calc/
    └── calc.go      → package calc
```

### 代码示例

**src/calc/calc.go**
```go
package calc  // 包名与目录名一致（规范）

// 首字母大写 → 可被外部包调用（导出）
func Add(x, y int) int {
    return x + y
}

// 首字母小写 → 仅包内可用（私有）
func minus(x, y int) int {
    return x - y
}
```

**src/main.go**
```go
package main

import (
    "fmt"
    "go-learn/src/calc"  // 模块名 + 包路径
)

func main() {
    fmt.Println(calc.Add(1, 2))   // ✓ 首字母大写，可以调用
    // fmt.Println(calc.minus(3,1)) // ✗ 首字母小写，编译报错
}
```

### 对比：同目录 vs 不同目录

| 对比项 | 同目录（同包） | 不同目录（不同包） |
|-------|------------|---------------|
| **包名** | 必须相同 | 各自独立 |
| **调用方式** | 直接调用 `func()` | `包名.Func()` |
| **可见性限制** | 无限制 | 首字母大写才可访问 |
| **是否需要 import** | 不需要 | 需要 |

---

## 三十六、init 函数与包的初始化顺序

### init 函数特点

- 每个包可以定义 `init()` 函数，**自动执行，无需手动调用**
- 用于包的初始化工作（初始化变量、注册驱动等）

### 包的初始化顺序

每个包内部按以下顺序执行：

```
const（常量）→ var（变量）→ init()
```

### 多包导入时的执行流程

根据依赖关系，**被依赖的包先初始化**，最后执行 `main()`：

```
main
 └── import pkg1
       └── import pkg2
             └── import pkg3
                   └── pkg3: const → var → init()
             └── pkg2: const → var → init()
       └── pkg1: const → var → init()
 └── main: const → var → init() → main()
 └── Exit
```

### 目录结构示例

```
src/
├── main.go
└── test/
    └── test.go
```

### 代码示例

**src/test/test.go**
```go
package test

import "fmt"

func init() {
    fmt.Println("test init()")  // 自动执行，早于 main()
}

func Hello() {
    fmt.Println("Hello from test pkg")
}
```

**src/main.go**
```go
package main

import (
    "fmt"
    "go-learn/src/test"
)

func init() {
    fmt.Println("main init()")
}

func main() {
    fmt.Println("main()")
    test.Hello()
}
```

**输出结果：**
```
test init()   ← 被依赖的包先执行 init
main init()   ← 然后执行 main 包的 init
main()        ← 最后执行 main()
Hello from test pkg
```

### 总结

| 执行阶段 | 顺序 | 说明 |
|---------|-----|-----|
| **依赖包优先** | 最先 | 最深层依赖的包最先初始化 |
| **const / var** | 包内最先 | 常量和变量先于 init 初始化 |
| **init()** | 自动调用 | 无需手动调用，不可被外部调用 |
| **main()** | 最后 | 所有包初始化完成后才执行 |

---

## 三十七、指针

### 变量的两层含义

每个变量有两层含义：
- **变量的内存**：存储的值
- **变量的地址**：变量在内存中的位置（用 `&` 取地址）

### 指针类型

- `*int`：保存 `int` 类型变量的地址
- `**int`：保存 `*int` 类型的地址（二级指针）

### 示例代码

```go
func main() {
    var a int = 10
    fmt.Printf("a = %d\n", a)   // 变量的值
    fmt.Printf("&a = %v\n", &a) // 变量的地址

    var p *int
    p = &a  // 指针变量指向谁，就把谁的地址赋值给指针变量

    fmt.Printf("p = %v, &a = %v\n", p, &a)
    *p = 666 // *p 操作的不是 p 的内存，是 p 所指向的内存（就是 a）
    fmt.Printf("*p = %v, a = %v\n", *p, a)
}
```

**输出结果：**
```
a = 10
&a = 0xc000018098
p = 0xc000018098    ← p 存的值就是 a 的地址
*p = 666, a = 666   ← 通过 *p 修改了 a 的值
```

### 核心操作

| 操作 | 含义 | 示例 |
|-----|-----|-----|
| `&a` | 取变量 a 的地址 | `p = &a` |
| `var p *int` | 声明指针变量，存 int 地址 | — |
| `p = &a` | 让指针 p 指向 a | — |
| `*p` | 解引用，访问 p 指向的内存（即 a） | `*p = 666` |

### 合法指向

指针使用前必须指向合法内存，否则触发 panic：

```go
var p *int
p = nil       // 默认 nil，没有合法指向
// *p = 666  // err！触发 panic

var a int
p = &a        // 有了合法指向
*p = 666      // ✓
```

### new 函数

`new(T)` 分配匿名内存并清零，返回 `*T`，无需手动管理生命周期：

```go
p1 := new(int)   // p1 为 *int，初始值为 0
*p1 = 111
```

| 方式 | 写法 | 说明 |
|-----|-----|-----|
| 取已有变量地址 | `p = &a` | 需要先声明变量 `a` |
| new 分配 | `p = new(int)` | 直接分配匿名内存，更简洁 |

---

## 三十八、数组

### 定义数组

```go
var a [10]int  // 长度为 10 的 int 数组
var b [5]int   // 长度为 5 的 int 数组
```

> `[10]int` 和 `[5]int` 是**不同类型**，数组长度是类型的一部分。

### 注意事项

```go
// n := 10
// var c [n]int  // err: non-constant array bound n
```

数组长度必须是**常量**，不能是变量。

### 操作数组元素

下标从 `0` 开始，到 `len()-1` 结束，可以是变量或常量：

```go
a[0] = 1    // 常量下标
i := 1
a[i] = 2    // 变量下标，a[1] = 2
```

### 遍历赋值与打印

```go
for i := 0; i < len(a); i++ {
    a[i] = i + 1
}
for i := 0; i < len(a); i++ {
    fmt.Printf("a[%d] = %d\n", i, a[i])
}
```

### 数组特点总结

| 特点 | 说明 |
|-----|-----|
| **长度固定** | 定义后长度不可变 |
| **长度是类型的一部分** | `[10]int` 与 `[5]int` 不同类型，不可互相赋值 |
| **下标从 0 开始** | 有效范围 `0 ~ len()-1`，越界会 panic |
| **长度必须是常量** | 不能用变量作为数组长度 |
| **len() 获取长度** | `len(a)` 返回数组元素个数 |

### 数组初始化方式

```go
// 完整初始化
var a [5]int = [5]int{1, 2, 3, 4, 5}

// 短变量声明
b := [5]int{6, 7, 8, 9, 10}

// 切片（不固定长度，区别于数组）
c := []int{1, 2, 3, 4, 5}

// 部分初始化，未指定的元素默认为 0
d := [5]int{1, 2, 3}  // [1 2 3 0 0]
```

### 二维数组

```go
// [行数][列数]类型
e := [3][4]int{1: {5, 6, 7, 8}}
fmt.Println("e=", e)
// 输出: e= [[0 0 0 0] [5 6 7 8] [0 0 0 0]]
```

- `[3][4]int`：3 行 4 列的二维数组
- `{1: {5, 6, 7, 8}}`：只初始化第 1 行（下标从 0 开始），其余行默认为 0
- 未初始化的行自动填充零值

### 数组的比较与赋值

```go
a := [5]int{1, 2, 3, 4, 5}
b := [5]int{1, 2, 3, 4, 5}
c := [5]int{1, 2, 3}

fmt.Println("a == b", a == b) // true，每个元素都相同
fmt.Println("a == c", a == c) // false，元素不同

// 同类型的数组可以直接赋值
var d [5]int
d = a
fmt.Println("d =", d) // [1 2 3 4 5]
```

**规则：**
- 数组只支持 `==` 和 `!=` 比较，比较的是**每个元素是否都相同**
- 参与比较的两个数组**类型必须一样**（长度+元素类型均相同）
- 同类型数组之间可以直接赋值（值拷贝）

### 数组作为函数参数（值传递）

数组作为参数传递给函数时，是**值传递**，函数内修改不影响原数组：

```go
func modify(a [5]int) {
    a[1] = 100  // 修改的是副本
}

func main() {
    a := [5]int{1, 2, 3, 4, 5}
    modify(a)
    fmt.Println(a) // [1 2 3 4 5]，原数组不变
}
```

> 若需要在函数内修改原数组，应传递**指针** `*[5]int` 或使用**切片**。

---

## 四十、切片（Slice）

切片是对数组的引用，**长度可变**，比数组更灵活。

### 从数组创建切片

语法：`array[low:high:max]`

- `low`：起始下标（包含）
- `high`：结束下标（不包含）
- `max`：容量上限下标（可省略）
- **len**（长度）= `high - low`
- **cap**（容量）= `max - low`

```go
a := [5]int{1, 2, 3, 4, 5}
b := a[0:3:5]

fmt.Println("b len", len(b)) // 3  （0~2，共3个元素）
fmt.Println("b cap", cap(b)) // 5  （容量从0到max=5）
fmt.Println(b)               // [1 2 3]
```

### 切片截取方式

| 操作 | 含义 |
|-----|-----|
| `s[n]` | 索引位置为 n 的元素 |
| `s[:]` | 从 0 到 len(s)-1，等同于完整切片 |
| `s[low:]` | 从 low 到 len(s)-1 |
| `s[:high]` | 从 0 到 high，len=high |
| `s[low:high]` | 从 low 到 high，len=high-low |
| `s[low:high:max]` | 从 low 到 high，len=high-low，cap=max-low |
| `len(s)` | 切片长度，总是 <= cap(s) |
| `cap(s)` | 切片容量，总是 >= len(s) |

### make 创建切片

```go
s1 := make([]int, 5, 10)   // len=5, cap=10，元素初始化为0
fmt.Println(s1)             // [0 0 0 0 0]
s1 = append(s1, 11)        // 追加元素
fmt.Println(s1)             // [0 0 0 0 0 11]，len变为6
```

### append 扩容规则

当 `append` 追加元素导致 len 超过 cap 时，Go 会自动扩容，**新 cap = 旧 cap × 2**：

```go
var s []int
for i := 0; i < 10; i++ {
    s = append(s, i)
    fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
}
// len=1  cap=1
// len=2  cap=2
// len=3  cap=4   ← cap不够，扩容为2倍
// len=4  cap=4
// len=5  cap=8   ← cap不够，扩容为2倍
// len=6  cap=8
// len=7  cap=8
// len=8  cap=8
// len=9  cap=16  ← cap不够，扩容为2倍
// len=10 cap=16
```

**注意：** 扩容时会分配新的底层数组，`append` 返回的切片与原切片**不再共享**底层数组，所以必须用返回值 `s = append(s, ...)`。

### copy 用法

`copy(dst, src)` 将 src 的元素复制到 dst，**复制数量 = min(len(dst), len(src))**，两个切片底层数组独立互不影响：

```go
a := []int{1, 2}
b := []int{6, 6, 6, 6, 6}
copy(b, a)
fmt.Println(b) // [1 2 6 6 6]  ← 只覆盖了前2个，a只有2个元素
```

| 对比 | 说明 |
|-----|-----|
| `b = a` | b 和 a 共享底层数组，修改互相影响 |
| `copy(b, a)` | 独立拷贝，修改互不影响 |

---

## 四十一、map

map 是键值对集合，key 和 value 类型各自指定。

### 声明与初始化

```go
// 方式1：声明（nil map，不能直接赋值）
var m1 map[int]string
fmt.Println(m1) // map[]

// 方式2：make 创建
m2 := make(map[int]string)

// 方式3：make 指定初始容量（可超过，会自动扩容）
m3 := make(map[int]string, 2)
m3[1] = "a"
m3[2] = "b"
m3[3] = "c"
```

### 遍历

```go
for key, value := range m3 {
    fmt.Printf("%d====>%s\n", key, value)
}
// map 是无序的，每次遍历顺序可能不同
```

### 判断 key 是否存在

```go
value, ok := m3[1]
if ok {
    fmt.Println("存在:", value)
} else {
    fmt.Println("不存在")
}
```

### 删除元素

```go
delete(m3, 1) // 删除 key=1 的元素
```

### map 是引用传递

map 作为函数参数时是**引用传递**，函数内修改会影响原 map：

```go
func test(m map[int]string) {
    delete(m, 1) // 删除的是原 map 的数据
}

test(m3)
fmt.Println(m3) // key=1 已被删除
```

> 与数组（值传递）不同，map 传参后共享同一份数据。

### 常用操作总结

| 操作 | 写法 |
|-----|-----|
| 创建 | `make(map[K]V)` |
| 赋值 | `m[key] = value` |
| 取值 | `v := m[key]` |
| 判断存在 | `v, ok := m[key]` |
| 删除 | `delete(m, key)` |
| 遍历 | `for k, v := range m` |

---

## 四十二、结构体

结构体是将多个不同类型的字段组合成一个自定义类型。

### 定义与初始化

```go
type Student struct {
    id   int
    name string
    sex  byte
    age  int
    addr string
}

// 顺序初始化（需按字段顺序）
s1 := Student{1, "xiaoming", 'm', 18, "bj"}

// 指定字段初始化（未指定字段自动为零值）
s2 := Student{name: "xiaohong", addr: "sh"}
```

### 结构体指针

```go
// 方式1：取地址
p1 := &Student{name: "mike", addr: "sh"}

// 方式2：先声明再取地址
var s Student
var p *Student = &s
p.id = 1          // p.id 与 (*p).id 完全等价
(*p).name = "mike"

// 方式3：new
p3 := new(Student)
p3.id = 1
p3.name = "linux"
```

### 可见性规则

- 字段名**首字母大写**：可被其他包访问（导出）
- 字段名**首字母小写**：仅限同一包内访问

---

## 四十三、匿名字段（结构体嵌套）

匿名字段：只写类型名，不写字段名，实现结构体的组合复用。

```go
type Person struct {
    name string
    age  int
}

type Student struct {
    Person        // 匿名字段，嵌入 Person
    school string
    id     int
}

// 自定义类型也可作为匿名字段
type mystr string
type Student2 struct {
    Person
    int
    mystr
}
```

### 初始化与访问

```go
s1 := Student{Person{"mike", 20}, "qinghua", 123}
fmt.Println(s1)

// %+v 显示字段名
fmt.Printf("%+v\n", s1) // {Person:{name:mike age:20} school:qinghua id:123}

// 指定字段初始化
s3 := Student{id: 1}
s4 := Student{Person: Person{name: "mike"}}

// 访问匿名字段的成员（直接访问，无需 s.Person.name）
fmt.Println(s1.name) // mike
fmt.Println(s1.Person.name) // 等价写法
```

---

## 四十四、面向过程 vs 面向对象函数

Go 通过**为类型添加方法**实现面向对象风格。

```go
// 面向过程：普通函数
func add1(a, b int) int {
    return a + b
}

// 面向对象：为类型绑定方法
type long int

func (tmp long) Add2(other long) long {
    return tmp + other
}

func main() {
    // 过程式调用
    result := add1(1, 2)

    // 对象式调用
    var n long = 2
    r := n.Add2(1) // n 作为接收者
}
```

**区别：** 方法有接收者 `(tmp long)`，将函数与类型绑定，通过 `变量.方法名()` 调用。

---

## 四十五、为结构体添加方法

### 值接收者（不修改原数据）

```go
func (tmp Person) PrintInfo() {
    fmt.Println(tmp)
}
```

### 指针接收者（修改原数据）

```go
func (p *Person) SetInfo(n string, s byte, a int) {
    p.name = n
    p.age = a
    p.sex = s
}
```

### 示例

```go
type Person struct {
    name string
    age  int
    sex  byte
}

p := Person{"张三", 20, 'm'}
p.PrintInfo()              // 打印原始信息

p.SetInfo("李四", 22, 'f') // 修改字段
p.PrintInfo()              // 打印修改后信息
```

### 值接收者 vs 指针接收者

| | 值接收者 `(t T)` | 指针接收者 `(t *T)` |
|--|----------------|------------------|
| **是否修改原数据** | 否（操作副本） | 是（操作原对象） |
| **适用场景** | 只读操作 | 需要修改字段 |

### 从切片生成切片 & 底层数组关系

**所有切片共享同一底层数组**，从切片再切片时，下标仍基于底层数组偏移：

```go
a := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

s3 := a[2:5]      // s3 = [3 4 5]，len=3，cap=7（底层数组从a[2]开始还有7个）
fmt.Println(s3[1]) // 4（s3的第1个元素，即a[3]）

s4 := s3[2:7]     // 从s3底层偏移继续截取，s4 = a[4:9] = [5 6 7 8 9]
fmt.Println(s4)    // [5 6 7 8 9]
```

底层数组示意图：

```
index:  0  1  2  3  4  5  6  7  8
  a:  [ 1  2  3  4  5  6  7  8  9 ]
              |--s3--|
                        |----s4----|
```

**关键规则：**
- 切片不拥有数据，只是底层数组的一个视图
- 修改切片元素会影响底层数组（及所有共享该数组的切片）
- `s3[2:7]` 的 high=7 是相对于 **s3 的底层数组起点**（a[2]），即 a[2+2 : 2+7] = a[4:9]

### 切片与数组的区别

| 对比 | 数组 | 切片 |
|-----|-----|-----|
| **长度** | 固定，类型的一部分 | 可变 |
| **传参** | 值传递（拷贝） | 引用传递（共享底层数组） |
| **声明** | `[5]int` | `[]int` 或 `a[0:3]` |

---

## 四十六、接口（interface）

接口是一组方法声明的集合。只要某个类型实现了接口中的所有方法，就算实现该接口。

```go
type Humaner interface {
    sayHi()
}
```

### 1) 接口的多态

同一个接口变量可以保存不同具体类型，只要它们都实现了接口方法。

```go
type Student struct {
    name string
    age  int
}

func (s *Student) sayHi() {
    fmt.Println("student say hi", s)
}

type Teacher struct {
    name    string
    address string
}

func (t *Teacher) sayHi() {
    fmt.Println("teacher say hi", t)
}

type MyStr string

func (m *MyStr) sayHi() {
    fmt.Printf("mystr say hi %s\n", *m)
}

func main() {
    var i Humaner

    s := &Student{"mike", 6}
    i = s
    i.sayHi()

    t := &Teacher{"bj", "go"}
    i = t
    i.sayHi()

    var str MyStr = "hello mike"
    i = &str
    i.sayHi()
}
```

要点：
- 接口变量本身不关心具体类型，只关心是否实现了接口方法。
- 这里 `sayHi()` 都是指针接收者，因此赋值给接口时要传指针。

### 2) 空接口 `interface{}`

空接口没有任何方法，因此任意类型都实现了空接口。它可以保存任意类型的数据。

```go
func aaa(arg ...interface{}) {
    // 可接收任意数量、任意类型参数
}

func main() {
    var i interface{} = 1
    fmt.Println("i=", i)

    i = "abc"
    fmt.Println("i=", i)
}
```

要点：
- `interface{}` 类似“通用容器”，可存放任意类型。
- 常用于：不确定参数类型的函数、可变参数、通用数据结构。

### 3) 总结

- 普通接口：用于定义行为规范（方法集合）。
- 空接口：用于接收任意类型值。
- 接口是 Go 实现多态的核心机制之一。

---

## 三十九、随机数

### 使用 math/rand 包

```go
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // 设置种子，只需设置一次
    // 如果种子参数一样，每次运行产生的随机数都一样
    rand.Seed(time.Now().UnixNano()) // 以当前时间纳秒作为种子，保证每次不同

    for i := 0; i < 5; i++ {
        // fmt.Println(rand.Int())      // 很大的随机数
        fmt.Println(rand.Intn(100))  // [0, 100) 范围内的随机整数
    }
}
```

### 常用函数

| 函数 | 说明 |
|-----|-----|
| `rand.Seed(n)` | 设置随机种子，程序启动时调用一次 |
| `rand.Int()` | 生成一个很大的非负随机整数 |
| `rand.Intn(n)` | 生成 `[0, n)` 范围内的随机整数 |

### 关键点

- **种子相同 → 随机数序列相同**：不设置种子或种子固定，每次运行结果一样
- **推荐种子**：`time.Now().UnixNano()`，以纳秒时间戳保证每次运行不同
- `rand.Intn(100)` 生成 0~99 的随机数（不含100）

---

## 30. error 接口的使用

### error 接口

Go 中 `error` 是一个内置接口，任何实现了 `Error() string` 方法的类型都满足该接口。

### 创建 error 的两种方式

**方式一：`fmt.Errorf()`**

```go
err := fmt.Errorf("%s", "this is normal err1")
fmt.Println("err=", err)
```

**方式二：`errors.New()`**

```go
import "errors"

err := errors.New("分母不能为0")
```

### 错误处理示例

```go
func myDiv(a, b int) (result int, err error) {
    err = nil
    if b == 0 {
        err = errors.New("分母不能为0")
    } else {
        result = a / b
    }
    return
}

func main() {
    result, err := myDiv(1, 0)
    if err != nil {
        fmt.Println("err=", err)
    } else {
        fmt.Println("result=", result)
    }
}
```

### 关键点

- 函数通常将 `error` 作为最后一个返回值
- 调用方通过 `if err != nil` 判断是否出错
- `fmt.Errorf` 支持格式化，`errors.New` 适合简单错误信息

---

## 31. panic 与 recover

### panic

`panic` 用于触发运行时错误（如数组越界、手动触发），程序会终止并打印堆栈信息。

```go
// 数组越界会自动触发 panic
var a [10]int
a[11] = 111 // panic: runtime error: index out of range
```

### recover

`recover` 用于捕获 `panic`，**必须配合 `defer` 使用**，否则无效。

```go
func testa(x int) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()
    var a [10]int
    if x < 0 || x >= len(a) {
        panic("index out of range")
    }
    a[x] = 111
}

func main() {
    testa(5)
}
```

### 关键点

- `recover()` 只在 `defer` 函数中有效
- `recover()` 返回 `nil` 表示没有发生 panic
- 捕获 panic 后程序可以继续运行，不会崩溃
- 建议在 `recover` 前主动做边界检查，避免依赖 panic 作为控制流

---

## 32. 字符串操作

### strings 包

```go
import "strings"

// 是否包含子串
strings.Contains("hellogo", "hello")   // true

// 连接切片
s := []string{"abc", "hello", "mike", "go"}
strings.Join(s, "x")                   // "abcxhelloxmikexgo"

// 查找子串位置（-1 表示不存在）
strings.Index("helloabc", "abc")       // 5

// 重复字符串
strings.Repeat("go", 3)               // "gogogo"

// 分割字符串
strings.Split("hello@abc@go", "@")    // ["hello", "abc", "go"]

// 去除首尾指定字符
strings.Trim("  hello  ", " ")        // "hello"

// 按空白符分割成单词切片
strings.Fields("  are you ok")        // ["are", "you", "ok"]
```

### strconv 包（类型转换）

```go
import "strconv"

// 整型 → 字符串（最常用）
str := strconv.Itoa(6666)               // "6666"

// 字符串 → 整型
a, _ := strconv.Atoi("567")             // 567

// 字符串 → bool
flag, _ := strconv.ParseBool("true")    // true

// 其他格式化转换
strconv.FormatBool(false)               // "false"
strconv.FormatFloat(3.14, 'f', -1, 64) // "3.14"

// 向 []byte 追加转换后的值
slice := make([]byte, 0, 1024)
slice = strconv.AppendBool(slice, true)
slice = strconv.AppendInt(slice, 1234, 10)
slice = strconv.AppendQuote(slice, "hello")
fmt.Println(string(slice)) // true1234"hello"
```

### 常用函数速查

| 函数 | 说明 |
|------|------|
| `strings.Contains(s, sub)` | 是否包含子串 |
| `strings.Join(slice, sep)` | 切片拼接为字符串 |
| `strings.Split(s, sep)` | 字符串分割为切片 |
| `strings.Index(s, sub)` | 子串首次出现位置 |
| `strings.Repeat(s, n)` | 重复字符串 n 次 |
| `strings.Trim(s, cutset)` | 去除首尾指定字符 |
| `strings.Fields(s)` | 按空白符分割 |
| `strconv.Itoa(n)` | int → string |
| `strconv.Atoi(s)` | string → int |
| `strconv.FormatFloat(f, fmt, prec, bitSize)` | float → string |

---

## 33. 正则表达式

### 基本匹配

```go
buf := "abc azc a7c aac 888 a9c tac"
reg1 := regexp.MustCompile(`a.c`)
re := reg1.FindAllStringSubmatch(buf, -1)
fmt.Println("result=", re) // [abc azc a7c aac a9c]
```

- `.` 表示匹配任意单个字符
- `FindAllStringSubmatch(..., -1)` 表示匹配全部结果

### 数字小数匹配示例

```go
buf := "3.14 567 agsdg 1.23 7. 8.9 1sdljgl 6.66 7.8 "
reg := regexp.MustCompile(`\d\.\d+`)
re := reg.FindAllStringSubmatch(buf, -1)
fmt.Println("re====", re) // [3.14 1.23 8.9 6.66 7.8]
```

- `\d` 表示数字
- `\d+` 表示一个或多个数字
- `\d\.\d+` 可用于提取形如 `x.xx` 的小数片段

---

## 34. JSON 结构体生成（结构体 -> JSON）

### 结构体转 JSON

```go
type IT struct {
    Company  string
    Subjects []string
    IsOK     bool
    Price    float64
}

s := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}
buf, _ := json.Marshal(s)
fmt.Println(string(buf))
```

### 美化输出

```go
buf2, _ := json.MarshalIndent(s, "", " ")
fmt.Println(string(buf2))
```

- 第二个参数是每行前缀（这里为空）
- 第三个参数是层级缩进字符串（这里为一个空格）

### 结构体字段标签（tag）

```go
type IT struct {
    Company  string   `json:"-"`
    Subjects []string `json:"subjects"`
    IsOK     bool     `json:",string"`
    Price    float64  `json:",string"`
}
```

- `json:"-"`：字段不参与 JSON 编解码
- `json:"subjects"`：指定 JSON 字段名
- `json:",string"`：以字符串形式编码/解码

### map 转 JSON

```go
m := make(map[string]interface{}, 4)
m["c"] = "hh"
m["s"] = []string{"Go", "C++", "python"}
m["ok"] = true
m["p"] = 55

result, _ := json.Marshal(m)
fmt.Println(string(result))
```

---

## 35. JSON 生成结构体（JSON -> 结构体）

### 反序列化示例

```go
type IT struct {
    Company  string   `json:"-"`
    Subjects []string `json:"subjects"`
    IsOK     bool     `json:",string"`
    Price    float64  `json:",string"`
}

jsonBuf := `
{
  "Company": "itcast",
  "Subjects": ["Go", "C++", "Python", "Test"],
  "IsOK": true,
  "Price": 666.666
}`

var tmp IT
json.Unmarshal([]byte(jsonBuf), &tmp)
fmt.Printf("tmp=%+v\n", tmp)
```

### 关键点

- 反序列化目标必须传指针：`&tmp`
- 字段首字母必须大写，反射才能访问
- tag 会影响 JSON 字段映射和编解码行为

---

## 36. 文件输入输出

### 控制台输入

```go
fmt.Println("请输入a:")
var a int
fmt.Scan(&a)
```

### 按行读取文件（`bufio.Reader`）

```go
f, _ := os.Open(path)
defer f.Close()
r := bufio.NewReader(f)
for {
    buf, err := r.ReadBytes('\n')
    if err != nil {
        if err == io.EOF {
            break
        }
    }
    fmt.Println(string(buf))
}
```

### 写文件 + 读文件

```go
f, err := os.Create(path)
if err != nil {
    return
}
defer f.Close()

for i := 0; i < 10; i++ {
    buf := fmt.Sprintf("i=%d\n", i)
    f.WriteString(buf)
}
```

```go
f, _ := os.Open(path)
defer f.Close()
buf := make([]byte, 1024*2)
n, _ := f.Read(buf)
fmt.Println(string(buf[:n]))
```

### 文件拷贝（命令行参数）

```go
list := os.Args
if len(list) != 3 {
    fmt.Println("输入有误，用法: copyFile <源文件> <目标文件>")
    return
}

srcFile := list[1]
dstFile := list[2]
if srcFile == dstFile {
    fmt.Println("源文件和目的文件名字不能相同")
    return
}

sf, err1 := os.Open(srcFile)
if err1 != nil {
    return
}
defer sf.Close()

df, err2 := os.Create(dstFile)
if err2 != nil {
    return
}
defer df.Close()

buf := make([]byte, 4*1024)
for {
    n, err := sf.Read(buf)
    if err != nil {
        if err == io.EOF {
            break
        }
    }
    df.Write(buf[:n])
}
```

### 关键点

- 命令行参数拷贝需传两个路径参数
- 目标文件应使用 `os.Create(dstFile)`，不要覆盖源文件
- 处理 `io.EOF` 作为读取结束条件
