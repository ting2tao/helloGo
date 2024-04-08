package main

import (
	"fmt"
	"github.com/volcengine/volc-sdk-golang/service/maas"
	"github.com/volcengine/volc-sdk-golang/service/maas/models/api"
	"github.com/xuri/excelize/v2"
	"os"
)

func main() {
	r := maas.NewInstance("maas-api.ml-platform-cn-beijing.volces.com", "cn-beijing")

	// fetch ak&sk from environmental variables
	r.SetAccessKey(os.Getenv("VOLC_ACCESSKEY"))
	r.SetSecretKey(os.Getenv("VOLC_SECRETKEY"))

	req := &api.ChatReq{
		Model: &api.Model{
			Name:    "moonshot-v1-8k",
			Version: "1.0", // use default version if not specified.
		},
		Messages: []*api.Message{
			{
				Role:    maas.ChatRoleOfUser,
				Content: "天为什么这么蓝？",
			},
			{
				Role:    maas.ChatRoleOfAssistant,
				Content: "因为有你",
			},
			{
				Role:    maas.ChatRoleOfUser,
				Content: "花儿为什么这么香？",
			},
		},
		Parameters: &api.Parameters{
			MaxNewTokens: 1000, // 输出文本的最大tokens限制，max_new_tokens + input_length <= max_input_size
			Temperature:  0.3,  // 用于控制生成文本的随机性和创造性，Temperature值越大随机性越大，取值范围0~1
			TopP:         0.9,  // 用于控制输出tokens的多样性，TopP值越大输出的tokens类型越丰富，取值范围0~1
		},
	}
	TestChat(r, req)
	TestStreamChat(r, req)
}

type xlsxData struct {
	No           string
	ID           string
	C            string
	UserQuestion string
	Label        string
	QType        string
	QAs          string
}

func getData(filePath string) ([]*xlsxData, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("source")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	data := make([]*xlsxData, 0, len(rows))
	for i, row := range rows {
		if i != 0 {
			data = append(data, &xlsxData{
				No:           row[0],
				ID:           row[1],
				C:            row[2],
				UserQuestion: row[3],
				Label:        row[4],
				QType:        row[5],
				QAs:          row[6],
			})
		}
	}

	return data, nil
}
