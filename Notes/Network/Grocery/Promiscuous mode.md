# Promiscuous mode

promiscuous mode（混杂模式）一般用于NIC，让流量不是经过NIC到目的NIC，而是先到CPU再到目的NIC。

通常用于流量嗅探，例如tcpdump，wireshark都会将NIC置于混杂模式