package git

import (
	"os/exec"
	"strings"
)

type Repo struct {
	Path string
}

func NewRepo(path string) (*Repo, error) {
	repo := &Repo{path}
	err := repo.fixRoot()
	return repo, err
}

func (r *Repo) fixRoot() error {
	output, err := r.Command("rev-parse", "--show-toplevel").Output()
	if err != nil {
		return err
	}

	r.Path = strings.TrimRight(string(output), "\r\n")
	return nil
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

func (r *Repo) Describe(matchPattern string) (string, error) {
	cmd := r.Command("describe", "--tags",
		"--match",
		matchPattern,
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
