# k8s 入门

## 安装kubectl

https://kubernetes.io/zh/docs/tasks/tools/install-kubectl

1. 下载kubectl

   ```
   curl -LO https://dl.k8s.io/release/v1.20.0/bin/linux/amd64/kubectl
   ```

2. 校验

   ```
   sha256 kubectl | sha256 -c 
   ```

3. [shell 补全](https://kubernetes.io/zh/docs/tasks/tools/install-kubectl/#%E5%90%AF%E7%94%A8-shell-%E8%87%AA%E5%8A%A8%E8%A1%A5%E5%85%A8%E5%8A%9F%E8%83%BD)

   如果是Zsh可以在`.zshrc`中配置持久补全

   ```
   source <(kubectl completion zsh)
   ```

## 安装minikube

https://minikube.sigs.k8s.io/docs/start/

1. 下载minikube，如果因为魔法可以先将binary下载后，通过SSH传输。或是使用aliyun提供的[镜像](https://developer.aliyun.com/article/221687)

2. 运行minikube

   如果出现如下错误，参考：https://github.com/kubernetes/minikube/issues/7903

   ```
   oot in /opt λminikube start
   😄  minikube v1.18.1 on Debian kali-rolling
   ✨  Automatically selected the docker driver. Other choices: none, ssh
   🛑  The "docker" driver should not be used with root privileges.
   💡  If you are running minikube within a VM, consider using --driver=none:
   📘    https://minikube.sigs.k8s.io/docs/reference/drivers/none/
   
   ❌  Exiting due to DRV_AS_ROOT: The "docker" driver should not be used with root privileges.
   ```

   切换到非root用户，执行`usermod -g docker <username>`。如果minikube启动成功显示如下

   ```
   chz@cyberpelican:/opt$ minikube start
   😄  minikube v1.18.1 on Debian kali-rolling
   ✨  Automatically selected the docker driver
   👍  Starting control plane node minikube in cluster minikube
   🚜  Pulling base image ...
   💾  Downloading Kubernetes v1.20.2 preload ...
       > preloaded-images-k8s-v9-v1....: 3.59 MiB / 491.22 MiB  0.73% 158.66 KiB 
   ```

   校验minikube是否安装成功

   ```
   
   ```

   
