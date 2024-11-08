---
createTime: 2024-10-28 15:00
tags:
  - "#hash1"
  - "#hash2"
---

# Cron 03 - Anacron

## 0x01 Preface

anacron 是 cronie(Modern Versions) 中的一部分，但是和 cron 不一样。cron 会认为主机是一直运行的，而 anacron 不认为主机是一直运行的[^1]

假设 cron 和 anacron 的使用方式都一致，有如下 scheduled task

```
0 4 * * * /usr/bin/local/trojan-shell -h 223.5.5.5
```

如果配置了该 scheduled task 的主机，每天的 23:00 - 06:00 都执行关机，在第二天的 10:00 开机

- cron 就不会在开机后执行该 scheduled task，因为 cron 认为主机是一直运行的，过了就不会执行
- anacron 会根据 timestamp files(存储在 `/var/spool/anacron`，后面会提到) 来判断开机后是否需要执行该 scheduled task

但是实际上 cron 和 anacron 的字段不一样，cron 有 5 个字段来表示 cron part，而 anacron 只有 2 个字段来表示 cron part

## 0x02 Anacrontab Files

anacron 默认会从 `/etc/anacrontab` 读取 anacron 任务

anacrontab files 的格式和 crontab files 不同

### 0x02a Comments and Nonsense Characters

