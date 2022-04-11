# dmidecode

ref:

https://en.wikipedia.org/wiki/System_Management_BIOS

## DMI/SMBIOS

System Management BIOS，规定了OS可以从BIOS中读取的信息(主要是硬件相关的信息)，DMI(Desktop management interface)有更多的信息，其中包含SMBIOS



## dmidecode

我们可以通过`dmidecode`来读取dmi table，每条entry由下面格式组成

```
       Handle 0x0002, DMI type 2, 8 bytes.  Base  Board  Informa‐
       tion
               Manufacturer: Intel
               Product Name: C440GX+
               Version: 727281-001
               Serial Number: INCY92700942

```

- A handler：唯一标识符，十六进制
- A type：SMBIOS定义了不同组件type不同
- A size：数据的大小，一般比实际要小
- decode value：根据type不同，显示的内容不同

可以展现的数据有大致如下：

- System

  包含机器型号，序列号等信息

  ```
  System Information
          Manufacturer: LENOVO
          Product Name: 82MS
          Version: Lenovo XiaoXinPro 14ACH 2021
          Serial Number: PF2PTW6L
          UUID: cb18d6e6-7fe9-11eb-80ec-38f3ab15da50
          Wake-up Type: Power Switch
          SKU Number: LENOVO_MT_82MS_BU_idea_FM_XiaoXinPro 14ACH 2021
          Family: XiaoXinPro 14ACH 2021
  ```

- Memory

  这里可以看出最大物理内存为16GB，由2条内存条组成双通

  ```
  Handle 0x0001, DMI type 16, 23 bytes
  Physical Memory Array
          Location: System Board Or Motherboard
          Use: System Memory
          Error Correction Type: None
          Maximum Capacity: 16 GB
          Error Information Handle: 0x0000
          Number Of Devices: 2
  
  ```

  可以使用如下命令来查看

  ```
  cpl in /etc λ sudo dmidecode -t memory 
  ```

- GPU cache

  可以使用如下命令来查看

  ```
  dmidecode -t cache
  ```

- CPU

  可以使用如下命令来查看

  ```
  cpl in /etc λ sudo dmidecode -t processor
  ```


## options

- `-q | --quite`

  只显示decode value

- `-s | --string KEYWORD`

  只显示指定字段，具体查看manual page

  ```
  ➜  /etc dmidecode -s bios-vendor
  LENOVO
  ```

- `-t | --type TYPE`

  只显示指定type的，具体可以使用的type查看man page DMI TYPES







