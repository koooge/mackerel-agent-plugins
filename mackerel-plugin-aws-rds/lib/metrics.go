package mpawsrds

import (
	mp "github.com/mackerelio/go-mackerel-plugin"
)

func (p RDSPlugin) rdsMetrics() (metrics []string) {
	for _, v := range p.GraphDefinition() {
		for _, vv := range v.Metrics {
			metrics = append(metrics, vv.Name)
		}
	}
	return
}

func mergeGraphDefs(a, b map[string]mp.Graphs) map[string]mp.Graphs {
	for k, v := range b {
		a[k] = v
	}
	return a
}

func (p RDSPlugin) mySQLGraphDefinition() map[string]mp.Graphs {
	return map[string]mp.Graphs{
		p.Prefix + ".BinLogDiskUsage": {
			Label: p.LabelPrefix + " BinLogDiskUsage",
			Unit:  "bytes",
			Metrics: []mp.Metrics{
				{Name: "BinLogDiskUsage", Label: "Usage"},
			},
		},
	}
}

func (p RDSPlugin) postgreSQLGraphDefinition() map[string]mp.Graphs {
	return map[string]mp.Graphs{
		p.Prefix + ".MaximumUsedTransactionIDs": {
			Label: p.LabelPrefix + " Maximum Used Transaction IDs",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "MaximumUsedTransactionIDs", Label: "IDs"},
			},
		},
		p.Prefix + ".OldestReplicationSlotLag": {
			Label: p.LabelPrefix + " Oldest Replication Slot Lag",
			Unit:  "bytes",
			Metrics: []mp.Metrics{
				{Name: "OldestReplicationSlotLag", Label: "SlotLag"},
			},
		},
		p.Prefix + ".ReplicationSlotDiskUsage": {
			Label: p.LabelPrefix + " Replication Slot Disk Usage",
			Unit:  "bytes",
			Metrics: []mp.Metrics{
				{Name: "ReplicationSlotDiskUsage", Label: "SlotDiskUsage"},
			},
		},
		p.Prefix + ".TransactionLogsDiskUsage": {
			Label: p.LabelPrefix + " Transaction Logs Disk Usage",
			Unit:  "bytes",
			Metrics: []mp.Metrics{
				{Name: "TransactionLogsDiskUsage", Label: "DiskUsage"},
			},
		},
		p.Prefix + ".TransactionLogsGeneration": {
			Label: p.LabelPrefix + " Transaction Logs Generation",
			Unit:  "bytes/sec",
			Metrics: []mp.Metrics{
				{Name: "TransactionLogsGeneration", Label: "Generation"},
			},
		},
	}
}

func (p RDSPlugin) auroraGraphDefinition() map[string]mp.Graphs {
	return map[string]mp.Graphs{
		p.Prefix + ".CPUUtilization": {
			Label: p.LabelPrefix + " CPU Utilization",
			Unit:  "percentage",
			Metrics: []mp.Metrics{
				{Name: "CPUUtilization", Label: "CPUUtilization"},
			},
		},
		// .CPUCreditBalance ...Only valid for T2 instances
		p.Prefix + ".CPUCreditBalance": {
			Label: p.LabelPrefix + " CPU CreditBalance",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "CPUCreditBalance", Label: "CPUCreditBalance"},
			},
		},
		// .CPUCreditUsage ...Only valid for T2 instances
		p.Prefix + ".CPUCreditUsage": {
			Label: p.LabelPrefix + " CPU CreditUsage",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "CPUCreditUsage", Label: "CPUCreditUsage"},
			},
		},
		p.Prefix + ".Deadlocks": {
			Label: p.LabelPrefix + " Dead Locks",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "Deadlocks", Label: "Deadlocks"},
			},
		},
		p.Prefix + ".DatabaseConnections": {
			Label: p.LabelPrefix + " Database Connections",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "DatabaseConnections", Label: "DatabaseConnections"},
			},
		},
		p.Prefix + ".FreeLocalStorage": {
			Label: p.LabelPrefix + " Free Local Storage",
			Unit:  "bytes",
			Metrics: []mp.Metrics{
				{Name: "FreeLocalStorage", Label: "FreeLocalStorage"},
			},
		},
		p.Prefix + ".FreeableMemory": {
			Label: p.LabelPrefix + " Freeable Memory",
			Unit:  "bytes",
			Metrics: []mp.Metrics{
				{Name: "FreeableMemory", Label: "FreeableMemory"},
			},
		},
		p.Prefix + ".EngineUptime": {
			Label: p.LabelPrefix + " Engine Uptime",
			Unit:  "integer",
			Metrics: []mp.Metrics{
				{Name: "EngineUptime", Label: "EngineUptime"},
			},
		},
		p.Prefix + ".AuroraReplicaLag": {
			Label: p.LabelPrefix + " Aurora ReplicaLag",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "AuroraReplicaLagMaximum", Label: "Maximum"},
				{Name: "AuroraReplicaLagMinimum", Label: "Minimum"},
			},
		},
		p.Prefix + ".NetworkThroughput": {
			Label: p.LabelPrefix + " Network Throughput",
			Unit:  "bytes/sec",
			Metrics: []mp.Metrics{
				{Name: "NetworkTransmitThroughput", Label: "Transmit"},
				{Name: "NetworkReceiveThroughput", Label: "Receive"},
			},
		},
	}
}

