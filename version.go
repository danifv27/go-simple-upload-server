package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"runtime"
)

var (
	//embed doesn't allow cross package boundaries, so version.json should be in this folder
	//go:embed version.json
	version string
	//Application Name
	Name string
)

type Version struct {
	versioninfo VersionInfo
}

type VersionInfo struct {
	//GitCommit The git commit that was compiled. This will be filled in by the compiler.
	GitCommit string
	//Version The main version number that is being run at the moment.
	Version  string
	Revision string
	Branch   string
	//BuildDate This will be filled in by the makefile
	BuildDate string
	BuildUser string
	//GoVersion The runtime version
	GoVersion string
	//OsArch The OS architecture
	OsArch string
}

func (v VersionInfo) String() string {

	return fmt.Sprintf("Version:\t%s\nGit commit:\t%s\nBuilt:\t\t%s (from %s by %s)",
		v.Version, v.GitCommit, v.BuildDate, v.Branch, v.BuildUser)
}

func NewVersion() (Version, error) {
	var j map[string]interface{}

	if err := json.Unmarshal([]byte(version), &j); err != nil {
		return Version{}, err
	}

	v := Version{}
	// The value of your map associated with key "git" is of type map[string]interface{}.
	// And we want to access the element of that map associated with the key "commit".
	// .(string) type assertion to convert interface{} to string
	v.versioninfo.GitCommit = j["git"].(map[string]interface{})["commit"].(string)
	v.versioninfo.Branch = j["git"].(map[string]interface{})["branch"].(string)
	v.versioninfo.Version = j["version"].(string)
	v.versioninfo.Revision = j["revision"].(string)
	v.versioninfo.BuildDate = j["build"].(map[string]interface{})["date"].(string)
	v.versioninfo.BuildUser = j["build"].(map[string]interface{})["user"].(string)
	v.versioninfo.GoVersion = runtime.Version()
	v.versioninfo.OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)

	return v, nil
}

// GetVersionInfo Returns the version information
func (m *Version) GetVersionInfo() (VersionInfo, error) {

	return m.versioninfo, nil
}
