# Linux socket

ref:

https://www.digitalocean.com/community/tutorials/understanding-sockets

https://en.wikipedia.org/wiki/Network_socket

https://en.wikipedia.org/wiki/Unix_domain_socket

https://unix.stackexchange.com/questions/16311/what-is-a-socket

https://stackoverflow.com/questions/14774668/what-is-raw-socket-in-socket-programming

http://www.kernel.org/doc/man-pages/online/pages/man7/packet.7.html

https://www.ibm.com/docs/en/i/7.1?topic=programming-how-sockets-work

## Socket Digest

> socket is a pseudo-file

Socket are a way to enable inter-process communication（IPC） between programs running on a server, or between programs running on separate servers.

Communication between servers relies on network sockets, wich use the Internel Protocol to encapsulate and hanle sending and receiving data

进程之间通过 socket 通信。如果是不同主机之间的进程通过 network sockets 来通信，通常封装在 IP 数据包中

wiki 上用 eletrical female connector 来形容 socket 比较形象，进程通过 socket 来通信

local socket 和 remote socket 的组合通常也被称为 socket pairs

## Type

### Stream Socket

==Stream sockets are connection oriented==, which means that packets sent to and received from a network socket are delivered by the host operating system in order for processing by an application. Network based stream sockets typicaly use the TCP to encapsulate and transmit data over a network interface

因为是面向连接的，所以一般用在使用 TCP 通信的进程中

TCP is designed to be a reliable network protocol that relies on a  stateful connection. Data that is sent by a program using a TCP-based  stream socket will be successfully received by a remote system (assuming there are no routing, firewall, or other connectivity issues). TCP  packets can arrive on a physical network interface in any order. In the  event that packets arrive out of order, the network adapter and host  operating system will ensure that they are reassembled in the correct  sequence for processing by an application. 

### Datagram Socket

==datagram sockets are connectionless==, which means that packets sent and received from a socket are processed individually by applications. Network-based datagram sockets typically use the UDP to encapsulate and transimit data

datagram sockets 不是面向连接的，所以通常用在使用 UDP 通信的进程中

UDP does not encode sequence information in packet headers, and there is no error correction built into the protocol. Programs that use  datagram-based network sockets must build in their own error handling  and data ordering logic to ensure successful data transmission.

### Unix domain socket

A unix domain socket aka UDS is communications endpoint for exchanging data between processes executing on the same host operating system

允许==同主机==进程间通信（IPC inter-process communication）的 socket，按照 socket types 可以分为如下几种

1.  SOCK_STREAM: for a stream-oriented socket
2. SOCK_DGRAM: for a datagram-oriented socket that preserves message boundaries 

UDS are used widely be database systems that do not need to connected to a network interface. For example, MYSQL on Ubuntu defaults to using a file named `/var/run/mysqld/mysql.sock` for communication with local clients. Clients read from adn write to the socket, as does the MYSQL server itself

### Raw socket

Allow direct sending and receiving of IP packets without any protocol-specific transpport layer formatting

简单的说就是可以自己定义 L2 L3数据包的 header 和 payload 的 socket

可以使用 scapy 来实现

## Socket vs Socket address

在很多场合中 Socket, Socket file descriptor, Socket address 都被混用，没有明确的定义。在这里为了自己方便记忆，我将

1. Socket 理解成 OS 实际存在文件
2. Socket address 理解成逻辑的地址，即 Protocol and IP and Port 

## How socket works

https://www.ibm.com/docs/en/i/7.1?topic=programming-how-sockets-work

以 stream socket 为例，下图是调用系统的 API 流图

![](https://www.ibm.com/docs/en/ssw_ibm_i_71/rzab6/rxab6500.gif)

1. The socket() API creates an endpoint for communications and returns a socket descriptor that represents the endpoint.
2. When an application has a socket descriptor, it can bind a unique name to the socket. Servers must bind a name to be accessible from the network.
3. The listen() API indicates a willingness to accept client connection requests. When a listen() API is issued for a socket, that socket cannot actively initiate connection requests. The listen() API is issued after a socket is allocated with a socket() API and the bind() API binds a name to the socket. A listen() API must be issued before an accept() API is issued.
4. The client application uses a connect() API on a stream socket to establish a connection to the server.
5. The server application uses the accept() API to accept a client connection request. The server must issue the bind() and listen() APIs successfully before it can issue an accept() API.
6. When a connection is established between stream sockets (between client and server), you can use any of the socket API data transfer APIs. Clients and servers have many data transfer APIs from which to choose, such as send(), recv(), read(), write(), and others.
7. When a server or client wants to stop operations, it issues a close() API to release any system resources acquired by the socket.