package client

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HBasePermission struct {
	User      string
	ACL       string
	TableName string
}

var (
	server_ip     = "docker-hbase"
	server_port   = "7654"
	serverAddress = server_ip + ":" + server_port
)

func getConnection(serverAddress string) *rpc.Client {
	conn, err := jsonrpc.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func ExecHBaseCommand(hbaseCommand string) {
	conn := getConnection(serverAddress)
	defer func(conn *rpc.Client) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
	output := ""
	e := conn.Call("HBaseAdmin.Operation", hbaseCommand, &output)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(output)
}

func ModifyTablePermission(p HBasePermission) {
	conn := getConnection(serverAddress)
	defer func(conn *rpc.Client) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
	output := ""
	e := conn.Call("HBaseAdmin.ModifyTablePermission", &p, &output)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(output)
}
