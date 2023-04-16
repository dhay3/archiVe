# CSS - Functions

> 具体 function 查看 MDN references 部分

如果 property’s value 是 simple keywords or numeric values, 通常支持 function

## calc

表示计算大小

```
<style>
.outer {
  border: 5px solid black;
}

.box {
  padding: 10px;
  width: calc(90% - 30px);
  background-color: rebeccapurple;
  color: white;
}
<style>
<div class="outer"><div class="box">The inner box is 90% - 30px.</div></div>
```

## rotate

一般用于 transform property, 表示旋转

```
<style>
.box {
  margin: 30px;
  width: 100px;
  height: 100px;
  background-color: rebeccapurple;
  transform: rotate(0.8turn);
}
<style>
<div class="box"></div>
```



**references**

1. https://developer.mozilla.org/en-US/docs/Learn/CSS/First_steps/How_CSS_is_structured

