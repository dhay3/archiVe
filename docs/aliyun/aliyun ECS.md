# aliyun ECS

## 概述

云服务器ECS(Elastic Computer Service)

### 基础概念

- 实例：等同于一台虚拟服务器

- 镜像：提供实例的操作系统、初始化应用数据及预装的软件

- 块存储：基于分布式存储架构的云盘以及基于物理机本地存储的本地盘。

  > 块存储包括云盘和本地盘
  >
  > https://help.aliyun.com/document_detail/63136.html?spm=a2c4g.11186623.2.20.314428cc66A999#concept-pl4-tzb-wdb

- 快照：snapshot

- 安全组：虚拟防火墙，类似iptables
- 专用网络：逻辑上彻底隔离的云上私有网络

## 实例元数据

使用echo命令让输出内容换行

