package command

const (
	taskVersion = "2.0.0"
)

type BaseTaskConfiguration struct {
	/**
	 * The type of a custom task. Tasks of type "shell" are executed
	 * inside a shell (e.g. bash, cmd, powershell, ...)
	 */
	Type string `json:"type"`
	/**
	 * The command to be executed. Can be an external program or a shell
	 * command.
	 */
	Command string `json:"command"`
	/**
	 * Specifies whether a global command is a background task.
	 */
	IsBackground bool `json:"isBackground"`
	/**
	 * The command options used when the command is executed. Can be omitted.
	 */
	Options CommandOptions `json:"options"`
	/**
	 * The arguments passed to the command. Can be omitted.
	 */
	Args []string `json:"args"`
	/**
	 * The presentation options.
	 */
	Presentation PresentationOptions
	/**
	 * The problem matcher to be used if a global command is executed (e.g. no tasks
	 * are defined). A tasks.json file can either contain a global problemMatcher
	 * property or a tasks property but not both.
	 */
	ProblemMatcher string `json:"problemMatcher"`
	/**
	 * The configuration of the available tasks. A tasks.json file can either
	 * contain a global problemMatcher property or a tasks property but not both.
	 */
	Tasks []TaskDescription `json:"tasks"`
}

type TaskConfiguration struct {
	//BaseTaskConfiguration
	/**
	 * The configuration's version number
	 */
	Version string `json:"version"`
	/**
	 * The command options used when the command is executed. Can be omitted.
	 */
	Options CommandOptions `json:"options"`
	/**
	 * The configuration of the available tasks. A tasks.json file can either
	 * contain a global problemMatcher property or a tasks property but not both.
	 */
	Tasks []TaskDescription `json:"tasks"`
}

type CommandOptions struct {
	/**
	 * The current working directory of the executed program or shell.
	 * If omitted the current workspace's root is used.
	 */
	Cwd string `json:"-"`
	/**
	 * The environment of the executed program or shell. If omitted
	 * the parent process' environment is used.
	 */
	Env map[string]string `json:"env"`
}

type OsCommandOptions struct {
	Args    []string       `json:"args"`
	Command string         `json:"command"`
	Options CommandOptions `json:"-"`
}

type TaskDescription struct {
	/**
	 * The task's name
	 */
	Label string `json:"label"`
	/**
	 * The type of a custom task. Tasks of type "shell" are executed
	 * inside a shell (e.g. bash, cmd, powershell, ...)
	 */
	Type string `json:"type"`
	/**
	 * The command to execute. If the type is "shell" it should be the full
	 * command line including any additional arguments passed to the command.
	 */
	Command string `json:"command"`
	/**
	 * Additional arguments passed to the command. Should be used if type
	 * is "process".
	 */
	Args []string `json:"-"`
	/**
	 * Whether the executed command is kept alive and runs in the background.
	 */
	IsBackground bool `json:"isBackground"`
	/**
	 * Defines the group to which this task belongs. Also supports to mark
	 * a task as the default task in a group.
	 */
	Group string `json:"group"`
	/**
	 * The presentation options.
	 */
	Presentation PresentationOptions `json:"presentation"`
	/**
	 * The problem matcher(s) to use to capture problems in the tasks
	 * output.
	 */
	ProblemMatcher string `json:"problemMatcher"`

	// OS command
	Windows OsCommandOptions `json:"windows"`
}

type PresentationOptions struct {
	/**
	 * Controls whether the task output is reveal in the user interface.
	 * Defaults to `always`.
	 */
	Reveal string `json:"reveal"`
	/**
	 * Controls whether the command associated with the task is echoed
	 * in the user interface.
	 */
	Echo bool `json:"echo"`
	/**
	 * Controls whether the panel showing the task output is taking focus.
	 */
	Focus bool `json:"focus"`
	/**
	 * Controls if the task panel is used for this task only (dedicated),
	 * shared between tasks (shared) or if a new panel is created on
	 * every task execution (new). Defaults to `shared`
	 */
	Panel string `json:"panel"`

	ShowReuseMessage bool `json:"showReuseMessage"`
}
