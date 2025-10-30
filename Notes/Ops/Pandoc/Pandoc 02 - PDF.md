---
createTime: 2025-10-30 16:33
license: cc by 4.0
tags: 
 - "#hash1" 
 - "#hash2"
---

# Pandoc 02 - PDF

## 0x01 Preface

pandoc 支持通过 Latex 转 PDF 文件，需要安装 latex

例如

```
cat test.md
## Test

- test1
- test2
- test3
- test4

pandoc test.md -o test.pdf
```

## 0x02 Trouble Shooting

### 0x02a I can't find the format file pdflatex.fmt!

需要安装 texlive-basic

### 0x02b ! LaTeX Error: File xcolor.sty not found.

需要安装 texlive-latexextra

### I can't find the format file xelatex.fmt!

需要安装 texlive-xetex

### ! LaTeX Error: Unicode character 问 (U+95EE)

pandoc 默认的 pdflatex engine 不支持 CJK，需要使用 xelatex `--pdf-engine=xelatex` 

### [WARNING] Missing character: There is no 涨 (U+6DA8) (U+6DA8) in font [lmroman10-regular]:mapping=t
   
需要指定 mainfont

```
pandoc test.md -o test.pdf --pdf-engine=xelatex -V mainfont="Noto Sans CJK SC"
```

---
*Value your freedom or you will lose it, teaches history. Don't bother us with politics, respond those who don't want to learn.*

***See also***



***References***


