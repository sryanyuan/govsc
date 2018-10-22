package command

var (
	// Settings template
	settingsTemplate = []byte(
		`{
	"go.gopath": "{{.GoPath}}",
	"go.buildOnSave": "package"
}`)

	// Build script 2.0.0
	buildScriptWindows200 = []byte(
		`@rem %1->build path | %2->action
@CD %1
@go %2

@if %errorlevel%==0 (echo %2 success) else (echo %2 failed)`)

	// Build script
	buildScriptWindows = []byte(
		`@rem param1 GOPATH
@rem param2 build path
@rem param3 action

@SET GOPATH=%1
@CD %2

@go %3

@rem succeed or failed
@if %errorlevel%==0 (echo %3 success) else (echo %3 failed)`)

	// Deprecated (version 0.1.0)
	// Tasks template
	tasksTemplate = []byte(
		`{
// See https://go.microsoft.com/fwlink/?LinkId=733558
// for the documentation about the tasks.json format
"version": "0.1.0",
"echoCommand": true,
"isShellCommand": true,
"showOutput": "always",
"windows": {
	"command": "cmd",
	"args": [
		"/C"
	]
},
"tasks": [
	{
		"taskName": "build",
		"isBuildCommand": true,
		"suppressTaskName": true,
		"args": [
			"tasks",
			"{{.GoPath}}",
			"${fileDirname}",
			"build"
		],
		"isWatching": false,
		"problemMatcher": {
			"owner": "go",
			"fileLocation": [
				"relative",
				"${fileDirname}"
			],
			"pattern": {
				"regexp": "^(.*):(\\d+):\\s+(.*)$",
				"file":1,
				"line": 2,
				"message": 3
			}
		}
	},
	{
		"taskName": "install",
		"suppressTaskName": true,
		"args": [
			"tasks",
			"{{.GoPath}}",
			"${fileDirname}",
			"install"
		],
		"isWatching": false,
		"problemMatcher": {
			"owner": "go",
			"fileLocation": [
				"relative",
				"${fileDirname}"
			],
			"pattern": {
				"regexp": "^(.*):(\\d+):\\s+(.*)$",
				"file":1,
				"line": 2,
				"message": 3
			}
		}
	}
]
}`)
)
