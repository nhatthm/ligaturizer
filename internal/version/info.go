package version

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strconv"
)

const (
	buildSettingsVCSRevision    = "vcs.revision"
	buildSettingsVCSModified    = "vcs.modified"
	buildSettingsDefaultVersion = "dev-0.0.0"
	buildSettingsDevelVersion   = "(devel)"
)

// Build information. Populated at build-time.
//
//nolint:gochecknoglobals
var (
	appName        = "program"
	repositoryName = ""
	version        = buildSettingsDefaultVersion
	isDirty        = false
	revision       string
	branch         string
	buildUser      string
	buildDate      string
	dependencies   []*debug.Module
)

// Information holds app version info.
type Information struct {
	AppName        string
	RepositoryName string
	Version        string
	Revision       Revision
	Branch         string
	BuildUser      string
	BuildDate      string
	GoVersion      string
	GoOS           string
	GoArch         string
	Dependencies   []*debug.Module
}

// Revision returns the current revision of the code.
type Revision struct {
	ID    string
	Dirty bool
}

// Short returns the short revision string.
func (r Revision) Short() string {
	if len(r.ID) == 0 {
		return ""
	}

	if r.Dirty {
		return fmt.Sprintf("%s-dirty", r.ID[:8])
	}

	return r.ID[:8]
}

// String returns the revision string.
func (r Revision) String() string {
	if r.Dirty {
		return fmt.Sprintf("%s-dirty", r.ID)
	}

	return r.ID
}

// Info returns app version info.
func Info() Information {
	return Information{
		AppName:        appName,
		RepositoryName: repositoryName,
		Version:        version,
		Revision:       Revision{ID: revision, Dirty: isDirty},
		Branch:         branch,
		BuildUser:      buildUser,
		BuildDate:      buildDate,
		GoVersion:      runtime.Version(),
		GoOS:           runtime.GOOS,
		GoArch:         runtime.GOARCH,
		Dependencies:   dependencies,
	}
}

//nolint:gochecknoinits
func init() {
	if repositoryName == "" {
		repositoryName = appName
	}

	if info, available := debug.ReadBuildInfo(); available {
		dependencies = info.Deps

		if len(version) == 0 && info.Main.Version != buildSettingsDevelVersion {
			version = info.Main.Version
		}

		if len(revision) == 0 {
			for _, setting := range info.Settings {
				switch setting.Key {
				case buildSettingsVCSRevision:
					revision = setting.Value

				case buildSettingsVCSModified:
					isDirty, _ = strconv.ParseBool(setting.Value) //nolint: errcheck
				}
			}
		}
	}
}
