# HTML Advanced text

ref

https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Advanced_text_formatting

## Description lists

除了 ordered lists 和 unordered lists 还有第 3 种 lists —— description lists `<dl>`

description lists 包含

- `<dt>` description term
- `<dd>` description definition

例如

```
<dl>
  <dt>soliloquy</dt>
  <dd>
    In drama, where a character speaks to themselves, representing their inner
    thoughts or feelings and in the process relaying them to the audience (but
    not to other characters.)
  </dd>
</dl>
```

## Quotes

> 实际在前端渲染中没有什么效果，只是标明引用出自那里

用于标明引用部分，cite attribute 标明 the source of the quote inside a cite

```
<blockquote  cite="https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote">
  <p>
  this is a quotation
  </p>
</blockquote>

```

除`<blockquote>`外还有 inline quotations `<q>`

```
<q cite="https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q">intended
for short quotations that don't require paragraph breaks.</q>
```

## Abbreviations

`<abbr>` is used to wrap around an abbreviation acronym

前端如果不添加 title attribute 不会显示任何效果

```
<abbr title="Reverend">Rev.</abbr> 
```

<abbr title="Reverend">Rev.</abbr> 

## address

`<address>` is for marking up contact details

```
<address>Chris Mills, Manchester, The Grim North, UK</address>
```

<address>Chris Mills, Manchester, The Grim North, UK</address>

效果是斜体显示

## superscript and subscript

superscript 上角标

subscript 下角标

```
<p>My birthday is on the 25<sup>th</sup> of May 2001.</p>
<p>
  Caffeine's chemical formula is
  C<sub>8</sub>H<sub>10</sub>N<sub>4</sub>O<sub>2</sub>.
</p>
<p>If x<sup>2</sup> is 9, x must equal 3 or -3.</p>
```

> tips markdonw 中不会显示使用 LaTeX 替代

My birthday is on the 25th of May 2001.

  Caffeine's chemical formula is  C8H10N4O2.

## snippet

HTML 还提供了代码片段的标签

1. `<code>`

   for making up computer code

2. `<pre>`

   for retaining whitespace

   ```
   <pre><code>const para = document.querySelector('p');
   
   para.onclick = function() {
     alert('Owww, stop poking me!');
   }</code></pre>
   ```

   <pre><code>const para = document.querySelector('p');
   para.onclick = function() {
     alert('Owww, stop poking me!');
   }</code></pre>

3. `<var>`

   for marking up variable

4. `<kbd>`

   for makring up keyboard (eg <kbd>ctrl</kbd>/<kbd>shift</kbd>)

## timedate

<time datetime="2016-01-20">20 January 2016</time>

没有啥实际效果，所以不记录，具体查看 MDN docs
