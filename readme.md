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

## 二十、 类型转换的兼容性和限制

### 类型转换的基本规则

Go 语言中，**不同类型之间不能直接进行转换**，必须遵循特定的规则。

#### 能够转换的类型对

**1) bool 类型的转换限制**
- **bool 类型不能转换为任何其他类型**
- **任何其他类型也不能转换为 bool**

**示例（错误）：**
```go
var flag bool
flag = true
// 错误：bool 类型不能转换为 int
fmt.Printf("flag = %d\n", int(flag))
```

**2) int 和 float 之间的转换**
- **int 可以转换为 float**
- **float 可以转换为 int**（但会丢失小数部分）

**示例（正确）：**
```go
var a int = 10
var b float64 = float64(a)  // int → float64
fmt.Println(b)  // 输出: 10

var d float64 = 3.14
var c int = int(d)  // float64 → int，丢失小数部分
fmt.Println(c)  // 输出: 3
```

**3) 字符和整数之间的转换**
- **字符（byte/rune）可以转换为整数**
- **整数可以转换为字符（byte/rune）**

**示例（正确）：**
```go
var ch byte = 'a'
var t int = int(ch)           // 字符 → 整数，得到 ASCII 值
fmt.Println("t =", t)         // 输出: 97

var num int = 97
var newCh byte = byte(num)    // 整数 → 字符
fmt.Println("newCh =", newCh) // 输出: a
```

### 类型转换的不兼容情况

#### 1) bool 和其他类型不兼容

**问题代码：**
```go
var flag bool = true
// 这种不能转换的类型，叫不兼容类型
t := int(flag)  // 编译错误！bool 不能转换为 int
```

**错误信息：**
```
cannot convert flag (type bool) to type int
```

**说明：**
- bool 类型在 Go 中是独立的类型
- bool 不能用数字 0 或 1 代表
- 不支持任何形式的 bool 类型转换

#### 2) 整数和 bool 不兼容

**问题代码：**
```go
var a int = 1
// 错误：int 也不能转换为 bool
var flag bool = bool(a)  // 编译错误！
```

**说明：**
- 虽然 0 代表假，非 0 代表真（在某些语言中）
- 但在 Go 中，这种隐式转换是不允许的

### 兼容性总结表

| 源类型 | 目标类型 | 兼容性 | 说明 |
|--------|---------|--------|------|
| **int** | float | ✓ 兼容 | 可以转换，float 可能精度损失 |
| **float** | int | ✓ 兼容 | 可以转换，int 会丢失小数部分 |
| **byte/rune** | int | ✓ 兼容 | 可以转换为对应的整数值 |
| **int** | byte/rune | ✓ 兼容 | 可以转换为字符 |
| **bool** | 其他类型 | ✗ 不兼容 | bool 不能转换为任何其他类型 |
| **其他类型** | bool | ✗ 不兼容 | 任何类型都不能转换为 bool |
| **int** | bool | ✗ 不兼容 | 即使 1 和 0，也不能转换 |
| **string** | 其他类型 | ✗ 不兼容 | 字符串不能直接强制转换（需要特殊方法） |

### 常见转换场景

#### 场景1：字符转整数（获取 ASCII 值）

```go
var ch byte = 'a'
var t int = int(ch)
fmt.Println("字符 'a' 的 ASCII 值:", t)  // 输出: 97
```

#### 场景2：整数转字符

```go
var num int = 97
var ch byte = byte(num)
fmt.Println("整数 97 对应的字符:", ch)  // 输出: a
```

#### 场景3：浮点数转整数（丢失小数）

```go
var d float64 = 3.14
var c int = int(d)
fmt.Println("3.14 转为整数:", c)  // 输出: 3（小数部分丢失）
```

### 类型转换的注意事项

| 注意事项 | 说明 |
|---------|------|
| **精度损失** | float → int 会丢失小数部分 |
| **范围溢出** | 大数值转换为小类型可能溢出 |
| **不兼容类型** | bool 与其他类型完全不兼容 |
| **显式转换** | Go 不支持隐式类型转换，必须显式转换 |
| **字符串转换** | string 类型需要使用 strconv 包的函数 |

### 类型转换的完整示例

