package app

import (
	"strings"

	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/knative-sample/tekton-proxy/cmd/app/options"
	"github.com/knative-sample/tekton-proxy/pkg/api"
	"github.com/knative-sample/tekton-proxy/pkg/version"
	"github.com/spf13/cobra"
)

// start edas api
func NewCommandStartServer(stopCh <-chan struct{}) *cobra.Command {
	ops := &options.Options{}
	mainCmd := &cobra.Command{
		Short: "hello world runner",
		Long:  "hello world runner",
		RunE: func(c *cobra.Command, args []string) error {
			glog.V(2).Infof("NewCommandStartServer main:%s", strings.Join(args, " "))
			run(stopCh, ops)
			return nil
		},
	}

	ops.SetOps(mainCmd)
	return mainCmd
}

func run(stopCh <-chan struct{}, ops *options.Options) {
	vs := version.Version().Info("tekton-proxy")
	if ops.Version {
		fmt.Println(vs)
		os.Exit(0)
	}

	if ops.ConfigPath == "" {
		glog.Fatalf("--config is empty")
	}

	config, err := loadConfig(ops.ConfigPath)
	if err != nil {
		glog.Fatalf("loadConfig: %s error:%s ", ops.ConfigPath, err.Error())
	}

	startApiArgs := &api.StartApiArgs{
		Port: config.Port,
	}
	api.StartApi(startApiArgs)

	<-stopCh
}
