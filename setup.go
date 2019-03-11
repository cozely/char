package char

import "golang.org/x/sys/unix"

var (
	oldStdin  *state
	oldStdout *state
)

func Setup() error {
	var err error

	oldStdin, err = makeRaw(unix.Stdin)
	if err != nil {
		return err
	}
	oldStdout, err = makeRaw(int(unix.Stdout))
	if err != nil {
		return err
	}

	err = resize()

	return err
}

func Cleanup() {
	unix.Write(unix.Stdout, clear)
	unix.Write(unix.Stdout, home)
	restore(int(unix.Stdout), oldStdout)
	restore(int(unix.Stdin), oldStdin)
}
