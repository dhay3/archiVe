# whois

## 0x01 Overview

whois 是一个 7 层协议(在 wireshark 中会被直接标识为 WHOIS protocol)，通过 object (可以是 域名，IP 或者是关键字)查询如下信息

1. domain names
2. IP address blocks
3. autonomous system

```
16:58:51.522365 enp4s0 Out IP 10.100.4.222.33612 > 192.34.234.30.43: Flags [S], seq 2988502373, win 24820, options [mss 1460,sackOK,TS val 703843859 ecr 0,nop,wscale 11], length 0
16:58:51.689238 enp4s0 Out IP 10.100.4.222.33612 > 192.34.234.30.43: Flags [.], ack 1749887454, win 24820, options [nop,nop,TS val 703844026 ecr 1585605014], length 0
16:58:51.689261 enp4s0 Out IP 10.100.4.222.33612 > 192.34.234.30.43: Flags [P.], seq 0:18, ack 1, win 24820, options [nop,nop,TS val 703844026 ecr 1585605014], length 18: WHOIS: domain baidu.com
16:58:51.846758 enp4s0 Out IP 10.100.4.222.33612 > 192.34.234.30.43: Flags [.], ack 1, win 24820, options [nop,nop,TS val 703844184 ecr 1585605014], length 0
16:58:51.847333 enp4s0 Out IP 10.100.4.222.33612 > 192.34.234.30.43: Flags [.], ack 1189, win 23760, options [nop,nop,TS val 703844184 ecr 1585605181], length 0
16:58:51.847453 enp4s0 Out IP 10.100.4.222.33612 > 192.34.234.30.43: Flags [.], ack 2377, win 23760, options [nop,nop,TS val 703844184 ecr 1585605181], length 0
16:58:51.847555 enp4s0 Out IP 10.100.4.222.33612 > 192.34.234.30.43: Flags [.], ack 3565, win 23760, options [nop,nop,TS val 703844184 ecr 1585605181], length 0
16:58:51.847585 enp4s0 Out IP 10.100.4.222.33612 > 192.34.234.30.43: Flags [.], ack 3588, win 23760, options [nop,nop,TS val 703844185 ecr 1585605181], length 0
16:58:51.847601 enp4s0 Out IP 10.100.4.222.33612 > 192.34.234.30.43: Flags [F.], seq 18, ack 3588, win 23760, options [nop,nop,TS val 703844185 ecr 1585605181], length 0
```

而 `whois` 命令是 whois 协议在 Unix 上的实现

## 0x02 Options

```
whois  [  {  -h  |  --host  }  HOST  ]  [  {  -p  |  --port  }  PORT ] [ -abBcdGHIKlLmMrRx ] [ -g SOURCE:FIRST-LAST ] [ -i ATTR[,ATTR]... ] [ -s SOURCE[,SOURCE]... ] [ -T TYPE[,TYPE]... ] [ --verbose ] [ --no-recursion ] OBJECT

whois -q KEYWORD

whois -t TYPE

whois -v TYPE

whois --help

whois --version

```

## 0x03 Example

根据查询的 object 不同，返回的结果也不同

例如 domain name

