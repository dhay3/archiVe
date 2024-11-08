# hugo - Template Syntax

hugo 使用 go template 的语法来构建 template
## Go template
> 在线编译器
> [https://go.dev/play/](https://go.dev/play/)

在 go template 中 `.` 表示调用 template 的实例 (可以理解成类似 Java 中的 `this`)，可以通过 `{{}}` 语法来获取实例中对应的值或者实例本身。整一个被称为 Action。==Action 外的内容会被原封不动的输出==
同样的在 go template 也是有作用域的，`.` 对应的实例成员及实际本身可以理解为全局变量，而在类似 `{{statement}}...{{end}}` 中的变量为局部变量
例如
```
type Inventory struct {
	Material string
	Count    uint
}
sweaters := Inventory{"wool", 17}
#定义一个 template 实例
tmpl, err := template.New("test").Parse("Here is {{.Count}} items are made of {{.Material}}. We could ues it.")
if err != nil { panic(err) }
#将 sweaters 中的传入 template，并按照 template 中的格式输出到 stdout
err = tmpl.Execute(os.Stdout, sweaters)
if err != nil { panic(err) }
```
`.` 就对应 `sweaters` 实例，`{{.Count}}` 等价于 `sweaters.Count`，`{{.Material}}` 等价于 `sweaters.Material`
最后 `tmpl.Execute(os.Stdout, sweaters)` 输出的结果为
```
Here is 17 items are made of wool. We could ues it.
```
`Here is` 和 `. We could ues it.` 部分会原样输出
### Text and spaces
在 go template 中可以通过 `{{- ` 或者 ` -}}`语法，来去除连续的空格 (注意 hyphen 前或者后是有空格的)
```
func main() {
	type stu struct {
		Name string
		Age  int
	}
	t, err := template.New("foo").Parse(`
		{{- .Name -}}   -     {{- .Age -}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, stu{"harry", 13})
	if nil != err {
		panic(err)
	}
}
```

- `{{- ` 表示去掉 action 的前导空格，即左侧空格
- ` -}}` 表示去掉 action 的后导空格，即右侧空格

所以上述会输出
```
harry-10
```
如果使用如下代码替代上述对应内容
```
	t, err := template.New("foo").Parse(`
		{{ .Name }}   -     {{ .Age }}
`)
```
就会输出
```
		harry   -     13