```go
package main

import "fmt"

func main() {
    // 示例1：bool 类型（不能转换）
    var flag bool = true
    fmt.Printf("flag = %t\n", flag)  // 输出: flag = true
    // fmt.Printf("flag = %d\n", int(flag))  // 错误！bool 不能转换
    
    // 示例2：字符转整数
    var ch byte = 'a'
    var t int = int(ch)
    fmt.Printf("字符 'a' → 整数: %d\n", t)  // 输出: 97
    
    // 示例3：整数转字符
    var num int = 97
    var newCh byte = byte(num)
    fmt.Printf("整数 97 → 字符: %c\n", newCh)  // 输出: a
    
    // 示例4：float 转 int（丢失小数）
    var d float64 = 3.14
    var c int = int(d)
    fmt.Printf("3.14 → int: %d\n", c)  // 输出: 3
}
```

---

## 二十一、 自定义类型（Type Definition）

### 什么是自定义类型？

**自定义类型是指用户自己定义的新类型，基于已有的基础类型**
- 给某个已存在的类型起一个新的别名
- 增强代码的可读性和语义

### 自定义类型的语法

```go
type 新类型名 基础类型
```

#### 语法说明

- **type**：Go 的关键字，用于定义新类型
- **新类型名**：自定义的类型名称（首字母通常大写）
- **基础类型**：已存在的类型（int、float64、string 等）

### 自定义类型的使用场景

#### 场景1：给 int64 起个别名

```go
// 给 int64 起一个别名叫 bigint
type bigint int64

var a bigint  // 等价于 var a int64
fmt.Printf("a type is %T\n", a)  // 输出: a type is main.bigint
```

**说明：**
- `bigint` 是一个新的类型名
- 使用 `bigint` 定义的变量在类型检查时被认为是 `bigint` 类型
- 不能直接赋值给 `int64` 类型的变量（需要显式转换）

#### 场景2：为多个类型定义别名

```go
type (
    long int64
    char byte
)

var b long = 11
var ch char = 'a'
fmt.Printf("b = %d, ch = %c\n", b, ch)
```

### 自定义类型与基础类型的关系

#### 重要区别

虽然自定义类型基于某个基础类型，但它们是**不同的类型**：

```go
type bigint int64

var a bigint = 10
var b int64 = 10

// 错误：不能直接赋值，因为类型不同
// a = b  // 编译错误
// b = a  // 编译错误

// 需要显式转换
a = bigint(b)
b = int64(a)
```

### 自定义类型的完整示例

```go
package main

import "fmt"

func main() {
    // 定义单个自定义类型
    type bigint int64
    
    var a bigint
    fmt.Printf("a type is %T\n", a)  // 输出: a type is main.bigint
    
    // 定义多个自定义类型
    type (
        long int64
        char byte
    )
    
    var b long = 11
    var ch char = 'a'
    fmt.Printf("b = %d, ch = %c\n", b, ch)  // 输出: b = 11, ch = a
}
```

### 自定义类型的优势

| 优势 | 说明 |
|------|------|
| **可读性** | 给类型起一个有意义的名字，代码更易理解 |
| **语义清晰** | `bigint` 比 `int64` 更能表达意图 |
| **类型安全** | 不同的自定义类型之间不能混用 |
| **代码组织** | 可以为特定业务场景定义专用类型 |

### 自定义类型的应用场景

**场景1：业务概念**
```go
type UserId int64
type ProductId int64

var userId UserId = 123
var productId ProductId = 456
// userId 和 productId 是不同的类型，即使都是 int64
```

**场景2：度量单位**
```go
type Temperature float64
type Distance float64

var temp Temperature = 36.5
var dist Distance = 100.5
// 类型不同，无法混用
```

**场景3：权限值**
```go
type Permission int

const (
    ReadOnly Permission = iota
    ReadWrite
    Execute
)
```

### 自定义类型 vs 类型别名

| 特性 | 自定义类型 | 类型别名 |
|------|----------|--------|
| **语法** | `type Name BaseType` | `type Name = BaseType` |
| **是否新类型** | 是 | 否 |
| **类型检查** | 严格不兼容 | 完全兼容 |
| **使用场景** | 需要类型安全 | 只是起个别名 |

