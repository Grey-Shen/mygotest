package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	yaml "gopkg.in/yaml.v2"
	"qiniupkg.com/api.v7/kodo"
)

const configpath = "/Users/chenyajun/Documents/goproject/src/download/base.yml"

type config struct {
	RSHost          string   `yaml:"rs_host"`
	RSFHost         string   `yaml:"rsf_host"`
	IoHost          string   `yaml:"io_host"`
	ApiHost         string   `yaml:"api_host"`
	UpHosts         []string `yaml:"up_hosts"`
	DomainHost      string   `yaml:"domain_host"`
	UrlLifetime     uint32   `yaml:"url_lifetime"`
	UptokenLifetime uint32   `yaml:"uptoken_lifetime"`
	AdminUser       string   `yaml:"admin_username"`
	AdminPassword   string   `yaml:"admin_password"`
	EmailDomain     string   `yaml:"email_domain"`
	BucketDomain    string   `yaml:"bucket_domain"`
	BucketName      string   `yaml:"bucket_name"`
	AccessKey       string   `yaml:"access_key"`
	SecretKey       string   `yaml:"secret_key"`
	MaxRetry        int      `yaml:"max_retry"`
}

func main() {

	var (
		baseConfig *config
		err        error
		fileKey    string
	)

	flag.StringVar(&fileKey, "k", "test.txt", "key of Pabos")
	flag.Parse()

	if baseConfig, err = loadConfig(configpath); err != nil {
		fmt.Println("Loadconfig failed err:", err)
		return
	}

	var kodoConfig = kodo.Config{
		AccessKey: baseConfig.AccessKey,
		SecretKey: baseConfig.SecretKey,
		RSHost:    baseConfig.RSHost,
		RSFHost:   baseConfig.RSFHost,
		IoHost:    baseConfig.IoHost,
		UpHosts:   baseConfig.UpHosts,
		Transport: http.DefaultTransport,
	}

	baseUrl := kodo.MakeBaseUrl(baseConfig.BucketDomain, fileKey)
	policy := kodo.GetPolicy{Expires: 60}
	c := kodo.New(0, &kodoConfig)
	privateUrl := c.MakePrivateUrl(baseUrl, &policy)
	fmt.Println(privateUrl)
}

func loadConfig(filepath string) (*config, error) {
	var baseConfig config

	buf, err := ioutil.ReadFile(filepath)

	if err != nil {
		return nil, fmt.Errorf("Failed to open qiniu config file: %s", err)
	}
	fmt.Println("buf is ", string(buf))
	err = yaml.Unmarshal(buf, &baseConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse qiniu config file: %s", err)
	}

	return &baseConfig, nil
}
