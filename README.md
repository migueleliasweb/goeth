# GoETH

**GoETH** is a Go program that aims to help automate scripts find a suitable IP address to bind to.

## Installing

All releases comes with a bundled compiled GoETH binary file. You can just grab it from github releases and add it to your *$PATH*:

````bash
$ wget https://github.com/migueleliasweb/goeth/raw/master/dist/goeth
$ mv goeth your_path
$ chmod a+x your_path/goeth
```

## Commandline arguments

All arguments can be passed with on "-" or two "--" dashes.

* *exclude-localhost*: Whether to exclude localhost from the result
* *exclude-docker-network*: Whether to exclude the docker network
* *ipv6*: Whether to allow ipv6 results
* *separator*: Separator to use on the output
* *only-public*: Whether to return only public addresses
* *only-private*: Whether to return only private addresses

## Examples

No arguments: returns all private IPs
```bash
$ goeth
127.0.0.1,192.168.0.16,172.17.0.1
```

Returns just your private IP
```bash
$ goeth --exclude-localhost --exclude-docker-network
192.168.0.16
```

Returns just public IP
```bash
$ goeth --exclude-localhost --exclude-docker-network --only-public
31.13.80.36
```

Returns also IPV6 addresses
```bash
$ goeth --exclude-localhost --exclude-docker-network --ipv6
31.13.80.36,fe80::a617:31ff:fefd:99,fe80::42:62ff:fe7e:90c9
```

Different separator
```bash
$ goeth --exclude-localhost --exclude-docker-network --ipv6 --separator=" "
31.13.80.36 fe80::a617:31ff:fefd:99 fe80::42:62ff:fe7e:90c9
```
