# SELinux概念

参考：

http://cn.linux.vbird.org/linux_basic/0440processcontrol_5.php

http://c.biancheng.net/view/3906.html

SELinux（Security Enhance Linux）

### DAC

系统的账号主要分为系统管理员（root）与一般用户，而两种身份能否使用系统上面的文件资源则与rwx的权限配置有关。==但是各种权限配置对root是无效的==。因此，当某个程序想要对文件进行存取时， 系统就会根据该程序的拥有者/群组，并比对文件的权限，若通过权限检查，就可以存取该文件了。这种存取文件系统的方式被称为==自主存取控制（Discretionary Access Control ，DAC）==，基本上，就是依据程序的拥有者与文件资源的rws权限来决定有无存取的能力。不过这种DAC的存取控制有几个困扰，那就是：

- root 具有最高的权限：如果不小心某支程序被有心人士取得， 且该程序属于 root 的权限，那么这支程序就可以在系统上进行任何资源的存取！真是要命！

- 使用者可以取得程序来变更文件资源的存取权限：如果你不小心将某个目录的权限配置为 777 ，由於对任何人的权限会变成 rwx ，因此该目录就会被任何人所任意存取！

### MAC

现在我们知道 DAC 的困扰就是当使用者取得程序后，他可以藉由这支程序与自己默认的权限来处理他自己的文件资源。 万一这个使用者对 Linux 系统不熟，那就很可能会有资源误用的问题产生。为了避免 DAC 容易发生的问题，==因此 SELinux 导入了委任式存取控制 (Mandatory Access Control, MAC) 的方法！==

委任式存取控制 (MAC) 有趣啦！他可以针对特定的程序与特定的文件资源来进行权限的控管！ 也就是说，==即使你是 root ，那么在使用不同的程序时，你所能取得的权限并不一定是 root ， 而得要看当时该程序的配置而定。如此一来，我们针对控制的『主体』变成了『程序』而不是使用者喔!==此外，这个主体程序也不能任意使用系统文件资源，因为每个文件资源也有针对该主体程序配置可取用的权限！ 如此一来，控制项目就细的多了！但整个系统程序那么多、文件那么多，一项一项控制可就没完没了！ 所以 SELinux 也提供一些默认的策略 (Policy) ，并在该策略内提供多个守则 (rule) ，让你可以选择是否激活该控制守则！

SElinux使用的是MAC方式来管控程序，它控制的主题是程序，而目标则是程序能否读取的【文件资源】。

- **主体Subject**

  就是想要访问文件或目录资源的进程（process）。想要得到资源，基本流程是这样的：==由用户调用命令，由命令产生进程，由进程去访问文件或目录资源==。

  DAC（Linux 默认权限中），靠权限控制的==主体是用户==；MAC（SELinux 中），靠策略规则控制的==主体则是进程==。

- **目标Object**

  就是需要访问的文件或目录资源

- **策略Policy**

  ```mermaid
  classDiagram
  Policy <| -- type
  Policy <| -- domain
  ```

  由于程序与文件数量庞大，因此SELinux会一句某些服务来指定基本的存取安全策略。这些策略还会有详细的规则（rule）来指定不同的服务开放某些资源的存取与否。==举个例子==，我们需要找对象，男人可以看作主体，女人就是目标了。而男人是否可以追到女人（主体是否可以访问目标），主要看两个人的性格是否合适（主体和目标的安全上下文是否匹配）。不过，两个人的性格是否合适，是需要靠生活习惯、为人处世、家庭环境等具体的条件来进行判断的（安全上下文是否匹配是需要通过策略中的规则来确定的）。

  - targeted：针对网络服务限制较多，针对本机限制较少，是默认策略
  - strict：完整的SELinux限制，限制方面较为严格

- 安全性文本security context

  我们刚刚谈到了主体、目标与策略面，但是主体能不能存取目标除了策略指定之外，主体与目标的安全性本文必须一致才能够顺利存取。 这个安全性本文 (security context) 有点类似文件系统的 rwx 啦！安全性本文的内容与配置是非常重要的！ 如果配置错误，你的某些服务(主体程序)就无法存取文件系统(目标资源)，当然就会一直出现『权限不符』的错误信息了！

