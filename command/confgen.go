package command

import (
	"encoding/json"
	"fmt"
)

const (
	version200 = "2.0.0"
)

func buildConfig(gopath string, version string) ([]byte, error) {
	switch version {
	case version200:
		{
			return buildConfig200(gopath)
		}
	default:
		{
			return nil, fmt.Errorf("Unsupport config version: %s", version)
		}
	}
}

func getTaskDescription(label string, args []string) TaskDescription {
	return TaskDescription{
		Label:        label,
		Group:        "build",
		Type:         "shell",
		IsBackground: true,
		Presentation: PresentationOptions{
			Echo:             true,
			Reveal:           "always",
			Focus:            false,
			Panel:            "shared",
			ShowReuseMessage: true,
		},
		ProblemMatcher: "$go",
		// Windows command
		Windows: OsCommandOptions{
			Command: "tasks.win",
			Args:    args,
		},
	}
}

func buildConfig200(gopath string) ([]byte, error) {
	var config TaskConfiguration
	config.Version = version200
	config.Options.Env = map[string]string{
		"GOPATH": gopath,
	}
	config.Tasks = make([]TaskDescription, 0, 2)
	// Build
	config.Tasks = append(config.Tasks, getTaskDescription("build", []string{"${fileDirname}", "build"}))
	// Install
	config.Tasks = append(config.Tasks, getTaskDescription("install", []string{"${fileDirname}", "install"}))

	jsonBytes, err := json.Marshal(&config)
	if nil != err {
		return nil, fmt.Errorf("Marshal json config error: %v", err)
	}
	return jsonBytes, nil
}
