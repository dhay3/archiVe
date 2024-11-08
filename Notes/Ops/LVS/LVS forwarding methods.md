# LVS forwarding methods

## LVS-NAT

```
                        ________
                       |        |
                       | client | (local or on internet)
                       |________|
                           |
                        (router)
                       DIRECTOR_GW
                           |
--                         |
L                      Virtual IP
i                      ____|_____
n                     |          | (director can have 1 or 2 NICs)
u                     | director |
x                     |__________|
                          DIP
V                          |
i                          |
r         -----------------+----------------
t         |                |               |
u         |                |               |
a        RIP1             RIP2            RIP3
l   ____________     ____________     ____________
   |            |   |            |   |            |
S  | realserver |   | realserver |   | realserver |
e  |____________|   |____________|   |____________|
r
v
e
r
```

## LVS-DR

```

                        ________
                       |        |
                       | client | (local or on internet)
                       |________|
                           |
                        (router)-----------
                           |    SERVER_GW  |
--                         |               |
L                         VIP              |
i                      ____|_____          |
n                     |          | (director can have 1 or 2 NICs)
u                     | director |         |
x                     |__________|         |
                          DIP              |
V                          |               |
i                          |               |
r         -----------------+----------------
t         |                |               |
u         |                |               |
a      RIP1,VIP         RIP2,VIP        RIP3,VIP
l   ____________     ____________     ____________
   |            |   |            |   |            |
S  | realserver |   | realserver |   | realserver |
e  |____________|   |____________|   |____________|
r
v
e
r
```

## LVS-Tun

```

                        ________
                       |        |
                       | client | (local or on internet)
                       |________|
                           |
                        (router)-----------
                           |    SERVER_GW  |
--                         |               |
L                         VIP              |
i                      ____|_____          |
n                     |          | (director can have 1 or 2 NICs)
u                     | director |         |
x                     |__________|         |
                          DIP              |
V                          |               |
i                          |               |
r         -----------------+----------------
t         |                |               |
u         |                |               |
a      RIP1,VIP         RIP2,VIP        RIP3,VIP
l   ____________     ____________     ____________
   |            |   |            |   |            |
S  | realserver |   | realserver |   | realserver |
e  |____________|   |____________|   |____________|
r
v
e
r
```

