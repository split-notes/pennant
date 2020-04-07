package bash

import (
	"fmt"
	"strings"
)

var (
	GoModVendor = "go mod vendor"
)

func GoProtoc(protoDir string) string { return fmt.Sprintf("protoc --go_out=plugins=grpc:. %s/*.proto", protoDir)}

func GoTest() string { return fmt.Sprintf("go test $(go list ./...)")}

// Mockgen has two versions
// SOURCE-based Mockgen looks at a file and generates a mock for an interface in the file
func GoMockSource(mockName string) string {
	titleCaseMockName := strings.Title(strings.ToLower(mockName))
	return fmt.Sprintf("mockgen" +
		" -source=./services/%s/interface.go" +
		" -destination=./mocks/services_mocks/%s_mock.go" +
		" -package=services_mocks" +
		" -mock_names Service=Mock%s",
		mockName, mockName, titleCaseMockName)
}

// REFLECT-based Mockgen generates a program that can understand the interface \
// 	and uses that knowledge to generate your mock.
func GoMockReflect(submodule string, mockName string) string {
	titleCaseMockName := strings.Title(strings.ToLower(mockName))
	return fmt.Sprintf("mockgen" +
		" -destination=./mocks/services_mocks/%s_mock.go" +
		" -package=services_mocks" +
		" github.com/split-notes/%s/services/grpc_service %sClient",
		mockName, submodule, titleCaseMockName)
}

