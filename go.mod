module github.com/yinliang-gl/golang_project

go 1.12

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/fwhezfwhez/errorx v0.0.0-20190917084916-7b3d6baebfde
	github.com/gin-gonic/gin v1.4.0 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/mock v1.2.0
	github.com/golang/protobuf v1.3.1
	github.com/json-iterator/go v1.1.6
	github.com/opentracing/opentracing-go v1.1.0
	github.com/robfig/cron/v3 v3.0.2-0.20200518143530-6a8421bcff44
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spaolacci/murmur3 v1.1.0
	github.com/tealeg/xlsx v1.0.5
	github.com/valyala/fasthttp v1.9.0
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297
	google.golang.org/grpc v1.21.0
	gopkg.in/gcfg.v1 v1.2.3
	gopkg.in/warnings.v0 v0.1.2 // indirect
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.37.4
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190506204251-e1dfcc566284
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190121172915-509febef88a4
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3
	golang.org/x/net => github.com/golang/net v0.0.0-20190514140710-3ec191127204
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190226205417-e64efc72b421
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190514135907-3a4b5fb9f71f
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20181108054448-85acf8d2951c
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190425222832-ad9eeb80039a
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.3.1
	google.golang.org/appengine => github.com/golang/appengine v1.4.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190513181449-d00d292a067c
	google.golang.org/grpc => github.com/grpc/grpc-go v1.19.1
	k8s.io/apimachinery => github.com/kubernetes/apimachinery v0.0.0-20190515023456-b74e4c97951f
	k8s.io/klog => github.com/kubernetes/klog v0.3.0
	k8s.io/kube-openapi => github.com/kubernetes/kube-openapi v0.0.0-20190228160746-b3a7cee44a30
	sigs.k8s.io/yaml => github.com/kubernetes-sigs/yaml v1.1.0
)
