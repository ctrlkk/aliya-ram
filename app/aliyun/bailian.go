package aliyun

import (
	"os"
	"regexp"

	bailian "github.com/alibabacloud-go/bailian-20231229/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Bailian struct {
	client      *bailian.Client
	workspaceId string
	indexId     string
}

func NewBalilian() (*Bailian, error) {
	config := &openapiutil.Config{
		AccessKeyId:     tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID")),
		AccessKeySecret: tea.String(os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")),
	}
	config.Endpoint = tea.String("bailian.cn-beijing.aliyuncs.com")

	client, err := bailian.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &Bailian{
		client:      client,
		workspaceId: os.Getenv("WORKSPACE_ID"),
		indexId:     os.Getenv("INDEX_ID"),
	}, nil
}

func retrieveIndex(client *bailian.Client, workspaceId, indexId, query string) (*bailian.RetrieveResponse, error) {
	headers := make(map[string]*string)
	request := &bailian.RetrieveRequest{
		IndexId: tea.String(indexId),
		Query:   tea.String(query),
	}
	runtime := &util.RuntimeOptions{}
	return client.RetrieveWithOptions(tea.String(workspaceId), request, headers, runtime)
}

func (bailian *Bailian) Query(query string) ([]*bailian.RetrieveResponseBodyDataNodes, error) {
	res, err := retrieveIndex(bailian.client, bailian.workspaceId, bailian.indexId, query)
	if err != nil {
		return nil, err
	}
	nodes := res.GetBody().GetData().Nodes
	reg1 := regexp.MustCompile(`AliyaMessage[^|]*\|`)
	reg2 := regexp.MustCompile(`PlayerChoice[^|]*\|`)
	for _, node := range nodes {
		text := reg1.ReplaceAllString(*node.GetText(), "Aliya:")
		text = reg2.ReplaceAllString(text, "COSMOS:")
		node.Text = &text
	}

	return nodes, nil
}
