# [MarkDown 中使用 LaTeX 数学式](https://www.cnblogs.com/nowgood/p/Latexstart.html)

参考:

https://www.cnblogs.com/nowgood/p/Latexstart.html

[TOC]

<img src="https://images2017.cnblogs.com/blog/1182370/201708/1182370-20170830201334046-1033111420.jpg"/>

最近看了些机器学习的书籍, 想写点笔记记录下. 由于需要使用到很多的数学推导, 所以就看了下如何在 Markdown 中插入数学式，发现在 Markdown 中可以直接插入 LaTeX 数学式.

排版数学公式是 TEXTEX 系统设计的初衷, 在 LATEXLATEX 中占有特殊地位, 是 LATEXLATEX 最为人称道的功能之一, 很多人就是冲着 LATEXLATEX 的公式输入功能来的:), 如我... 下面简要介绍下 MarkDown 中如何使用 LATEXLATEX 输入数学公式.

## 数学模式

在 LaTeX 中，最常用到的主要有文本模式和数学模式这两种模式。数学模式又可分为行内公式{inline math)和行间公式 (display math) 两种形式。

行内公式形式是将数学式插入文本行之内，使之与文本融为一体，这种形式适合编写简 短的数学式。

行间公式形式是将数学式插在文本行之间，自成一行或一个段落，与上下文附加一段垂 直空白，使数学式突出醒目。多行公式、公式组和微积分方程等复杂的数学式都是采用行间 公式形式编写。

行内公式 `$ ... $`
行间公式 `$$ ... $$`

```
函数 ${f(x)=a_nx^n+a_{n-1}x^{n-1}+a_{n-2}x^{n-2}}+\cdots$

函数 $${f(x)=a_nx^n+a_{n-1}x^{n-1}+a_{n-2}x^{n-2}}+\cdots \tag{1.1}$$
```

函数 ${f(x)=a_nx^n+a_{n-1}x^{n-1}+a_{n-2}x^{n-2}}+\cdots$

函数

 $${f(x)=a_nx^n+a_{n-1}x^{n-1}+a_{n-2}x^{n-2}}+\cdots \tag{1.1}$$



> LaTeX 注释符号为 %%

### 输入上下标

`^` 表示上标, `_` 表示下标。如果上下标的内容多于一个字符，要用大括号 { } 把这些内容括起来当成一个整体。上下标是可以嵌套的，也可以同时使用。

$\sum_i^na_i$

```
$\sum_i^na_i$
```

### 输入分数

分数的输入形式为 `\frac{分子}{分母}`

$P(v)=\frac{1}{1+exp(-v/T)}$

```
$P(v)=\frac{1}{1+exp(-v/T)}$
```

### 上下划线与花括号


$$
\begin{array}
\overline{a+b+c} \\
\underline{a+b+c} \\
\overleftarrow{a+b} \\
\underleftarrow{a+b} \\
\underleftrightarrow{a+b} \\
\vec x = \vec{AB} \\
\overbrace {a+b}^\text{a,b} \\
a+\rlap{\overbrace{\phantom{b+c+d}}^m}b+\underbrace{c+d+e}_n+f
\end{array}
$$


```
$$
\begin{array}
\overline{a+b+c} \\
\underline{a+b+c} \\
\overleftarrow{a+b} \\
\underleftarrow{a+b} \\
\underleftrightarrow{a+b} \\
\vec x = \vec{AB} \\
\overbrace {a+b}^\text{a,b} \\
a+\rlap{\overbrace{\phantom{b+c+d}}^m}b+\underbrace{c+d+e}_n+f
\end{array}
$$
```

### 输入根号

 
$$
\begin{align*} \sqrt {12} \\ \sqrt[n]{12}  \end{align*}
$$


```
$$
\begin{align*}
\sqrt {12} \\
\sqrt[n]{12} 
\end{align*}
$$
```

### 输入括号和分隔符

`(), [] , |` 分别表示原尺寸的形状，由于大括号 {} 在 LaTeX 中有特定含义, 所以使用需要转义, 即`\{` 和 `\}` 分别表示表示{ }。当需要显示大尺寸的上述符号时, 在上述符号前加上 `\left` 和 `\right` 命令.

$\{a\}$
$f(x,y,z) = 3y^2z  3+(\frac{7x+5}{1+y^2}) $
$f(x,y,z) = 3y^2z + \left( 3 +\frac{7x+5}{1+y^2} \right)$

```
$\{a\}$
$f(x,y,z) = 3y^2z3+(\frac{7x+5}{1+y^2}) $
$f(x,y,z) = 3y^2z + \left( 3 +\frac{7x+5}{1+y^2} \right)$
```

