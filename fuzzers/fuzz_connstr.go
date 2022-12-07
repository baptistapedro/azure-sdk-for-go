package myfuzz

import "github.com/Azure/azure-sdk-for-go/sdk/data/aztables"

func Fuzz(data []byte) int {
	connStr := string(data)
	_, err := aztables.NewServiceClientFromConnectionString(connStr, nil)
	if err != nil { return 1 }
	return 0
}
