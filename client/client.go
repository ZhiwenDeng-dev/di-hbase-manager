package main

import (
	"di-hbase-manager/utils"
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//err := execHBaseCommand("create_namespace 'paredose_credit_live_ph'")
	//err := execHBaseCommand("create 'paredose_credit_live_ph:ph_metrics_kredit_f_get_spl_total_limit_utilization_core_v2'," +
	//	"{NAME => 'cf',COMPRESSION => 'SNAPPY', DATA_BLOCK_ENCODING => 'FAST_DIFF',TTL => 'FOREVER'}, {NUMREGIONS => 4, SPLITALGO => 'HexStringSplit'}")
	//execHBaseCommand("list_namespace_tables 'paredose_credit_live_ph'")
	//execHBaseCommand("list_namespace_tables 'paredose_credit_live_ph'")
	//execHBaseCommand("disable 'paredose_credit_live_ph:ph_metrics_kredit_f_get_spl_total_limit_utilization_core_v2'")
	//execHBaseCommand("enable 'paredose_credit_live_ph:ph_metrics_kredit_f_get_spl_total_limit_utilization_core_v2'")
	//execHBaseCommand("put 'paredose_credit_live_ph:ph_metrics_kredit_f_get_spl_total_limit_utilization_core_v2','r1','cf','v1' ")
	//execHBaseCommand("put 'paredose_credit_live_ph:ph_metrics_kredit_f_get_spl_total_limit_utilization_core_v2','r2','cf','v2' ")
	execHBaseCommand("scan 'paredose_credit_live_ph:ph_metrics_kredit_f_get_spl_total_limit_utilization_core_v2'")
}

func execHBaseCommand(hbaseCommand string) {
	conn, err := jsonrpc.Dial("tcp", "docker-hbase:7654")
	defer func(conn *rpc.Client) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
	if err != nil {
		log.Fatal(err)
	}
	output := ""
	command := utils.ConcatHbaseCommandStr(hbaseCommand)
	e := conn.Call("HBaseAdmin.Operation", command, &output)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(output)
}
