package stack

import "github.com/go_datastructure/common/log"
import "fmt"

func init() {
	//Init Configuraiton, Ideally should read form Config File
	fmt.Printf("Initializing Stack Package")
	log.SetLogFormatter("JSON")
	log.SetLogOutputMode("info.log", "error.log")
}
