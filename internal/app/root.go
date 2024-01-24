package app

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"

	"github.com/bool64/ctxd"
	"github.com/bool64/zapctxd"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var rootCfg = appConfig{}

func rootCommand() *cobra.Command {
	defer func() {
		if e := recover(); e != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%s\n\n%s",
				color.HiRedString(fmt.Sprintf("%s", e)),
				debug.Stack(),
			)
		}
	}()

	logger := ctxd.Logger(ctxd.NoOpLogger{})

	cmd := ligaturizeCommand(&logger)

	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	cmd.PersistentPreRun = func(cmd *cobra.Command, _ []string) {
		logger = makeLogger(cmd.ErrOrStderr())
	}

	cmd.PersistentFlags().BoolVarP(&rootCfg.Debug, "debug", "d", rootCfg.Debug, "debug output")
	cmd.PersistentFlags().BoolVarP(&rootCfg.Verbose, "verbose", "v", rootCfg.Verbose, "verbose output")

	cmd.AddCommand(
		versionCommand(),
	)

	return cmd
}

type appConfig struct {
	Verbose    bool
	Debug      bool
	ConfigFile string
}

func logLevel() zapcore.Level {
	if rootCfg.Debug {
		return zap.DebugLevel
	}

	if rootCfg.Verbose {
		return zap.InfoLevel
	}

	return zap.WarnLevel
}

func makeLogger(w io.Writer) ctxd.Logger {
	logCfg := zapctxd.Config{
		DevMode: true,
		Level:   logLevel(),
		Output:  io.Discard,
	}

	if rootCfg.Verbose || rootCfg.Debug {
		logCfg.Output = w
		logCfg.ColoredOutput = true
	}

	return zapctxd.New(logCfg)
}
