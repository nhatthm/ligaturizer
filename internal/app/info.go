package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/Masterminds/semver/v3"
	"github.com/bool64/ctxd"
	"github.com/spf13/cobra"

	"go.nhat.io/ligaturizer/internal/fontforge"
)

func infoCommand(logger *ctxd.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Show font information",
		Long:  "Show font information",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInfo(cmd.OutOrStdout(), args[0], *logger)
		},
	}

	return cmd
}

func runInfo(stdout io.Writer, fontFile string, logger ctxd.Logger) error {
	// Prepare.
	font, err := fontforge.Open(fontFile)
	if err != nil {
		return fmt.Errorf("failed to open font: %w", err)
	}

	defer font.Close() //nolint: errcheck

	ctx := ctxd.AddFields(context.Background(), "font", fontFile)

	logger.Debug(ctx, "loaded input font")

	i := fontInfo{
		FontName:   font.FontName(),
		FullName:   font.FullName(),
		FamilyName: font.FamilyName(),
		Version:    font.Version(),
	}

	if err := json.NewEncoder(stdout).Encode(i); err != nil {
		return fmt.Errorf("failed to encode font info: %w", err)
	}

	return nil
}

type fontInfo struct {
	FontName   string          `json:"font_name"`
	FullName   string          `json:"full_name"`
	FamilyName string          `json:"family_name"`
	Version    *semver.Version `json:"version"`
}