<img src="..\..\..\imgs\_Linux\Snipaste_2020-10-27_20-03-04.png"/>

解释一下这张示意图：当主体想要访问目标时，如果系统中启动了 SELinux，则主体的访问请求首先需要和 SELinux 中定义好的策略进行匹配(查询当前Policy中是否有该规则)。如果进程符合策略中定义好的规则，则允许访问，==这时进程的安全上下文就可以和目标的安全上下文进行匹配==；如果比较失败，则拒绝访问，并通过 AVC（Access Vector Cache，访问向量缓存，主要用于记录所有和 SELinux 相关的访问统计信息）生成拒绝访问信息。如果安全上下文匹配，则可以正常访问目标文件。当然，最终是否可以真正地访问到目标文件，还要匹配产生进程（主体）的用户是否对目标文件拥有合理的读、写、执行权限。

我们在进行 SELinux 管理的时候，==一般只会修改文件或目录的安全上下文==，使其和访问进程的安全上下文匹配或不匹配，用来控制进程是否可以访问文件或目录资源；而很少会去修改策略中的具体规则，因为规则实在太多了，修改起来过于复杂。不过，我们是可以人为定义规则是否生效，用以控制规则的启用与关闭的。

## SELinux的工作模式

### Disable关闭模式

SELinux 被关闭，==默认的 DAC 访问控制方式被使用==。对于那些不需要增强安全性的环境来说，该模式是非常有用的。

例如，若从你的角度看正在运行的应用程序工作正常，但是却产生了大量的 SELinux AVC 拒绝消息，最终可能会填满日志文件，从而导致系统无法使用。在这种情况下，最直接的解决方法就是禁用 SELinux，当然，你也可以在应用程序所访问的文件上设置正确的安全上下文。

需要注意的是，在禁用 SELinux 之前，需要考虑一下是否可能会在系统上再次使用 SELinux，如果决定以后将其设置为 Enforcing 或 Permissive，那么当下次重启系统时，系统将会通过一个自动 SELinux 文件重新进程标记。

关闭 SELinux 的方式也很简单，只需编辑配置文件` /etc/selinux/config`，并将文本中` SELINUX= `更改为` SELINUX=disabled `即可，==重启系统后==，SELinux 就被禁用了。

### Permissive宽容模式

在 Permissive 模式中，SELinux 被启用，但安全策略规则并没有被强制执行。==当安全策略规则应该拒绝访问时，访问仍然被允许==。然而，此时会向日志文件发送一条消息，表示该访问应该被拒绝。

SELinux Permissive 模式主要用于以下几种情况：
==审核当前的 SELinux 策略规则；==
测试新应用程序，看看将 SELinux 策略规则应用到这些程序时会有什么效果；
解决某一特定服务或应用程序在 SELinux 下不再正常工作的故障。

某些情况下，可使用 audit2allow 命令来读取 SELinux 审核日志并生成新的 SELinux 规则，从而有选择性地允许被拒绝的行为，而这也是一种在不禁用 SELinux 的情况下，让应用程序在 Linux 系统上工作的快速方法。

### Enforcing强制模式

 SELinux 启动，并强制执行所有的安全策略规则。

## 安全性上下文

==安全性本文你就将他想成 SELinux 内必备的 rwx 就是了==

==安全性本文是放置到文件的 inode 内的，==因此主体程序想要读取目标文件资源时，同样需要读取 inode ， 这 inode 内就可以比对安全性本文以及 rwx 等权限值是否正确，而给予适当的读取权限依据。

SELinux 管理过程中，进程是否可以正确地访问文件资源，取决于它们的安全上下文。==进程和文件都有自己的安全上下文，==SELinux 会为进程和文件添加安全信息标签，比如 SELinux 用户、角色、类型、类别等，当运行 SELinux 后，所有这些信息都将作为访问控制的依据。

