# HTML Table Advanced Features

ref

https://developer.mozilla.org/en-US/docs/Learn/HTML/Tables/Advanced

这里只补充一些未学习到或者忘记的知识

## structure

在 table 中通常使用`<thead>`,`<tfoot>`,`<tbody>`使结构更加清晰

thead 表示表头

tfoot 通常用在最后一行 previous rows summed

tbody 除 thead 和 tfoot 外的部分

## col/colgroup

HTML 中可以通过`<col>`和`<colgroup>`标签来为 table columns 设置 CSS（其实也可是使用伪类`:nth-child`实现）

例如只对第 2 列赋值 yellow bg color。这里需要注意的是即使第一列不做修改，也需要使用一个`<col/>`表示当前列同时做定位的功能

```
<table>
  <colgroup>
    <col />
    <col style="background-color: yellow" />
  </colgroup>
  <tr>
    <th>Data 1</th>
    <th>Data 2</th>
  </tr>
  <tr>
    <td>Calcutta</td>
    <td>Orange</td>
  </tr>
  <tr>
    <td>Robots</td>
    <td>Jazz</td>
  </tr>
</table>
```

<table>
  <colgroup>
    <col />
    <col style="background-color: yellow" />
  </colgroup>
  <tr>
    <th>Data 1</th>
    <th>Data 2</th>
  </tr>
  <tr>
    <td>Calcutta</td>
    <td>Orange</td>
  </tr>
  <tr>
    <td>Robots</td>
    <td>Jazz</td>
  </tr>
</table>

