package cmd

import (
	"flag"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

var version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:   "celery",
	Short: "CEL expression evaluator for Kubernetes YAML",
	Long: `Celery evaluates Common Expression Language (CEL) expressions against 
Kubernetes Resource Model (KRM) YAML files.

Use it to validate resources with custom rules, check compliance with policies,
or query resource properties using CEL's expression language.`,
	Version: version,
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}

func init() {
	klogFlags := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlags)
	rootCmd.PersistentFlags().AddGoFlag(klogFlags.Lookup("v"))
}
