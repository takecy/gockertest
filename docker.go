package gockertest

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// Cli is struct
type Cli struct {
	// ID is container id
	ID string
}

// Arguments is arguments for docker commands
type Arguments struct {
	// Ports is `-p` option
	Ports map[int]int

	// Envs is `-e` option
	Envs map[string]string

	// Name is `--name` option
	Name string

	// Network is `--net` option
	Network string

	// RequireLogin is required login
	RequireLogin bool

	// Login is login info
	Login Login

	// Foreground is add `--rm` option instead of `-d`
	Foreground bool
}

// Login is struct for login registry
type Login struct {
	User     string
	Password string
	Registry string
}

// Run is run images
func Run(image string, args Arguments) *Cli {
	if args.RequireLogin {
		out, err := run("docker", "login", "-u", args.Login.User, "-p", args.Login.Password, args.Login.Registry)
		if err != nil {
			panic(err)
		}
		fmt.Printf("login.result:%+v\n", out)
	}

	if len(args.Ports) == 0 {
		panic(fmt.Errorf("ports.required"))
	}

	// update image
	out, err := run("docker", "pull", image)
	if err != nil {
		panic(err)
	}
	fmt.Printf("update.result:%+v\n", out)

	// ports mapping
	ports := make([]string, 0, len(args.Ports))
	for k, v := range args.Ports {
		ports = append(ports, fmt.Sprintf("%d:%d", k, v))
	}

	// environment variables
	envs := make([]string, 0, len(args.Envs))
	for k, v := range args.Envs {
		envs = append(envs, fmt.Sprintf("%s=%s", k, v))
	}

	// build arguments for [docker] command
	cmdArgs := make([]string, 0, len(ports)+1)
	cmdArgs = append(cmdArgs, "run")

	if args.Foreground {
		cmdArgs = append(cmdArgs, "--rm")
	} else {
		cmdArgs = append(cmdArgs, "-d")
	}

	for _, p := range ports {
		cmdArgs = append(cmdArgs, "-p")
		cmdArgs = append(cmdArgs, p)
	}

	for _, e := range envs {
		cmdArgs = append(cmdArgs, "-e")
		cmdArgs = append(cmdArgs, e)
	}

	if args.Name != "" {
		cmdArgs = append(cmdArgs, "--name", args.Name)
	}

	if args.Network != "" {
		cmdArgs = append(cmdArgs, "--net", args.Network)
	}

	cmdArgs = append(cmdArgs, image)

	fmt.Printf("cmdArgs:%+v\n", cmdArgs)

	// run container
	continerID, err := run("docker", cmdArgs...)
	if err != nil {
		panic(err)
	}

	fmt.Printf("start.container.id:%s\n", continerID)

	return &Cli{
		ID: continerID,
	}
}

// Cleanup is cleanup image
func (cli *Cli) Cleanup() {
	fmt.Printf("cleanup.container.id:%s\n", cli.ID)
	run("docker", "rm", "-f", cli.ID)
}

func run(name string, args ...string) (out string, err error) {
	cmd := exec.Command(name, args...)

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err = cmd.Run(); err != nil {
		return
	}

	if cmd.ProcessState.Success() {
		return strings.TrimSpace(stdout.String()), nil
	}

	err = errors.New("command execution failed " + stderr.String())
	return
}
