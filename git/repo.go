package git

import "os/exec"

type Repo struct {
	Path string
}

func NewRepo(path string) *Repo {
	return &Repo{path}
}

func (r *Repo) AddAll() error {
	return r.Command("add", "--all").Run()
}

func (r *Repo) Commit(message string) error {
	return r.Command("commit", "-m", message).Run()
}

func (r *Repo) Tag(tag, message string) error {
	return r.Command("tag", "-a", tag, "-m", message).Run()
}

func (r *Repo) Init() error {
	return r.Command("init").Run()
}

func (r *Repo) Describe() (string, error) {
	cmd := r.Command("describe", "--all",
		"--match",
		"v*",
		"--abbrev=0",
	)
	output, err := cmd.Output()
	return string(output), err
}

func (r *Repo) Command(args ...string) *exec.Cmd {
	cmd := exec.Command("git", args...)
	cmd.Dir = r.Path
	return cmd
}
