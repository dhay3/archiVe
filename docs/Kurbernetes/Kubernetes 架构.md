# Kubernetes 架构

![](D:\asset\note\imgs\_Kubernetes\Snipaste_2020-12-31_17-59-33.png)

- **kube-apiserver**

  如果需要与kubernetes集群进行交互，就要通过Kube-apiserver(Kubernetes控制平面的前端)，来调用API。

  可以通过REST调用，kubectl 命令行界面或其他命令行工具来访问 API。

  