```
 whois baidu.com
   Domain Name: BAIDU.COM
   Registry Domain ID: 11181110_DOMAIN_COM-VRSN
   Registrar WHOIS Server: whois.markmonitor.com
   Registrar URL: http://www.markmonitor.com
   Updated Date: 2023-11-30T06:00:19Z
   Creation Date: 1999-10-11T11:05:17Z
   Registry Expiry Date: 2026-10-11T11:05:17Z
   Registrar: MarkMonitor Inc.
   Registrar IANA ID: 292
   Registrar Abuse Contact Email: abusecomplaints@markmonitor.com
   Registrar Abuse Contact Phone: +1.2086851750
   Domain Status: clientDeleteProhibited https://icann.org/epp#clientDeleteProhibited
   Domain Status: clientTransferProhibited https://icann.org/epp#clientTransferProhibited
   Domain Status: clientUpdateProhibited https://icann.org/epp#clientUpdateProhibited
   Domain Status: serverDeleteProhibited https://icann.org/epp#serverDeleteProhibited
   Domain Status: serverTransferProhibited https://icann.org/epp#serverTransferProhibited
   Domain Status: serverUpdateProhibited https://icann.org/epp#serverUpdateProhibited
   Name Server: NS1.BAIDU.COM
   Name Server: NS2.BAIDU.COM
   Name Server: NS3.BAIDU.COM
   Name Server: NS4.BAIDU.COM
   Name Server: NS7.BAIDU.COM
   DNSSEC: unsigned
   URL of the ICANN Whois Inaccuracy Complaint Form: https://www.icann.org/wicf/
>>> Last update of whois database: 2024-05-06T09:17:44Z <<<

For more information on Whois status codes, please visit https://icann.org/epp

NOTICE: The expiration date displayed in this record is the date the
registrar's sponsorship of the domain name registration in the registry is
currently set to expire. This date does not necessarily reflect the expiration
date of the domain name registrant's agreement with the sponsoring
registrar.  Users may consult the sponsoring registrar's Whois database to
view the registrar's reported date of expiration for this registration.

TERMS OF USE: You are not authorized to access or query our Whois
database through the use of electronic processes that are high-volume and
automated except as reasonably necessary to register domain names or
modify existing registrations; the Data in VeriSign Global Registry
Services' ("VeriSign") Whois database is provided by VeriSign for
information purposes only, and to assist persons in obtaining information
about or related to a domain name registration record. VeriSign does not
guarantee its accuracy. By submitting a Whois query, you agree to abide
by the following terms of use: You agree that you may use this Data only
for lawful purposes and that under no circumstances will you use this Data
to: (1) allow, enable, or otherwise support the transmission of mass
unsolicited, commercial advertising or solicitations via e-mail, telephone,
or facsimile; or (2) enable high volume, automated, electronic processes
that apply to VeriSign (or its computer systems). The compilation,
repackaging, dissemination or other use of this Data is expressly
prohibited without the prior written consent of VeriSign. You agree not to
use electronic processes that are automated and high-volume to access or
query the Whois database except as reasonably necessary to register
domain names or modify existing registrations. VeriSign reserves the right
to restrict your access to the Whois database in its sole discretion to ensure
operational stability.  VeriSign may restrict or terminate your access to the
Whois database for failure to abide by these terms of use. VeriSign
reserves the right to modify these terms at any time.

The Registry database contains ONLY .COM, .NET, .EDU domains and
Registrars.
Domain Name: baidu.com
Registry Domain ID: 11181110_DOMAIN_COM-VRSN
Registrar WHOIS Server: whois.markmonitor.com
Registrar URL: http://www.markmonitor.com
Updated Date: 2023-11-30T04:07:58+0000
Creation Date: 1999-10-11T11:05:17+0000
Registrar Registration Expiration Date: 2026-10-11T07:00:00+0000
Registrar: MarkMonitor, Inc.
Registrar IANA ID: 292
Registrar Abuse Contact Email: abusecomplaints@markmonitor.com
Registrar Abuse Contact Phone: +1.2086851750
Domain Status: clientUpdateProhibited (https://www.icann.org/epp#clientUpdateProhibited)
Domain Status: clientTransferProhibited (https://www.icann.org/epp#clientTransferProhibited)
Domain Status: clientDeleteProhibited (https://www.icann.org/epp#clientDeleteProhibited)
Domain Status: serverUpdateProhibited (https://www.icann.org/epp#serverUpdateProhibited)
Domain Status: serverTransferProhibited (https://www.icann.org/epp#serverTransferProhibited)
Domain Status: serverDeleteProhibited (https://www.icann.org/epp#serverDeleteProhibited)
Registrant Organization: Beijing Baidu Netcom Science Technology Co., Ltd.
Registrant State/Province: Beijing
Registrant Country: CN
Registrant Email: Select Request Email Form at https://domains.markmonitor.com/whois/baidu.com
Admin Organization: Beijing Baidu Netcom Science Technology Co., Ltd.
Admin State/Province: Beijing
Admin Country: CN
Admin Email: Select Request Email Form at https://domains.markmonitor.com/whois/baidu.com
Tech Organization: Beijing Baidu Netcom Science Technology Co., Ltd.
Tech State/Province: Beijing
Tech Country: CN
Tech Email: Select Request Email Form at https://domains.markmonitor.com/whois/baidu.com
Name Server: ns1.baidu.com
Name Server: ns7.baidu.com
Name Server: ns4.baidu.com
Name Server: ns2.baidu.com
Name Server: ns3.baidu.com
DNSSEC: unsigned
URL of the ICANN WHOIS Data Problem Reporting System: http://wdprs.internic.net/
>>> Last update of WHOIS database: 2024-05-06T09:16:30+0000 <<<

For more information on WHOIS status codes, please visit:
  https://www.icann.org/resources/pages/epp-status-codes

If you wish to contact this domain’s Registrant, Administrative, or Technical
contact, and such email address is not visible above, you may do so via our web
form, pursuant to ICANN’s Temporary Specification. To verify that you are not a
robot, please enter your email address to receive a link to a page that
facilitates email communication with the relevant contact(s).

Web-based WHOIS:
  https://domains.markmonitor.com/whois

If you have a legitimate interest in viewing the non-public WHOIS details, send
your request and the reasons for your request to whoisrequest@markmonitor.com
and specify the domain name in the subject line. We will review that request and
may ask for supporting documentation and explanation.

The data in MarkMonitor’s WHOIS database is provided for information purposes,
and to assist persons in obtaining information about or related to a domain
name’s registration record. While MarkMonitor believes the data to be accurate,
the data is provided "as is" with no guarantee or warranties regarding its
accuracy.

By submitting a WHOIS query, you agree that you will use this data only for
lawful purposes and that, under no circumstances will you use this data to:
  (1) allow, enable, or otherwise support the transmission by email, telephone,
or facsimile of mass, unsolicited, commercial advertising, or spam; or
  (2) enable high volume, automated, or electronic processes that send queries,
data, or email to MarkMonitor (or its systems) or the domain name contacts (or
its systems).

MarkMonitor reserves the right to modify these terms at any time.

By submitting this query, you agree to abide by this policy.

MarkMonitor Domain Management(TM)
Protecting companies and consumers in a digital world.

Visit MarkMonitor at https://www.markmonitor.com
Contact us at +1.8007459229
In Europe, at +44.02032062220
--

```

