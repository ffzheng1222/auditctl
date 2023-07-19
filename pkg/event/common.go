package event

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/viper"
)

// 构建全局k8s audit record数据存储结构体
type K8sAuditLog struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"apiVersion"`
	Level      string `json:"level"`
	AuditID    string `json:"auditID"`
	Stage      string `json:"stage"`
	RequestURI string `json:"requestURI"`
	Verb       string `json:"verb"`
	User       struct {
		Username string   `json:"username"`
		Groups   []string `json:"groups"`
	} `json:"user"`
	SourceIPs []string `json:"sourceIPs"`
	UserAgent string   `json:"userAgent"`
	ObjectRef struct {
		Resource    string `json:"resource"`
		Namespace   string `json:"namespace"`
		Name        string `json:"name"`
		APIVersion  string `json:"apiVersion"`
		Subresource string `json:"subresource"`
	} `json:"objectRef"`
	ResponseStatus struct {
		Metadata struct {
		} `json:"metadata"`
		Code int `json:"code"`
	} `json:"responseStatus"`
	RequestReceivedTimestamp time.Time `json:"requestReceivedTimestamp"`
	StageTimestamp           time.Time `json:"stageTimestamp"`
	Annotations              struct {
		AuthorizationK8SIoDecision string `json:"authorization.k8s.io/decision"`
		AuthorizationK8SIoReason   string `json:"authorization.k8s.io/reason"`
	} `json:"annotations"`
}

// 全局记录k8s audit 审计日志
var auditRecords []K8sAuditLog

func NewAuditRecordArr() []K8sAuditLog {
	defAuditLogfile := viper.GetString("logFile")
	readAuditLogfile(defAuditLogfile)

	return auditRecords
}

func readAuditLogfile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err !", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		convertToGOStruct(str)
	}
}

func convertToGOStruct(jsonStr string) {
	var k8sAuditLog K8sAuditLog

	err := json.Unmarshal([]byte(jsonStr), &k8sAuditLog)
	if err != nil {
		fmt.Printf("jsonStr unmarshal map err=%v\n", err)
	}

	//保存所有已经序列化的audit log到全局数据结构auditRecords中.
	auditRecords = append(auditRecords, k8sAuditLog)
}
