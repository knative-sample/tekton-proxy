package options

import (
	"github.com/spf13/cobra"
)

type Options struct {
	ConfigPath string
	Version    bool
}

func (s *Options) SetOps(ac *cobra.Command) {
	ac.Flags().StringVar(&s.ConfigPath, "config", s.ConfigPath, "config file path")
	ac.Flags().BoolVar(&s.Version, "version", s.Version, "Print version information")
}
