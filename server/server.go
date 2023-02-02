package main

import (
	"di-hbase-manager/utils"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/exec"
)

type HBaseAdmin struct {
}

type HBasePermission struct {
	User      string
	ACL       string
	TableName string
}

func (admin *HBaseAdmin) Operation(hbaseCommand string, output *string) error {
	commandStr := utils.EscapeHbaseCommandStr(hbaseCommand)
	fmt.Println(commandStr)
	cmd := exec.Command("bash", "-c", commandStr)
	f, _ := os.OpenFile("server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer f.Close()
	cmdStdoutPipe, _ := cmd.StdoutPipe()
	cmdStderrPipe, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	go utils.SyncLog(cmdStdoutPipe, f, output)
	go utils.SyncLog(cmdStderrPipe, f, output)
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		*output = *output + "\n============ HBaseAdmin operation Failed ============"
		return nil
	}
	*output = *output + "\n============ HBaseAdmin operation Successful ============"
	return nil
}

func (admin *HBaseAdmin) ModifyTablePermission(p HBasePermission, output *string) error {
	user := p.User
	acl := p.ACL
	tb := p.TableName
	hbaseCommand := "grant '" + user + "','" + acl + "','" + tb + "'"
	err := admin.Operation(hbaseCommand, output)
	return err
}

func main() {
	rpc.Register(new(HBaseAdmin))
	lis, err := net.Listen("tcp", ":7654")
	if err != nil {
		log.Panicln(err)
	}
	pid := os.Getpid()
	fmt.Printf("Server PID: %d \n", pid)
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("new client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}

}
