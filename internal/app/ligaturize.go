package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bool64/ctxd"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"go.nhat.io/fontforge"

	"go.nhat.io/ligaturizer/internal/ligaturizer"
	firacodev3 "go.nhat.io/ligaturizer/internal/ligaturizer/firacode/v3"
	"go.nhat.io/ligaturizer/internal/version"
)

const (
	extOTF = ".otf"
	extTTF = ".ttf"
)

var ligaturizerCfg = ligaturizerConfig{}

var (
	errInputFontFileNotSpecified          = errors.New("input font file must be specified")
	errLigatureFontFileAndDirNotSpecified = errors.New("either ligature font file or ligature font dir must be specified")
	errLigatureFontDirNotFound            = errors.New("ligature font dir not found")
	errLigatureFontUnsupported            = errors.New("unsupported ligature font")
	errLigatureFontVersionNotFound        = errors.New("could not find version of ligature font")
	errLigatureFontNotFound               = errors.New("ligature font not found")
)

func init() { //nolint: gochecknoinits
	if err := loadLigatureConfig(); err != nil {
		panic(err)
	}
}

func ligaturizeCommand(logger *ctxd.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "ligaturizer [flags] input-font-file",
		Short:        "Ligaturize a font",
		SilenceUsage: true,
		Args: func(_ *cobra.Command, args []string) error {
			if l := len(args); l == 0 {
				return errInputFontFileNotSpecified
			} else if l > 1 {
				return fmt.Errorf("accepts only 1 arg, received %d", l) //nolint: goerr113
			}

			return nil
		},
		PreRunE: func(*cobra.Command, []string) error {
			if ligaturizerCfg.LigatureFontFile == "" && ligaturizerCfg.LigatureFontDir == "" {
				return errLigatureFontFileAndDirNotSpecified
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ligaturizerCfg.InputFontFile = args[0]

			return runLigaturize(cmd.Context(), ligaturizerCfg, *logger)
		},
	}

	cmd.Flags().SortFlags = false

	cmd.Flags().StringVarP(&ligaturizerCfg.LigatureFontFile, "ligature-font-file", "L", ligaturizerCfg.LigatureFontFile,
		`the file to copy ligatures from.
if unspecified, ligaturize will attempt to pick a suitable one from --ligature-font-dir based on the input font's weight`,
	)
	cmd.Flags().StringVarP(&ligaturizerCfg.LigatureFontDir, "ligature-font-dir", "D", ligaturizerCfg.LigatureFontDir,
		`the dir to search for ligature fonts when --ligature-font-file is unspecified.`,
	)
	cmd.Flags().StringVarP(&ligaturizerCfg.OutputDir, "output-dir", "O", ligaturizerCfg.OutputDir,
		`the directory to save the ligaturized font in.
the actual filename will be automatically generated based on the input font name and the --prefix and --output-name flags`,
	)
	cmd.Flags().StringVarP(&ligaturizerCfg.OutputName, "output-name", "N", ligaturizerCfg.OutputName,
		`name of the generated font.
completely replaces the original.`,
	)
	cmd.Flags().StringVarP(&ligaturizerCfg.OutputNamePrefix, "prefix", "P", ligaturizerCfg.OutputNamePrefix,
		`string to prefix the name of the generated font with.`,
	)
	cmd.Flags().BoolVar(&ligaturizerCfg.CopyCharacterGlyphs, "copy-character-glyphs", ligaturizerCfg.CopyCharacterGlyphs,
		`copy glyphs for (some) individual characters from the ligature font as well.
this will result in punctuation that matches the ligatures more closely, but may not fit in as well with the rest of the font`,
	)
	cmd.Flags().Float64Var(&ligaturizerCfg.ScaleCharacterGlyphThreshold, "scale-character-glyph-threshold", ligaturizerCfg.ScaleCharacterGlyphThreshold,
		`when copying character glyphs, if they differ in width from the width of the input font by at least this much, scale them horizontally to match the input font even if this noticeably changes their aspect ratio. the default (0.1) means to scale if they are at least 10%% wider or narrower. a value of 0 will scale all copied character glyphs; a value of 2 effectively disables character glyph scaling`,
	)
	cmd.Flags().StringVar(&ligaturizerCfg.BuildID, "build-id", ligaturizerCfg.BuildID, "build id to be added to the font version")

	cmd.MarkFlagFilename("ligature-font-file", "otf", "ttf") //nolint: errcheck,gosec
	cmd.MarkFlagDirname("ligature-font-dir")                 //nolint: errcheck,gosec

	return cmd
}

