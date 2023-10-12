CNBlog美化日记

[TOC]

## Markdown

### 修改Markdown文本风格

参考：

https://www.cnblogs.com/nowgood/p/3bokejiemain.html#_nav_7

https://www.cnblogs.com/liuxianan/p/custom-cnblogs-highlight-theme.html

在[hightlight.js](https://highlightjs.org/static/demo/)挑选一个喜爱的主题，到[github](https://github.com/highlightjs/highlight.js/tree/master/src/styles)下载自爱的风格, 将css中的 .hljs 替换全部成 `.cnblogs-markdown .hljs`。我这里使用Shades-of-Purple，注意的是如果无法显示背景颜色，给background最高优先级即可

示例代码：

```css
/**
 * Shades of Purple Theme — for Highlightjs.
 *
 * @author (c) Ahmad Awais <https://twitter.com/mrahmadawais/>
 * @link GitHub Repo → https://github.com/ahmadawais/Shades-of-Purple-HighlightJS
 * @version 1.5.0
 */
pre {
    /*控制代码不换行*/
    white-space: pre;
    word-wrap: normal;
}

.cnblogs-markdown .hljs {
  display: block;
  overflow-x: auto;
  /* Custom font is optional */
  /* font-family: 'Operator Mono', 'Fira Code', 'Menlo', 'Monaco', 'Courier New', 'monospace';  */
  padding: 0.5em;
  background: #2d2b57 !important; 
  font-weight: normal;
}

.cnblogs-markdown .hljs-title {
  color: #fad000;
  font-weight: normal;
}

.cnblogs-markdown .hljs-name {
  color: #a1feff;
}

.cnblogs-markdown .hljs-tag {
  color: #ffffff;
}

.cnblogs-markdown .hljs-attr {
  color: #f8d000;
  font-style: italic;
}

.cnblogs-markdown .hljs-built_in,
.cnblogs-markdown .hljs-selector-tag,
.cnblogs-markdown .hljs-section {
  color: #fb9e00;
}

.cnblogs-markdown .hljs-keyword {
  color: #fb9e00;
}

.cnblogs-markdown .hljs,
.cnblogs-markdown .hljs-subst {
  color: #e3dfff;
}

.cnblogs-markdown .hljs-string,
.cnblogs-markdown .hljs-attribute,
.cnblogs-markdown .hljs-symbol,
.cnblogs-markdown .hljs-bullet,
.cnblogs-markdown .hljs-addition,
.cnblogs-markdown .hljs-code,
.cnblogs-markdown .hljs-regexp,
.cnblogs-markdown .hljs-selector-class,
.cnblogs-markdown .hljs-selector-attr,
.cnblogs-markdown .hljs-selector-pseudo,
.cnblogs-markdown .hljs-template-tag,
.cnblogs-markdown .hljs-quote,
.cnblogs-markdown .hljs-deletion {
  color: #4cd213;
}

.cnblogs-markdown .hljs-meta,
.cnblogs-markdown .hljs-meta-string {
  color: #fb9e00;
}

.cnblogs-markdown .hljs-comment {
  color: #ac65ff;
}

.cnblogs-markdown .hljs-keyword,
.cnblogs-markdown .hljs-selector-tag,
.cnblogs-markdown .hljs-literal,
.cnblogs-markdown .hljs-name,
.cnblogs-markdown .hljs-strong {
  font-weight: normal;
}

.cnblogs-markdown .hljs-literal,
.cnblogs-markdown .hljs-number {
  color: #fa658d;
}

.cnblogs-markdown .hljs-emphasis {
  font-style: italic;
}

.cnblogs-markdown .hljs-strong {
  font-weight: bold;
}
```

### 添加复制功能

参考:

https://www.cnblogs.com/byho/p/13180288.html

```css
/*markdown添加按钮*/
.cnblogs-markdown pre {
    position: relative;
}

.cnblogs-markdown pre > span {
    position: absolute;
    top: 0;
    right: 0;
    border-radius: 2px;
    padding: 0 10px;
    font-size: 12px;
    background: rgba(0, 0, 0, 0.4);
    color: #fff;
    cursor: pointer;
    display: none;
}

.cnblogs-markdown pre:hover > span {
    display: block;
}

.cnblogs-markdown pre > .copyed {
    background: #67c23a;
}
```

添加JS代码

```js
<script  src="https://blog-static.cnblogs.com/files/kikochz/clipboard.js"></script>
<script  src="https://blog-static.cnblogs.com/files/kikochz/cp.js"></script>

```

### Latex渲染

在页首HTML代码中添加

==有一个自动换行的细节，如果你的一些列数学公式通过&的方式在某个地方对其（比如在=这里对齐），自动换行会不起作用。==

```js
<script type="text/javascript" async src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.1/MathJax.js?config=TeX-MML-AM_CHTML">
<script type="text/x-mathjax-config" src="https://blog-static.cnblogs.com/files/kikochz/latex.js#"></script>
```

### 标题

参考：

https://www.cnblogs.com/ZhaoxiCheung/p/CustomizeBlog.html

```css
#cnblogs_post_body h2 {
    padding-bottom: 4px;
    border-bottom: 2px solid #999;
    color: #008891;
    font-family: "Monaco", "微软雅黑" ,Arial;
    font-size: 24px;
    font-weight: bold;
    line-height: 24px;
    margin: 20px 0 !important;
    padding: 10px 0px 10px 5px;
    text-shadow: 2px 1px 2px lightgrey;
}

#cnblogs_post_body h2:hover {
    color: rgb(255, 102, 0);
}
#cnblogs_post_body h3 {
    padding-bottom: 4px;
    color: #016e8c;
    font-family: "Monaco", "微软雅黑" ,Arial;
    font-size: 18px;
    font-weight: bold;
    line-height: 23px;
    margin: 0px 0 !important;
    padding: 5px 0px 5px 0px;
    text-shadow: 2px 1px 2px lightgrey;
}
#cnblogs_post_body h4 {
    padding-bottom: 4px;
    color: #016e8c;
    font-family: "Monaco", "微软雅黑" ,Arial;
    font-size: 16px;
    font-weight: bold;
    line-height: 23px;
    margin: 0px 0 !important;
    padding: 5px 0px 5px 0px;
    text-shadow: 2px 1px 2px lightgrey;
}
```

## 背景图片

参考：

https://www.cnblogs.com/miluluyo/p/setites.html

这里使用bing的随机图片

```css
body:after {
    background: url(http://api.dujin.org/bing/1366.php) center/cover no-repeat;
    content: '';
    background-repeat: no-repeat;
    background-position: center;
    opacity: 0.4;
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: -1;
    -webkit-filter: grayscale(40%);
    -moz-filter: grayscale(40%);
    -ms-filter: grayscale(40%);
    -o-filter: grayscale(40%);
    filter: grayscale(40%);
    filter: gray;
}
```

## 背景特效

### 矩阵

```js
<div><canvas id="matrixBG"></canvas></div>
<script language="javascript" type="text/javascript" src="https://blog-static.cnblogs.com/files/kikochz/matrix.js"></script>
```

![gif5](https://github.com/dhay3/image-repo/raw/master/20210518/gif5.4agknhvfrcs0.gif)

### 彩带

```js
<script  src="https://blog-static.cnblogs.com/files/kikochz/ribbon.js"></script>
```

效果图:

![1](https://github.com/dhay3/image-repo/raw/master/20210518/1.18l2spkys5z4.gif)

### 粒子球

```js
<script src="https://blog-static.cnblogs.com/files/kikochz/canvas.js"></script>
```

![gif2](https://github.com/dhay3/image-repo/raw/master/20210518/gif2.2ejt6yj1ktq8.gif)

### 流星

```js
<script src="https://blog-static.cnblogs.com/files/kikochz/vendors.js"></script>
<script src="https://blog-static.cnblogs.com/files/kikochz/borealsky.js"></script>
```

![gif3](https://github.com/dhay3/image-repo/raw/master/20210518/gif3.1a9xfadzwzgg.gif)

## 载入条

https://github.com/rstacruz/nprogress

CSS

```css
/* Make clicks pass-through */
#nprogress {
    pointer-events: none;
}

#nprogress .bar {
    background: deeppink;

    position: fixed;
    z-index: 1031;
    top: 0;
    left: 0;

    width: 100%;
    height: 10px;
}

/* Fancy blur effect */
#nprogress .peg {
    display: block;
    position: absolute;
    right: 0px;
    width: 100px;
    height: 100%;
    box-shadow: 0 0 10px #29d, 0 0 5px #29d;
    opacity: 1.0;

    -webkit-transform: rotate(3deg) translate(0px, -4px);
    -ms-transform: rotate(3deg) translate(0px, -4px);
    transform: rotate(3deg) translate(0px, -4px);
}

/* Remove these to get rid of the spinner */
#nprogress .spinner {
    display: block;
    position: fixed;
    z-index: 1031;
    top: 15px;
    right: 15px;
}

#nprogress .spinner-icon {
    width: 18px;
    height: 18px;
    box-sizing: border-box;

    border: solid 2px transparent;
    border-top-color: deeppink;
    border-left-color: deeppink;
    border-radius: 50%;

    -webkit-animation: nprogress-spinner 400ms linear infinite;
    animation: nprogress-spinner 400ms linear infinite;
}

.nprogress-custom-parent {
    overflow: hidden;
    position: relative;
}

.nprogress-custom-parent #nprogress .spinner,
.nprogress-custom-parent #nprogress .bar {
    position: absolute;
}

@-webkit-keyframes nprogress-spinner {
    0%   { -webkit-transform: rotate(0deg); }
    100% { -webkit-transform: rotate(360deg); }
}
@keyframes nprogress-spinner {
    0%   { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
}

```

JS

```js
<script  src="https://blog-static.cnblogs.com/files/kikochz/loader.js"></script>
<script>
$(document).ready(function(){
	    NProgress.start();
	    NProgress.done();
})
</script>
```

## 滚动条

```css
/*滚动条整体样式*/
body::-webkit-scrollbar {
    width: 5px;
    height: 1px;
}
/*滚动条滑块*/
body::-webkit-scrollbar-thumb {
    border-radius: 5px;
    -webkit-box-shadow: inset 0 0 5px rgba(0,0,0,0.2);
    background: deeppink;
}
/*滚动条轨道*/
body::-webkit-scrollbar-track {
    -webkit-box-shadow: inset 0 0 1px rgba(0,0,0,0);
    border-radius: 5px;
    background: transparent;
}
```

## 目录

参考：

https://www.cnblogs.com/sakuraph/p/5814060.html

```js
<div class="fixedIndexs" style="position: fixed;bottom: 40px;display: none"></div>
<script language="javascript" type="text/javascript" src="https://blog-static.cnblogs.com/files/kikochz/list.js"></script>
```

![gif4](https://github.com/dhay3/image-repo/raw/master/20210518/gif4.4ili5b75iqe0.gif)

## Github Corner/ribbon

左上角

```js
<a href="https://github.com/dhay3" title="我的站点" target="_Blank" class="github-corner" aria-label="View source on Github"><svg width="80" height="80" viewBox="0 0 250 250" style="fill:#64CEAA; color:#fff; position: absolute; top: 0; border: 0; left: 0; transform: scale(-1, 1);" aria-hidden="true"><path d="M0,0 L115,115 L130,115 L142,142 L250,250 L250,0 Z"></path><path d="M128.3,109.0 C113.8,99.7 119.0,89.6 119.0,89.6 C122.0,82.7 120.5,78.6 120.5,78.6 C119.2,72.0 123.4,76.3 123.4,76.3 C127.3,80.9 125.5,87.3 125.5,87.3 C122.9,97.6 130.6,101.9 134.4,103.2" fill="currentColor" style="transform-origin: 130px 106px;" class="octo-arm"></path><path d="M115.0,115.0 C114.9,115.1 118.7,116.5 119.8,115.4 L133.7,101.6 C136.9,99.2 139.9,98.4 142.2,98.6 C133.8,88.0 127.5,74.4 143.8,58.0 C148.5,53.4 154.0,51.2 159.7,51.0 C160.3,49.4 163.2,43.6 171.4,40.1 C171.4,40.1 176.1,42.5 178.8,56.2 C183.1,58.6 187.2,61.8 190.9,65.4 C194.5,69.0 197.7,73.2 200.1,77.6 C213.8,80.2 216.3,84.9 216.3,84.9 C212.7,93.1 206.9,96.0 205.4,96.6 C205.1,102.4 203.0,107.8 198.3,112.5 C181.9,128.9 168.3,122.5 157.7,114.1 C157.9,116.9 156.7,120.9 152.7,124.9 L141.0,136.5 C139.8,137.7 141.6,141.9 141.8,141.8 Z" fill="currentColor" class="octo-body"></path></svg></a><style>.github-corner:hover .octo-arm{animation:octocat-wave 560ms ease-in-out}@keyframes octocat-wave{0%,100%{transform:rotate(0)}20%,60%{transform:rotate(-25deg)}40%,80%{transform:rotate(10deg)}}@media (max-width:500px){.github-corner:hover .octo-arm{animation:none}.github-corner .octo-arm{animation:octocat-wave 560ms ease-in-out}}</style>
```

右上角

https://tholman.com/github-corners/

```js
<a href="https://github.com/dhay3"><img style="position: absolute; top: 76px; right: 0; border: 0" alt="Fork me on GitHub" src="https://cdn.jsdelivr.net/gh/yanglr/Beautify-cnblogs/images/github-pendant-rightCorner.svg?sanitize=true"></a>
```

![Snipaste_2020-08-27_15-12-12](https://github.com/dhay3/image-repo/raw/master/20210518/Snipaste_2020-08-27_15-12-12.2gwhmzzk65w0.png)

## 修改图标

```js
<script>
$("head").append('<link type="image/x-icon" rel="icon"  href="https://images.cnblogs.com/cnblogs_com/kikochz/1831554/o_200828132513favicon.png">')
</script>
```

## 在公告栏添加滚动文字

```html
<marquee><a href="#"><font color="#0066FF" size="4">假如可以带粉笔进入迷宫，以纯蓝标记每一处通往灾祸的岔口：“我到这儿必将不再受诱”，它们将变得可以承受。</marquee>
```

## 人体时钟

```js
<embed wmode="transparent" src="https://files.cnblogs.com/files/enjoy233/honehone_clock_tr.swf" quality="high" bgcolor="#FDF6E3" width="200" height="120" name="honehoneclock" align="middle" allowscriptaccess="always"type="application/x-shockwave-flash" pluginspage="http://www.macromedia.com/go/getflashplayer">
```

## 隐藏默认项

隐藏推荐按钮

```css
.diggit {
	display: none;
}

```

隐藏反对按钮

```css
.buryit {
	display: none;
}

.comment_bury {
	display: none;
}
```

隐藏整个框

```css
#div_digg {
	display: none;
}
```

隐藏阅读排行榜

```css
#sidebar_topviewedposts {
	display:none;
}
```

隐藏推荐排行榜

```css
#sidebar_topdiggedposts{
	display:none;
}
```

隐藏随笔档案

```css
.catListPostArchive {
	display:none;
}
```

## 正文图片放大

### 悬停放大

```css
.post img {
    cursor: pointer;
    transition: all 0.5s;
}
.post img:hover {
    transform: scale(1.3);
}
```

### 点击放大

```js
<script type="text/javascript" src="https://blog-static.cnblogs.com/files/kikochz/imgerlarge.js"></script>
<div id="outerdiv" style="text-align: center;position: fixed;z-index: 1000;top: 0;left: 0;
    width: 100%;height: 100%;background-color: rgba(255,255,255,.9);">
    <img id="bigimg" style="height: auto;width: 65%;border: 5px solid #7e8c8d; 
        margin: auto;position: absolute;top: 0;bottom: 0;left: 0;right: 0;" src="" />
</div>
```

## 侧栏

参考:

https://www.cnblogs.com/0x4D75/p/8965227.html

```css
/* 定制自己导航栏的样式 */
#shwtop ul {
    margin: 0;
    padding: 0;
    list-style-type: none; /*去除li前的标注*/
    background-color: #333;
    overflow: hidden; /*隐藏溢出的部分，保持一行*/
}
#shwtop li {
    float: left; /*左浮动*/
}
#shwtop li a, .dropbtn {
    display: inline-block; /*设置成块*/
    color: white;
    text-align: center;
    text-decoration: none;
    padding: 14px 16px;
}
/*鼠标移上去，改变背景颜色*/
#shwtop li a:hover, .dropdown:hover .dropbtn { 
    /* 当然颜色你可以自己改成自己喜欢的，我还是挺喜欢蓝色的 */
    background-color: blue;
}
#shwtop .dropdown {
    /*
    display:inline-block将对象呈递为内联对象，
    但是对象的内容作为块对象呈递。
    旁边的内联对象会被呈递在同一行内，允许空格。
    */
    display: inline-block;
}
#shwtop .dropdown-content {
    display: none;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
}
#shwtop .dropdown-content a {
    display: block;
    color: black;
    padding: 8px 10px;
    text-decoration:none;
}
#shwtop .dropdown-content a:hover {
    background-color: #a1a1a1;
}
#shwtop .dropdown:hover .dropdown-content{
    display: block;
}
```

### 侧栏内容修改

```js
<script>
$(function () {
  $('#sidebar_search .catListTitle').text('C:/')
  $('#sidebar_shortcut .catListTitle').text('D:/')
  $('#sidebar_postcategory .catListTitle').text('E:/')
})
</script>
```

## 导航栏渐变样式

```css
/* 头部 */
#header {
	position: relative;
	height: 280px;
	margin: 0;
	background: #020031;
	background: -moz-linear-gradient(45deg,#020031 0,#6d3353 100%);
	background: -webkit-gradient(linear,left bottom,right top,color-stop(0%,#020031),color-stop(100%,#6d3353));
	background: -webkit-linear-gradient(45deg,#020031 0,#6d3353 100%);
	background: -o-linear-gradient(45deg,#020031 0,#6d3353 100%);
	background: -ms-linear-gradient(45deg,#020031 0,#6d3353 100%);
	background: linear-gradient(45deg,#020031 0,#6d3353 100%);
	filter: progid:DXImageTransform.Microsoft.gradient(startColorstr='#020031', endColorstr='#6d3353', GradientType=1);
	-webkit-box-shadow: inset 0 3px 7px rgba(0,0,0,.2),inset 0 -3px 7px rgba(0,0,0,.2);
	-moz-box-shadow: inset 0 3px 7px rgba(0,0,0,.2),inset 0 -3px 7px rgba(0,0,0,.2);
	box-shadow: inset 0 3px 7px rgba(0,0,0,.2),inset 0 -3px 7px rgba(0,0,0,.2);
}
```

## 评论样式

### 样式一

```css
    .blog_comment_body {
        background: #B2E866;
        float: left;
        border-radius: 5px;
        position: relative;
        overflow: visible;
        margin-left: 33px;
        max-width: 700px;
    }
 
    .feedbackListSubtitle a.layer {
        background: #B2E866;
        color: #414141 !important;
        padding: 2px 4px;
        border-radius: 2px;
    }
```

### 样式二

CSS

```css
.feedbackCon img:hover {
-webkit-transform: rotateZ(360deg);
-moz-transform: rotateZ(360deg);
-ms-transform: rotateZ(360deg);
-o-transform: rotateZ(360deg);
transform: rotateZ(360deg);
}
 
.feedbackCon img {
border-radius: 40px;
-webkit-transition: all 0.6s ease-out;
-moz-transition: all 0.5s ease-out;
-ms-transition: all 0.5s ease-out;
-o-transition: all 0.5s ease-out;
transition: all 0.5s ease-out;
}
```

JS

```js
<script  src="https://blog-static.cnblogs.com/files/kikochz/comment.js"></script>
```

### 样式三

参考:

https://www.cnblogs.com/miluluyo/p/11683773.html

```css
/*评论区*/
#commentform_title, .feedback_area_title {
    font: normal normal 16px/35px "Microsoft YaHei";
    margin: 10px 0 30px;
    border-bottom: 2px solid #ccc;
    background-image: none;
    padding: 0;
    border-bottom: 0;
}

#commentform_title:after, .feedback_area_title:after {
    content: '';
    display: block;
    width: 100%;
    text-align: center;
    position: relative;
    bottom: 16px;
    left: 110px;
    border-bottom: 1px dashed #e9e9e9;
}

#tbCommentAuthor {
    padding-left: 10px;
    color: #555;
    border: 1px solid #ddd;
    border-radius: 3px;
    -moz-border-radius: 3px;
    -webkit-border-radius: 3px;
    width: 320px;
    height: 20px;
    background: #fff;
}

.commentbox_title {
    width: 100%;
}

div.commentform p {
    margin-bottom: 20px
}

textarea#tbCommentBody {
    width: calc(100% - 20px);
    border-radius: 10px;
    outline: 0;
    padding: 10px;
    height: 200px;
    position: relative;
    background: black;
    background-size: contain;
    background-repeat: no-repeat;
    background-position: right;
    resize: vertical;
}

/*评论列表*/
.feedbackItem {
    margin-top: 30px;
}

.feedbackListSubtitle {
    clear: both;
    color: #a8a8a8;
    padding: 8px 5px;
}

.feedbackManage {
    width: 200px;
    text-align: right;
    float: right;
}

.feedbackListSubtitle a:link, .feedbackListSubtitle a:visited, .feedbackListSubtitle a:active {
    color: #777;
    font-weight: 450;
}


.feedbackCon {
    border-bottom: 1px solid #EEE;
    padding: 10px 20px 10px 5px;
    min-height: 35px;
    _height: 35px;
    margin-bottom: 1em;
    line-height: 1.5;
}

.comment-avatar {
    width: 48px;
    height: 48px;
    border: 1px solid #dcd6b3;
    padding: 3px;
    border-radius: 50%;
    -webkit-transition: all .6s ease-out;
    -moz-transition: all .5s ease-out;
    -ms-transition: all .5s ease-out;
    -o-transition: all .5s ease-out;
    transition: all .5s ease-out;
}

.blog_comment_body {
    display: inline-block;
    width: 70%;
    margin-left: 15px;
    vertical-align: initial !important;
    font-family: Lato, Helvetica, Arial, sans-serif;
}

.comment_vote {
    padding-right: 10px;
}

.comment_vote a {
    color: #999;
}

.blog_comment_body a {
    color: #2daebf;
}

.comment-avatar:hover {
    transform: rotateZ(360deg);
}

#comment_nav {
    padding-top: 10px;
}

.blog_comment_body img {
    max-width: 100px;
}

/*提交评论*/
.comment_btn {
    width: 180px;
    height: 38px;
    padding: 8px 20px;
    text-align: center;
    font-size: 14px;
    color: #fff;
    border: 0;
    background: #7396a7 !important;
    border-radius: 3px;
    -moz-border-radius: 3px;
    -webkit-border-radius: 3px;
    -webkit-transition: all .4s ease;
    -moz-transition: all .4s ease;
    -o-transition: all .4s ease;
    -ms-transition: all .4s ease;
    transition: all .4s ease;
    cursor: pointer;
    display: inline-block;
    vertical-align: middle;
    outline: 0;
    text-decoration: none;
}

.comment_btn:hover {
    background: #8cb7cc !important;
}

p#commentbox_opt {
    text-align: center;
}
```

JS

```js
<link rel="stylesheet" href="https://blog-static.cnblogs.com/files/elkyo/OwO.min.css" />
<script src="https://blog-static.cnblogs.com/files/elkyo/OwO.min.js"></script>
<script>
  /*文章评论*/
  var le = $(".feedbackItem").length
  for(var i = 0;i < le;i++){
    var src = $(".feedbackItem").eq(i).find(".feedbackCon").find("span").text()
  }
  $("#tbCommentBody").attr("placeholder","%userprofile%/AppData/Roaming/Microsoft/Windows/Start Menu/Programs/Startup>")
  $("#commentbox_opt").nextAll().remove()
  $("#btn_comment_submit").val("提交评论 (Ctrl + Enter)")
</script>
```

