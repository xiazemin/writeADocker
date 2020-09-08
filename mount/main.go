package main
import(
	"syscall"
	"os"
	"fmt"
	"os/exec"
)

func main(){
	switch os.Args[1]{
	case "run":
		run()
	case "child":
		child()
	default:
		panic("bad command")
	}
}

//func run(){
//	fmt.Printf("running %v\n",os.Args[2:])
//	cmd:=exec.Command(os.Args[2],os.Args[3:]...)
//	cmd.Stdin=os.Stdin
//	cmd.Stdout=os.Stdout
//	cmd.Stderr=os.Stderr
//	cmd.Run()
//}

func run()  {
	fmt.Printf("running %v\n",os.Args[2:])
	cmd:=exec.Command("/proc/self/exe",append([]string{"child"},os.Args[2:]...)...)
	cmd.Stdin=os.Stdin
	cmd.Stdout=os.Stdout
	cmd.Stderr=os.Stderr
	cmd.SysProcAttr=&syscall.SysProcAttr{
		Cloneflags:syscall.CLONE_NEWUTS|syscall.CLONE_NEWPID,
	}

	cmd.Run()
}

func child(){
	fmt.Println("Runing %v\n",os.Args[2:])
	syscall.Sethostname([]byte("container"))
	syscall.Chroot("/vagrant/ubuntu-fs")
	syscall.Chdir("/")
	syscall.Mount("proc","proc","proc",0,"")

	cmd:=exec.Command(os.Args[2],os.Args[3:]...)
	cmd.Stdin=os.Stdin
	cmd.Stdout=os.Stdout
	cmd.Stderr=os.Stderr
	cmd.SysProcAttr=&syscall.SysProcAttr{
		Cloneflags:syscall.CLONE_NEWUTS,
	}
	syscall.Sethostname([]byte("container"))
	cmd.Run()
	syscall.Unmount("/proc",0)
}
