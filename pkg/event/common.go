package event

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/viper"
)

// 构建全局k8s audit record数据存储结构体
type K8sAuditLog struct {
	AuditID   string `json:"auditID"`
	Stage     string `json:"stage"`
	Verb      string `json:"verb"`
	ObjectRef struct {
		Resource  string `json:"resource"`
		Namespace string `json:"namespace"`
		Name      string `json:"name"`
	} `json:"objectRef"`
}

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
