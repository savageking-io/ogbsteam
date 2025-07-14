module github.com/savageking-io/ogbsteam

go 1.23.4

require (
	github.com/savageking-io/ogbcommon v0.0.0-20250712075748-d83991b6c5c9
	github.com/savageking-io/ogbsteam/proto v0.0.0
	github.com/sirupsen/logrus v1.9.3
	github.com/urfave/cli v1.22.17
	google.golang.org/grpc v1.73.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.7 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/savageking-io/ogbsteam/proto v0.0.0 => ./proto