关于各种数学符号写法, 详见[Cmd Markdown 公式指导手册](https://www.zybuluo.com/codeep/note/163962#cmd-markdown-公式指导手册), 下面主要介绍下常用的 矩阵和多行公式输入 做详细的记录.

## 矩阵

矩阵中, 不同的列使用 `&` 分割, 行使用 `\\` 分隔

下面展示一系列矩阵环境排版, 区别在于外面的括号不同


$$
\begin{pmatrix}
a & b & c \\
d & e & f \\
g & h & i 
\end{pmatrix}
$$

$$
\chi(\lambda) =  
\begin{vmatrix}
\lambda - a & -b & -c \\
-d & \lambda - e & -f \\
-g & -h & \lambda - i 
\end{vmatrix}
$$



```
$$
\begin{pmatrix}
a & b & c \\
d & e & f \\
g & h & i 
\end{pmatrix} 
$$

$$
\chi(\lambda) =  
\begin{vmatrix}
\lambda - a & -b & -c \\
-d & \lambda - e & -f \\
-g & -h & \lambda - i 
\end{vmatrix}
$$
```

### 省略号




$$ {eqnarray*}
\\ \ldots \\ \cdots \\ \vdots \\ \ddots \\
$$ {eqnarray*}




```
$$
\begin{eqnarray*} \\
\ldots \\
\cdots \\
\vdots \\
\ddots \\
\end{eqnarray*}
$$
```

## 单行公式与多行公式

`equation` 环境用来输入单行公式, 自动生成编号, 也可以使用 \tag{...} 自己对公式编号; 使用 `equation*` 环境, 不会自动生成公式编号, 后续介绍的公式输入环境都是在自动编号后面加上 `*` 便是不自动编号环境.



$\begin{equation} (a+b) \times c = a\times c + b \times c \\ \end{equation}$



```
\begin{equation}
(a+b) \times c = a\times c + b \times c \\
\end{equation}
```

`\[ ... \]` 是 `equation*` 环境的简写


$$
(a+b) \times c = a\times c + b \times c
$$


```
\\[
(a+b) \times c = a\times c + b \times c \\
\\] 
```

`eqnarray` 环境用来输入按照等号(或者其他关系符)对齐的方程组, 编号



$$ \begin{eqnarray} f(x) = a_nx^n \\ g(x) = x^2 \end{eqnarray} $$



```
$$
\begin{eqnarray}
f(x) = a_nx^n \\
g(x) = x^2
\end{eqnarray}
$$
```

输入多行公式, `gather` 环境得到的公式是每行居中的, `align`环境则允许公式按照等号或者其他关系符对齐, 在关系符前加`&`表示对齐


$$
\begin{gather}
(a+b) \times c = a\times c + b \times c \notag \\
ac= a\times c \\
\end{gather}
$$



$$
\begin{align}
y &= \cos t + 1 \\
y &= 2sin t \\
\end{align}
$$




```
$$
\begin{gather}
(a+b) \times c = a\times c + b \times c \notag \\
ac= a\times c \\
\end{gather}
$$

$$
\begin{align}
y &= \cos t + 1 \\
y &= 2sin t \\
\end{align}
$$
```

`align` 环境还允许排列多列对齐公式, 列与列之间使用`&`分割
$$
\begin{align*}
 x &= t & x &= \cos t &  x &= t \\
 y &= 2t & y &= \sin (t+1) & y &= \sin t \\
\end{align*}
$$

$$
\begin{align*}
& (a+b)(a^2-ab+b^2) \\
= {}& a^3-a^2b+ab^2+a^2b-ab^2+b^2 \\
= {}& a^3 + b^3
\end{align*}
$$

```
$$
\begin{align*}
 x &= t & x &= \cos t &  x &= t \\
 y &= 2t & y &= \sin (t+1) & y &= \sin t \\
\end{align*}
$$

$$
\begin{align*}
& (a+b)(a^2-ab+b^2) \\
= {}& a^3-a^2b+ab^2+a^2b-ab^2+b^2 \\
= {}& a^3 + b^3
\end{align*}
$$
```

align 环境中列分隔符 & 一般放在关系符前面, 如果个别需要再关系符后面或者别的地方对齐的, 则应该注意使用的符号类型


$$
% 关系符后对齐，需要使用空的分组
% 代替关系符右侧符号，保证间距
\begin{align*}
    & (a+b)(a^2-ab+b^2) \notag \\
={ } & a^3 - a^2b + ab^2 + a^2b
      - ab^2 + b^2 \notag \\
={ } & a^3 + b^3 \label{eq:cubesum}
\end{align*}
$$


```
$$
% 关系符后对齐，需要使用空的分组
% 代替关系符右侧符号，保证间距
\begin{align*}
    & (a+b)(a^2-ab+b^2) \notag \\
={ } & a^3 - a^2b + ab^2 + a^2b
      - ab^2 + b^2 \notag \\
={ } & a^3 + b^3 \label{eq:cubesum}
\end{align*}
$$
```

### 跨多行的单个公式

单个公式很长的时候需要换行，但仅允许生成一个编号时，可以用 split 环境包围公式代码，在需要转行的地方使用 \. split 环境一般用在 equation, gather 环境里面, 可以把单个公式拆成多行, 同时支持 align 那样对齐公式.

split 环境不产生编号, 编号由外面的数学环境产生; 每行需要使用1个&来标识对齐的位置，结束后可使用 \tag{...} 标签编号。 如果 split 环境中某一行不是在二元关系符前面对齐, 需要通过 \quad 等手段设置间距或对齐方式.
$$
% 注意 \tag{...} 编号的位置
\begin{equation}
\begin{split}
\cos 2x &= \cos^2 x - \sin^2 x \\
        &= 2\cos^2 x - 1  
\end{split} \tag{3.1}
\end{equation}
$$

$$
\begin{equation}
\begin{split}
\frac12 (\sin(x+y) + \sin(x-y))
  &= \frac12(\sin x\cos y + \cos x\sin y) \\
  & \quad + \frac12(\sin x\cos y - \cos x\sin y) \\
  &= \sin x\cos y
\end{split}
\end{equation}
$$



```
$$
% 注意 \tag{...} 编号的位置
\begin{equation}
\begin{split}
\cos 2x &= \cos^2 x - \sin^2 x \\
        &= 2\cos^2 x - 1  
\end{split} \tag{3.1}
\end{equation}  
$$

$$
\begin{equation}\label{eq:trigonometric}
\begin{split}
\frac12 (\sin(x+y) + \sin(x-y))
  &= \frac12(\sin x\cos y + \cos x\sin y) \\
  & \quad + \frac12(\sin x\cos y - \cos x\sin y) \\
  &= \sin x\cos y
\end{split}
\end{equation}
$$
```

### 将公式组合为块

最常见的是 `case 环境`, 他在几行公式前面用花括号括起来, 表示几种不同的情况; 每行公式使用 & 分隔, 便是表达式与条件, 例如
$$
\begin{equation}
D(x) = \begin{cases}
1, & \text{if } x \in \mathbb{Q}; \\
0, & \text{if } x \in
     \mathbb{R}\setminus\mathbb{Q}.
\end{cases}
\end{equation}
$$


```
$$
\begin{equation}
D(x) = \begin{cases}
1, & \text{if } x \in \mathbb{Q}; \\
0, & \text{if } x \in
     \mathbb{R}\setminus\mathbb{Q}.
\end{cases}
\end{equation}
$$
```

`gathered环境` 将几行公式居中排列, 组合为一个整体;


$$
\left. \begin{gathered}
S \subseteq T \\
S \supseteq T
\end{gathered} \right\}
\implies S = T
$$


