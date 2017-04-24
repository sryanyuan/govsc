# govsc
Tool for generate vscode go project

# Usage

govsc init -p [project-dir]

eg:

* govsc init -p github.com/sryanyuan/test

This will create a project named test in $GOPATH/src/github.com/sryanyuan/test directory.

If you want to use $GOPATH as project root path, it is also supported.

Then you can drag your project directory to visual studio code.

To build your project, open your .go source file has main function and press ctrl+shift+b, then the binary output file will be generated in your project path.

To install your project, press ctrl+shift+p, and press backspace, enter task ,select install command, then press enter.

You can extend your own command in .vscode/tasks.json to meet your own needs.