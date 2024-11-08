# HTML hyperlink

ref:

https://developer.mozilla.org/en-US/docs/Learn/HTML/Introduction_to_HTML/Creating_hyperlinks

## Document fragments

link 可以链接到 fragment(中文也被翻译成 锚点)，例如

```
<h2 id="Mailing_address">Mailing address</h2>

#跨文件链接到 fragment
<a href="contacts.html#Mailing_address">mailing address</a>.

#本文件链接到 fragment
<a href="#Mailing_address">company mailing address</a>
```

## Linking to Downloaded resource

当链接到一些可下载的资源时(例如 PDF,Word document) browser 会自动下载, 这时需要一些文字上的提示以减少用户 confusion, 例如

```
  <a href="https://www.example.com/large-report.pdf">
    Download the sales report (PDF, 10MB)
  </a>
```

还可以设置 download attribute, 用来设置  default download filename

```
<a
  href="https://download.mozilla.org/?product=firefox-latest-ssl&os=win64&lang=en-US"
  download="firefox-latest-64bit-installer.exe">
  Download Latest Firefox for Windows (64-bit) (English, US)
</a>
```

## Email links

点击链接时通过 `mailto:` 打开 mail 客户端以发送邮件

```
<a href="mailto:nowhere@mozilla.org">Send email to nowhere</a>
```