例如 IP

```
 whois 39.156.66.10
% [whois.apnic.net]
% Whois data copyright terms    http://www.apnic.net/db/dbcopyright.html

% Information related to '39.128.0.0 - 39.191.255.255'

% Abuse contact for '39.128.0.0 - 39.191.255.255' is 'abuse@chinamobile.com'

inetnum:        39.128.0.0 - 39.191.255.255
netname:        CMNET
descr:          China Mobile Communications Corporation
descr:          Mobile Communications Network Operator in China
descr:          Internet Service Provider in China
country:        CN
org:            ORG-CM1-AP
admin-c:        ct74-AP
tech-c:         HL1318-AP
abuse-c:        AC2006-AP
status:         ALLOCATED PORTABLE
remarks:        service provider
remarks:        --------------------------------------------------------
remarks:        To report network abuse, please contact mnt-irt
remarks:        For troubleshooting, please contact tech-c and admin-c
remarks:        Report invalid contact via www.apnic.net/invalidcontact
remarks:        --------------------------------------------------------
mnt-by:         APNIC-HM
mnt-lower:      MAINT-CN-CMCC
mnt-routes:     MAINT-CN-CMCC
mnt-irt:        IRT-CHINAMOBILE-CN
last-modified:  2020-10-20T00:58:36Z
source:         APNIC

irt:            IRT-CHINAMOBILE-CN
address:        China Mobile Communications Corporation
address:        29, Jinrong Ave., Xicheng District, Beijing, 100032
e-mail:         abuse@chinamobile.com
abuse-mailbox:  abuse@chinamobile.com
admin-c:        CT74-AP
tech-c:         CT74-AP
auth:           # Filtered
remarks:        abuse@chinamobile.com was validated on 2024-02-06
mnt-by:         MAINT-CN-CMCC
last-modified:  2024-02-06T13:22:19Z
source:         APNIC

organisation:   ORG-CM1-AP
org-name:       China Mobile
org-type:       LIR
country:        CN
address:        29, Jinrong Ave.
phone:          +86-10-5268-6688
fax-no:         +86-10-5261-6187
e-mail:         hostmaster@chinamobile.com
mnt-ref:        APNIC-HM
mnt-by:         APNIC-HM
last-modified:  2023-09-05T02:14:48Z
source:         APNIC

role:           ABUSE CHINAMOBILECN
address:        China Mobile Communications Corporation
address:        29, Jinrong Ave., Xicheng District, Beijing, 100032
country:        ZZ
phone:          +000000000
e-mail:         abuse@chinamobile.com
admin-c:        CT74-AP
tech-c:         CT74-AP
nic-hdl:        AC2006-AP
remarks:        Generated from irt object IRT-CHINAMOBILE-CN
remarks:        abuse@chinamobile.com was validated on 2024-02-06
abuse-mailbox:  abuse@chinamobile.com
mnt-by:         APNIC-ABUSE
last-modified:  2024-02-06T13:23:22Z
source:         APNIC

role:           chinamobile tech
address:        29, Jinrong Ave.,Xicheng district
address:        Beijing
country:        CN
phone:          +86 5268 6688
fax-no:         +86 5261 6187
e-mail:         hostmaster@chinamobile.com
admin-c:        HL1318-AP
tech-c:         HL1318-AP
nic-hdl:        ct74-AP
notify:         hostmaster@chinamobile.com
mnt-by:         MAINT-cn-cmcc
abuse-mailbox:  abuse@chinamobile.com
last-modified:  2016-11-29T09:37:27Z
source:         APNIC

person:         haijun li
nic-hdl:        HL1318-AP
e-mail:         hostmaster@chinamobile.com
address:        29,Jinrong Ave, Xicheng district,beijing,100032
phone:          +86 1052686688
fax-no:         +86 10 52616187
country:        CN
mnt-by:         MAINT-CN-CMCC
abuse-mailbox:  abuse@chinamobile.com
last-modified:  2016-11-29T09:38:38Z
source:         APNIC

% This query was served by the APNIC Whois Service version 1.88.25 (WHOIS-AU1)
```

**references**

[^1]:https://en.wikipedia.org/wiki/WHOIS
[^2]:https://github.com/rfc1036/whois

