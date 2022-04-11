# k8s 环境安装

https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/

使用kubeadm，kubectl，kubelet

## 配置源 & 安装

### 前置条件

1. docker加速

2. 硬件，实际告诉我2核不行

   https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/#before-you-begin

   ```
   [root@k8smaster ~]# kubeadm init
   [init] Using Kubernetes version: v1.20.4
   [preflight] Running pre-flight checks
   	[WARNING Firewalld]: firewalld is active, please ensure ports [6443 10250] are open or your cluster may not function correctly
   	[WARNING Hostname]: hostname "k8smaster" could not be reached
   	[WARNING Hostname]: hostname "k8smaster": lookup k8smaster on 8.8.8.8:53: no such host
   error execution phase preflight: [preflight] Some fatal errors occurred:
   	[ERROR NumCPU]: the number of available CPUs 1 is less than the required 2
   	[ERROR Mem]: the system RAM (972 MB) is less than the minimum 1700 MB
   	[ERROR Swap]: running with swap on is not supported. Please disable swap
   [preflight] If you know what you are doing, you can make a check non-fatal with `--ignore-preflight-errors=...`
   To see the stack trace of this error execute with --v=5 or higher
   ```

3. 关闭交换分区，还需要修改`/etc/fstab`

   ```
   [root@k8smaster ~]# swapon
   NAME      TYPE      SIZE USED PRIO
   /dev/dm-1 partition   2G   0B   -2
   [root@k8smaster ~]# swapoff  /dev/dm-1
   ```

4. check the compability of version

   https://kubernetes.io/docs/setup/release/version-skew-policy/#supported-version-skew

   ```
   [root@k8smaster ~]# kubeadm version
   kubeadm version: &version.Info{Major:"1", Minor:"20", GitVersion:"v1.20.4", GitCommit:"e87da0bd6e03ec3fea7933c4b5263d151aafd07c", GitTreeState:"clean", BuildDate:"2021-02-18T16:09:38Z", GoVersion:"go1.15.8", Compiler:"gc", Platform:"linux/amd64"}
   
   [root@k8smaster ~]# kubectl version 
   Client Version: version.Info{Major:"1", Minor:"20", GitVersion:"v1.20.4", GitCommit:"e87da0bd6e03ec3fea7933c4b5263d151aafd07c", GitTreeState:"clean", BuildDate:"2021-02-18T16:12:00Z", GoVersion:"go1.15.8", Compiler:"gc", Platform:"linux/amd64"}
   ```

5. check ports and firewall

   ```
   [root@k8smaster ~]# firewall-cmd --permanent --new-service=kubernetes
   success
   [root@k8smaster ~]# firewall-cmd --permanent --service=kubernetes --add-port=6443/tcp --add-port=10250/tcp
   success
   [root@k8smaster ~]# firewall-cmd --reload 
   success
   [root@k8smaster ~]# firewall-cmd --permanent --add-service=kubernetes 
   success
   [root@k8smaster ~]# firewall-cmd --reload 
   success
   [root@k8smaster ~]# firewall-cmd --list-all
   public (active)
     target: default
     icmp-block-inversion: no
     interfaces: ens33
     sources: 
     services: dhcpv6-client kubernetes ssh
     ports: 
     protocols: 
     masquerade: no
     forward-ports: 
     source-ports: 
     icmp-blocks: 
     rich rules: 
   ```
   
   考虑到后面需要对外开放端口，为了方便我直接使用truested zone，实际不可采用这种策略
   
   ```
   [root@k8smaster opt]# firewall-cmd --set-default-zone=trusted
   success
   [root@k8smaster opt]# firewall-cmd --reload
   success
   [root@k8smaster opt]# firewall-cmd --list-all
   trusted (active)
     target: ACCEPT
     icmp-block-inversion: no
     interfaces: ens33
     sources:
     services:
     ports:
     protocols:
     masquerade: no
     forward-ports:
     source-ports:
     icmp-blocks:
     rich rules:
   ```

### 安装

- tuna源

  不推荐，tuna k8s的镜像没有同步完整

  ```
  cat <<EOF | sudo tee /etc/yum.repos.d/kubernetes.repo
  [kubernetes]
  name=Kubernetes
  baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-\$basearch
  enabled=1
  gpgcheck=1
  repo_gpgcheck=1
  gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
  exclude=kubelet kubeadm kubectl
  EOF
  
  #修改源为tuna源，关闭gpg校验tuna没有同步asc签名文件
  sed -i 's/packages.cloud.google.com/mirrors.tuna.tsinghua.edu.cn\/kubernetes/' /etc/yum.repos.d/kubernetes.repo 
  sed -i '5,7d' kubernetes.repo /etc/yum.repos.d/kubernetes.repo 
  
  # Set SELinux in permissive mode (effectively disabling it)
  sudo setenforce 0
  sudo sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config
  
  yum clean all && yum makecache
  sudo yum install -y kubelet kubeadm kubectl --disableexcludes=kubernetes
  
  sudo systemctl enable --now kubelet
  ```

- huawei源

  https://mirrors.huaweicloud.com/

  ```
  cat <<EOF > /etc/yum.repos.d/kubernetes.repo
  [kubernetes]
  name=Kubernetes
  baseurl=https://repo.huaweicloud.com/kubernetes/yum/repos/kubernetes-el7-$basearch
  enabled=1
  gpgcheck=1
  repo_gpgcheck=1
  gpgkey=https://repo.huaweicloud.com/kubernetes/yum/doc/yum-key.gpg https://repo.huaweicloud.com/kubernetes/yum/doc/rpm-package-key.gpg
  EOF
  
  # Set SELinux in permissive mode (effectively disabling it)
  sudo setenforce 0
  sudo sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config
  
  yum clean all && yum makecache
  sudo yum install -y kubelet kubeadm kubectl --disableexcludes=kubernetes
  
  sudo systemctl enable --now kubelet
  ```

### 配置补全脚本

`command completion --help`

kubectl

https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/#enable-shell-autocompletion

kubeadm

```
[root@k8smaster ~]# kubeadm completion bash > /etc/bash_completion.d/kubeadm
```

## cgroup driver

如果使用docker做为container runtime interface时，k8s会自动探测。如果CRI

















