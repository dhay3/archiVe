# go template

## dot

通过`{{}}`获取传入Action的值

```
data is: {{ . }}
```

如果模板中传入的数据的数据是基本类型，可以使用`.`来表示

```go
func main() {
   t, err := template.ParseFiles("src/6.template/t.txt")
   if nil !=err {
      panic(err)
   }
   err = t.Execute(os.Stdout, "hell world")
   if nil != err {
      panic(err)
   }
}
```

如果传入的数据是复合对象，可以通过`.field`来取得(只能获取一个属性)。如果是属性还是复合对象，可以通过`{{.field.field}}`来获得，中间没有空格。

```
type stu struct {
	Name string
	Age  int
}

func main() {
	t, err := template.New("foo").Parse(`
		{{.Name}}-{{.Age}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, stu{"zs", 10})
	if nil != err {
		panic(err)
	}
}
```

## comment

`{{/*comment*/}}`

## trim

curly brackets + dash + space 表示去掉空格连续

- `{{- `

  表示去掉前一个action的后缀空格

- ` -}}`

  表示去掉后一个action的前缀空格

```
type stu struct {
	Name string
	Age  int
}

func main() {
	//t, err := template.ParseFiles("src/6.template/t.txt")
	t, err := template.New("foo").Parse(`
		{{- .Name -}}   -     {{- .Age -}}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, stu{"zs", 10})
	if nil != err {
		panic(err)
	}
}
#output
zs-10
```

## Action

Action之外的内容原样输出，空值为0，false，nil，interface

1. `{{if pipeline}} T1 {{end}}`

   如果data不为空执行T1

2. `{{if pipeline}} T1 {{else}} T0 {{end}}`

   如果data不为空执行T1否则执行end

   ```
   func main() {
   	t, err := template.New("foo").Parse(`
   			{{- if .Name}}111{{else}}222{{end -}}
   `)
   	if nil != err {
   		panic(err)
   	}
   	err = t.Execute(os.Stdout, stu{"zs", 10})
   	if nil != err {
   		panic(err)
   	}
   }
   
   #output
   111
   ```

3. `{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}`

4. `{{range pipeline}} T1 {{end}}`

   遍历+T，pipeline的数据类型为array，slice，map or channel，否则不输出

5. `{{range pipeline}} T1 {{else}} T0 {{end}}`

   如果7不能执行就执行T0

   ```
   func main() {
   	t, err := template.New("foo").Parse(`
   			{{- range .Hobbies}}1{{else}}2{{end -}}
   `)
   	if nil != err {
   		panic(err)
   	}
   	err = t.Execute(os.Stdout, stu{"zs", 10, []string{"football","basketball","ping pong"}})
   	if nil != err {
   		panic(err)
   	}
   }
   ```

6. `{{with pipeline}} T1 {{end}}`

   如果pipeline的值不为空，dot设置为pipeline的值，执行T1。否则将没有数据输出

7. `{{with pipeline}} T1 {{else}} T0 {{end}}`

   如果pipeline的值不为空，dot设置为pipeline的值，执行T1。否则执行T0



## pipeline

和Unix中的管道符一样，用于接受stdout转为stdin

```
func main() {
	t, err := template.New("foo").Parse(`
		{{.Name | printf "%s"}}
`)
	if nil != err {
		panic(err)
	}

	err = t.Execute(os.Stdout, stu{"zs", 10})
	if nil != err {
		panic(err)
	}
}
```

## template

```
func main() {
	//t, err := template.ParseFiles("src/6.template/t.txt")
	t, err := template.New("foo").Parse(`
		{{define "T1"}}define T1{{end}}
		{{- .Name -}}-{{- .Age}}
		{{template "T1" }}
`)
	if nil != err {
		panic(err)
	}
	err = t.Execute(os.Stdout, stu{"zs", 10})
	if nil != err {
		panic(err)
	}
}

#output
		zs-10
		define T1
		
----
可以通过dot传参

type names struct {
	Name string
}
type stu struct {
	Names *names
	Age int
	Hobbies map[string]string
}

func main() {
	t, err := template.New("foo").Funcs(template.FuncMap{"show": Show}).Parse(`
		{{- define "default"}}{{.Name}}{{end}}
		{{template "default" .Names}}
`)
	if nil != err {
		panic(err)
	}

	err = t.Execute(os.Stdout, stu{&names{"zs"}, 10,map[string]string{
		"a":"foo",
		"b":"bar"}})
	if nil != err {
		panic(err)
	}
}
```

- define：

  `{{define "T1"}}ONE{{end}}`

  和全局变量一样，~~必须定义在首部~~，类似于vue中的template

- invoke

  ```
  {{define "T1"}}ONE{{end}}
  {{define "T2"}}TWO{{end}}
  {{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
  {{template "T3"}}
  ```

  使用`{{template definedName}}`来调用模板，这里定义了T3。通过T3调用T1和T2

## built-in

- 数学运算同shell

- `and x y z`

  ```
  func main() {
  	t, err := template.New("foo").Parse(`
  		{{- and .Name  .Age .Gender -}}
  `)
  	if nil != err {
  		panic(err)
  	}
  
  	err = t.Execute(os.Stdout, stu{"zs", 0,true})
  	if nil != err {
  		panic(err)
  	}
  }
  #output
  0
  ```

  返回第一个为空的值或者返回最后一个值，`if !x then x elif !y then y else z`

- index

  如果没有对应的值就取空
  
  ```
  func main() {
     t, err := template.New("foo").Parse(`
        {{- index .Hobbies 0}}
        {{- index .Hobbies 1}}
        {{- index .Hobbies 2}}
  `)
     if nil != err {
        panic(err)
     }
  
     err = t.Execute(os.Stdout, stu{"zs", 1, []string{"a","b","c"}})
     if nil != err {
        panic(err)
     }
  }
  -------------
  type stu struct {
  	Name string
  	Hobbies map[string]string
  }
  
  func main() {
  	t, err := template.New("foo").Funcs(template.FuncMap{"show": Show}).Parse(`
  		{{index .Hobbies "a"}}-{{index .Hobbies "c"}}
  `)
  	if nil != err {
  		panic(err)
  	}
  
  	err = t.Execute(os.Stdout, stu{"zs", map[string]string{
  		"a":"foo",
  		"b":"bar"}})
  	if nil != err {
  		panic(err)
  	}
  }
  ```

...

## ``

表示字符常量

## func

go template中支持自定义函数

```
type stu struct {
	Name string
	Age  int
}
//必须要有返回值
func  Show(name string) string {
	return name
}
func main() {
	t, err := template.New("foo").Funcs(template.FuncMap{"show": Show}).Parse(`
		{{show .Name}}
`)
	if nil != err {
		panic(err)
	}

	err = t.Execute(os.Stdout, stu{"zs", 1})
	if nil != err {
		panic(err)
	}
}
#output
zs
```

