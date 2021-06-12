module github.com/PonyWilliam/go-ProductWeb

go 1.14

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/PonyWilliam/go-area v1.0.2
	github.com/PonyWilliam/go-arealogs v1.0.3
	github.com/PonyWilliam/go-borrow v1.0.2
	github.com/PonyWilliam/go-borrow-logs v1.1.0
	github.com/PonyWilliam/go-category v1.0.2
	github.com/PonyWilliam/go-common v1.0.7
	github.com/PonyWilliam/go-product v1.0.4
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190808125512-07798873deee
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/breaker/hystrix/v2 v2.9.1
	github.com/prometheus/client_golang v1.5.1 // indirect
	github.com/tencentcloud/tencentcloud-sdk-go v1.0.127
)
