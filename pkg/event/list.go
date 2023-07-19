package event

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var ListTableHeader = []string{"NAMESPACE", "RESOURCE", "NAME", "VERB", "STAGE"}
var ObjectRefHeader = []string{"NAMESPACE", "RESOURCE", "NAME"}

type AuditLogInfo struct {
	allAuditLogData []map[string]AuditLogData
}

type AuditLogData struct {
	OneColumDescribe map[string]ColumDescribe
}

type ColumDescribe struct {
	ColumnIndex      int
	ColumnValue      string
	ColumnSpaceCount int
}

type ObjectRef struct {
	Resource    string `json:"resource"`
	Namespace   string `json:"namespace"`
	Name        string `json:"name"`
	APIVersion  string `json:"apiVersion"`
	Subresource string `json:"subresource"`
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

func NewAuditLogInfo() *AuditLogInfo {
	var auditLogInfo AuditLogInfo
	oneAuditLogDataArr := make([]map[string]AuditLogData, 0)
	oneColumDescribe := make(map[string]ColumDescribe)

	for inx, thName := range ListTableHeader {
		oneColumDescribe[thName] = ColumDescribe{inx, "", 0}
	}

	auditLogInfo.allAuditLogData = oneAuditLogDataArr
	return &auditLogInfo
}

func judgeContainRelation(orStr string, strSlice *[]string) bool {

	for _, str := range *strSlice {
		if str == orStr {
			return true
		}
	}
	return false
}

func reflectGetObjectRef(thName string, objectRef ObjectRef) (*reflect.StructField, *reflect.Value) {
	var tf reflect.StructField
	var tv reflect.Value
	reflectType := reflect.TypeOf(objectRef)
	reflectValue := reflect.ValueOf(objectRef)

	for i := 0; i < reflectType.NumField(); i++ {
		tf = reflectType.Field(i)

		if ok := strings.EqualFold(thName, tf.Name); ok {
			//fmt.Printf("*** %v \n", tf.Name)
			tv = reflectValue.FieldByName(tf.Name)
		}
	}

	//fmt.Printf("tonyfan3 %v\n", tv.Interface())
	return &tf, &tv
}

func reflectGetAuditRecord(thName string, auditRecord K8sAuditLog) (*reflect.StructField, *reflect.Value) {
	var tf reflect.StructField
	var tv reflect.Value
	reflectType := reflect.TypeOf(auditRecord)
	reflectValue := reflect.ValueOf(auditRecord)

	// 可以规范化处理
	if ok := judgeContainRelation(thName, &ObjectRefHeader); ok {
		thName = "ObjectRef"
	}

	for i := 0; i < reflectType.NumField(); i++ {
		tf = reflectType.Field(i)

		if ok := strings.EqualFold(thName, tf.Name); ok {
			tv = reflectValue.FieldByName(tf.Name)
		}
	}

	//fmt.Printf("tonyfan2 %v\n", tv.Interface())
	return &tf, &tv
}

func interfaceToStruct(rfV *reflect.Value) (*ObjectRef, error) {
	toInterface := rfV.Interface()
	resByre, resByteErr := json.Marshal(toInterface)
	if resByteErr != nil {
		fmt.Printf("%v", resByteErr)
		return nil, resByteErr
	}

	var objectRef ObjectRef
	jsonResErr := json.Unmarshal(resByre, &objectRef)
	if jsonResErr != nil {
		fmt.Printf("%v", jsonResErr)
		return nil, jsonResErr
	}
	//fmt.Printf("使用 json: %v \n", objectRef)
	return &objectRef, nil
}

func listK8sAuditLog() {
	//NewAuditRecordArr 记录着全量的audit审计日志
	auditRecordarr := NewAuditRecordArr()

	var auditLogData AuditLogData
	tableColumDescribe := make(map[string]ColumDescribe)

	// 初始化audit log list表头header
	auditLogInfo := NewAuditLogInfo()

	var (
		auditRecordValue  *reflect.Value
		columnFieldLength int
		columnValue       string
	)
	for _, auditRecord := range auditRecordarr {
		oneAuditLogData := make(map[string]AuditLogData, 0)
		auditID := auditRecord.AuditID

		for inx, thName := range ListTableHeader {
			_, auditRecordValue = reflectGetAuditRecord(thName, auditRecord)

			if ok := judgeContainRelation(thName, &ObjectRefHeader); ok {
				objectRef, err := interfaceToStruct(auditRecordValue)
				if err != nil {
					fmt.Printf("%v", err)
				}
				_, objectRefValue := reflectGetObjectRef(thName, *objectRef)

				columnValue = objectRefValue.Interface().(string)
				columnFieldLength = len(objectRefValue.Interface().(string))
			} else {
				columnValue = auditRecordValue.Interface().(string)
				columnFieldLength = len(auditRecordValue.Interface().(string))
			}
			tableColumDescribe[thName] = ColumDescribe{inx, columnValue, columnFieldLength}
			auditLogData.OneColumDescribe = tableColumDescribe
		}
		oneAuditLogData[auditID] = auditLogData
		auditLogInfo.allAuditLogData = append(auditLogInfo.allAuditLogData, oneAuditLogData)
	}

	fmt.Printf("%v \n", auditLogInfo)

	auditLogListPrint(auditLogInfo)
}

func auditLogListPrint(alInfo *AuditLogInfo) {
	var nsMax, resourceMax, nameMax, verbMax, statgeMax string
	//var printRecordInfo AuditLogData

	for _, auditLogDatas := range (*alInfo).allAuditLogData {
		for _, recordInfo := range auditLogDatas {
			var cloumPrinter = recordInfo.OneColumDescribe
			nsMax = strconv.Itoa(cloumPrinter["NAMESPACE"].ColumnSpaceCount)
			resourceMax = strconv.Itoa(cloumPrinter["RESOURCE"].ColumnSpaceCount)
			nameMax = strconv.Itoa(cloumPrinter["NAME"].ColumnSpaceCount)
			verbMax = strconv.Itoa(cloumPrinter["VERB"].ColumnSpaceCount)
			statgeMax = strconv.Itoa(cloumPrinter["STAGE"].ColumnSpaceCount)

			if cloumPrinter["NAMESPACE"].ColumnValue == "" {
				var cdStrucu ColumDescribe
				cdStrucu.ColumnIndex = cloumPrinter["NAMESPACE"].ColumnIndex
				cdStrucu.ColumnSpaceCount = cloumPrinter["NAMESPACE"].ColumnSpaceCount
				cdStrucu.ColumnValue = "--"
				cloumPrinter["NAMESPACE"] = cdStrucu
			}
			//printRecordInfo = recordInfo
		}

		fmt.Printf("nsMax=%s , resourceMax=%s , nameMax=%s , verbMax=%s statgeMax=%s \n",
			nsMax, resourceMax, nameMax, verbMax, statgeMax)

		// fmt.Printf("%-"+nsMax+"s %-"+resourceMax+"s %-"+nameMax+"s %-"+verbMax+"s %-"+statgeMax+"s\n",
		// 	printRecordInfo.OneColumDescribe["NAMESPACE"], printRecordInfo.OneColumDescribe["RESOURCE"],
		// 	printRecordInfo.OneColumDescribe["NAME"], printRecordInfo.OneColumDescribe["VERB"],
		// 	printRecordInfo.OneColumDescribe["STAGE"])
	}
}
