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
curl localhost:8000/orders -d '{"userId":1, "selling":true, "quantity": 3, "price":10}'; sleep 0.1
curl localhost:8000/orders -d '{"userId":2, "selling":false, "quantity": 1, "price":11}'; sleep 0.1
curl localhost:8000/orders -d '{"userId":2, "selling":false, "quantity": 1, "price":11}'; sleep 0.1
curl localhost:8000/orders -d '{"userId":2, "selling":false, "quantity": 2, "price":11}'; sleep 0.1
```

it returns:

```bash
{"userID":1,"selling":true,"quantity":3,"remain":3,"price":10,"createAt":1544676693}
{"userID":2,"selling":false,"quantity":1,"remain":0,"price":11,"createAt":1544676693}
{"userID":2,"selling":false,"quantity":1,"remain":0,"price":11,"createAt":1544676693}
{"userID":2,"selling":false,"quantity":2,"remain":1,"price":11,"createAt":1544676693}
```

the final screen in server logs:

```bash
2018/12/13 12:58:29 orders: &[]*models.Order{(*models.Order)(0xc0000b63c0), (*models.Order)(0xc0000b6550), (*models.Order)(0xc00013e0a0), (*models.Order)(0xc00013e2d0)}

2018/12/13 12:58:29 orders[0]: models.Order{UserID:0x1, Selling:true, Quantity:3, Remain:0, Price:10, Matchs:[]*models.Match(nil), CreatedAt:1544677109}
2018/12/13 12:58:29 orders[1]: models.Order{UserID:0x2, Selling:false, Quantity:1, Remain:0, Price:11, Matchs:[]*models.Match{(*models.Match)(0xc00009e680)}, CreatedAt:1544677109}
2018/12/13 12:58:29 orders[1].Matchs[0]: models.Match{Order:(*models.Order)(0xc0000b63c0), Quantity:1, Price:10}
2018/12/13 12:58:29 orders[2]: models.Order{UserID:0x2, Selling:false, Quantity:1, Remain:0, Price:11, Matchs:[]*models.Match{(*models.Match)(0xc00000c080)}, CreatedAt:1544677109}
2018/12/13 12:58:29 orders[2].Matchs[0]: models.Match{Order:(*models.Order)(0xc0000b63c0), Quantity:1, Price:10}
2018/12/13 12:58:29 orders[3]: models.Order{UserID:0x2, Selling:false, Quantity:2, Remain:1, Price:11, Matchs:[]*models.Match{(*models.Match)(0xc00000c1c0)}, CreatedAt:1544677109}
2018/12/13 12:58:29 orders[3].Matchs[0]: models.Match{Order:(*models.Order)(0xc0000b63c0), Quantity:1, Price:10}
```
