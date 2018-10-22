package command

import (
	"html/template"
	"log"
	"os"
	"runtime"

	"fmt"
	"path"

	"strings"

	"io"

	"github.com/spf13/cobra"
)

const (
	osWindows = "windows"
	osLinux   = "linux"
)

type initProjectOptions struct {
	projectDir string
}

type vscodeGoSetting struct {
	GoPath string `json:"go.gopath"`
}

func renderTemplate(templateData []byte, ctx map[string]string, writer io.Writer) error {
	tpl, err := template.New("tpl").Parse(string(templateData))
	if nil != err {
		return err
	}

	return tpl.Execute(writer, ctx)
}

func cleanPath(dir string) string {
	dir = strings.Replace(dir, "\\", "/", -1)
	dir = strings.Trim(dir, "/")
	dirRuneBytes := []rune(dir)
	finalDirRuneBytes := make([]rune, 0, len(dirRuneBytes))
	finalDirRuneBytes = append(finalDirRuneBytes, dirRuneBytes[0])
	for i := 1; i < len(dir); i++ {
		if dirRuneBytes[i] == dirRuneBytes[i-1] &&
			dirRuneBytes[i] == '/' {
			// Do not append
		} else {
			finalDirRuneBytes = append(finalDirRuneBytes, dirRuneBytes[i])
		}
	}
	return string(finalDirRuneBytes)
}

func initProject(cmd *cobra.Command, options *initProjectOptions) {
	// Load GOPATH env
	gopath := os.Getenv("GOPATH")
	if "" == gopath {
		log.Println("GOPATH not set")
		return
	}

	// Output directory must be a relative path, if projectPath is empty
	// .vscode is in $GOPATH/.vscode
	if path.IsAbs(options.projectDir) {
		log.Println("output directory must be a relative path relative to GOPATH")
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
		log.Println("Create setting directory failed with error:", err)
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
		log.Println("Create settings file error:", err)
		return
	}
	defer settingFile.Close()
	if err = renderTemplate(settingsTemplate, renderContext, settingFile); nil != err {
		log.Println("Render settings.json error:", err)
		return
	}
	// Render tasks.json
	tasksFile, err := os.Create(fmt.Sprintf("%s/.vscode/tasks.json", workingDir))
	if nil != err {
		log.Println("Create tasks file error:", err)
		return
	}
	defer tasksFile.Close()
	if err = renderTemplate(tasksTemplate, renderContext, tasksFile); nil != err {
		log.Println("Render tasks.json error:", err)
		return
	}

	// Output build scripts
	if runtime.GOOS == osWindows {
		buildFile, err := os.Create(fmt.Sprintf("%s/tasks.cmd", workingDir))
		if nil != err {
			log.Println("Create build script error:", err)
			return
		}
		defer buildFile.Close()
		buildFile.Write(buildScriptWindows)
	} else {
		log.Println("Sorry,", runtime.GOOS, "not support now")
	}

	log.Println("Done ...")
}

// NewInitCommand execute the govsc -init command
func NewInitCommand() *cobra.Command {
	// Arguments
	var options initProjectOptions

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize go vscode project",
		Long:  "Initialize go vscode project",
		Run: func(cmd *cobra.Command, args []string) {
			initProject(cmd, &options)
		},
	}

	initCmd.Flags().StringVarP(&options.projectDir, "project", "p", "", "project directory, eg:github.com/xxx/testproject")

	return initCmd
}