anacrontab files 注释，空行 和 crontab files 一致([Cron 02 - Cron#0x03a Comments and Nonsense Characters](Cron%2002%20-%20Cron.md#0x03a%20Comments%20and%20Nonsense%20Characters))

### 0x02b Anacron Entries

Anacron Entries 每一行代表一个任务

```
#period in days   delay in minutes   job-identifier   command
1       5       cron.daily              nice run-parts /etc/cron.daily
```

细分为 4 部分

#### period in days

以天为单位，执行的频率。可以是一个 integer 也可以是一个 marco 例如

- `@daily`
	等价于 1
- `@weekly`
	等价于 7
- `@month`
	根据月份来判断

#### delay in minutes

anacron 会在 delay in minutes 后执行 scheduled job

#### job identifier

对应 schedule job 的唯一标识符

#### command

实际执行的命令或者是脚本，逻辑和 [Cron 02 - Cron#command part](Cron%2002%20-%20Cron.md#command%20part) 一样

### 0x02c Environments

anacrontab files 中的环境变量大体上和 crontab files 一致([Cron 02 - Cron#0x04c Environments](#0x04c%20Environments))，但额外增加一个变量

- `START_HOURS_RANGE`
	指定了 anacron scheduled job 允许在什么时间段内运行，如果 scheduled job 不在这个时间段内就不会运行
- `RANDOM_DELAY`
	为每个 schedule job 额外随机增加的最大的 delay(in mintues)，例如 `RANDOM_DELAY=12` 表示为每个 schedule job 额外随机增加 0 - 12 minutes delay。如果值为 0 表示不额外增加

## 0x03 `/etc/anacrontab`

从上面可以知道 anacron 默认会从 `/etc/anacrontab` 读取 anacron 任务，通常默认内容如下

```
SHELL=/bin/sh
PATH=/sbin:/bin:/usr/sbin:/usr/bin
MAILTO=root
# the maximal random delay added to the base delay of the jobs
RANDOM_DELAY=45
# the jobs will be started during the following hours only
START_HOURS_RANGE=3-22

#period in days   delay in minutes   job-identifier   command
1       5       cron.daily              nice run-parts /etc/cron.daily
7       25      cron.weekly             nice run-parts /etc/cron.weekly
@monthly 45     cron.monthly            nice run-parts /etc/cron.monthly
```

`START_HOURS_RANGE` 表示 anacron 只会在每天的 03:00 - 22:00 执行 scheduled tasks

- id `cron.daily` 每天都会执行，当 $current time - timestamp file stores time >= 1 day$ 时，任务会在 5 mins 内执行
- id `cron.weekly` 每周都会执行，当 $current time - timestamp file stores time >= 1 week$ 时，任务会在 25 mins 内执行
- id `cron.monthly` 每月都会执行，当 $current time - timestamp file stores time >= 1 month$ 时，仍会会在 45 mins 内执行

> [!NOTE] 
> 没有使用 `nice` 默认优先级值为 0，使用 `nice` 但是没有指定 `-n` 参数默认优先级值为 10，值越大优先级越低

`nice` 降低 `run-parts` 运行的优先级，以确保系统的一些核心应用不会因为 scheduled task 被影响

从 `run-parts` 可以确定，`/etc/cron.{daily,weekly,month}` 中的内容就是一些 executable file

```
[root@vbox etc]# tree cron*
cron.allow [error opening dir]
cron.d
└── 0hourly
cron.daily
├── logrotate
└── man-db.cron
cron.deny [error opening dir]
cron.hourly
└── 0anacron
cron.monthly
crontab [error opening dir]
cron.weekly
```

在 RHel 中通常会有一个对应 `logrotate` 的脚本，来定时执行 `logrotate` 的切割任务

```
#!/bin/sh

/usr/sbin/logrotate -s /var/lib/logrotate/logrotate.status /etc/logrotate.conf
EXITVALUE=$?
if [ $EXITVALUE != 0 ]; then
    /usr/bin/logger -t logrotate "ALERT exited abnormally with [$EXITVALUE]"
fi
exit 0
```

## 0x04 Anacron Cmd[^5]

Anacron 没有类似 `crontab` 的命令，用于管理 anacrontab files。但是提供了一个告诉 anacron 该如何执行 scheduled task 的指令 `anacron`

```
 anacron [-s] [-f] [-n] [-d] [-q] [-t anacrontab] [-S spooldir] [job]
 anacron [-S spooldir] -u [-t anacrontab] [job]
 anacron [-V|-h]
 anacron -T [-t anacrontab]
```

### 0x04a Optional args

- `-f`
	不考虑 timestamp files，立刻执行 `/etc/anacrontab` 中的内容，delay in time 还是生效的
- `-u`
	将 timestamp files
- `-s`
	以串行的方式，执行 `/etc/anacrontab` 中的内容，只有前一个任务执行完了，才可以执行下一个任务
- `-n`
	立即执行 `/etc/anacrontab` 中的内容，delay in time 不生效，但是会考虑 timestamp files。同时也会默认启用 `-s`
	通常和 `-f` 一起使用，立即执行 anacron 的 scheduled task
- `-d`
	理解成 debug 模式，会将执行的 job identifier 输出到 stdout
- `-t <anacrontab_file>`
	使用指定的 anacrontab file 而不是默认的 `/etc/anacrontab`
- `-T`
	校验 anacrontab file 的语法是否准确

## 0x05 Anacron Timestamp Files

Anacron timestamp files 是 anacron 判断 scheduled tasks 是否执行的因子。通常储存在 `/var/spool/anacron`，文件名对应 job-identifier(如果 schedule task 被删除了，这些 timestamp files 不会被 anacron 自动删除)

例如 默认的 `/etc/anacrontab` 中有如下内容

```
1       5       cron.daily              nice run-parts /etc/cron.daily
7       25      cron.weekly             nice run-parts /etc/cron.weekly
@monthly 45     cron.monthly            nice run-parts /etc/cron.monthly
```

那么 `/var/spool/anacron` 就会有如下三个文件，每个文件记录了前一次任务的时间戳(因为 anacron 最小的精度就是 daily，所以没有必要记录 时分秒)

```
[root@vbox anacron]# for f in $(ls /var/spool/anacron);do echo ${f} && cat ${f};done
cron.daily
20241031
cron.monthly
20241030
cron.weekly
20241030
```

假设现在，我们自己自定义了一个 anacrontab file，内容如下

```
cat anacrontab_test
@monthly 0 anacrontab_test echo "hello world" && fzf
```

使用 `anacron -t anacrontab_test` (也可以和 `-d` 一起使用方便 debug)就会生成一个
`/var/spool/anacron/anacrontab_test` timestamp file，内容为 current time(即对应的 schedule task 上一次执行的时间)

```
[root@vbox ~]# cat /var/spool/anacron/anacrontab_test
20241101
```

同时会将对应的执行结果以 mail 的形式记录到 `/var/spool/mail/root` 中

```
[root@vbox ~]# tail -f /var/spool/mail/root

From root@vbox.localdomain  Fri Nov  1 11:39:43 2024
Return-Path: <root@vbox.localdomain>
X-Original-To: root
Delivered-To: root@vbox.localdomain
Received: by vbox.localdomain (Postfix, from userid 0)
        id 50EAE600CFB0; Fri,  1 Nov 2024 11:39:43 +0800 (CST)
From: Anacron <root@vbox.localdomain>
To: root@vbox.localdomain
Content-Type: text/plain; charset="UTF-8"
Subject: Anacron job 'anacrontab_test' on vbox
Message-Id: <20241101033943.50EAE600CFB0@vbox.localdomain>
Date: Fri,  1 Nov 2024 11:39:43 +0800 (CST)

hello world
/bin/sh: fzf: command not found
```

如果我们再次使用 `ancron` 执行这个 anacrontab file，scheduled task 就不会被调用(也可以使用 `tail -f /var/spool/mail/root` 来观察)

```
[root@vbox ~]# anacron -d  -t anacrontab_test
Anacron started on 2024-11-01
Checking against 0 with 31
Normal exit (0 jobs run)
[root@vbox ~]# anacron -d  -t anacrontab_test
Anacron started on 2024-11-01
Checking against 0 with 31
Normal exit (0 jobs run)
```

我们先使用 `watch -n 0.1 'stat /var/spool/anacron/anacrontab_test && cat /var/spool/anacron/anacron_test'` 监视 timestamp file 的变化

```
Every 0.1s: stat /var/spool/anacron/anacrontab_test && cat /var/spool/anacron...  Fri Nov  1 12:02:47 2024

  File: ‘/var/spool/anacron/anacrontab_test’
  Size: 9               Blocks: 8          IO Block: 4096   regular file
Device: fd00h/64768d    Inode: 100716455   Links: 1
Access: (0600/-rw-------)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:system_cron_spool_t:s0
Access: 2024-11-01 12:02:40.732000000 +0800
Modify: 2024-11-01 12:02:40.618000000 +0800
Change: 2024-11-01 12:02:40.618000000 +0800
 Birth: -
20241101
```

这时我们手动更改 `/var/spool/anacron/anacron_tab` 中记录的时间戳

```
[root@vbox ~]# echo 20231101 > /var/spool/anacron/anacrontab_test
```

然后再次使用 `anacron` 就会发现，scheduled task 被执行了。同时 timestamp file 中的时间戳就会变成当前的

```
Every 0.1s: stat /var/spool/anacron/anacrontab_test && cat /var/spool/anacron...  Fri Nov  1 12:04:24 2024

  File: ‘/var/spool/anacron/anacrontab_test’
  Size: 9               Blocks: 8          IO Block: 4096   regular file
Device: fd00h/64768d    Inode: 100716455   Links: 1
Access: (0600/-rw-------)  Uid: (    0/    root)   Gid: (    0/    root)
Context: unconfined_u:object_r:system_cron_spool_t:s0
Access: 2024-11-01 12:04:22.296000000 +0800
Modify: 2024-11-01 12:04:22.275000000 +0800
Change: 2024-11-01 12:04:22.275000000 +0800
 Birth: -
20241101
```

所以我们可以得出如下结论

> [!important] 
> anacron 会根据 timestamp file 中的时间戳来判断是否执行 scheduled task
> - $timestamp + period in days <= current time$ 执行对应 job identifier 的 schedule task
> - $timestamp + peroid in days > current time$ 不执行对应 job identifier 的 schedule task


## 0x06 Source Code Analyzing

上面我们知道 anacron 判断 scheduled task 是否需要执行，是根据 timestamp file 中的时间戳来的。那么为了加深印象，我们看一下源码

[入口函数](https://github.com/cronie-crond/cronie/blob/master/anacron/main.c#L439)，主要的逻辑就是运行 `launch_job`，`wait_children`

```c
int
main(int argc, char *argv[])
{
	...
	parse_opts(argc, argv);
	...
    if (!no_daemon && !testing_only)
	go_background();
    else
	primary_pid = getpid();

    record_start_time();
    read_tab(cwd);
    close(cwd);
    arrange_jobs();
    ...
    explain_intentions();
    set_signal_handling();
    running_jobs = running_mailers = 0;
    for(j = 0; j < njobs; ++j)
    {
	xsleep(time_till(job_array[j]));
	if (serialize) wait_jobs();
	launch_job(job_array[j]);
    }
    wait_children()
    ...
}
```

[record_start_time()](https://github.com/cronie-crond/cronie/blob/master/anacron/main.c#L363) 会对 `start_sec` 赋值(后面会用到，值为 current time)

```c
static void
record_start_time(void)
{
    struct tm *tm_now;

    start_sec = time(NULL);
    tm_now = localtime(&start_sec);
    year = tm_now->tm_year + 1900;
    month = tm_now->tm_mon + 1;
    day_of_month = tm_now->tm_mday;
    day_now = day_num(year, month, day_of_month);
    if (day_now == -1) die("Invalid date (this is really embarrassing)");
    if (!update_only && !testing_only)
	explain("Anacron started on %04d-%02d-%02d",
		year, month, day_of_month);
}

```

[read_tab()](https://github.com/cronie-crond/cronie/blob/master/anacron/readtab.c#L384) 会对 [job_array]([job_array](https://github.com/cronie-crond/cronie/blob/master/anacron/global.h#L80))  赋值(后面会用到，实际上就是一个链表，每一个 node 对应 anacrontab files 中的一个 scheduled task 和 timestamp file)

```c
...
job_rec **job_array;
...
void
read_tab(int cwd)
/* Read the anacrontab file into memory */
{
    char *tab_line;

    first_job_rec = last_job_rec = NULL;
    first_env_rec = last_env_rec = NULL;
    jobs_read = 0;
    line_num = 0;
    /* Open the anacrontab file */
    if (fchdir(cwd)) die_e("Can't chdir to original cwd");
    tab = fopen(anacrontab, "r");
    if (chdir(spooldir)) die_e("Can't chdir to %s", spooldir);

    if (tab == NULL) die_e("Error opening %s", anacrontab);
    /* Initialize the obstacks */
    obstack_init(&input_o);
    obstack_init(&tab_o);
    while ((tab_line = read_tab_line()) != NULL)
    {
	line_num++;
	parse_tab_line(tab_line);
	obstack_free(&input_o, tab_line);
    }
    if (fclose(tab)) die_e("Error closing %s", anacrontab);
}
...
void
arrange_jobs(void)
/* Make an array of pointers to jobs that are going to be executed,
 * and arrange them in the order of execution.
 * Also lock these jobs.
 */
{
    job_rec *j;

    j = first_job_rec;
    njobs = 0;
    while (j != NULL)
    {
	if (j->arg_num != -1 && (update_only || testing_only || consider_job(j)))
	{
	    njobs++;
	    obstack_grow(&tab_o, &j, sizeof(j));
	}
	j = j->next;
    }
    job_array = obstack_finish(&tab_o);

    /* sort the jobs */
    qsort(job_array, (size_t)njobs, sizeof(*job_array),
	  (int (*)(const void *, const void *))execution_order);
}
```

`arrange_jobs()` 不仅会对 `job_array` 赋值，还会对 `njobs` 赋值。其中 `j->arg_num` 值为 `job_arg_num(ident)` (通常是 0)，由 [job_arg_num()](https://github.com/cronie-crond/cronie/blob/master/anacron/readtab.c#L112) 生成，而 `job_nargs` 和 `job_args` 在 [parse_opts()](https://github.com/cronie-crond/cronie/blob/master/anacron/main.c#L101) 被定义

```c
	...
	char *defarg = "*";
	...
	static void
	parse_opts(int argc, char *argv[])
	 {  
		 ... 
		if (optind == argc)
	    {
		/* no arguments. Equivalent to: `*' */
		job_nargs = 1;
		job_args = &defarg;
	    }
	    ...
	}
```

所以就会运行 [consider_job()](https://github.com/cronie-crond/cronie/blob/master/anacron/lock.c#L77)，而该函数就会根据 timestamp file 来判断 anacron 是否运行 scheduled task

- 当 $day_now - timestamp >=0 && day_now - timestamp < period in days$ 时就不会执行当前的 scheduled task
- 当 $day_now - timestamp >=0 && day_now - timestamp >= period in days$ 就会执行当前的 scheduled task 这部分逻辑没有在 `consider_job()` 中体现，但是如果 task 满足运行的条件就会加入到 `job_array` 这个链表中

这里没有考虑 macro，具体逻辑看源码

```c
int
consider_job(job_rec *jr)
{
	char timestamp[9];
	open_tsfile(jr);
    int ts_year, ts_month, ts_day, dn;
    ssize_t b;
    b = read(jr->timestamp_fd, timestamp, 8);
	...
	int day_delta;
	...
	if (sscanf(timestamp, "%4d%2d%2d", &ts_year, &ts_month, &ts_day) == 3)
	    dn = day_num(ts_year, ts_month, ts_day);
	else
	    dn = 0;
	...
	day_delta = day_now - dn;
	if (day_delta >= 0 && day_delta < jr->period)
	{
            /* yes, skip job */
	    xclose(jr->timestamp_fd);
	    return 0;
	}

}
```

这里的 `jr->timestamp_fd` 在 [open_tsfile](https://github.com/cronie-crond/cronie/blob/master/anacron/lock.c#L40) 中赋值。`jr->ident` 在 [register_\[period\]_job](https://github.com/cronie-crond/cronie/blob/master/anacron/readtab.c#L161) 中赋值，值即为 anacron file 对应 schedule tasked 的 job-identifier 

> [!NOTE]
> 需要注意的一点是这里读取的文件是在 `/var/spool/anacron` 下的，因为在 `readtab()` 中使用 `if (chdir(spooldir))` 更改了 working directory
> `spooldir` 在 `configure.ac` 中定义

```c
static void
open_tsfile(job_rec *jr)
/* Open the timestamp file for job jr */
{
    jr->timestamp_fd = open(jr->ident, O_RDWR | O_CREAT, S_IRUSR | S_IWUSR);
    if (jr->timestamp_fd == -1)
	die_e("Can't open timestamp file for job %s", jr->ident);
    fcntl(jr->timestamp_fd, F_SETFD, 1);    /* set close-on-exec flag */
    /* We want to own this file, and set its mode to 0600. This is necessary
     * in order to prevent other users from putting locks on it. */
    if (fchown(jr->timestamp_fd, getuid(), getgid()))
	die_e("Can't chown timestamp file %s", jr->ident);
    if (fchmod(jr->timestamp_fd, S_IRUSR | S_IWUSR))
	die_e("Can't chmod timestamp file %s", jr->ident);
}
```

回到主函数，剩下的核心逻辑就是执行 [launch_job(job_array)](https://github.com/cronie-crond/cronie/blob/master/anacron/runjob.c#L289)

```c
    for(j = 0; j < njobs; ++j)
    {
	xsleep(time_till(job_array[j]));
	if (serialize) wait_jobs();
	launch_job(job_array[j]);
    }
```

***References***

- [anacron - Wikipedia](https://en.wikipedia.org/wiki/Anacron)
- `man anacron.8`
- `man anacrontab.5`

***FootNotes***

[^1]:[Confused about relationship between cron and anacron - Ask Ubuntu](https://askubuntu.com/questions/848610/confused-about-relationship-between-cron-and-anacron)

[^4]:[anacron(8) - Linux manual page](https://www.man7.org/linux/man-pages/man8/anacron.8.html)
[^5]:[anacrontab(5) - Linux manual page](https://www.man7.org/linux/man-pages/man5/anacrontab.5.html)


