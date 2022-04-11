# WebStorm 使用webpack

**webpack**主要来处理我们写的`js`代码, 并且`webpack`会自动处理`js`之间相关的依赖

参考:

https://www.webpackjs.com

[TOC]

`webpack`默认配置文件`webpack.config.js`, 固定写法

As follows:

```json
//需要通过npm init来初始化
var path = require('path');

module.exports = {
  //导入
  entry: './src/main.js',
  //导出
  output: {
    //导出路径, _dirname相当于java的user.dir
    path: path.resolve(__dirname, 'dist'),
    //导出的文件名
    filename: 'bundle.js'
  }
};
```

然后输入`npm init`(推荐在项目生成时就使用)生成`package.json`文件, 重新输入`webpack`可以打包当前文件到

`dist\bundle.js`

```json
{
  //项目名
  "name": "meetwebpack",
  //版本号
  "version": "1.0.0",
  //描述
  "description": "",
  "main": "index.js",
  //脚本, 用于映射npm和webpack的命令
  "scripts": {
      //在改配置中执行命令优先使用本地的命令, 如果本地不存在才会使用全局的
      //执行npm run test就是调用后面的脚本
    "test": "echo \"Error: no test specified\" && exit 1",
      //输入npm run build 相当于使用webpack
    "build": "webpack"

  },
  //作者
  "author": "",
  //协议,如果项目不开源不需要改属性
  "license": "ISC"
}

```

`npm install webpack@3.6.0 --save-dev`本地安装`webpack`

安装后`package.json`会添加如下代码片段

```json
  //开发时依赖
  "devDependencies": {
    "webpack": "^3.6.0"
  }
```

==凡是在terminal中输入的命令行都是在执行全局命令==



如果想要将css样式也通过模块化的方式加入到bundle.js中需要另外安装

`npm install css-loader --save-dev`

`npm install style-loader --save-dev`

然后可以通过`npm run build`就可以调用`webpack`将css样式加入到`bundle.js`
