package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os/exec"
)

type HBaseAdmin struct {
}

func (admin *HBaseAdmin) Operation(hbaseCommand string, output *string) error {
	cmd := exec.Command("bash", "-c", hbaseCommand)
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("[ERROR] HBase operation error:\n", err)
		*output = "[ERROR] HBase operation error:\n" + string(stdout)
		return err
	}
	fmt.Println("[INFO] HBase operation success:", string(stdout))
	*output = "[INFO] HBase operation success:\n" + string(stdout)
	return nil
}

func main() {

	rpc.Register(new(HBaseAdmin))
	lis, err := net.Listen("tcp", ":7654")
	if err != nil {
		log.Panicln(err)
	}
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
