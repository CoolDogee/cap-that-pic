module github.com/cooldogee/cap-that-pic

require (
	github.com/Azure/azure-sdk-for-go v36.1.0+incompatible
	github.com/Azure/go-autorest/autorest v0.9.2
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.4.0
	github.com/gobuffalo/packr/v2 v2.2.0
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/onsi/ginkgo v1.7.0
	github.com/onsi/gomega v1.4.3
	github.com/ugorji/go/codec v0.0.0-20181204163529-d75b2dcb6bc8 // indirect
	go.mongodb.org/mongo-driver v1.5.1
)

go 1.13

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
