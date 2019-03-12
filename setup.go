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
	unix.Write(unix.Stdout, escSmcup)

	err = resize()
	//unix.Write(unix.Stdout, cursorShape(6))

	return err
}

func Cleanup() {
	unix.Write(unix.Stdout, escClear)
	unix.Write(unix.Stdout, escRmcup)
	restore(int(unix.Stdout), oldStdout)
	restore(int(unix.Stdin), oldStdin)
}
