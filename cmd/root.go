package cmd

import (
	"crypto/tls"
	"encoding/base64"
	"github.com/Jeffail/gabs"
	"github.com/aiziyuer/registryV2/impl/common"
	"github.com/aiziyuer/registryV2/impl/registry"
	"github.com/aiziyuer/registryV2/impl/util"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"path"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:          "registryV2",
	SilenceUsage: false,
}

var level string
var outputFormat string
var registryV2ConfigDir string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}

func init() {

	// detect the log level
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {

		level, err := logrus.ParseLevel(level)
		if err != nil {
			return err
		}

		logrus.SetLevel(level)

		return nil
	}

	rootCmd.PersistentFlags().StringVarP(
		&level,
		"verbose",
		"v",
		logrus.WarnLevel.String(),
		"Log level (trace, debug, info, warn, error, fatal, panic)",
	)

	rootCmd.PersistentFlags().StringVarP(
		&outputFormat,
		"output", "o", "table",
		"options output format: table, yaml, json",
	)

	home, err := homedir.Dir()
	if err != nil {
		logrus.Fatal(err)
	}

	rootCmd.PersistentFlags().StringVarP(
		&registryV2ConfigDir,
		"config", "", path.Join(home, ".registryV2"),
		"location of config files like $REGISTRY_V2_CONFIG ",
	)
}

func getClient(host string) (*registry.Registry, error) {

	var client *registry.Registry = nil

	// 尝试读取存储密码
	jsonParsed, _ := gabs.ParseJSONFile(path.Join(registryV2ConfigDir, ".registryV2/config.json"))
	encodedAuth := jsonParsed.
		Search("auths").
		Search(util.GetEnvAnyWithDefault("registry-1.docker.io", "REGISTRY_V2_HOST")).
		Search("auth").
		Data()

	if encodedAuth != nil {
		ret, err := base64.StdEncoding.DecodeString(encodedAuth.(string))
		if err == nil {
			nameAndPass := string(ret)

			client = registry.NewClient(&http.Client{
				Transport: &http.Transport{
					Proxy: http.ProxyFromEnvironment,
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				},
			}, &registry.Endpoint{
				Schema: util.GetEnvAnyWithDefault("https", "REGISTRY_V2_SCHEMA"),
				Host:   util.GetEnvAnyWithDefault(host, "REGISTRY_V2_HOST"),
			}, &common.Auth{
				UserName: strings.SplitN(nameAndPass, ":", 2)[0],
				PassWord: strings.SplitN(nameAndPass, ":", 2)[1],
			})
		}

	}

	// 实在不行就采用无认证客户端
	if client == nil {
		client = registry.NewClient(&http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}, &registry.Endpoint{
			Schema: util.GetEnvAnyWithDefault("https", "REGISTRY_V2_SCHEMA"),
			Host:   util.GetEnvAnyWithDefault(host, "REGISTRY_V2_HOST"),
		}, nil)

	}

	return client, nil
}