**注意：** 本章重点是自定义类型（第一种），类型别名（第二种）通常在特定场景才使用。

### 自定义类型的常见错误

| 错误 | 原因 | 解决方案 |
|------|------|---------|
| **直接赋值** | 自定义类型与基础类型不兼容 | 使用显式类型转换 |
| **混淆概念** | 把自定义类型当作基础类型 | 记住自定义类型是新类型 |
| **遗漏 type** | 忘记使用 `type` 关键字 | 确保定义时包含 `type` |

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

## 二十三、 switch 的高级用法

### 1) switch 后面写条件表达式

**switch 后面不一定要写变量，也可以直接写表达式**
- 这样 case 后面就可以写条件表达式
- 相当于简化的 if-else 语句

#### 语法

```go
switch 表达式/变量 {
case 值1:
    // 执行代码
case 值2:
    // 执行代码
default:
    // 默认执行
}
```

#### 示例1：没有条件的 switch

```go
score := 85
switch {  // 没有条件表达式
case score > 90:
    fmt.Println("优秀")
case score > 80:
    fmt.Println("良好")  // 会执行这里
case score > 70:
    fmt.Println("一般")
default:
    fmt.Println("不及格")
}
// 输出: 良好
```

#### 示例2：switch 后面写表达式

```go
num := 1
switch num {  // switch 后面写变量本身
case 1:
    fmt.Println("按下的是1楼")
case 2:
    fmt.Println("按下的是2楼")
case 3:
    fmt.Println("按下的是3楼")
case 4:
    fmt.Println("按下的是4楼")
default:
    fmt.Println("按下的是xxx楼")
}
// 输出: 按下的是1楼
```

### 2) case 后面写条件表达式

**case 后面可以写条件，而不仅仅是值**
- 当 switch 后面为空（没有表达式）时，case 后面可以写条件判断
- case 会依次判断条件是否为真

#### 语法

```go
switch {
case 条件1:
    // 如果条件1为真执行
case 条件2:
    // 如果条件2为真执行
default:
    // 都不满足时执行
}
```

#### 完整示例

```go
package main

import "fmt"

func main() {
    score := 85
    
    // 方法1：switch 后面没有表达式，case 后面写条件
    switch {
    case score > 90:
        fmt.Println("优秀")
    case score > 80:
        fmt.Println("良好")  // 会执行
    case score > 70:
        fmt.Println("一般")
    default:
        fmt.Println("不及格")
    }
    
    // 输出: 良好
}
```

### 3) 不同数据类型的 switch

#### 字符串类型的 switch

```go
var day string = "Monday"
switch day {
case "Monday":
    fmt.Println("星期一")
case "Tuesday":
    fmt.Println("星期二")
case "Wednesday":
    fmt.Println("星期三")
default:
    fmt.Println("其他日期")
}
// 输出: 星期一
```

#### 类型判断的 switch（类型断言）

```go
var x interface{} = 10

switch v := x.(type) {
case int:
    fmt.Println("是整数:", v)
case string:
    fmt.Println("是字符串:", v)
case float64:
    fmt.Println("是浮点数:", v)
default:
    fmt.Println("其他类型")
}
// 输出: 是整数: 10
```

### switch 的执行逻辑总结

#### 不同场景的 switch 执行

| 场景 | switch 表达式 | case 后面 | 说明 |
|------|-------------|---------|------|
| **值比较** | 有表达式（变量） | 具体值 | 与各个值依次比较 |
| **条件判断** | 无表达式 | 条件表达式 | 依次判断条件是否为真 |
| **多值匹配** | 有表达式 | 多个值用逗号分隔 | 匹配任意一个值 |
| **类型判断** | interface{} | type 关键字 | 判断具体数据类型 |

### switch 与 if-else 的对比应用

#### 场景1：分数等级（用 switch）

```go
score := 85
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")  // 执行
case score >= 70:
    fmt.Println("C")
default:
    fmt.Println("F")
}
```

#### 场景2：分数等级（用 if-else）

```go
score := 85
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")  // 执行
} else if score >= 70 {
    fmt.Println("C")
} else {
    fmt.Println("F")
}
```

**对比：** switch 方式代码更清晰，结构更规范

### switch 的实际应用示例