type ligaturizerConfig struct {
	LigatureFontFile             string  `envconfig:"LIGATURE_FONT_FILE"`
	LigatureFontDir              string  `envconfig:"LIGATURE_FONT_DIR"`
	InputFontFile                string  `envconfig:"INPUT_FONT_FILE"`
	OutputDir                    string  `envconfig:"OUTPUT_DIR"`
	OutputName                   string  `envconfig:"OUTPUT_NAME"`
	OutputNamePrefix             string  `envconfig:"OUTPUT_NAME_PREFIX" default:""`
	CopyCharacterGlyphs          bool    `envconfig:"COPY_CHARACTER_GLYPHS" default:"false"`
	ScaleCharacterGlyphThreshold float64 `envconfig:"SCALE_CHARACTER_GLYPH_THRESHOLD" default:"0.1"`
	BuildID                      string  `envconfig:"BUILD_ID"`
}

func loadLigatureConfig() error {
	err := envconfig.Process("", &ligaturizerCfg)
	if err != nil {
		return fmt.Errorf("failed to load ligature config: %w", err)
	}

	if ligaturizerCfg.LigatureFontDir == "" || ligaturizerCfg.OutputDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current working directory: %w", err)
		}

		if ligaturizerCfg.LigatureFontDir == "" {
			ligaturizerCfg.LigatureFontDir = cwd
		}

		if ligaturizerCfg.OutputDir == "" {
			ligaturizerCfg.OutputDir = cwd
		}
	}

	return nil
}

func runLigaturize(ctx context.Context, cfg ligaturizerConfig, logger ctxd.Logger) error { //nolint: funlen
	// Prepare.
	inputFont, err := fontforge.Open(cfg.InputFontFile)
	if err != nil {
		return fmt.Errorf("failed to open input font: %w", err)
	}

	defer inputFont.Close() //nolint: errcheck

	logger.Debug(ctx, "loaded input font", "file", cfg.InputFontFile)

	ligFontFile, err := getLigatureFontFile(ctx, cfg, inputFont.FontName(), logger)
	if err != nil {
		return err
	}

	ligFont, err := fontforge.Open(ligFontFile)
	if err != nil {
		return fmt.Errorf("failed to open ligature font: %w", err)
	}

	defer ligFont.Close() //nolint: errcheck

	logger.Debug(ctx, "loaded ligature font", "file", ligFontFile)

	if cfg.OutputName == "" {
		cfg.OutputName = inputFont.FamilyName()
	}

	if cfg.OutputNamePrefix != "" {
		cfg.OutputName = fmt.Sprintf("%s %s", cfg.OutputNamePrefix, cfg.OutputName)
	}

	l, err := makeLigaturizer(ligFont, logger)
	if err != nil {
		return err
	}

	logger.Info(ctx, "ligaturize font", "input_font", cfg.InputFontFile, "ligature_font", ligFontFile)

	// Ligaturize.
	if err := ligaturizer.Ligaturize(ctx, l, ligFont, inputFont, cfg.OutputName); err != nil {
		return err //nolint: wrapcheck
	}

	// Footprint.
	updateCopyright(ligFont, inputFont)
	updateVersion(inputFont, cfg.BuildID)

	// Output.
	outputType := extTTF

	ext := strings.ToLower(filepath.Ext(inputFont.Path()))
	if ext == extOTF {
		outputType = extOTF
	}

	outputFile := filepath.Join(cfg.OutputDir, fmt.Sprintf("%s%s", inputFont.FontName(), outputType))

	logger.Info(ctx, "generate ligaturized font", "file", outputFile)

	if err = fontforge.Generate(inputFont, outputFile); err != nil {
		return fmt.Errorf("failed to generate font %s: %w", outputFile, err)
	}

	return nil
}

