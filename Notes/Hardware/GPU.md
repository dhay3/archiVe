# GPU

## Digest

> the GPU is used in a wide range of applications, including graphics and video rendering  

Graphics processing unit(GPU)，中文通常叫图形处理器。区别与 CPU, GPU 主要用在 computer graphics 和 image processing，==通常和显卡绑定在一起（但是也有专业的GPU是独立于显卡的）。==这也意味着 如果需要替换 GPU 就需要替换显卡。==另外需要注意的一点是， GPU 同样归 CPU 调度，所以当 CPU 性能不行时，GPU 再好，渲染还是一样不行==

GPU 只要支持如下几个功能

1. for gaming
2. for video editing and content creation
3. for machine learning

## Terms

> 可以对比参考 CPU Terms 部分

- Overclocking：超频

- VRAM（video memory）：显卡的显存（不占用主存，自产自销）

- GTT：显卡可以访问的CPU主存

- Core clock：主频。GPU芯片的速度，影响FPS

- Shader clock：着色器运行的频率

- Memory clock：显存频率。决定访问GPU RAM的速度

  

**references**

[^1]:https://www.easypc.io/gpu-memory-clock-speed/
[^2]:https://happyseeker.github.io/kernel/2016/03/01/about-Video-Memory.html
[^3]:https://zhuanlan.zhihu.com/p/217881237

[^4]:https://www.intel.com/content/www/us/en/products/docs/processors/what-is-a-gpu.html
