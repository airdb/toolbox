package oskit

import (
	"encoding/json"
	"os"
	"runtime"
	"time"

	"github.com/airdb/toolbox/oskit"
	"github.com/airdb/toolbox/utilkit"
)

// Build version info.
type BuildInfo struct {
	GoVersion string
	Env       string
	Repo      string
	Version   string
	Swagger   string `json:"swagger,omitempty"`
	Build     string
	BuildTime string
	Uptime    string
	CreatedAt time.Time
}

var (
	Env       string
	Repo      string
	Version   string
	Swagger   string
	Build     string
	BuildTime string
	Uptime    string
	CreatedAt time.Time
)

func Init() {
	// Init the loc.
	loc, _ := time.LoadLocation("Asia/Shanghai")

	// Set timezone.
	CreatedAt = time.Now().In(loc)

	Env = os.Getenv("ENV")
	if Env == "" {
		Env = oskit.DeployStageDev
	}
}

func GetBuildInfo() *BuildInfo {
	return &BuildInfo{
		GoVersion: runtime.Version(),
		Env:       oskit.GetDeployStage(),
		Repo:      Repo,
		Version:   Version,
		Swagger:   Swagger,
		Build:     Build,
		BuildTime: BuildTime,
		CreatedAt: CreatedAt,
		Uptime:    time.Since(CreatedAt).String(),
	}
}

func (info *BuildInfo) ToString() string {
	out, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func (info *BuildInfo) ToProject() string {
	return utilkit.LastStringWithSplit(info.Repo, "/")
}
