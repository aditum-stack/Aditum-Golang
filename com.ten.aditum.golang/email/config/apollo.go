package config

import "fmt"

//"github.com/zouyx/agollo"

//// 从配置中心获取数据
//func getConfigFromApollo(apolloConfig *ApolloConfig) {
//	agollo.Start()
//	apolloConfig = &ApolloConfig{
//		agollo.GetStringValue("Host", "/mail"),
//		agollo.GetStringValue("Port", "12001"),
//		agollo.GetStringValue("Api", "/")}
//}

// 配置中心数据结构
type ApolloConfig struct {
	Host string `config:"host"`
	Port string `config:"port"`
	Api  string `config:"api"`
}

// 默认配置数据
func getConfigDefault(apolloConfig *ApolloConfig) {
	apolloConfig = &ApolloConfig{
		"/mail",
		"12001",
		"/"}
}

func init() {
	fmt.Print("use agollo fail")
}
