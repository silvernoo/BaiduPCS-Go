package pcscommand

import (
	"bytes"
	"net/http"
	"fmt"
	"github.com/iikira/BaiduPCS-Go/internal/pcsconfig"
)

func SendToRPC(link string) {
	client := &http.Client{}
	rawJson := `{
    "jsonrpc": "2.0",
    "method": "aria2.addUri",
    "id": "1",
    "params": [
        [
            "%s"
        ],
        {
            "user-agent": "%s"
        }
    ]
}
`
	postData := fmt.Sprintf(rawJson, link, pcsconfig.Config.UserAgent())
	println(postData)
	req, err := http.NewRequest("POST", pcsconfig.Config.RPCAddress(), bytes.NewBuffer([]byte(postData)))
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	_, err = client.Do(req)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
