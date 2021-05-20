# rack-detect

kube-vip init image for rack awareness and BGP peering

## Overview
Uses hostIP and environment variables to determine BGP peers.

This has only been used to peer with top of rack switches

**NOTE** Currently requires modification of kube-vip to work.
**NOTE** Only supports up to 3 racks/subnets/peerlists.

## Features
Echos a script to use to start kube-vip with BGP_PEER and BGP_AS environment variables.

## Requirements to work as designed for kube-vip
container environment variables for each network cidr - NET_CIDR1, NET_CIDR2, NET_CID3
container environment variables for each BGP "group" - BGP_PEERS1, BGP_PEERS2, BGP_PEERS3
container environment variables for BGP_AS - BGP_AS1, BGP_AS2, BGP_AS3
downward api environment variable status.hostIP

container command/args used to create start script for kube-vip

rack-detect init container example
```
spec:
  initContainers:
  - name: rack-detect
    command: ["/bin/sh","-c"]
    args: ["/rack-detect > /tmp/shared/kube-vip-start.sh; chmod 755 /tmp/shared/kube-vip-start.sh"]
    env:
    - name: NET_CIDR1
      value: 192.168.30.0/24
    - name: NET_CIDR2
      value: 174.0.2.0/24
    - name: NET_CIDR3
      value: 174.0.3.0/24
    - name: BGP_AS1
      value: "65001"
    - name: BGP_AS2
      value: "22222"
    - name: BGP_AS3
      value: "33333"
    - name: BGP_PEERS1
      value: 192.168.10.140:65000::true
    - name: BGP_PEERS2
      value: 2.1.1.1:12345::false,2.1.1.2:12345::false
    - name: BGP_PEERS3
      value: 3.1.1.1:54321::false,3.1.1.2:54321::false
    - name: IP
      valueFrom:
        fieldRef:
          fieldPath: status.hostIP
    ...
    volumeMounts:
    - mountPath: /tmp/shared
      name: shared-vol
```


kube-vip example
```
containers:
  - command: ["/bin/sh","-c"]
    args: ["/tmp/shared/kube-vip-start.sh"]
    env:
    - name: vip_arp
      value: "false"
    - name: vip_interface
      value: eth0
    - name: port
      value: "6443"
    - name: vip_cidr
      value: "24"
    - name: cp_enable
      value: "true"
    - name: cp_namespace
      value: kube-system
    - vip_startleader
      value: "false"
    - name: bgp_enable
      value: "true"
    - name: bgp_routerid
      valueFrom:
        fieldRef:
          fieldPath: status.hostIP
    - name: bgp_peeraddress
    - name: bgp_peerpass
    - name: bgp_peers
    - name: vip_address
      value: 7.8.9.10
    imagePullPolicy: IfNotPresent
    name: kube-vip
    ...
    volumeMounts:
    - mountPath: /tmp/shared
      name: shared-vol

```

