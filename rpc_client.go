package main

import (
	"os/exec"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Create struct {
	Cpu  int
	Mem  int
	Disk int
}

type Arith int

func (t *Arith) Add(args *Create, reply *int) error {
	*reply = args.Cpu + args.Mem + args.Disk
	return nil
}

func (t *Arith) Create(args *Create, reply *bool) error {
	fmt.Println(args.Cpu)
	fmt.Println(args.Mem)
	fmt.Println(args.Disk)
//	input := fmt.Sprintf("/bin/fgrep 'vcpu' /etc/libvirt/qemu/%s.xml | awk '{split($0,a,\">|<\");print a[3]}'", vm)
//	cmd := exec.Command("/bin/bash", "-c", input)
	input := fmt.Sprintf("/tmp/create.sh %d %d %d",args.Cpu,args.Mem,args.Disk)
	fmt.Println(input);
	cmd := exec.Command("/bin/bash", "-c", input)
//	cmd := exec.Command("/bin/bash","-c", "/tmp/create.sh" args.Cpu args.Mem args.Disk)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}else{
//	fmt.Println(string(out))
		*reply = true
	}
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":5978")
	if e != nil {
		fmt.Println("listen error", e.Error())
	}
	go http.Serve(l, nil)
	time.Sleep(500 * time.Second)
}
