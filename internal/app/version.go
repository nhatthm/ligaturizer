package app

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"go.nhat.io/ligaturizer/internal/version"
)

func versionCommand() *cobra.Command {
	var showFull bool

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show version",
		Long:  "Show version information",
		Run: func(cmd *cobra.Command, _ []string) {
			runVersion(cmd.OutOrStdout(), cmd.OutOrStderr(), showFull)
		},
	}

	cmd.Flags().BoolVarP(&showFull, "full", "f", false, "Show full information")

	return cmd
}

func runVersion(stdout, stderr io.Writer, showFull bool) {
	info := version.Info()

	rev := info.Revision.String()
	if !showFull {
		rev = info.Revision.Short()
	}

	_, _ = fmt.Fprintf(stdout, "%s (rev: %s; %s; %s/%s)\n",
		info.Version,
		rev,
		info.GoVersion,
		info.GoOS, info.GoArch,
	)

	if !showFull {
		return
	}

	_, _ = fmt.Fprintln(stderr)
	_, _ = fmt.Fprintf(stderr, "build user: %s\n", info.BuildUser)
	_, _ = fmt.Fprintf(stderr, "build date: %s\n", info.BuildDate)
	_, _ = fmt.Fprintln(stderr)
	_, _ = fmt.Fprintln(stderr, "dependencies:")

	for _, dep := range info.Dependencies {
		_, _ = fmt.Fprintf(stderr, "  %s: %s\n", dep.Path, dep.Version)
	}
}
