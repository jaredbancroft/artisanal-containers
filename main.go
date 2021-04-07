package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic("Unrecognized Command")
	}
}

func parent() {
	fmt.Println("Running ", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Println("Running ", os.Args[2:])

	cg()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("container")))
	must(syscall.Chroot("/workspaces/artisanal-containers/filesystem"))
	must(syscall.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(syscall.Mount("tmp", "tmp", "tmpfs", 0, ""))
	must(cmd.Run())

	must(syscall.Unmount("proc", 0))
	must(syscall.Unmount("tmp", 0))
}

func cg() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	os.Mkdir(filepath.Join(pids, "basec"), 0755)

	must(ioutil.WriteFile(filepath.Join(pids, "basec/pids.max"), []byte("20"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, "basec/notify_on_release"), []byte("1"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, "basec/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
