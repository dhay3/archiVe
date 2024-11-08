                         chkrootkit V. 0.54
    
          Nelson Murilo <nelson@pangeia.com.br> (main author)
            Klaus Steding-Jessen <jessen@cert.br> (co-author)
    
          This program locally checks for signs of a rootkit.
         chkrootkit is available at: http://www.chkrootkit.org/


                 No illegal activities are encouraged!
         I'm not responsible for anything you may do with it.
    
           This tool includes software developed by the
           DFN-CERT, Univ. of Hamburg (chklastlog and chkwtmp),
           and small portions of ifconfig developed by
           Fred N. van Kempen, <waltje@uwalt.nl.mugnet.org>.


 1. What's chkrootkit?
---------------------

 chkrootkit is a tool to locally check for signs of a rootkit.  It
 contains:

 * chkrootkit: a shell script that checks system binaries for
   rootkit modification.

 * ifpromisc.c: checks if the network interface is in promiscuous
   mode.

 * chklastlog.c: checks for lastlog deletions.

 * chkwtmp.c: checks for wtmp deletions.

 * check_wtmpx.c: checks for wtmpx deletions.  (Solaris only)

 * chkproc.c: checks for signs of LKM trojans.

 * chkdirs.c: checks for signs of LKM trojans.

 * strings.c: quick and dirty strings replacement.

 * chkutmp.c: checks for utmp deletions.

 chkwtmp and chklastlog *try* to check for deleted entries in the wtmp
 and lastlog files, but it is *not* guaranteed that any modification
 will be detected.

 Aliens tries to find sniffer logs and rootkit config files.  It looks
 for some default file locations -- so it is also not guaranteed it
 will succeed in all cases.

 chkproc checks if /proc entries are hidden from ps and the readdir
 system call.  This could be the indication of a LKM trojan.  You can
 also run this command with the -v option (verbose).


 2. Rootkits, Worms and LKMs detected
------------------------------------

 For an updated list of rootkits, worms and LKMs detected by
 chkrootkit please visit: http://www.chkrootkit.org/


 3. Supported Systems
--------------------

 chkrootkit has been tested on: Linux 2.0.x, 2.2.x, 2.4.x and 2.6.x,
 FreeBSD 2.2.x, 3.x, 4.x and 5.x, OpenBSD 2.x, 3.x and 4.x., NetBSD
 1.6.x, Solaris 2.5.1, 2.6, 8.0 and 9.0, HP-UX 11, Tru64, BSDI and Mac
 OS X.


 4. Package Contents
-------------------

 README
 README.chklastlog
 README.chkwtmp
 COPYRIGHT
 chkrootkit.lsm

 Makefile
 chklastlog.c
 chkproc.c
 chkdirs.c
 chkwtmp.c
 check_wtmpx.c
 ifpromisc.c
 strings.c
 chkutmp.c

 chkrootkit


 5. Installation
---------------

 To compile the C programs type:

 # make sense

 After that it is ready to use and you can simply type:

 # ./chkrootkit


 6. Usage
--------

 chkrootkit must run as root.  The simplest way is:

 # ./chkrootkit

 This will perform all tests.  You can also specify only the tests you
 want, as shown below:

 Usage: ./chkrootkit [options] [testname ...]
 Options:
         -h                show this help and exit
         -V                show version information and exit
         -l                show available tests
         -d                debug
         -q                quiet mode
         -x                expert mode
         -r dir            use dir as the root directory
         -p dir1:dir2:dirN path for the external commands used by chkrootkit
         -n                skip NFS mounted dirs

 Where testname stands for one or more from the following list:

 aliens asp bindshell lkm rexedcs sniffer w55808 wted scalper slapper
 z2 chkutmp amd basename biff chfn chsh cron crontab date du dirname
 echo egrep env find fingerd gpm grep hdparm su ifconfig inetd
 inetdconf identd init killall ldsopreload login ls lsof mail mingetty
 netstat named passwd pidof pop2 pop3 ps pstree rpcinfo rlogind rshd
 slogin sendmail sshd syslogd tar tcpd tcpdump top telnetd timed
 traceroute vdir w write

 For example, the following command checks for trojaned ps and ls
 binaries and also checks if the network interface is in promiscuous
 mode.

   # ./chkrootkit ps ls sniffer

 The `-q' option can be used to put chkrootkit in quiet mode -- in
 this mode only output messages with `infected' status are shown.

 With the `-x' option the user can examine suspicious strings in the
 binary programs that may indicate a trojan -- all the analysis is
 left to the user.

 Lots of data can be seen with:

   # ./chkrootkit -x | more

 Pathnames inside system commands:

   # ./chkrootkit -x | egrep '^/'

 chkrootkit uses the following commands to make its tests: awk, cut,
 egrep, find, head, id, ls, netstat, ps, strings, sed, uname.  It is
 possible, with the `-p' option, to supply an alternate path to
 chkrootkit so it won't use the system's (possibly) compromised
 binaries to make its tests.

 To use, for example, binaries in /cdrom/bin:

   # ./chkrootkit -p /cdrom/bin

 It is possible to add more paths with a `:'

   # ./chkrootkit -p /cdrom/bin:/floppy/mybin

 Sometimes is a good idea to mount the disk from a compromised machine
 on a machine you trust.  Just mount the disk and specify a new
 rootdir with the `-r' option.

 For example, suppose the disk you want to check is mounted under
 /mnt, then:

   # ./chkrootkit -r /mnt


 7. Output Messages
------------------

 The following messages are printed by chkrootkit (except with the -x
 and -q command options) during its tests:

   "INFECTED": the test has identified a command probably modified by
   a known rootkit;

   "not infected": the test didn't find any known rootkit signature.

   "not tested": the test was not performed -- this could happen in
   the following situations:
     a) the test is OS specific;
     b) the test depends on an external program that is not available;
     c) some specific command line options are given. (e.g. -r ).

   "not found": the command to be tested is not available;

   "Vulnerable but disabled": the command is infected but not in use.
   (not running or commented in inetd.conf)


 8. A trojaned command has been found.  What should I do now?
------------------------------------------------------------

 Your biggest problem is that your machine has been compromised and
 this bad guy has root privileges.

 Maybe you can solve the problem by just replacing the trojaned
 command -- the best way is to reinstall the machine from a safe media
 and to follow your vendor's security recommendations.


 9. Reports and questions
------------------------

 Please send comments, questions and bug reports to
 nelson@pangeia.com.br and jessen@cert.br.

 A simple FAQ and Related information about rootkits and security can
 be found at chkrootkit's homepage, http://www.chkrootkit.org.


 10. ACKNOWLEDGMENTS
-------------------

 See the ACKNOWLEDGMENTS file.
