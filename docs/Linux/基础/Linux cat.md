# Linux cat

参考：

https://blog.csdn.net/zongshi1992/article/details/71693045

concatenate，会将stdin输入到testcat中

```
root in /opt λ cat > testcat 
111
111
root in /opt λcat testcat
  File: testcat
  111
```

这里通过`ctrl+D`手动挂起cat进程

使用`-n`参数打印内容的同时打印序号，对sed非常有帮助

```
[root@chz yum.repos.d]# cat -n kubernetes.repo 
     1	[kubernetes]
     2	name=Kubernetes
     3	baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-$basearch
     4	enabled=1
     5	gpgcheck=1
     6	repo_gpgcheck=1
     7	gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
     8	exclude=kubelet kubeadm kubectl
```

## 特殊用法

执行脚本是，需要往一个文件中输入N行内容，如果使用echo追加的方式，效率极低。这时候就可以和EOF结合使用，EOF必须以一对出现

```
#这种方式会覆盖testcat
root in /opt λcat > testcat <<EOF
heredoc> catEOF
heredoc> EOF
root in /opt λcat testcat 
  File: testcat
  catEOF

#这种方式会追加
root in /opt λ cat >> testcat << EOF     
heredoc> testEOF
heredoc> EOF
root in /opt λ cat testcat 
  File: testcat
  catEOF
  testEOF
```

EOF的位置可以调换

```
cat << EOF >> testcat 
>...
>EOF
```

这里将`br_netfilter`写入到管道里然后通过tee到`/etc/modules-load.d/k8s.conf`

```
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
br_netfilter
EOF
```