```
$$
\left. \begin{gathered}
S \subseteq T \\
S \supseteq T
\end{gathered} \right\}
\implies S = T  
$$
```

## 括号的其他用法

| 功能           | 语法                                         | 显示         |
| -------------- | -------------------------------------------- | ------------ |
| 圆括号，小括号 | \left( \frac{a}{b} \right)                   | (ab)(ab)     |
| 方括号，中括号 | \left[ \frac{a}{b} \right]                   | [ab][ab]     |
| 花括号，大括号 | \left\{ \frac{a}{b} \right\}                 | {ab}{ab}     |
| 尖括号         | \left \langle \frac{a}{b} \right \rangle     | ⟨ab⟩⟨ab⟩     |
| 单竖线，绝对值 | \left \| \frac{a}{b} \right\|                | 丨abab丨     |
| 双竖线，范式   | \left \| \frac{a}{b} \right \|               | ∥∥ab∥∥‖ab‖   |
| 取整函数       | \left \lfloor \frac{a}{b} \right \rfloor     | ⌊ab⌋⌊ab⌋     |
| 取顶函数       | \left \lceil \frac{c}{d} \right \rceil       | ⌈cd⌉⌈cd⌉     |
| 斜线与反斜线   | \left / \frac{a}{b} \right \backslash        | /ab\/ab\     |
| 上下箭头       | \left \uparrow \frac{a}{b} \right \downarrow | ↑⏐⏐ab⏐↓⏐↑ab↓ |
| 混合括号1      | \left [ 0,1 \right )                         | [0,1)[0,1)   |
| 混合括号2      | \left \langle \psi \right\|                  | ⟨ψ∥⟨ψ‖       |
| 单左括号       | \left \{ \frac{a}{b} \right .                | {ab{ab       |
| 单右括号       | \left . \frac{a}{b} \right \}                | ab}ab}       |

