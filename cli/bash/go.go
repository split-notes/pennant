package bash

import "fmt"

var (
	GoModVendor = "go mod vendor"
)

func GoProtoc(protoDir string) string { return fmt.Sprintf("protoc --go_out=plugins=grpc:. %s/*.proto", protoDir)}
func GoMock(mockName string) string {
	return fmt.Sprintf("mockgen" +
		" -source=./services/%s/interface.go" +
		" -destination=./mocks/services_mocks/%s_mock.go" +
		" -package=services_mocks" +
		" -mock_names Service=Mock_%s",
	mockName, mockName, mockName)
}
func GoTest() string { return fmt.Sprintf("go test $(go list ./...)")}

