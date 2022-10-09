# CSS Selectors

ref

https://developer.mozilla.org/en-US/docs/Learn/CSS/First_steps/Getting_started

> 两两以上的组合逻辑相同，所以不在这里展示

## Symbolic Selectors

`*`

all elements

## Single Selectors

### ID selector

```

# {
  color: orange;
  font-weight: bold;
}
```

### Class selector

```
li.special,
span.special {
  color: orange;
  font-weight: bold;
}
```

### element selector

### attribute selector

## Combination Selectors

### element.class

any `<li>` element that has a class of special

```
li.special,
span.special {
  color: orange;
  font-weight: bold;
}
```

### element element

any `<em>` element that is inside any `<li>`

```
li em {
  color: rebeccapurple;
}
```

### element + element

the `<p>` element(only effect one) just directly after any `<h1>` element

```
h1 + p {
  font-size: 200%;
}
```

### element > element

```

```



## State Selectors

具体 State  Selectors 查看 MDN docs

```
/* all link haven't been visited */
a:link {
  color: pink;
}

/* all link element has been visted */
a:visited {
  color: green;
}

/* all link when users hovers over it */
a:hover {
  text-decoration: none;
}
```

