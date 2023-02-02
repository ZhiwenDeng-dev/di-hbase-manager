package main

import "di-hbase-manager/client"

func main() {
	//err := execHBaseCommand("create_namespace 'paredose_credit_live_ph'")
	//execHBaseCommand("create 'paredose_credit_live_ph:ph_metrics_kredit_f_get_spl_total_limit_utilization_core_v3'," + "{NAME => 'cf',COMPRESSION => 'SNAPPY', DATA_BLOCK_ENCODING => 'FAST_DIFF',TTL => 'FOREVER'}, {NUMREGIONS => 4, SPLITALGO => 'HexStringSplit'}")
	//execHBaseCommand("list_namespace_tables 'paredose_credit_live_ph'")
	//execHBaseCommand("list_namespace_tables 'paredose_credit_live_ph'")
	//execHBaseCommand("disable 'paredose_credit_live_ph:ph_metrics_kredit_f_get_spl_total_limit_utilization_core_v2'")
	//execHBaseCommand("enable 'paredose_credit_live_ph:ph_metrics_kredit_f_get_spl_total_limit_utilization_core_v2'")
	//execHBaseCommand("put 'paredose_credit_live_ph:ph_metrics_kredit_f_get_spl_total_limit_utilization_core_v2','r1','cf','v1' ")
	//execHBaseCommand("put 'paredose_credit_live_ph:ph_metrics_kredit_f_get_spl_total_limit_utilization_core_v2','r2','cf','v2' ")
	//execHBaseCommand("scan 'paredose_credit_live_ph:ph_metrics_kredit_f_get_spl_total_limit_utilization_core_v2'")
	p := client.HBasePermission{
		"hbase",
		"CRW",
		"paredose_credit_live_ph:ph_metrics_kredit_f_get_spl_total_limit_utilization_core_v2",
	}
	client.ModifyTablePermission(p)
	//execHBaseCommand("user_permission 'paredose_credit_live_ph:ph_metrics_kredit_f_get_spl_total_limit_utilization_core_v2'")
}