func getLigatureFontFile(ctx context.Context, cfg ligaturizerConfig, inputFontName string, logger ctxd.Logger) (string, error) {
	if cfg.LigatureFontFile != "" {
		return cfg.LigatureFontFile, nil
	}

	fs := afero.NewOsFs()
	if ok, err := afero.DirExists(fs, cfg.LigatureFontDir); err != nil {
		return "", fmt.Errorf("failed to check ligature font dir: %w", err)
	} else if !ok {
		return "", fmt.Errorf("%w: %s", errLigatureFontDirNotFound, cfg.LigatureFontDir)
	}

	fileName := filepath.Join(cfg.LigatureFontDir, fmt.Sprintf("FiraCode%s", getFontWeight(inputFontName)))
	extensions := []string{extOTF, extTTF}

	for _, ext := range extensions {
		file := fmt.Sprintf("%s%s", fileName, ext)

		if ok, err := afero.Exists(fs, file); err != nil {
			return "", fmt.Errorf("failed to check ligature font file: %w", err)
		} else if ok {
			logger.Debug(ctx, "found ligature font in dir", "file", file)

			return file, nil
		}
	}

	logger.Error(ctx, "no ligature fonts found", "dir", cfg.LigatureFontDir)

	return "", errLigatureFontNotFound
}

func getFontWeight(fontName string) string {
	fontName = strings.ToLower(fontName)

	for _, weight := range []string{"-Bold", "-Retina", "-Medium", "-Regular", "-Light"} {
		if strings.HasSuffix(fontName, strings.ToLower(weight)) {
			return weight
		}
	}

	if strings.Contains(fontName, "bold") || strings.Contains(fontName, "heavy") {
		return "-Bold"
	}

	return "-Regular"
}

func makeLigaturizer(ligFont *fontforge.Font, logger ctxd.Logger) (ligaturizer.Ligaturizer, error) {
	fontName, _, _ := strings.Cut(ligFont.FontName(), "-")

	if fontName != "FiraCode" {
		return nil, fmt.Errorf("%w: %s", errLigatureFontUnsupported, fontName)
	}

	v := ligFont.Version()
	if v == nil {
		return nil, errLigatureFontVersionNotFound
	}

	if v.Major() != 3 {
		return nil, fmt.Errorf("%w: %s %s", errLigatureFontUnsupported, fontName, v.String())
	}

	return firacodev3.NewLigaturizer(firacodev3.Config{}, logger)
}

const (
	copyrightTool        = `Ligaturized by Ligaturizer %s (https://github.com/nhatthm/ligaturizer) using %s %s`
	copyrightInspiration = `Inspired by ToxicFrog's Ligaturizer (https://github.com/ToxicFrog/Ligaturizer)`
)

func updateCopyright(src, dst *fontforge.Font) {
	srcVersion := ""
	srcFontName, _, _ := strings.Cut(src.FontName(), "-")

	if v := src.Version(); v != nil {
		srcVersion = v.String()
	}

	toolCopyright := fmt.Sprintf(copyrightTool, version.Info().Version, srcFontName, srcVersion)
	fontCopyright := dst.Copyright()
	fontSFNTCopyright := dst.SFNTNames().Find("Copyright")

	dst.SetCopyright(formatCopyright(fontCopyright, toolCopyright))
	dst.SetSFNTNames("Copyright", formatCopyright(fontSFNTCopyright, toolCopyright))
}

func updateVersion(f *fontforge.Font, buildID string) {
	ov := f.Version()
	if ov == nil {
		return
	}

	m := buildID
	if cm := ov.Metadata(); cm != "" {
		m = fmt.Sprintf("%s.%s", cm, m)
	}

	v, _ := ov.SetMetadata(m) //nolint: errcheck

	f.SetVersion(v)
	f.SetSFNTNames("Version", fmt.Sprintf("Version %s", v.String()))
}

func formatCopyright(copyright, toolCopyright string) string {
	copyright = fmt.Sprintf("%s\n%s\n%s", copyright, toolCopyright, copyrightInspiration)
	copyright = strings.Trim(copyright, "\n")

	return copyright
}
