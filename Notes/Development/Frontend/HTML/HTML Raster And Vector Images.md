# HTML Raster And Vector Images

ref

https://developer.mozilla.org/en-US/docs/Learn/HTML/Multimedia_and_embedding/Adding_vector_graphics_to_the_Web

## Raster/vector images

- Raster images

  a grid of pixels - a raster image file contains information showing exactly where each pixel is to be placed, and exactly what color it should be placed, and exactly what color it should be.

  Popular web raster formats include Bitmap (`.bmp`), PNG (`.png`), JPEG (`.jpg`), and GIF (`.gif`)

  中文通常也叫做位图(按照音译比较奇怪，理解成栅栏图即可)

- Vector images

  a picture of algorithms - a vector image file contains shape and path definitions that the computer can use to work out what the image should look like when rendered on the screen

  the `.svg` format allows us to create powerful vector graphics for use on the web

  中文通常也叫做矢量图

两者最主要的区别在于

when raster images are zoomed, each pixel is increased in size to fill multiplse pixels on screen, so the image starts to look blocky. The vector image however continues to look nice and crisp, becuase no matter what isze it is, the algorithms are used to work out the shapes in the image, with the values being scaled as it gets bigger

简单的说就是 raster image 放大会模糊，但是 vector image 不会，具体例子参考 MDN docs

## SVG

> 这是个很复杂的 element，具体查看 MDN docs
>
> 一般通过 illustrator 或者 photoshop 来生成复杂的 svg 文件

SVG is an XML - based language for describing vector images

