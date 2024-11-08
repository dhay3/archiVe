# HTML Video Audio

> HTML 还提供了 substitle 的类功能, 具体查看 `<track>` 部分

ref

https://developer.mozilla.org/en-US/docs/Learn/HTML/Multimedia_and_embedding/Video_and_audio_content

历史上 web 通过 Flash 和 sliverlight 插件来运用音频, 但是 Flash 和 sliverlight 都是不完美的软件有大量的漏洞, 所以 HTML 使用了内嵌的 `<video>` 和 `<audio>` element tag 来插入音频

## video

```
<video
  controls
  width="400"
  height="400"
  autoplay
  loop
  muted
  preload="auto"
  poster="poster.png">
  <source src="rabbit320.mp4" type="video/mp4" />
  <source src="rabbit320.webm" type="video/webm" />
  <p>
    Your browser doesn't support this video. Here is a
    <a href="rabbit320.mp4">link to the video</a> instead.
  </p>
</video>

```

<video
  controls
  width="400"
  height="400"
  autoplay
  loop
  muted
  preload="auto"
  poster="poster.png">
  <source src="rabbit320.mp4" type="video/mp4" />
  <source src="rabbit320.webm" type="video/webm" />
  <p>
    Your browser doesn't support this video. Here is a
    <a href="rabbit320.mp4">link to the video</a> instead.
  </p>
</video>

- `src`

  a path to the video you wanna embed

- `controls`

  a default control interface

- `width/height`

  video size

- `autoplay`

  playing right away while the reset of the page is loading

- `loop`

  makes the video or audio start play again whenever it finishes

- `muted`

  playing the media with the sound tunred off by default

- `poster`

  the layer of the video before it is played

- `preload`

  used for buffering large files, it can take one of three values

  - `none` dose not buffer the file
  - `auto` buffers the media file
  - `metadata` buffers only the metadata for the file

## audio

`<audio>` element works just like the `<video>`  except:

1. the audio element doesn’t support the width/height attributes

2. also doesn’t support the poseter attribute