func (p RDSPlugin) auroraMySQLGraphDefinition() map[string]mp.Graphs {
	return map[string]mp.Graphs{
		p.Prefix + ".Transaction": {
			Label: p.LabelPrefix + " Transaction",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "ActiveTransactions", Label: "Active"},
				{Name: "BlockedTransactions", Label: "Blocked"},
			},
		},
		p.Prefix + ".Latency": {
			Label: p.LabelPrefix + " Latency [msec]",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "SelectLatency", Label: "Select"},
				{Name: "InsertLatency", Label: "Insert"},
				{Name: "UpdateLatency", Label: "Update"},
				{Name: "DeleteLatency", Label: "Delete"},
				{Name: "CommitLatency", Label: "Commit"},
				{Name: "DDLLatency", Label: "DDL"},
				{Name: "DMLLatency", Label: "DML"},
			},
		},
		p.Prefix + ".Throughput": {
			Label: p.LabelPrefix + " Throughput [ops/sec]",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "SelectThroughput", Label: "Select"},
				{Name: "InsertThroughput", Label: "Insert"},
				{Name: "UpdateThroughput", Label: "Update"},
				{Name: "DeleteThroughput", Label: "Delete"},
				{Name: "CommitThroughput", Label: "Commit"},
				{Name: "DDLThroughput", Label: "DDL"},
				{Name: "DMLThroughput", Label: "DML"},
			},
		},
		p.Prefix + ".Queries": {
			Label: p.LabelPrefix + " Queries",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "Queries", Label: "Queries"},
			},
		},
		p.Prefix + ".LoginFailures": {
			Label: p.LabelPrefix + " Login Failures",
			Unit:  "integer",
			Metrics: []mp.Metrics{
				{Name: "LoginFailures", Label: "LoginFailures"},
			},
		},
		p.Prefix + ".CacheHitRatio": {
			Label: p.LabelPrefix + " Cache Hit Ratio",
			Unit:  "percentage",
			Metrics: []mp.Metrics{
				{Name: "ResultSetCacheHitRatio", Label: "ResultSet"},
				{Name: "BufferCacheHitRatio", Label: "Buffer"},
			},
		},
		p.Prefix + ".AuroraBinlogReplicaLag": {
			Label: p.LabelPrefix + " Aurora Binlog ReplicaLag",
			Unit:  "integer",
			Metrics: []mp.Metrics{
				{Name: "AuroraBinlogReplicaLag", Label: "BinlogReplicaLag"},
			},
		},
	}
}

func (p RDSPlugin) auroraPostgreSQLGraphDefinition() map[string]mp.Graphs {
	return map[string]mp.Graphs{
		p.Prefix + ".DiskQueueDepth": {
			Label: p.LabelPrefix + " Disk Queue Depth",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "DiskQueueDepth", Label: "Depth"},
			},
		},
		p.Prefix + ".SwapUsage": {
			Label: p.LabelPrefix + " Swap Usage",
			Unit:  "bytes",
			Metrics: []mp.Metrics{
				{Name: "SwapUsage", Label: "SwapUsage"},
			},
		},
		p.Prefix + ".IOPS": {
			Label: p.LabelPrefix + " IOPS",
			Unit:  "iops",
			Metrics: []mp.Metrics{
				{Name: "ReadIOPS", Label: "Read"},
				{Name: "WriteIOPS", Label: "Write"},
			},
		},
		p.Prefix + ".Latency": {
			Label: p.LabelPrefix + " Latency [msec]",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "ReadLatency", Label: "Read"},
				{Name: "WriteLatency", Label: "Write"},
				{Name: "CommitLatency", Label: "Commit"},
			},
		},
		p.Prefix + ".Throughput": {
			Label: p.LabelPrefix + " Throughput",
			Unit:  "bytes/sec",
			Metrics: []mp.Metrics{
				{Name: "ReadThroughput", Label: "Read"},
				{Name: "WriteThroughput", Label: "Write"},
				{Name: "CommitThroughput", Label: "Commit"},
			},
		},
		p.Prefix + ".CacheHitRatio": {
			Label: p.LabelPrefix + " Cache Hit Ratio",
			Unit:  "percentage",
			Metrics: []mp.Metrics{
				{Name: "ResultSetCacheHitRatio", Label: "ResultSet"},
			},
		},
		p.Prefix + ".MaximumUsedTransactionIDs": {
			Label: p.LabelPrefix + " Maximum Used Transaction IDs",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "MaximumUsedTransactionIDs", Label: "IDs"},
			},
		},
		p.Prefix + ".TransactionLogsDiskUsage": {
			Label: p.LabelPrefix + " Transaction Logs Disk Usage",
			Unit:  "bytes",
			Metrics: []mp.Metrics{
				{Name: "TransactionLogsDiskUsage", Label: "DiskUsage"},
			},
		},
	}
}
