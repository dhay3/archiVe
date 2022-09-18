# Autonomous system

reference:

https://www.cisco.com/c/en/us/td/docs/switches/datacenter/nexus3000/sw/unicast/503_u1_2/nexus3000_unicast_config_gd_503_u1_2/l3_overview.html#wp1114179

https://www.omnisecu.com/cisco-certified-network-associate-ccna/what-is-autonomous-system.php

Autonomous system（AS）指由一个机构完全管理的network（例如Internet Service Provider ISP 或者是大型的企业）。AS将Internet分隔成独立的routing domains。

AS内部的Route由Interior Gateway Protocol（IGP）管理，例如RIP，IGRP，EIGRP ，OSPF

AS之间的Route由Exterior Gateway Protocol（EGP）管理，例如BGP

ASN（Autonomous system number）有一个16bit的数字组成，即10进制表示最大只有65535。

0和65535都是保留的ASN，只有64512到65534被分配为私用

| 2-Byte Numbers | 4-Byte Numbers in AS.dot Notation | 4-Byte Numbers in plaintext Notation | Purpose                                                      |
| -------------- | --------------------------------- | ------------------------------------ | ------------------------------------------------------------ |
| 1 to 64511     | 0.1 to 0.64511                    | 1 to 64511                           | Public AS (assigned by RIR)[1](https://www.cisco.com/c/en/us/td/docs/switches/datacenter/nexus3000/sw/unicast/503_u1_2/nexus3000_unicast_config_gd_503_u1_2/l3_overview.html#wp1114223) |
| 64512 to 65534 | 0.64512 to 0.65534                | 64512 to 65534                       | Private AS (assigned by local administrator)                 |
| 65535          | 0.65535                           | 65535                                | Reserved                                                     |
| N/A            | 1.0 to 65535.65535                | 65536 to 4294967295                  | Public AS (assigned by RIR)                                  |

