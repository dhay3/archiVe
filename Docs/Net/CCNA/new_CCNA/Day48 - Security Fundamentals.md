# Day48 - Security Fundamentals

## CIA

安全中有一个 CIA Triad 准则

1. Confidentiality(保密)

   - Only authorized users should be able to access data

   - Some information/data is public and can be accessed by anyone, some is secret and should only be accessed by specific people

2. Integrity(完整)

   - Data should not be tampered with(modified) by unauthorized users

   - Data should be correct and authentic

3. Availability(可用)

   - The network/systems should be operational and accessible to authorized users

## Vulnerability/Exploit/Threat/Mitigation

- A **vulnerability** is any potential weakness that can compromise the CIA of a system/info

  漏洞

- An **exploit** is something that can potentially be used to exploit the vulnerability

  攻击漏洞的东西

- A **threat** is the potential of a vulnerability to be exploited

  可以被用于攻击的漏洞

- A **mitigation** technique is somthing that can protect against threats

  保护漏洞的技术

## Common Attacks

- DoS(denial-of-service) attacks
- Spoofing attacks
- Reflection/amplification attacks
- Man-in-the-middle attacks
- Reconnaissance attack
- Malware
- Social engineering attacks
- Password-related attacks

### DoS

最常见的 DoS 攻击就是 TCP SYN flood

1. The attacker sends countless TCP SYN messages to the target
2. The target sends a SYN-ACK message in response to each SYN it receives
3. The attacker never replies with the final ACK of the TCP three-way handshake
4. The incomplete connections fill up the target’s TCP connection table
5. The attacker continues sending SYN messages
6. The target is no longer able to make legitimate TCP connections

例如

通常 Attacker 都会使用代理，所以通常 Target 回送的 SYN-ACK 并不会直接到 Attacker

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_11-29.1jvxjecqbfwg.webp)

但是一般会由多台机器组成 botnet 来攻击，也被称为 DDoS(Distributed Denial-Of-Service)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_11-34.4gvyqbnq85q8.webp)

### Spoofing attacks

使用假的 IP 或者 MAC 地址，例如

DHCP exhaustion

- An attacker uses spoofed MAC addresses to flood DHCP Discover messages
- The target server’s DHCP pool becomes full, resulting in a denial-of-service to other devices

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_11-38.37hgn4fjd1z4.webp)

> 仅仅伪装源 MAC 地址，报文其他的部分值和 DHCP Discover 一样

还有一种就是 ARP Spoofing，伪装是目的主机回送 ARP Reply

### Reflection/Amplification

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_12-18.a260qyguc1s.webp)

### Man-in-the-middle attack

ARP Spoofing 就是一种中间人攻击

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_12-18.a260qyguc1s.webp)

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_12-21.1tix2qv2bhsw.webp)

### Reconnaissance attacks

Reconnaissance(侦查) 攻击，顾名思义就是收集攻击目标的信息

![](https://cdn.staticaly.com/gh/dhay3/image-repo@master/20230718/2023-07-18_12-23.28zqdm80neio.webp)

### Malware

Malware(malicious software) 即有害的程序，例如

- Viruses

  viruses infect other software(a ‘host program’). The virus spreads as the software is shared by users. Typically they corrupt or modify files on the target computer

- Worms

  Worms do not require a host program. They are standalone malware and they are able to spread on their own, without user interaction. The spread of worms can congest the network, but the ‘payload’ of a worm can cause additional harm to target devices

- Trojan Horses

  Trojna Horses are harmful software that is disguised as legitimate software. They are spread through user interaction such as opening email attachments, or donwloading a file from the Internet

### Social Engineering attack

Social engineering attacks target the most vulnerable part of any system  - people

社会工程学就是通过“人”这个漏洞来攻击系统的

- Phising

  typically involves fraudulent emails that appear to come from a legitimate business and contain links to a fradulent website that seems legitimate. Users are told to login to the fraudulent website, providing their login credentials to the attacker

  钓鱼

- Vishing

  voice phishing

  诈骗电话

- Smishing

  SMS phishing

  诈骗短信

- Watering hole

- Tailgating

  尾随

### Password-related attack

针对密码的攻击

- Dictionary attack

  使用字典来破解密码

- Brute force attack

  暴力破解密码

## Authtication

### Multi-factor atuthentication

Multi-factor authentication(MFA) 提供多种除账户密码认证外的方式，例如短信、指纹

即使攻击者知道了你的账号密码，也不能直接使用你的账号，还需要通过 MFA

### Digital certificates

通过数字证书认证

## AAA

AAA(triple-A) Stands for 

Authentication,

*Authentication is the process of verifying a user’s identify*

*logging in = authentication*

Authorization,

*Authorization is the process of granting the user the appropriate access and permission*

*granting the user access to some files/services, restring access to other files/server = authorization*

Accounting

审计

*Accounting is the process of recording the user’s ativities on the System.*

*Logging when a user make a change to a file = accounting*

通常企业会使用 AAA server 来提供 AAA 服务，例如 Cisco ISE 就是 Cisco 旗下的 AAA server

AAA servers 通常支持两种 AAA 协议

1. RADUIS

   an open standard protocol. Uses UDP port 1812 and 1813

2. TACACS+

   A cisco propriety protocol. Uses TCP port 49

## Security Program Element

这里 Program 并不是指程序，而是活动

1. User awareness program

   例如模拟钓鱼邮件的活动，如果员工点击了，就通报，以此来提高员工的安全意识

2. User training program

   例如安全考试

3. Physical access control

   例如门禁卡、指纹锁

## LAB

https://www.youtube.com/watch?v=EBs47-0ZD-A&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=94

**references**

1. [^https://www.youtube.com/watch?v=VvFuieyTTSw&list=PLxbwE86jKRgMpuZuLBivzlM8s2Dk5lXBQ&index=93]