#### 应用1：根据用户输入返回相应的提示

```go
var num int
fmt.Printf("请选择操作（1:新增, 2:删除, 3:修改, 4:查询）: ")
fmt.Scan(&num)

switch num {
case 1:
    fmt.Println("执行新增操作")
case 2:
    fmt.Println("执行删除操作")
case 3:
    fmt.Println("执行修改操作")
case 4:
    fmt.Println("执行查询操作")
default:
    fmt.Println("请选择有效的操作")
}
```

#### 应用2：成绩判定

```go
var score int
fmt.Printf("请输入成绩: ")
fmt.Scan(&score)

switch {
case score >= 90:
    fmt.Println("等级: 优秀")
case score >= 80:
    fmt.Println("等级: 良好")
case score >= 70:
    fmt.Println("等级: 及格")
case score >= 60:
    fmt.Println("等级: 勉强及格")
default:
    fmt.Println("等级: 不及格")
}
```

### switch 的注意事项

| 注意事项 | 说明 |
|---------|------|
| **break 默认存在** | 每个 case 执行后自动退出，无需手动 break |
| **fallthrough 用法** | 需要穿透到下一个 case 时使用 fallthrough |
| **表达式类型** | switch 表达式和 case 值的类型必须兼容 |
| **default 位置** | default 可以在任何位置，不一定在最后 |
| **switch 嵌套** | 可以在 case 中嵌套另一个 switch |

---

## 二十四、 循环结构 - for 循环

### 什么是 for 循环？

**for 循环是一种反复执行某段代码的结构**
- 用于重复执行相同或相似的代码
- Go 语言中只有 for 循环，没有 while 循环

### for 循环的三种使用方法

## 方法1：传统的 for 循环（C 风格）

#### 语法

```go
for 初始化语句; 条件表达式; 迭代语句 {
    // 循环体
}
```

#### 语法说明

- **初始化语句**：循环前执行一次，通常初始化循环变量
- **条件表达式**：每次循环前检查，为 true 则继续，为 false 则退出
- **迭代语句**：每次循环体执行后执行，通常用于更新循环变量
- **循环体**：每次条件为 true 时执行的代码

#### 执行流程

1. 执行初始化语句（只执行一次）
2. 判断条件表达式
3. 条件为 true，执行循环体
4. 执行迭代语句
5. 回到第 2 步，重复直到条件为 false

