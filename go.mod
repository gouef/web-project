module github.com/gouef/web-project

go 1.23.4

require (
	github.com/gin-contrib/multitemplate v1.1.0
	github.com/gin-gonic/gin v1.10.0
	github.com/gouef/diago v1.2.4
	github.com/gouef/finder v1.1.2
	github.com/gouef/renderer v0.4.1
	github.com/gouef/router v1.2.5
)

replace (
	github.com/gouef/web-project/app => ./app
	github.com/gouef/web-project/controllers => ./controllers
	github.com/gouef/web-project/handlers => ./handlers
	github.com/gouef/web-project/middleware => ./middleware
	github.com/gouef/web-project/models => ./models
)

require (
	github.com/bytedance/sonic v1.13.2 // indirect
	github.com/bytedance/sonic/loader v0.2.4 // indirect
	github.com/cloudwego/base64x v0.1.5 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/gin-contrib/sse v1.0.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.25.0 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/gouef/mode v1.0.7 // indirect
	github.com/gouef/utils v1.9.4 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.10 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.15.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
