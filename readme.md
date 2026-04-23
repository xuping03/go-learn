# 笔记

## JSX的本质

### 核心概念
JSX并不是标准的JS语法，它是**JS的语法扩展**，浏览器本身不能识别，需要通过解析工具（如Babel）做解析之后才能在浏览器中运行。

### JSX转换过程
```
<div>                      BABEL                  import { jsx as _jsx } from
  this is div        ───────────>              "react/jsx-runtime";
</div>                                          /*#__PURE__*/_jsx("div", {
                                                  children: "this is div"
                                                });
```

JSX代码经过Babel转换后，会变成React的函数调用（_jsx函数），该函数会创建React元素对象。

## JSX中使用JS表达式

在JSX中可以通过**大括号语法{}**识别JavaScript中的表达式，比如常见的变量、函数调用、方法调用等等

### 使用场景

1. **使用引号传递字符串**
   - 直接在JSX属性中使用引号传递字符串字面量

2. **使用JavaScript变量**
   - 通过{}在JSX中嵌入JavaScript变量

3. **函数调用和方法调用**
   - 在{}中可以调用函数或对象方法

4. **使用JavaScript对象**
   - 可以在JSX中使用JavaScript对象

注意：只有表达式可以放在{}中，if语句、switch语句、变量声明属于语句，不是表达式，不能出现在{}中

```jsx
function App() {
  const name = 'lilei';
  return (
    <div className="App">
      this is app
      {/* 是用引号传递字符串 */}
      {'this is message'}
      {/* 识别js变量 */}
      {name}
      {/* 函数调用 */}
      {(() => 'hello world')()}
      {/* 使用js对象 */}
      {{
        name: 'lilei',
        age: 18
      }.name}
    </div>
  );
}
```

## 列表渲染

### 使用map方法进行列表渲染

在React中，通常使用数组的`map`方法来遍历数组并渲染列表。

```jsx
const list = [
  { id: 1001, name: 'Vue' },
  { id: 1002, name: 'React' },
  { id: 1003, name: 'Angular' }
];

function App() {
  return (
    <div className="App">
      this is App
      {/* 渲染列表 */}
      {/* map 循环哪个结构 return结构 */}
      {/* 注意事项: 加上一个独一无二的key 字符串或者number id */}
      {/* key的作用: React框架内部使用 提升更新性能的 */}
      <ul>
        {list.map(item => <li key={item.id}>{item.name}</li>)}
      </ul>
    </div>
  )
}
```

### 关键注意事项

1. **使用map方法**：遍历数组并返回需要渲染的结构
2. **key属性是必需的**：每个列表项需要添加一个独一无二的key（字符串或number类型的id）
3. **key的作用**：React框架内部使用key来提升更新性能，帮助React识别哪些元素发生了变化

## JSX中实现条件渲染

### 条件渲染概念

在React中，可以根据不同的条件渲染不同的JSX结构。例如：
- 当`isLogin`为`true`时，显示"Jack"
- 当`isLogin`为`false`时，显示"请登录"

### 实现方式

在React中，可以通过以下方式实现基础的条件渲染：

1. **逻辑与运算符 &&**
2. **三元表达式 (?:)**

### 代码示例

```jsx
// 1. 使用逻辑与运算符 &&
{flag && <span>this is span</span>}

// 2. 使用三元表达式
{loading ? <span>loading...</span> : <span>this is span</span>}
```

### 使用说明

- **逻辑与 &&**：当条件为true时渲染后面的内容，为false时不渲染
- **三元表达式 ?:**：可以根据条件渲染两种不同的内容，适用于有明确"真"和"假"两种情况的场景

## 复杂条件渲染

### 使用函数实现多分支条件渲染

当条件分支较多时，可以将条件判断逻辑抽离成一个单独的函数，根据不同条件返回不同的JSX模板，再在组件中调用该函数进行渲染。

### 代码示例

```jsx
// 定义文章类型
const articleType = 3  // 0 1 3

// 定义核心函数（根据文章类型返回不同的JSX模板）
function getArticleTem () {
  if (articleType === 0) {
    return <div>我是无图文章</div>
  } else if (articleType === 1) {
    return <div>我是单图模式</div>
  } else {
    return <div>我是三图模式</div>
  }
}

function App () {
  return (
    <div className="App">
      {/* 调用函数渲染不同的模版 */}
      {getArticleTem()}
    </div>
  )
}

export default App
```

### 使用说明

- 将复杂的多分支条件逻辑封装到独立函数中，保持JSX结构清晰简洁
- 函数根据条件返回对应的JSX，在组件中通过 `{函数名()}` 的方式调用渲染

