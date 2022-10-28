package piston

type ExecutionTask struct {
	Language           string    `json:"language"`
	Version            string    `json:"version"`
	Files              []JobFile `json:"files"`
	Stdin              string    `json:"stdin"`
	Args               []string  `json:"args"`
	RunTimeout         int       `json:"run_timeout,omitempty"`
	CompileTimeout     int       `json:"compile_timeout,omitempty"`
	RunMemoryLimit     int       `json:"run_memory_limit,omitempty"`
	CompileMemoryLimit int       `json:"compile_memory_limit,omitempty"`
}

type ExecutionResult struct {
	Run      JobOutput `json:"run"`
	Language string    `json:"language"`
	Version  string    `json:"version"`
}

type JobOutput struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Code   *int   `json:"code"`
	Signal *int   `json:"signal"`
	Output string `json:"output"`
}

type JobFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
