# 深入理解CAS算法原理

转载:

https://www.jianshu.com/p/21be831e851e

## 1、什么是CAS？

CAS：Compare and Swap，即比较再交换。

jdk5增加了并发包java.util.concurrent.*,其下面的类使用CAS算法实现了区别于synchronouse同步锁的一种乐观锁。JDK 5之前Java语言是靠synchronized关键字保证同步的，这是一种独占锁，也是是悲观锁。

## 2、CAS算法理解

对CAS的理解，CAS是一种无锁算法，CAS有3个操作数，内存值V，旧的预期值A，要修改的新值B。当且仅当预期值A和内存值V相同时，将内存值V修改为B，否则什么都不做。

CAS比较与交换的伪代码可以表示为：

do{

备份旧数据；

基于旧数据构造新数据；

}while(!CAS( 内存地址，备份的旧数据，新数据 ))



![img](https://upload-images.jianshu.io/upload_images/5954965-b88918b03518f254?imageMogr2/auto-orient/strip|imageView2/2/w/320/format/webp)

注：t1，t2线程是同时更新同一变量56的值

因为t1和t2线程都同时去访问同一变量56，所以他们会把主内存的值完全拷贝一份到自己的工作内存空间，所以t1和t2线程的预期值都为56。

假设t1在与t2线程竞争中线程t1能去更新变量的值，而其他线程都失败。（失败的线程并不会被挂起，而是被告知这次竞争中失败，并可以再次发起尝试）。t1线程去更新变量值改为57，然后写到内存中。此时对于t2来说，内存值变为了57，与预期值56不一致，就操作失败了（想改的值不再是原来的值）。

（上图通俗的解释是：CPU去更新一个值，但如果想改的值不再是原来的值，操作就失败，因为很明显，有其它操作先改变了这个值。）

就是指当两者进行比较时，如果相等，则证明共享数据没有被修改，替换成新值，然后继续往下运行；如果不相等，说明共享数据已经被修改，放弃已经所做的操作，然后重新执行刚才的操作。容易看出 CAS 操作是基于共享数据不会被修改的假设，采用了类似于数据库的commit-retry 的模式。当同步冲突出现的机会很少时，这种假设能带来较大的性能提升。

## 3、CAS开销

前面说过了，CAS（比较并交换）是CPU指令级的操作，只有一步原子操作，所以非常快。而且CAS避免了请求操作系统来裁定锁的问题，不用麻烦操作系统，直接在CPU内部就搞定了。但CAS就没有开销了吗？不！有cache miss的情况。这个问题比较复杂，首先需要了解CPU的硬件体系结构：

![img](https://upload-images.jianshu.io/upload_images/5954965-a866fcf5501b54c1?imageMogr2/auto-orient/strip|imageView2/2/w/562/format/webp)

上图可以看到一个8核CPU计算机系统，每个CPU有cache（CPU内部的高速缓存，寄存器），管芯内还带有一个互联模块，使管芯内的两个核可以互相通信。在图中央的系统互联模块可以让四个管芯相互通信，并且将管芯与主存连接起来。数据以“缓存线”为单位在系统中传输，“缓存线”对应于内存中一个 2 的幂大小的字节块，大小通常为 32 到 256 字节之间。当 CPU 从内存中读取一个变量到它的寄存器中时，必须首先将包含了该变量的缓存线读取到 CPU 高速缓存。同样地，CPU 将寄存器中的一个值存储到内存时，不仅必须将包含了该值的缓存线读到 CPU 高速缓存，还必须确保没有其他 CPU 拥有该缓存线的拷贝。

比如，如果 CPU0 在对一个变量执行“比较并交换”（CAS）操作，而该变量所在的缓存线在 CPU7 的高速缓存中，就会发生以下经过简化的事件序列：

CPU0 检查本地高速缓存，没有找到缓存线。

请求被转发到 CPU0 和 CPU1 的互联模块，检查 CPU1 的本地高速缓存，没有找到缓存线。

请求被转发到系统互联模块，检查其他三个管芯，得知缓存线被 CPU6和 CPU7 所在的管芯持有。

请求被转发到 CPU6 和 CPU7 的互联模块，检查这两个 CPU 的高速缓存，在 CPU7 的高速缓存中找到缓存线。

CPU7 将缓存线发送给所属的互联模块，并且刷新自己高速缓存中的缓存线。

CPU6 和 CPU7 的互联模块将缓存线发送给系统互联模块。

系统互联模块将缓存线发送给 CPU0 和 CPU1 的互联模块。

CPU0 和 CPU1 的互联模块将缓存线发送给 CPU0 的高速缓存。

CPU0 现在可以对高速缓存中的变量执行 CAS 操作了

以上是刷新不同CPU缓存的开销。最好情况下的 CAS 操作消耗大概 40 纳秒，超过 60 个时钟周期。这里的“最好情况”是指对某一个变量执行 CAS 操作的 CPU 正好是最后一个操作该变量的CPU，所以对应的缓存线已经在 CPU 的高速缓存中了，类似地，最好情况下的锁操作（一个“round trip 对”包括获取锁和随后的释放锁）消耗超过 60 纳秒，超过 100 个时钟周期。这里的“最好情况”意味着用于表示锁的数据结构已经在获取和释放锁的 CPU 所属的高速缓存中了。锁操作比 CAS 操作更加耗时，是因深入理解并行编程

为锁操作的数据结构中需要两个原子操作。缓存未命中消耗大概 140 纳秒，超过 200 个时钟周期。需要在存储新值时查询变量的旧值的 CAS 操作，消耗大概 300 纳秒，超过 500 个时钟周期。想想这个，在执行一次 CAS 操作的时间里，CPU 可以执行 500 条普通指令。这表明了细粒度锁的局限性。

以下是cache miss cas 和lock的性能对比：

![img](https://upload-images.jianshu.io/upload_images/5954965-288a861ec12d6b93?imageMogr2/auto-orient/strip|imageView2/2/w/320/format/webp)

## 4、CAS算法在JDK中的应用

在原子类变量中，如java.util.concurrent.atomic中的AtomicXXX，都使用了这些底层的JVM支持为数字类型的引用类型提供一种高效的CAS操作，而在java.util.concurrent中的大多数类在实现时都直接或间接的使用了这些原子变量类。

## 问题引出

```java
public final class Test {
    static int i = 10;
    public static void main(String[] args) throws InterruptedException {
        //存在两个线程运行完毕后结果不定, 由于i++没有原子性
            new Thread(() -> {
                for (int j = 0; j < 20000; j++) {
                    i++;
                }
            }).start();
            new Thread(() -> {
                for (int j = 0; j < 20000; j++) {
                    i++;
                }
            }).start();
            TimeUnit.SECONDS.sleep(6);
            System.out.println(i);
        }
}
```

## 解决

```java
public final class Test {
    static AtomicInteger i = new AtomicInteger(10);
    public static void main(String[] args) throws InterruptedException {
        //使用AtomicInteger替代普通变量, 在写回数据前和主存中的数据进行比较,如果不相同就放弃本次操作,继续争抢cpu
            new Thread(() -> {
                for (int j = 0; j < 20000; j++) {
                    i.getAndIncrement();
                }
            }).start();
            new Thread(() -> {
                for (int j = 0; j < 20000; j++) {
                    i.getAndIncrement();
                }
            }).start();
            TimeUnit.SECONDS.sleep(6);
            System.out.println(i);
        }
}
```
## 问题引出

```java
public final class Test {
    static boolean flag = true;
    public static void main(String[] args) throws InterruptedException {
        //主线程拷贝一份flag到主线程自己的工作区, 当该线程修改值后, 写回主存, 但是主存中的值不变
        new Thread(() -> {
            try {
                TimeUnit.SECONDS.sleep(1);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
            flag = false;
        }).start();
        while (flag) {
            ;
        }
    }
}
```

## 解决

```java
public final class Test {
    static AtomicBoolean flag = new AtomicBoolean(true);
    public static void main(String[] args) throws InterruptedException {
        new Thread(() -> {
            try {
                TimeUnit.SECONDS.sleep(1);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
             flag.set(false);
        }).start();
       //每次读取的时候会重新从主存中读取
        while (flag.get()) {
            ;
        }
    }
}
```



Java 1.7中AtomicInteger.incrementAndGet()的实现源码为：

![img](https://upload-images.jianshu.io/upload_images/5954965-565839533a3ff3d6?imageMogr2/auto-orient/strip|imageView2/2/w/524/format/webp)

![img](https://upload-images.jianshu.io/upload_images/5954965-effbf420acf8de75?imageMogr2/auto-orient/strip|imageView2/2/w/670/format/webp)

由此可见，AtomicInteger.incrementAndGet的实现用了乐观锁技术，调用了类sun.misc.Unsafe库里面的 CAS算法，用CPU指令来实现无锁自增。所以，AtomicInteger.incrementAndGet的自增比用synchronized的锁效率倍增。