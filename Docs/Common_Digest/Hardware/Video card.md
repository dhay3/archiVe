# Video card

ref:

https://en.wikipedia.org/wiki/Graphics_card

## Digest

![Video card](https://www.computerhope.com/cdn/video-card.jpg)

Video card or Graphics card 中文通常叫显卡，主要做图形化处理。同时显卡分成几大类

## Types

### on-board video/integrated graphics

集成显卡 指的是显卡集成在主板或者CPU上，不可拆解。性能一般较差(because the graphics proessor shares sysem resources with CPU)，价格比较低廉，一般用于笔记本。功耗较低（意味着发热量小）。

但是一般的主板都支持关闭集显，通过PCIE槽外接一张独立显卡

### discrete card

独立显卡 使用独立插槽连接至主板，可拆卸。有自己的 RAM, cooling system, and dedicated power regulatros

![PCIe video card](https://www.computerhope.com/issues/pictures/pcie-video-card.png)

上图是较为常见的 PCI 接口的显卡

## Terms

> 可以对比参考 CPU Terms 部分

- Overclocking：超频

- VRAM(Video memory)

  显存，数值越大性能越强(不占用CPU的主存)。提一嘴的是还需要看显存的类型，是 DDRx 多少的

- GTT：显卡可以访问的CPU主存

- Core clock：主频。GPU芯片的速度，影响FPS

- Shader clock：着色器运行的频率

- Memory clock：显存频率。决定访问GPU RAM的速度