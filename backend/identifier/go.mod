module github.com/athomecomar/athome/backend/identifier

go 1.14

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/andybalholm/cascadia v1.2.0 // indirect
	github.com/athomecomar/envconf v1.1.0
	github.com/athomecomar/semantic v0.1.2
	github.com/athomecomar/xerrors v1.2.1
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/gocolly/colly v1.2.0
	github.com/gocolly/colly/v2 v2.0.1
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.4.1 // indirect
	github.com/pkg/errors v0.9.1
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2 // indirect
	golang.org/x/sys v0.0.0-20200515095857-1151b9dac4a9 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200515170657-fc4c6c6a6587 // indirect
	google.golang.org/grpc v1.29.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.29.1