```
### Actions
#### pipeline
pipeline 指的是数据，例如 `.`，`.Name`, `"test"`, 函数 都是 pipeline (只要能产生数据的都是 pipeline)
`{{pipeline}}` 等价于 `fmt.Print(pipeline)`
```
func main() {
	type stu struct {
		Name string
		Age  int
	}
	t, err := template.New("foo").Parse(`{{"this is a pipeline"}}`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, stu{"harry", 13})
	if nil != err {
		panic(err)
	}
}
```
输出
```
this is a pipeline
```
go template 中还可以使用类似 shell 中的 pipeline (通常也被称为管道符)
```
func main() {
	t, err := template.New("foo").Parse(`{{"this is a pipeline" | len}}`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, nil)
	if nil != err {
		panic(err)
	}
}
```
输出对应的字长
更多例子
```
{{"put" | printf "%s%s" "out" | printf "%q"}}
{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
```
#### comments
`{{/*comments*/}}`
表示注释，不会被 go template 解析输出
```
func main() {
	type stu struct {
		Name string
		Age  int
	}
	t, err := template.New("foo").Parse(`
		{{/*this is comment*/}}
		{{- .Name -}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, stu{"harry", 13})
	if nil != err {
		panic(err)
	}
}
```
输出
```

		harry
```
注意这里是有空格的，注释也是 action，action 外的部分是原样输出的
如果需要去掉注释的空格可以使用 `{{- /*comment*/ -}}` 语法
#### variables
在 go template 可以通过如下方式声明变量并赋值
```
{{$variable := pipeline}}
```
变量声明后可以通过如下方式来赋值
```
{{$variable = pipeline}}
```
然后可以通过 `{{$variable}}` 的方式来获取变量
例如
```
func main() {
	t, err := template.New("foo").Parse(`
{{$my_name:=.Name}}{{$my_power:=.Power}}
My name is {{$my_name}}, and my power is {{$my_power}}.
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, map[string]string{"Name": "batmat", "Power": "money"})
	if nil != err {
		panic(err)
	}
}
```
输出
```
My name is batmat, and my power is money.
```
同样也可以在 `range` 使用 ( 不支持在 `if` 中声明 )
```
range $index, $element := pipeline
range $value := pipeline
```
例如
```
func main() {
	t, err := template.New("foo").Parse(`
{{range $k,$v:=.}}{{$v}}{{"\t"}}{{end}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, []string{"hommie", "amigo", "bro"})
	if nil != err {
		panic(err)
	}
}
```
输出
```
hommie	amigo	bro	
```
同样的 variable 也是有作用域的。例如
```
func main() {
	t, err := template.New("foo").Parse(`
{{if .}}{{$myvar:=.}}if statement {{$myvar}}{{end}}
out of if statement {{$myvar}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, "this is a variable")
	if nil != err {
		panic(err)
	}
}
```
输出
```
panic: template: foo:3: undefined variable "$myvar"
```
因为 `{{$myvar}}` 只在 `{{if pipeline}} ... {{end}}` 中生效
#### string
`{{"string"}}`
在 go template 中如果字符串在 `{{}}` 内，同样也会被原样输出，特殊字符会被转译
```
func main() {
	t, err := template.New("foo").Parse(`
	{{"this is a string\n"}}
  	{{.}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, "this is a template")
	if nil != err {
		panic(err)
	}
}
```
输出
```

	this is a string

  this is a template
```
#### conditon flow

- `{{if pipeline}} T1 {{end}} `

If the value of the pipeline is empty, no output is generated; 
otherwise, T1 is executed. The empty values are false, 0, any nil pointer or interface value, and any array, slice, map, or string of length zero. 	Dot is unaffected.
```
if pipeline {
	T1
}
```

- `{{if pipeline}} T1 {{else}} T0 {{end}}`

If the value of the pipeline is empty, T0 is executed; 	
otherwise, T1 is executed. Dot is unaffected.
```
if pipeline {
	T1
} else {
	T0
}
```

- `{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}`

	To simplify the appearance of if-else chains, the else action of an if may include another if directly; the effect is exactly the same as writing `{{if pipeline}} T1 {{else}}{{if pipeline}} T0 {{end}}{{end}}`
```
if pipeline {
	T1
} else if pipeline {
	T0
}
```

- `{{break}}`

和其他编程语言中的 `break` 相同

- `{{continue}}`

和其他编程语言中的 `continue` 相同
#### loop

- `{{range pipeline}} T1 {{end}}`

	The value of the pipeline must be an array, slice, map, or channel. If the value of the pipeline has length zero, nothing is output; otherwise, dot is set to the successive elements of the array, slice, or map and T1 is executed. If the value is a map and the keys are of basic type with a defined order, the elements will be visited in sorted key order.
```
func main() {
	type stu struct {
		Name string
		Age  int
	}
	t, err := template.New("foo").Parse(`
{{range .}}<h1>{{.Name}}</h1>{{end}}`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, []stu{{"harry", 13}, {"henry", 14}})
	if nil != err {
		panic(err)
	}
}
```
输出
```
<h1>harry</h1><h1>henry</h1>
```
这里的 `{{range .}}` 中的 dot 是调用的 template 的实例，即例子中的 `[]stu{{"harry", 13}, {"henry", 14}}`
```
for _, v := range []stu{{"harry", 13}, {"henry", 14}} {
		fmt.Print(v.Name)
	}
}
```

- `{{range pipeline}} T1 {{else}} T0 {{end}}`

The value of the pipeline must be an array, slice, map, or channel. If the value of the pipeline has length zero, dot is unaffected and T0 is executed; otherwise, dot is set to the successive elements of the array, slice, or map and T1 is executed.
```
func main() {
	type stu struct {
		Name    string
		Age     int
		Hobbies []string
	}
	t, err := template.New("foo").Parse(`
			{{- range .Hobbies}}{{.}} {{else}}2{{end -}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, stu{"jackie", 10, []string{"football", "basketball", "ping pong"}})
	if nil != err {
		panic(err)
	}
}

```
输出
```
football basketball ping pong 
```
注意这里的 `{{.}}` 类似 for loop 中的 item
```
for item in hobbies:
	if item:
  	print(item)
	else:
  	print(2)
```
#### with

- `{{with pipeline}} T1 {{end}}`

当 pipeline 的值不为空，dot 的值为 pipeline，然后执行 T1
```
func main() {
	type stu struct {
		Name string
		Age  int
	}
	t, err := template.New("foo").Parse(`
			{{- with .Name}}My name is {{.}}{{end -}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, stu{"dasiy", 10})
	if nil != err {
		panic(err)
	}
}
```
输出
```
My name is dasiy
```
例子中即 `{{.Name}}` 的值不为空，则在作用域内的 `{{.}}` 值为 `{{.Name}}`
对应的 go 中没有 `with` 关键字语法，上面的内容等价于
```
if pipeline:
	temp := pipeline
```
temp 指代 `{{.}}`

- `{{with pipeline}} T1 {{else}} T0 {{end}}`

当 pipeline 的值不为空，dot 的值为 pipeline，然后执行 T1，否则执行 T0
```
func main() {
	type stu struct {
		Name   string
		Age    int
		Gender int
	}
	t, err := template.New("foo").Parse(`
			{{- with .Gender}}I is a {{.}}{{else}}Mind your own bussiness{{end -}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, stu{"dasiy", 10, 0})
	if nil != err {
		panic(err)
	}
}

```
输出
```
Mind your own bussiness
```
这里会执行 T1 的结果，是因为在 go 中和 python 一样， `0` 同样表示空值
#### template
类似于 Vue 中的 template，预定义后可以被调用
在 go template 中通过 `{{define "templateName"}}T1{{end}}`来预定义 template (可以抽象成函数)
然后通过如下几种方式来调用 template

- `{{template "templateName"}}`

调用指定的 template
```
func main() {
	t, err := template.New("foo").Parse(`
		{{- define "T1"}}define T1{{end -}}
		{{template "T1"}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, []int{1})
	if nil != err {
		panic(err)
	}
}
```
输出
```
define T1
```

- `{{template "templateName" pipeline}}`

调用指定的 template，templateName 中对应的 `{{.}}` 的值就为 pipeline。
```
func main() {
	t, err := template.New("foo").Parse(`
		{{- define "T1"}}define T1, {{.}}{{end -}}
  	{{template "T1" "echo T1"}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, []int{1})
	if nil != err {
		panic(err)
	}
}
```
输出
```
define T1, echo T1
```
可以将其抽象成一个函数，pipeline 就是一个入参
```
function templateName (x): {{.}} = x
templateName(pipeline)
```
go template 还有一种特殊的用法 `{{block "name" pipeline}}T1{{end}}`  ，等价于先声明 `{{define "name"}}T1{{end}}` template，然后执行 `{{template "name" pipeline}}` 
```
func main() {
	t, err := template.New("foo").Parse(`
		{{- block "T1" "echo block"}}this is a block template, {{.}}{{end}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, []int{1})
	if nil != err {
		panic(err)
	}
}
```
输出
```
this is a block template, echo block
```
#### functions
在 go template 中的可以通过如下方式来调用函数
```
{{functionName args1 args2}}
```
如果 args 也是函数，需要使用 parentheses
```
{{functionName (func1 arg1 arg2) (func2 arg1 arg2)}}
```
例如
```
type stu struct {
	Name   string
	Age    int
	Gender int
}

func (s stu) Say(content string) string {
	return s.Name + content
}

func main() {

	t, err := template.New("foo").Parse(`{{.Say ", hello world"}}`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, stu{"marry", 10, 0})
	if nil != err {
		panic(err)
	}
}
```
输出
```
marry, hello world
```
需要注意的一点是，function body 有且只能含有有一句 `return` 时才有效，即一下函数均不能被调用
```
//panic: template: foo:1:2: executing "foo" at <.F1>: can't call method/function "F2" with 0 results
func (s stu) F1(content string) {
	fmt.Print(content)
}

//./prog.go:18:9: syntax error: unexpected d at end of statement
func (s stu) F2(content string) {
	string d = s.Name + content
	return d
}
```
##### built-in
go template 有一系列内建的函数

- `{{not}}`

取反
```
func main() {
	t, err := template.New("foo").Parse(`
		{{- if not .}}1{{else}}2{{end}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, nil)
	if nil != err {
		panic(err)
	}
}
```
输出 `1`

- `{{or}}`

逻辑或，
```
func main() {
	t, err := template.New("foo").Parse(`
		{{- range .}}{{if (gt . 0) or (lt . 3)}}{{.}}{{end}}{{end}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, []int{0,-2,2,5})
	if nil != err {
		panic(err)
	}
}
```
输出 `0-225`

- `{{and}}`

逻辑与
```
func main() {
	t, err := template.New("foo").Parse(`
		{{- range .}}{{if and (gt . 0) (lt . 3)}}{{.}}{{end}}{{end}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, []int{0, -2, 2, 5})
	if nil != err {
		panic(err)
	}
}
```
输出 `2`

- `{{print}}`

等价于 `fmt.Print`

- `{{printf}}`

等价于 `fmt.Printf`

- `{{println}}`

等价于 `fmt.Println`

- `{{slice elements start end}}`

截取切片，逻辑如下
```
function Slice(slice T[],start int,end int) T[]{
	return slice[start:end]
}
```

- `{{len elements}}`

获取长度，逻辑如下
```
function Len(elements T[]) int {
	return len(elements)
}
```

- `{{html}}`
- `{{js}}`
#### comparision operators
在 go template 中使用类似 shell 中的 comparision operators，实际也可以归类于 function

- `{{eq arg1 arg2}}`

equal，`==`

- `{{ne args1 arg2}}`

not equal，`!=`

- `{{lt args1 arg2}}`

less than，`<`

- `{{le args1 arg2}}`

less equal，`<=`

- `{{gt args1 arg2}}`

greater than，`>`

- `{{ge args1 arg2}}`

greater equal，`>=`
对于 `eq` ，还支持多个参数对比
```
eq arg1 arg2 arg3 ...
```
类似等价于
```
arg1==arg2 || arg1==arg3 || arg1==arg4 ...
```
需要注意的是 `||` 和 go 中的不一样，并不表示短路或，而是表示逻辑或，所有的 comparision 都会校验
## Hugo template
hugo template 在 go template 的基础上实现了自己的一些特殊功能
### Partial
导入额外的 template。需要注意的一点是，导入的内容需要在 `layouts/partials` 目录下
例如，导入 `layouts/partials/header.html` 中的内容到当前 template
```
{{ partial "header.html" . }}
```
当然 hugo 也支持 go template 中的 `{{template "templateName" .}}` 语法，下面例子即表示导入 `"_internal/opengraph.html"` 中的内容到当前目录
```
{{ template "_internal/opengraph.html" . }}
```



**references**

1. [https://gohugo.io/templates/introduction/](https://gohugo.io/templates/introduction/)
2. [https://pkg.go.dev/text/template#pkg-overview](https://pkg.go.dev/text/template#pkg-overview)
3. [https://cloud.tencent.com/developer/article/1683688](https://cloud.tencent.com/developer/article/1683688)