## React 事件绑定

### 四种常见场景

**1. 基础绑定**

```jsx
const handleClick = () => {
  console.log('button被点击了')
}
```

**2. 获取事件参数 e**

```jsx
const handleClick = (e) => {
  console.log('button被点击了', e)
}
```

**3. 传递自定义参数**

```jsx
const handleClick = (name) => {
  console.log('button被点击了', name)
}
```

**4. 既传递自定义参数，又获取事件对象 e**

```jsx
const handleClick = (name, e) => {
  console.log('button被点击了', name, e)
}

return (
  <div className="App">
    <button onClick={(e) => handleClick('jack', e)}>click me</button>
  </div>
)
```

### 使用说明

- 基础绑定：`onClick={handleClick}`，直接传函数引用，不加括号
- 需要自定义参数或同时获取事件对象时，用箭头函数包裹：`onClick={(e) => handleClick('jack', e)}`
- 事件对象 `e` 需要从外层箭头函数显式传入处理函数，不能直接在函数参数中省略
  - 只获取参数e的时候，不需要额外显式传入


## React 组件

### 概念

组件是 React 应用的基本构建单元，可以将 UI 拆分为独立、可复用的部分。

> **注意：自定义组件名称首字母必须大写**，React 通过首字母大小写区分原生 HTML 标签（小写）和自定义组件（大写）。

### 定义与使用

```jsx
// 1. 定义组件（首字母必须大写）
const Button = () => {
  // 业务逻辑/组件逻辑
  return <button>click me!</button>
}

// 2. 使用组件（渲染组件）
function App () {
  return (
    <div className="App">
      {/* 自闭合 */}
      <Button />
      {/* 成对标签 */}
      <Button></Button>
    </div>
  )
}

export default App
```

### 使用说明

- 组件本质是一个**返回 JSX 的函数**，支持箭头函数或普通函数写法
- 使用时支持**自闭合** `<Button />` 和**成对标签** `<Button></Button>` 两种写法
- 组件可以在其他组件中嵌套使用，实现 UI 的模块化拆分

## useState

### 概念

`useState` 是 React 提供的 Hook，用于在函数组件中添加**状态变量**。状态变量发生变化时，React 会自动重新渲染组件。

### 语法

```jsx
const [状态变量, 修改状态的方法] = useState(初始值)
```

- **状态变量**：保存当前状态值（只读，不能直接修改）
- **修改状态的方法**：调用后会用新值替换旧值，并触发组件重新渲染

### 示例：计数器按钮

```jsx
// useState 实现一个计数器按钮
import { useState } from 'react'

function App () {
  // 1. 调用 useState 添加一个状态变量
  // count    状态变量
  // setCount 修改状态变量的方法
  const [count, setCount] = useState(0)

  // 2. 点击事件回调
  const handleClick = () => {
    // 作用：1. 用传入的新值修改 count
    //       2. 重新使用新的 count 渲染 UI
    setCount(count + 1)
  }

  return (
    <div>
      <button onClick={handleClick}>{count}</button>
    </div>
  )
}

export default App
```

### 使用说明

- `useState` 需要从 `react` 中导入：`import { useState } from 'react'`
- **不能直接修改状态变量**（如 `count = count + 1`），必须通过 `setCount` 方法修改,复杂数据（对象）也不行，必须用set方法修改，使用新值替换旧值
- 调用 `setCount` 后，React 会用新值重新渲染组件 UI

## JSX 样式控制

### 行内样式

通过 `style` 属性传入一个 **JS 对象**来控制样式，CSS 属性名采用**驼峰命名法**。

### className 类名控制

通过 `className` 属性引用外部 CSS 文件中定义的类名来控制样式（推荐方式）。

### 完整示例

```jsx
// 导入样式
import './index.css'

const style = {
  color: 'red',
  fontSize: '50px'
}

function App () {
  return (
    <div>
      {/* 行内样式控制 */}
      <span style={style}>this is span</span>
      {/* 通过 class 类名控制 */}
      <span className="foo">this is class foo</span>
    </div>
  )
}

export default App
```

### 使用说明

- **行内样式**：`style` 值为 JS 对象，CSS 属性名用**驼峰命名**（`fontSize`、`backgroundColor`）
- **className 类名**：JSX 中用 `className` 代替 HTML 的 `class`，样式写在独立 CSS 文件中，通过 `import` 导入
- 实际项目中推荐使用 `className` 方式，结构与样式分离，便于维护

## lodash 工具库

### 常用方法：`_.orderBy`

对数组按指定字段和顺序排序，返回排序后的新数组（不修改原数组）。