#### 示例1：打印 1 到 5

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
// 输出:
// 1
// 2
// 3
// 4
// 5
```

#### 示例2：遍历字符串

```go
str := "abc"
for i := 0; i < len(str); i++ {
    fmt.Println(str[i])
}
// 输出:
// 97  (字符 'a' 的 ASCII 值)
// 98  (字符 'b' 的 ASCII 值)
// 99  (字符 'c' 的 ASCII 值)
```

#### 示例3：求和

```go
sum := 0
for i := 1; i <= 100; i++ {
    sum += i
}
fmt.Println("1 到 100 的和:", sum)
// 输出: 1 到 100 的和: 5050
```

---

## 方法2：range 循环（遍历字符串/数组）

#### 语法

```go
for 索引, 值 := range 容器 {
    // 循环体
}
```

或者只要索引：
```go
for 索引 := range 容器 {
    // 循环体
}
```

或者只要值：
```go
for _, 值 := range 容器 {
    // 循环体
}
```

#### 语法说明

- **索引**：当前元素的位置（从 0 开始）
- **值**：当前元素的值
- **容器**：要遍历的数据结构（字符串、数组、切片等）
- **_**：匿名变量，用来忽略不需要的值

#### 示例1：遍历字符串（获取索引和值）

```go
str := "abc"
for i, data := range str {
    fmt.Println(i, data)
}
// 输出:
// 0 97  (索引 0，值 'a')
// 1 98  (索引 1，值 'b')
// 2 99  (索引 2，值 'c')
```

**说明：** 注意 `data` 是字符的 Unicode 值（ASCII 码），不是字符本身

#### 示例2：只获取索引

```go
str := "abc"
for i := range str {
    fmt.Println(i)
}
// 输出:
// 0
// 1
// 2
```

#### 示例3：只获取值

```go
str := "abc"
for _, data := range str {
    fmt.Printf("%c ", data)
}
// 输出: a b c
```

#### 示例4：打印字符（使用 %c）

```go
str := "abc"
for i, data := range str {
    fmt.Printf("索引:%d, 字符:%c, ASCII值:%d\n", i, data, data)
}
// 输出:
// 索引:0, 字符:a, ASCII值:97
// 索引:1, 字符:b, ASCII值:98
// 索引:2, 字符:c, ASCII值:99
```

---

## 方法3：简化的 for 循环（类似 while）

#### 语法

```go
for 条件表达式 {
    // 循环体
}
```

#### 说明

- 只有条件，没有初始化和迭代
- 相当于其他语言的 while 循环
- 当条件为 false 时退出循环

#### 示例

```go
i := 1
for i <= 5 {
    fmt.Println(i)
    i++
}
// 输出:
// 1
// 2
// 3
// 4
// 5
```

---

## 方法4：无限循环

#### 语法

```go
for {
    // 循环体
    // 通常需要 break 语句来退出
}
```

#### 说明

- 条件表达式为空，始终为 true
- 必须使用 break 语句才能退出循环

#### 示例

```go
for {
    fmt.Println("这是无限循环")
    break  // 立即退出循环
}
// 输出: 这是无限循环
```

---

### 循环控制语句

#### 1) break - 退出循环

```go
for i := 1; i <= 10; i++ {
    if i == 5 {
        break  // 当 i 等于 5 时退出循环
    }
    fmt.Println(i)
}
// 输出:
// 1
// 2
// 3
// 4
```

#### 2) continue - 跳过本次迭代

```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue  // 跳过 i=3，继续下一次循环
    }
    fmt.Println(i)
}
// 输出:
// 1
// 2
// 4
// 5
```

### for 循环的对比

| 特性 | 传统 for | range | 简化 for | 无限 for |
|------|---------|-------|---------|----------|
| **语法** | `for i:=0; i<n; i++` | `for i, v := range x` | `for 条件` | `for` |
| **用途** | 精确控制循环次数 | 遍历数据结构 | 简单的条件循环 | 需要手动控制 |
| **循环变量** | 需要声明 | 自动声明 | 需要声明 | 无 |
| **推荐场景** | 固定次数、索引访问 | 遍历数组/字符串 | 动态条件 | 特殊逻辑 |

### 完整的 for 循环示例

```go
package main

import "fmt"

func main() {
    // 方法1：传统 for 循环
    fmt.Println("方法1：传统 for 循环")
    str := "abc"
    for i := 0; i < len(str); i++ {
        fmt.Printf("索引:%d, ASCII值:%d\n", i, str[i])
    }
    
    // 方法2：range 循环
    fmt.Println("\n方法2：range 循环")
    for i, data := range str {
        fmt.Printf("索引:%d, 字符:%c, ASCII值:%d\n", i, data, data)
    }
    
    // 方法3：简化的 for 循环
    fmt.Println("\n方法3：简化的 for 循环")
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i++
    }
}
```

### for 循环的最佳实践

| 实践 | 说明 |
|------|------|
| **选择合适的方法** | 不同场景选择不同的 for 循环方式 |
| **避免无限循环** | 确保循环条件最终会变为 false |
| **合理使用 break** | 用于提前退出循环 |
| **合理使用 continue** | 用于跳过不需要处理的迭代 |
| **避免深层嵌套** | 过深的循环嵌套会降低代码可读性 |

---

## 二十五、 goto 语句和标签

### 什么是 goto 语句？

**goto 是一种无条件跳转语句**
- 可以将程序控制流转移到被标签标记的位置
- 用于跳出复杂的嵌套结构
- Go 中很少使用（通常被认为是不好的编程习惯）

### goto 语句的基本语法

```go
// 定义标签
标签名:
    代码块

// 使用 goto 跳转
goto 标签名
```

#### 语法说明

- **标签**：由用户自定义的标识符，后面跟冒号 `:`
- **goto**：Go 的关键字，用于无条件跳转
- 标签可以在 goto 前面或后面定义

### goto 语句的执行流程

当执行 `goto 标签名` 时：
1. 程序会立即跳转到标签所在的位置
2. 从标签处继续执行代码
3. 标签之间的代码会被跳过

### goto 语句的使用示例

#### 示例1：简单的 goto 跳转

```go
package main