- 文件和目录安全性文本

  ```
  [root@chz network-scripts]# ls -Z
  -rw-r--r--. root root system_u:object_r:net_conf_t:s0  ifcfg-ens33
  -rw-r--r--. root root system_u:object_r:net_conf_t:s0  ifcfg-ens34
  -rw-r--r--. root root system_u:object_r:net_conf_t:s0  ifcfg-lo
  lrwxrwxrwx. root root system_u:object_r:bin_t:s0       ifdown -> ../../../usr/sbin/ifdown
  -rwxr-xr-x. root root system_u:object_r:bin_t:s0       ifdown-bnep
  -rwxr-xr-x. root root system_u:object_r:bin_t:s0       ifdown-eth
  -rwxr-xr-x. root root system_u:object_r:bin_t:s0       ifdown-ib
  ```

- 进程安全性文本

  添加`Z`参数

  ```
  [root@chz opt]# ps -efZ|more
  LABEL                           UID         PID   PPID  C STIME TTY          TIME CMD
  system_u:system_r:init_t:s0     root          1      0  0 13:13 ?        00:00:04 /usr/lib/systemd/systemd --switched-root --s
  ystem --deserialize 22
  system_u:system_r:kernel_t:s0   root          2      0  0 13:13 ?        00:00:00 [kthreadd]
  system_u:system_r:kernel_t:s0   root          4      2  0 13:13 ?        00:00:00 [kworker/0:0H]
  ```

安全上下文看起来比较复杂，它使用“：”分隔为 4 个字段，其实共有 5 个字段，只是最后一个“类别”字段是可选的，例如：

```
system_u：object_r：httpd_sys_content_t：s0：[类别]
#身份字段：角色：类型：灵敏度：[类别]
```

- **身份标识user**

  用于标识该数据被哪个身份所拥有，相当于权限中的用户身份。user 字段只用于标识数据或进程被哪个身份所拥有，==一般系统数据的 user 字段就是 system_u，而用户数据的 user 字段就是 user_u。==

  1. root：表示安全上下文的身份是root
  2. user_u：表示与一般用户账户相关的身份，其中“_u”代表 user。
  3. system_u：表示系统用户身份，其中“_u”代表 user。

- **角色role**

  主要用来表示此数据是进程还是文件或目录。

  1. object_r：代表该数据是文件或目录，这里的“_r”代表 role。
  2. system_r：代表该数据是进程，这里的“_r”代表 role。

- **类型type**

  类型字段是安全上下文中最重要的字段，进程是否可以访问文件，主要就是看==进程（Subject）的安全上下文类型字段是否和文件（Object）的安全上下文类型字段相匹配（具体查看seinfo）==，如果匹配则可以访问。

  1. domain：进程（主体Subject）安全上下文称为域（domain）
  2. type：文件资源（目标Object）安全上下文称为类（type）

  注意，类型字段在文件或目录的安全上下文中被称作类型（type），但是在进程的安全上下文中被称作域（domain）。也就是说，==在主体（Subject）的安全上下文中，这个字段被称为域；在目标（Object）的安全上下文中，这个字段被称为类型（type）。域和类型需要匹配（进程的类型要和文件的类型相匹配），才能正确访问==

  不过，我们已知 apache 进程可以访问 /var/www/html/（此目录为 RPM 包安装的 apache 的默认网页主目录）目录中的网页文件，所以 apache 进程的域和 /var/www/html/ 目录的类型应该是匹配的，我们查询一下，命令如下：

  ```
  [root@chz ~]# ps -efZ|grep httpd
  system_u:system_r:httpd_t:s0    root       3584      1  0 22:54 ?        00:00:00 /usr/sbin/httpd -DFOREGROUND
  
  
  [root@chz html]# ls -Z
  -rw-r--r--. root root unconfined_u:object_r:httpd_sys_content_t:s0 index.html
  ```

  apache 进程的域是 httpd_t，/var/www/html/ 目录的类型是 httpd_sys_content_t，这个主体的安全上下文类型经过策略规则的比对，是和目标的安全上下文类型匹配的，所以 apache 进程可以访问 /var/www/html/ 目录。

- **灵敏度**

  灵敏度一般是用 s0、s1、s2 来命名的，数字代表灵敏度的分级。数值越大，代表灵敏度越高。

- **类别**

  类别字段不是必须有的，所以我们使用 ls 和 ps 命令查询的时候并没有看到类别字段。