**语法**

```js
_.orderBy(collection, fields, orders)
```

- `collection`：要排序的数组
- `fields`：排序字段，字符串或字符串数组
- `orders`：排序方向，`'asc'`（升序）或 `'desc'`（降序）

**示例：根据 tab 类型切换排序**

```js
const clickType = (type) => {
  console.log(type)
  setType(type)
  if (type === 'hot') {
    // 按点赞数 like 降序排列
    setList(_.orderBy(list, 'like', 'desc'))
  } else {
    // 按发布时间 ctime 降序排列
    setList(_.orderBy(list, 'ctime', 'desc'))
  }
}
```

### 使用说明

- 安装：`npm install lodash`，使用时 `import _ from 'lodash'`
- `_.orderBy` 返回**新数组**，原数组不变，符合 React 状态不可直接修改的原则
- 支持多字段排序：`_.orderBy(list, ['like', 'ctime'], ['desc', 'asc'])`

## classnames 库

### 作用

用于**动态拼接 className 字符串**，避免手动字符串拼接的繁琐，在需要条件性添加类名时非常实用。

### 语法

```js
classNames('固定类名', { '条件类名': 布尔表达式 })
```

- 固定类名：始终生效
- 条件类名：值为 `true` 时添加该类名，`false` 时不添加

### 示例：Tab 高亮切换

```jsx
<li className="nav-sort">
  {/* 高亮类名: active */}
  {tabs.map(item => (
    <span
      key={item.type}
      className={classNames('nav-item', { active: type === item.type })}
      onClick={() => clickType(item.type)}
    >
      {item.text}
    </span>
  ))}
</li>
```

当 `type === item.type` 为 `true` 时，该 `span` 的类名为 `"nav-item active"`，否则为 `"nav-item"`。

### 使用说明

- 安装：`npm install classnames`，使用时 `import classNames from 'classnames'`
- 支持多个条件类名同时判断：`classNames('a', { b: true, c: false })` → `"a b"`
- 是 React 中处理**动态类名**的标准做法，替代模板字符串拼接

## React 中获取 DOM

在 React 组件中获取/操作 DOM，需要使用 `useRef` 钩子函数，分为两步：

**步骤 1：使用 useRef 创建 ref 对象，并与 JSX 绑定**

```jsx
const inputRef = useRef(null)

<input type="text" ref={inputRef} />
```

**步骤 2：在 DOM 可用时，通过 `inputRef.current` 拿到 DOM 对象**

```jsx
console.log(inputRef.current)
```

### 完整示例

```jsx
import { useRef } from 'react'

function App () {
  const inputRef = useRef(null)
// 渲染完毕之后dom生成之后才可用
  const handleClick = () => {
    // 通过 inputRef.current 获取真实 DOM
    console.log(inputRef.current)
    inputRef.current.focus()
  }

  return (
    <div>
      <input type="text" ref={inputRef} />
      <button onClick={handleClick}>获取input DOM</button>
    </div>
  )
}
```

### 使用说明

- `useRef(null)` 初始值为 `null`，组件挂载后 `current` 会自动指向绑定的真实 DOM 节点
- 必须在**组件挂载后**（如事件回调、`useEffect` 中）才能访问 `ref.current`，渲染阶段访问仍为 `null`
- 修改 `ref.current` **不会触发组件重新渲染**，适合存储不需要响应式更新的值

## 组件通信 — 父传子（props）

### 核心步骤

1. **父组件传递数据**：在子组件标签上绑定属性
2. **子组件接收数据**：通过函数参数 `props` 接收

### 代码示例

```jsx
// 子组件：通过 props 参数接收数据
// props 是一个对象，包含父组件传递过来的所有数据
function Son (props) {
  console.log(props)
  return <div>this is son, {props.name}</div>
}

// 父组件：在子组件标签上绑定属性传递数据
function App () {
  const name = 'this is app name'
  return (
    <div>
      <Son name={name} />
    </div>
  )
}

export default App
```

### 使用说明

- `props` 是一个普通对象，所有绑定在子组件标签上的属性都会合并进来
- 可以传递任意类型的数据：字符串、数字、数组、对象、函数、JSX 等
- `props` 是**只读的**，子组件不能直接修改 props 中的值
- 可以通过解构简化写法：`function Son ({ name }) { ... }`

### 特殊的 prop — children

**场景**：当把内容嵌套在子组件标签中时，子组件会自动在名为 `children` 的 prop 属性中接收该内容。

```jsx
// 父组件：将内容嵌套在子组件标签中
<Son>
  <span>this is span</span>
</Son>

// 子组件：通过 props.children 接收嵌套内容
function Son (props) {
  return <div>{props.children}</div>
}
```

