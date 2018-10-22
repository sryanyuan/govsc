package command

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// NewInit2Command execute the govsc -init command
func NewInit2Command() *cobra.Command {
	// Arguments
	var options initProjectOptions

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize go vscode project",
		Long:  "Initialize go vscode project",
		Run: func(cmd *cobra.Command, args []string) {
			initProject2(cmd, &options)
		},
	}

	initCmd.Flags().StringVarP(&options.projectDir, "project", "p", "", "project directory, eg:github.com/xxx/testproject")

	return initCmd
}

func initProject2(cmd *cobra.Command, options *initProjectOptions) {
	// Load GOPATH env
	gopath := os.Getenv("GOPATH")
	if "" == gopath {
		fmt.Println("GOPATH not set")
		return
	}

	// Output directory must be a relative path, if projectPath is empty
	// .vscode is in $GOPATH/.vscode
	if path.IsAbs(options.projectDir) {
		fmt.Println("output directory must be a relative path relative to GOPATH")
		return
	}

	workingDir := gopath
	if 0 != len(options.projectDir) {
		workingDir = fmt.Sprintf("%s/src/%s", gopath, options.projectDir)
	}
	workingDir = strings.Replace(workingDir, "\\", "/", -1)
	workingDir = strings.Trim(workingDir, "/")
	cmd.Println("Create project in directory", workingDir)

	// Create .vscode directory
	err := os.MkdirAll(fmt.Sprintf("%s/.vscode", workingDir), 0666)
	if nil != err {
		fmt.Println("Create setting directory failed with error:", err)
		return
	}
	// Get gopath relative path
	goPathRelative := "${workspaceRoot}/"
	if 0 != len(options.projectDir) {
		splitterCnt := strings.Count(cleanPath(options.projectDir), "/")
		for i := 0; i < splitterCnt+2; i++ {
			goPathRelative += "../"
		}
	}
	renderContext := map[string]string{
		"GoPath": goPathRelative,
	}
	// Render settings.json
	settingFile, err := os.Create(fmt.Sprintf("%s/.vscode/settings.json", workingDir))
	if nil != err {
		fmt.Println("Create settings file error:", err)
		return
	}
	defer settingFile.Close()
	if err = renderTemplate(settingsTemplate, renderContext, settingFile); nil != err {
		fmt.Println("Render settings.json error:", err)
		return
	}
	// Render tasks.json
	tasksFile, err := os.Create(fmt.Sprintf("%s/.vscode/tasks.json", workingDir))
	if nil != err {
		fmt.Println("Create tasks file error:", err)
		return
	}
	defer tasksFile.Close()
	taskBytes, err := buildConfig(goPathRelative, version200)
	if nil != err {
		fmt.Printf("Get config content error: %v\n", err)
		return
	}
	if _, err = tasksFile.Write(taskBytes); nil != err {
		fmt.Printf("Write config file error: %v\n", err)
		return
	}

	// Output build scripts
	if runtime.GOOS == osWindows {
		buildFile, err := os.Create(fmt.Sprintf("%s/tasks.win.cmd", workingDir))
		if nil != err {
			fmt.Println("Create build script error:", err)
			return
		}
		defer buildFile.Close()
		buildFile.Write(buildScriptWindows200)
	} else {
		fmt.Println("Sorry,", runtime.GOOS, "not support now")
	}

	fmt.Println("Done ...")
}