import "fmt"

func main() {
    fmt.Println("开始")
    
    goto End  // 跳转到 End 标签
    
    fmt.Println("这行代码会被跳过")
    fmt.Println("这行代码也会被跳过")
    
    End:
    fmt.Println("跳转到这里了")
}

// 输出:
// 开始
// 跳转到这里了
```

#### 示例2：在循环中使用 goto

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        if i == 3 {
            goto Skip  // 当 i=3 时跳转
        }
        fmt.Println(i)
    }
    
    Skip:
    fmt.Println("循环已完成")
}

// 输出:
// 1
// 2
// 跳转到这里了
```

### break 和 continue 的限制

#### 1) break 的限制

- **break 只能在循环或 switch 中使用**
- break is not in a loop, switch, or select（编译错误）
- 不能在循环外使用 break

**错误示例：**
```go
if true {
    break  // 编译错误！break 不在循环中
}
```

#### 2) continue 的限制

- **continue 只能在循环中使用**
- continue is not in a loop（编译错误）
- 不能在 switch 或循环外使用 continue

**错误示例：**
```go
switch x {
case 1:
    continue  // 编译错误！continue 不在循环中
}
```

### goto vs break/continue

| 特性 | goto | break | continue |
|------|------|-------|----------|
| **用途** | 无条件跳转到标签位置 | 退出循环或 switch | 跳过本次迭代 |
| **使用位置** | 任何地方 | 只在循环或 switch 中 | 只在循环中 |
| **跳转范围** | 无限制 | 只能退出当前循环 | 只能跳过当前迭代 |
| **代码可读性** | 差（容易造成混乱） | 好 | 好 |
| **推荐使用** | 少用 | 常用 | 常用 |

### goto 的应用场景

#### 场景1：跳出嵌套循环

```go
package main

import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 1 && j == 1 {
                goto Exit  // 一次性跳出两层循环
            }
            fmt.Printf("(%d, %d) ", i, j)
        }
    }
    
    Exit:
    fmt.Println("\n循环结束")
}

// 输出:
// (0, 0) (0, 1) (0, 2) (1, 0) 
// 循环结束
```

**对比：** 用 break 和 continue 实现相同的功能会更复杂

#### 场景2：错误处理中的跳转

```go
package main

import "fmt"

func main() {
    // 模拟错误处理流程
    value := -5
    
    if value < 0 {
        fmt.Println("错误：值不能为负数")
        goto ErrorHandler
    }
    
    fmt.Println("处理正常数据")
    
    ErrorHandler:
    fmt.Println("执行错误处理逻辑")
}

// 输出:
// 错误：值不能为负数
// 执行错误处理逻辑
```

### 完整的 goto 示例

```go
package main

import "fmt"

func main() {
    // break 和 continue 的限制
    // break 只能在循环或 switch 中
    // continue 只能在循环中
    
    // goto 可以在任何地方，但不能穿函数使用
    fmt.Println("111111111111111")
    
    goto End  // goto 是关键字，End 是标签（用户起的名字）
    
    fmt.Println("222222222222222")
    
    End:
    fmt.Println("333333333333333")
}

// 输出:
// 111111111111111
// 333333333333333
```

### goto 的注意事项

| 注意事项 | 说明 |
|---------|------|
| **避免过度使用** | goto 会使代码流程难以理解，一般避免使用 |
| **不能跨函数** | goto 只能在同一函数内跳转，不能跳转到其他函数 |
| **标签命名** | 标签名应该有意义，便于理解跳转的目的 |
| **代码可读性** | 使用 goto 会降低代码可读性，优先使用 break/continue |
| **特殊场景** | 仅在复杂的嵌套循环中才考虑使用 |

### Go 编程建议

**虽然 Go 支持 goto，但强烈建议：**
1. 尽量避免使用 goto
2. 使用 break 和 continue 来控制循环流程
3. 使用函数调用来组织代码逻辑
4. 使用 if-else 来处理条件分支

**总结：** goto 是一个"危险的"工具，虽然存在但不推荐在日常编程中使用。