## 希腊字母

| 希腊字母(小写) | 输入                  | 希腊字母(大写) | 输入     |
| -------------- | --------------------- | -------------- | -------- |
| α              | \alpha                | Α              | A        |
| β              | \beta                 | Β              | B        |
| γ              | \gamma                | Γ              | \Gamma   |
| δ              | \delta                | Δ              | \Delta   |
| ε或ϵ           | \epsilon或\varepsilon | Ε              | E        |
| ζ              | \zeta                 | Ζ              | Z        |
| η              | \eta                  | Η              | H        |
| θ或ϑ           | \theta或\vartheta     | Θ              | \Theta   |
| ι              | \iota                 | Ι              | I        |
| κ              | \kappa                | Κ              | K        |
| λ              | \lambda               | Λ              | \Lambda  |
| μ              | \mu                   | Μ              | M        |
| ν              | \nu                   | Ν              | N        |
| ξ              | \xi                   | Ξ              | \Xi      |
| ο              | o                     | Ο              | O        |
| π或ϖ           | \pi或\varpi           | Π              | \Pi      |
| ρ或ϱ           | \rho或\varrho         | Ρ              | P        |
| σ或ς           | \sigma或\varsigma     | Σ              | \Sigma   |
| τ              | \tau                  | Τ              | T        |
| υ              | \upsilon              | Υ              | \Upsilon |
| φ或φ           | \phi或\varphi         | Φ              | \Phi     |
| χ              | \chi                  | Χ              | X        |
| ψ              | \psi                  | Ψ              | \Psi     |
| ω              | \omega                | Ω              | \Omega   |

## 三角函数与逻辑数学字符

| 数学字符 | 输入            | 数学字符 | 输入           |
| -------- | --------------- | -------- | -------------- |
| ±        | \pm             | ×        | \times         |
| ÷        | \div            | \|       | \mid           |
| ∤∤       | \nmid           | ⋅        | \cdot          |
| ∘        | \circ           | ∗        | \ast           |
| ⨀        | \bigodot        | ⨂        | \bigotimes     |
| ⨁        | \bigoplus       | ≤        | \leq           |
| ≥        | \geq            | ≠        | \neq           |
| ≈        | \approx         | ≡        | \equiv         |
| ∑        | \sum            | ∏        | \prod          |
| ∐        | \coprod         | ∅        | \emptyset      |
| ∈        | \in             | ∉        | \notin         |
| ⊂        | \subset         | ⊃        | \supset        |
| ⊆        | \subseteq       | ⊇        | \supseteq      |
| ⋂        | \bigcap         | ⋃        | \bigcup        |
| ⋁        | \bigvee         | ⋀        | \bigwedge      |
| ⨄        | \biguplus       | ⨆        | \bigsqcup      |
| log      | \log            | lg       | \lg            |
| ln       | \ln             | ⊥        | \bot           |
| ∠        | \angle          | 30^∘     | 30 ^ \circ     |
| sin      | \sin            | cos      | \cos           |
| tan      | \tan            | cot      | \cot           |
| ′        | \prime          | ∫        | \int           |
| ∬        | \iint           | ∭        | \iiint         |
| ⨌        | \iiiint         | ∮        | \oint          |
| lim      | \lim            | ∞        | \infty         |
| ∇        | \nabla          | ∵        | \because       |
| ∴        | \therefore      | ∀        | \forall        |
| ∃        | \exists         | ≠        | \not=          |
| ≯        | \not>           | ⊄        | \not\subset    |
| ŷ        | \hat{y}         | yˇ       | \check{y}      |
| y˘       | \breve{y}       | sec      | \sec           |
| ↑        | \uparrow        | ↓        | \downarrow     |
| ⇑        | \Uparrow        | ⇓        | \Downarrow     |
| →        | \rightarrow     | ←        | \leftarrow     |
| ⇒        | \Rightarrow     | ⇐        | \Leftarrow     |
| ⟶        | \longrightarrow | ⟵        | \longleftarrow |
| ⟹        | \Longrightarrow | ⟸        | \Longleftarrow |
|          | \quad           | #        | #              |

参考

[Markdown中编写LaTeX数学公式](http://blog.csdn.net/fzch_struggling/article/details/44998901)
[Markdown下LaTeX公式、编号、对齐](https://www.zybuluo.com/fyywy520/note/82980)
＜＜LaTeX入门＞＞ 刘海洋
