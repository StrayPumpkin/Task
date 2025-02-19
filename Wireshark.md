# Wireshark

## 布局

![](https://raw.githubusercontent.com/StrayPumpkin/img/refs/heads/main/cd2ae876bbae9e5cac0b6bb64da29c3d.png)

## 捕获过滤器
用于确定什么样的信息记录会显示在捕获结果中。

**注意**：需要在开始捕获前设置。设置捕获过滤器

**过滤器格式**：`协议 方向 主机端口(host/port/portrange)`

**protocol**：http、https、ftp、udp、tcp、ipv4、ipv6、arp、icmp等协议。

**direction**： src(源地址）、dst（目标地址）、src and dst、src or dst 默认src or dst

**host(s)**:net（过滤与特定网络（IP地址段）相关的流量）、port（过滤与特定端口相关的流量）、host（过滤与特定主机（IP地址或域名）相关的流量）、portrange（过滤与特定端口范围相关的流量） 默认host

**logical operation**（逻辑运算）：not and or not具有最高优先级，and or优先级相同，运算从左向右

```
显示ip协议，来源ip为10.1.1.1的封包:ip src host 10.1.1.1
显示ip协议，来源目标ip为10.1.1.1的封包:ip host 10.1.1.
```

## 显示过滤器

显示过滤器是wireshark的第二层过滤器（第一层为捕获过滤器），他能很快的过滤出我们需要的封包数据。

**比较操作符**：==、！=、<、>、>=、=
**逻辑操作符**：and、or、not(没有条件满足)、xor(有且仅有一个条件满足)
I**P地址**：ip.addr（来源ip地址或者目标ip地址）、ip.src（来源ip地址限制）、ip.dst（目标ip地址限制）
**协议过滤**：arp、ip、icmp、udp、tcp、bootp、dns

```
显示snmp或者dns或者icmp协议的封包:snmp || dns || icmp
显示来源或者目标地址为10.1.1.1的封包:ip.addr == 10.1.1.1
```

## Packet List Pane(数据包列表)

![](https://raw.githubusercontent.com/StrayPumpkin/img/refs/heads/main/013799a29e1238d1e787267f631d133c.png)

- **No.**：封包的序号，按照捕获的顺序递增。
- **Time**：封包的时间戳，显示了封包捕获的时间。
- **Source**：封包的源地址，即发送封包的主机的 IP 地址。
- **Destination**：封包的目标地址，即接收封包的主机的 IP 地址。
- **Protocol**：封包使用的协议，例如 TCP、UDP 或 ICMP。
- **Length**：封包的长度，以字节为单位。
- **Info**：封包的简要信息，通常包括一些关键字段的值，如 seq、ack、win、mss 等。

##  Packet Details Pane(数据包详细信息)左下角
在数据包列表中选择指定数据包，在数据包详细信息中会显示数据包的所有详细信息内容。数据包详细信息面板是最重要的，用来查看协议中的每一个字段。各行信息分别为

- **Frame**: 物理层的数据帧概况
- **Ethernet II**: 数据链路层以太网帧头部信息
- **Internet Protocol Version 4**: 互联网层IP包头部信息
- **Transmission Control Protocol**: 传输层T的数据段头部信息，此处是TCP
- **Hypertext Transfer Protocol**: 应用层的信息，此处是HTTP协议