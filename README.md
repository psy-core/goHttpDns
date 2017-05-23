## goHttpDns

[![Go Report Card](https://goreportcard.com/badge/psy-core/goHttpDns)](https://goreportcard.com/report/github.com/psy-core/goHttpDns)

A HttpDns Server Written by Go, In order to avoid Dns hijacking and cache resolve answer

Go HttpDns 服务, 抵抗 DNS 劫持污染，并带有缓存功能 。

Fork From：[http://github.com/zheng-ji/goHttpDns](http://github.com/zheng-ji/goHttpDns)

### How To Compile

```
cd $GOPATH;
git clone http://github.com/psy-core/goHttpDns;
./build.sh
```

### How To Configure

```
# redis connect config
redis:
  host: 127.0.0.1:6379
  network: tcp
  db: 0
  connectTimeout: 5000
  readTimeout: 10000
  writeTimeout: 10000
  maxActive: 3
  maxIdle: 5
  idleTimeout: 10
  wait: false

# seelog config 
log_config: ../etc/logger.xml

# ip & port & answer cache TTL
listen: 0.0.0.0
port: 9999
ttl: 100

# DnsServer lists
dnsservers:
    - 202.96.128.86
    - 202.96.128.166
    - 8.8.8.8
    - 8.8.4.4
```

### How To Run

After `./build.sh`
```
$ cd dist
$ bin/start.sh
```

### How To Use

```
$ curl http://127.0.0.1:9999/d?url=http://zheng-ji.info

Resp:
{
    "c":0,
    "targetip":"http://106.185.48.24",
    "host":"zheng-ji.info",
    "msg":""
}
```

### Dependence Third Part Lib

Thanks to:

* [launchpad/goyaml](https://launchpad.net/goyaml)
* [cihub/seelog](github.com/cihub/seelog)
* [miekg/dns](github.com/miekg/dns)
* [redisgo/redis](github.com/garyburd/redigo/redis")

You need to `go get` the list above


License
-------

Copyright (c) 2015 by [zheng-ji](zheng-ji.info) released under a MIT style license.
