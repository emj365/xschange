# xschange

[![CircleCI](https://circleci.com/gh/emj365/xschange/tree/master.svg?style=svg)](https://circleci.com/gh/emj365/xschange/tree/master)
[![](https://images.microbadger.com/badges/image/emj365/xschange.svg)](https://microbadger.com/images/emj365/xschange "Get your own image badge on microbadger.com")

An extreme simple ["Trade matching engine"](http://marketswiki.com/wiki/Trade_matching_engine) with Golang.

> A trade matching engine is the core software and hardware component of an electronic exchange. It matches up bids and offers to complete trades. Matching engines use one or several algorithms to allocate trades among competing bids and offers at the same price.

-- http://marketswiki.com/wiki/Trade_matching_engine

**Work in progress. Just for Fun.**

## Get start

### Method A: Go1.11 module

#### Install Dependecies

trun on Go1.11 module support:

```bash
export GO111MODULE=on
```

get modules:

```bash
go mod download
```

#### Run

```bash
go run main.go
```

### Method B: docker

```bash
# https://hub.docker.com/r/emj365/xschange/
docker run -p 8000:8000 xschange
```

## Run test

### Install dependencies

```bash
go install github.com/onsi/ginkgo/ginkgo
```

### Run

```bash
$ ginkgo -r

Running Suite: Xschange Models
==============================
Random Seed: 1544674849
Will run 5 of 5 specs

•••••
Ran 5 of 5 Specs in 0.000 seconds
SUCCESS! -- 5 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS

Ginkgo ran 1 suite in 1.124721406s
Test Suite Passed
```

## Demo with curl

```bash
curl localhost:8000/orders -d '{"userId":0, "selling":true, "quantity": 3, "price":10}'; sleep 0.1
curl localhost:8000/orders -d '{"userId":1, "selling":false, "quantity": 1, "price":11}'; sleep 0.1
curl localhost:8000/orders -d '{"userId":1, "selling":false, "quantity": 1, "price":11}'; sleep 0.1
curl localhost:8000/orders -d '{"userId":1, "selling":false, "quantity": 2, "price":11}'; sleep 0.1
```

it returns:

```bash
{"userID":0,"selling":true,"quantity":3,"remain":3,"price":10,"createAt":1544700262}
{"userID":1,"selling":false,"quantity":1,"remain":0,"price":11,"createAt":1544700262}
{"userID":1,"selling":false,"quantity":1,"remain":0,"price":11,"createAt":1544700262}
{"userID":1,"selling":false,"quantity":2,"remain":1,"price":11,"createAt":1544700262}
```

the final screen in server logs:

```bash
2018/12/14 14:49:16 orders: &[0xc0000b23c0 0xc00013c0a0 0xc0000b2550 0xc0000b2780]

2018/12/14 14:49:16 orders[0]: {0 true 3 0 10 [] 1544770156008362000}
2018/12/14 14:49:16 orders[1]: {1 false 1 0 11 [0xc00015a040] 1544770156130807000}
2018/12/14 14:49:16 orders[1].Matchs[0]: {0xc0000b23c0 1 10}
2018/12/14 14:49:16 orders[2]: {1 false 1 0 11 [0xc000098680] 1544770156253797000}
2018/12/14 14:49:16 orders[2].Matchs[0]: {0xc0000b23c0 1 10}
2018/12/14 14:49:16 orders[3]: {1 false 2 1 11 [0xc000098780] 1544770156376116000}
2018/12/14 14:49:16 orders[3].Matchs[0]: {0xc0000b23c0 1 10}
2018/12/14 14:49:16 users[0]: {97 130}
2018/12/14 14:49:16 users[1]: {103 70}
```
