package event

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

type AuditLogListShower struct {
	Stage     string
	Verb      string
	Resource  string
	Namespace string
	Name      string
}

type ColumDescribe struct {
	ColumnIndex      int
	ColumnSpaceCount int
}

var defaultTableWidhPind = 4

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "auditctl event list",
	Long:  "list k8s apiserver audit event",
	Run: func(cmd *cobra.Command, args []string) {
		//defAuditConfigFile := viper.ConfigFileUsed()
		//fmt.Printf("Load k8s audit log file of %s \n\n", defAuditConfigFile)

		listK8sAuditLog()
	},
}

func initAuditLogListHeader(alListShowerArr *[]AuditLogListShower, columDescribe *map[string]ColumDescribe) {
	// 初始化audit log list表头header
	var auditLogListShower AuditLogListShower
	auditLogListShower.Namespace = "NAMESPACE"
	auditLogListShower.Resource = "RESOURCE"
	auditLogListShower.Name = "NAME"
	auditLogListShower.Verb = "VERB"
	auditLogListShower.Stage = "STAGE"
	*alListShowerArr = append(*alListShowerArr, auditLogListShower)

	(*columDescribe)["NAMESPACE"] = ColumDescribe{0, 0}
	(*columDescribe)["RESOURCE"] = ColumDescribe{1, 0}
	(*columDescribe)["NAME"] = ColumDescribe{2, 0}
	(*columDescribe)["VERB"] = ColumDescribe{3, 0}
	(*columDescribe)["STAGE"] = ColumDescribe{4, 0}
}

func listK8sAuditLog() {
	var auditLogListShower AuditLogListShower
	auditLogListShowerArrs := make([]AuditLogListShower, 0)

	//NewAuditRecordArr 记录着全量的audio审计日志
	auditRecordarr := NewAuditRecordArr()

	tableColumDescribe := make(map[string]ColumDescribe)
	initAuditLogListHeader(&auditLogListShowerArrs, &tableColumDescribe)

	for _, auditRecord := range auditRecordarr {
		auditLogListShower.Namespace = auditRecord.ObjectRef.Namespace
		auditLogListShower.Resource = auditRecord.ObjectRef.Resource
		auditLogListShower.Name = auditRecord.ObjectRef.Name
		auditLogListShower.Verb = auditRecord.Verb
		auditLogListShower.Stage = auditRecord.Stage

		auditLogListShowerArrs = append(auditLogListShowerArrs, auditLogListShower)
	}

	auditLogListPrintPrepare(&auditLogListShowerArrs, &tableColumDescribe)
	//fmt.Printf("%v \n", tableColumDescribe)

	for _, auditLogListShower := range auditLogListShowerArrs {
		//fmt.Printf("%v \n", auditLogListShower)
		auditLogListPrint(auditLogListShower, &tableColumDescribe)
	}
}

func auditLogListPrintPrepare(alListShowerArr *[]AuditLogListShower, columDescribe *map[string]ColumDescribe) {
	var index = 0
	var nsStrWidth, resourceStrWidh, nameStrWidth, verbStrWidth, stageStrWidth int

	for _, alListShower := range *alListShowerArr {
		nsStrWidth = len(alListShower.Namespace)
		resourceStrWidh = len(alListShower.Resource)
		nameStrWidth = len(alListShower.Name)
		verbStrWidth = len(alListShower.Verb)
		stageStrWidth = len(alListShower.Stage)

		if nsStrWidth > (*columDescribe)["NAMESPACE"].ColumnSpaceCount {
			(*columDescribe)["NAMESPACE"] = ColumDescribe{0, nsStrWidth}
		}

		if resourceStrWidh > (*columDescribe)["RESOURCE"].ColumnSpaceCount {
			(*columDescribe)["RESOURCE"] = ColumDescribe{1, resourceStrWidh}
		}

		if nameStrWidth > (*columDescribe)["NAME"].ColumnSpaceCount {
			(*columDescribe)["NAME"] = ColumDescribe{2, nameStrWidth}
		}

		if verbStrWidth > (*columDescribe)["VERB"].ColumnSpaceCount {
			(*columDescribe)["VERB"] = ColumDescribe{3, verbStrWidth}
		}

		if stageStrWidth > (*columDescribe)["STAGE"].ColumnSpaceCount {
			(*columDescribe)["STAGE"] = ColumDescribe{4, stageStrWidth}
		}
	}

	for key, value := range *columDescribe {
		index = value.ColumnIndex
		(*columDescribe)[key] = ColumDescribe{index, value.ColumnSpaceCount + defaultTableWidhPind}
	}
}

func auditLogListPrint(aLLShower AuditLogListShower, columDescribe *map[string]ColumDescribe) {
	if aLLShower.Namespace == "" {
		aLLShower.Namespace = "--"
	}

	var cloumPrinter = *columDescribe
	nsMax := strconv.Itoa(cloumPrinter["NAMESPACE"].ColumnSpaceCount)
	resourceMax := strconv.Itoa(cloumPrinter["RESOURCE"].ColumnSpaceCount)
	nameMax := strconv.Itoa(cloumPrinter["NAME"].ColumnSpaceCount)
	verbMax := strconv.Itoa(cloumPrinter["VERB"].ColumnSpaceCount)
	StatgeMax := strconv.Itoa(cloumPrinter["STAGE"].ColumnSpaceCount)

	fmt.Printf("%-"+nsMax+"s %-"+resourceMax+"s %-"+nameMax+"s %-"+verbMax+"s %-"+StatgeMax+"s\n",
		aLLShower.Namespace, aLLShower.Resource, aLLShower.Name, aLLShower.Verb, aLLShower.Stage)
}
