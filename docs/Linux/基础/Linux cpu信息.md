# Linux 获取cpu

1. 系统文件，使用`|grep cores`查看具体核数

   ```
   root in /etc/ssh λ cat /proc/cpuinfo 
          File: /proc/cpuinfo
      1   processor   : 0
      2   vendor_id   : GenuineIntel
      3   cpu family  : 6
      4   model       : 158
      5   model name  : Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
      6   stepping    : 9
      7   microcode   : 0xb4
      8   cpu MHz     : 2808.001
      9   cache size  : 6144 KB
     10   physical id : 0
     11   siblings    : 1
     12   core id     : 0
     13   cpu cores   : 1
     14   apicid      : 0
     15   initial apicid  : 0
     16   fpu     : yes
     17   fpu_exception   : yes
     18   cpuid level : 22
     19   wp      : yes
     20   flags       : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ss syscall nx pdpe1gb rdtscp 
          lm constant_tsc arch_perfmon nopl xtopology tsc_reliable nonstop_tsc cpuid pni pclmulqdq ssse3 fma cx16 pcid sse4_1 sse4_2 x2apic movbe popcn
          t tsc_deadline_timer aes xsave avx f16c rdrand hypervisor lahf_lm abm 3dnowprefetch cpuid_fault invpcid_single pti ssbd ibrs ibpb stibp fsgsb
          ase tsc_adjust bmi1 avx2 smep bmi2 invpcid mpx rdseed adx smap clflushopt xsaveopt xsavec xsaves arat md_clear flush_l1d arch_capabilities
     21   bugs        : cpu_meltdown spectre_v1 spectre_v2 spec_store_bypass l1tf mds swapgs itlb_multihit srbds
     22   bogomips    : 5616.00
     23   clflush size    : 64
     24   cache_alignment : 64
     25   address sizes   : 43 bits physical, 48 bits virtual
     26   power management:
     27   
     28   processor   : 1
     29   vendor_id   : GenuineIntel                                                                  
   ```

2. 命令`lscpu`，使用`grep | cores`查看cpu核数

   ```
   root in /etc/ssh λ lscpu                    
   Architecture:                    x86_64
   CPU op-mode(s):                  32-bit, 64-bit
   Byte Order:                      Little Endian
   Address sizes:                   43 bits physical, 48 bits virtual
   CPU(s):                          4
   On-line CPU(s) list:             0-3
   Thread(s) per core:              1
   Core(s) per socket:              1
   Socket(s):                       4
   NUMA node(s):                    1
   Vendor ID:                       GenuineIntel
   CPU family:                      6
   Model:                           158
   Model name:                      Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
   Stepping:                        9
   CPU MHz:                         2808.001
   BogoMIPS:                        5616.00
   Hypervisor vendor:               VMware
   Virtualization type:             full
   L1d cache:                       128 KiB
   L1i cache:                       128 KiB
   L2 cache:                        1 MiB
   L3 cache:                        24 MiB
   NUMA node0 CPU(s):               0-3
   Vulnerability Itlb multihit:     KVM: Vulnerable
   Vulnerability L1tf:              Mitigation; PTE Inversion
   Vulnerability Mds:               Mitigation; Clear CPU buffers; SMT Host state unknown
   Vulnerability Meltdown:          Mitigation; PTI
   Vulnerability Spec store bypass: Mitigation; Speculative Store Bypass disabled via prctl and seccomp
   Vulnerability Spectre v1:        Mitigation; usercopy/swapgs barriers and __user pointer sanitization
   Vulnerability Spectre v2:        Mitigation; Full generic retpoline, IBPB conditional, IBRS_FW, STIBP disabled, RSB filling
   Vulnerability Srbds:             Unknown: Dependent on hypervisor status
   Vulnerability Tsx async abort:   Not affected
   Flags:                           fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ss syscall nx pdp
                                    e1gb rdtscp lm constant_tsc arch_perfmon nopl xtopology tsc_reliable nonstop_tsc cpuid pni pclmulqdq ssse3 fma cx16
                                     pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand hypervisor lahf_lm abm 3dnowpr
                                    efetch cpuid_fault invpcid_single pti ssbd ibrs ibpb stibp fsgsbase tsc_adjust bmi1 avx2 smep bmi2 invpcid mpx rdse
                                    ed adx smap clflushopt xsaveopt xsavec xsaves arat md_clear flush_l1d arch_capabilities                       /0.0s
   ```

   
