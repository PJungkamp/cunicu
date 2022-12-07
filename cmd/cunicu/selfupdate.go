package main

// derived from http://github.com/restic/restic

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stv0g/cunicu/pkg/selfupdate"
	"go.uber.org/zap"
)

type selfUpdateOptions struct {
	output string
}

func init() { //nolint:gochecknoinits
	opts := &selfUpdateOptions{}
	cmd := &cobra.Command{
		Use:   "selfupdate",
		Short: "Update the cunīcu binary",
		Long: `Downloads the latest stable release of cunīcu from GitHub and replaces the currently running binary.
After download, the authenticity of the binary is verified using the GPG signature on the release files.`,
		Run: func(cmd *cobra.Command, args []string) {
			selfUpdate(cmd, args, opts)
		},
	}

	rootCmd.AddCommand(cmd)

	selfPath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	self := filepath.Base(selfPath)
	if strings.Contains(selfPath, "go-build") {
		self = "cunicu"
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.output, "output", "o", self, "Save the downloaded file as `filename`")

	if err := cmd.MarkFlagFilename("output"); err != nil {
		panic(err)
	}
}

func selfUpdate(_ *cobra.Command, _ []string, opts *selfUpdateOptions) {
	logger := logger.Named("self-update")

	rel, err := selfupdate.SelfUpdate(opts.output, logger)
	if err != nil {
		logger.Fatal("Self-update failed", zap.Error(err))
	}

	logger.Info("Successfully updated cunicu",
		zap.String("version", rel.Version),
		zap.String("filename", opts.output),
	)
}