在 React DevTools 中可以看到，`props` 中自动出现了 `children: <span />` 字段。

- `children` 可以是任意内容：文本、JSX 元素、组件、数组等
- 常用于封装**容器类组件**（如弹窗、卡片、布局组件）

## 组件通信 — 子传父

### 核心思路

> 在子组件中调用父组件中的函数并传递实参

父组件将**回调函数**通过 props 传给子组件，子组件调用该函数时把数据作为参数传回，父组件即可拿到子组件的数据。

### 代码示例

```jsx
import { useState } from 'react'

// 子组件：解构接收父组件传来的回调函数
function Son ({ onGetSonMsg }) {
  // Son 组件中的数据
  const sonMsg = 'this is son msg'
  return (
    <div>
      this is Son
      {/* 点击按钮时调用父组件函数，将数据传递出去 */}
      <button onClick={() => onGetSonMsg(sonMsg)}>sendMsg</button>
    </div>
  )
}

function App () {
  const [msg, setMsg] = useState('')

  const getMsg = (msg) => {
    console.log(msg)
    setMsg(msg)
  }

  return (
    <div>
      this is App, {msg}
      {/* 将回调函数作为 prop 传给子组件 */}
      <Son onGetSonMsg={getMsg} />
    </div>
  )
}
```

### 使用说明

- 父组件定义函数 → 通过 props 传给子组件 → 子组件在合适时机调用并传入数据
- 命名惯例：回调 prop 通常以 `on` 开头，如 `onGetSonMsg`、`onChange`、`onClose`
- 子组件只负责**触发**，数据的存储和处理由父组件的 `useState` 管理

## 组件通信 — 兄弟组件（状态提升）

### 实现思路

借助**状态提升**机制，将共享状态提升到最近的公共父组件中，通过父组件进行兄弟组件之间的数据传递。

```
        App (state)
       /           \
    子传父          父传子
     /                 \
   A 组件            B 组件
(发送数据)          (接收数据)
```

### 核心步骤

1. **A → 父**：子组件 A 通过回调函数（子传父）将数据传给父组件
2. **父存储**：父组件用 `useState` 保存从 A 接收到的数据
3. **父 → B**：父组件将数据通过 props（父传子）传给兄弟组件 B

### 代码示例

```jsx
import { useState } from 'react'

// 组件 A：发送数据
function A ({ onGetName }) {
  const name = 'this is A name'
  return (
    <div>
      <button onClick={() => onGetName(name)}>send</button>
    </div>
  )
}

// 组件 B：接收数据
function B ({ name }) {
  return <div>{name}</div>
}

// 父组件 App：中间桥梁
function App () {
  const [name, setName] = useState('')
  return (
    <div>
      <A onGetName={setName} />
      <B name={name} />
    </div>
  )
}
```

### 使用说明

- 兄弟组件**不能直接通信**，必须借助共同的父组件中转
- 状态提升的本质：把状态放到需要共享它的组件的**最近公共祖先**中管理
- 适用于组件层级不深的场景；层级较深时推荐使用 Context 或状态管理库

## 组件通信 — 跨层级（Context）

### 场景

当组件层级较深时（如 App → A → B），用 props 逐层传递繁琐，使用 **Context** 可以让顶层组件直接向任意深度的子组件传递数据，中间组件无需参与。

```
App（提供数据）
  └── A（中间层，无需传递）
        └── B（直接消费数据）
```

### 实现步骤

1. 使用 `createContext` 方法创建一个上下文对象 `Ctx`
2. 在顶层组件（App）中通过 `Ctx.Provider` 组件提供数据
3. 在底层组件（B）中通过 `useContext` 钩子函数获取消费数据

### 代码示例

```jsx
import { createContext, useContext } from 'react'

// 1. 创建上下文对象
const MsgCtx = createContext()

// 底层组件 B：通过 useContext 消费数据
function B () {
  const msg = useContext(MsgCtx)
  return <div>{msg}</div>
}

// 中间组件 A：无需关心数据传递
function A () {
  return <B />
}

// 顶层组件 App：通过 Provider 提供数据
function App () {
  const msg = 'this is app name'
  return (
    <MsgCtx.Provider value={msg}>
      <A />
    </MsgCtx.Provider>
  )
}
```

### 使用说明

- `createContext()` 通常在组件文件外部创建，可跨文件导入共享
- `Provider` 的 `value` 变化时，所有消费该 Context 的组件会自动重新渲染
- 适合主题、语言、登录用户信息等**全局共享数据**的传递
