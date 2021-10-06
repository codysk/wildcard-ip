# wildcard-ip
a simple wildcard DNS.

## Usage
wildcard-ip allows you to do that by mapping any IP Address to a hostname using the following formats:

```
// simple
10.0.0.1.domain.name -> 10.0.0.1
192-168-1-1.domain.name -> 192.168.1.1

// with subdomain

subdomain.10.0.0.1.domain.name -> 10.0.0.1
anylevel.subdomain.10.0.0.1.domain.name -> 10.0.0.1

subdomain.192-168-1-1.domain.name -> 192.168.1.1
anylevel.subdomain.192-168-1-1.domain.name -> 192.168.1.1
```

## Installation

1. pull the image

`docker pull ghcr.io/codysk/wildcard-ip:latest`

2. run and expose port 53

`docker run -d --name wildcardip -p 53:53/tcp -p 53:53/udp ghcr.io/codysk/wildcard-ip:latest`

3. set ns record to your host at domain management
