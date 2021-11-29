# pip

参考：

https://pip.pypa.io/en/stable/quickstart/

如果是从python.org下载的python2 >=2.79 或 python3 >= 3.4, pip默认在捆绑包中。如果是linux使用[包管理器](https://packaging.python.org/guides/installing-using-linux-tools/)更新或下载pip

可以使用如下命令查看是否安装了pip。==如果是在windows上使用py替代python==

```
cpl in / λ python -m pip --version
pip 20.3.1 from /usr/lib/python3.9/site-packages/pip (python 3.9)
```

这里使用`-m`指定使用pip模块

## command completion

```
python -m pip completion --zsh >> ~/.zprofile;source ~/.zprofile
python -m pip completion --bash >> ~/.profile;source ~/.profile
```

如果使用pip3，手动修改文件中pip为pip3

## command

> 由于pypi的服务器被发送了大量的恶意请求，现在永久关闭了`pip search`功能，如果需要查询包通过如下链接
>
> https://pypi.org/

可以使用`pip help <command>`来查看command的具体用法

### install

从如下源按照包

- PyPI
- VCS project urls
- Local project directories
- Local or remote source archives
- requiresments files

#### options

- `-r | --requirement <file>`

  从requirement.txt中安装package

- `--no-deps`

  不安装依赖

- `--pre`

  包含pre-release和development versions，pip默认只安装stable version

- `-U | --upgrade`

  将依赖更新到最新的

- `--force-reinstall`

  强制重新安装

## requirement.txt/constraints.txt

> 文件支持的operator
>
> https://pip.pypa.io/en/stable/user_guide/#understanding-your-error-message

requirements.txt用于描述需要的依赖，constraints.txt用于描述不能安装的依赖。每一行表示一个依赖

```
cpl in / λ pip freeze 
apparmor==3.0.1
appdirs==1.4.4
application-utility==1.3.2
attrs==21.2.0
Brlapi==0.8.2
btrfsutil==5.12.1
CacheControl==0.12.6
ceph==1.0.0
ceph-volume==1.0.0
cephfs==2.0.0
cephfs-shell==0.0.1
certifi==2020.12.5
cffi==1.14.5
chardet==4.0.0
colorama==0.4.4
```

通常用于将pip freeze的内容导出，以实现重复安装

```
#将pip freeze的内容写入到文件
pip freeze > requirements.txt
#读取requirementes.txt，并安装依赖
pip install -r requirements.txt
```

