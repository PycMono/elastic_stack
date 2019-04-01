package elastic

import (
	"context"
	"fmt"
	elastic_v5 "gopkg.in/olivere/elastic.v5"
	"moqikaka.com/elastic_stack/src/model"
)

// Elastic对象
type Elastic struct {
	// 客服端对象
	client *elastic_v5.Client

	// 保存的地址
	index string

	// elasticURL地址
	url string
}

// 消息保存
// 参数：
// msg：消息信息
// id：消息Id
// msgType：消息类型
// 返回值：
// 1.IndexResponse is the result of indexing a document in Elasticsearch.
// 2.错误信息
func (this *Elastic) save(msg *model.MessageObj, msgType string) (*elastic_v5.IndexResponse, error) {
	indexRes, err := this.client.Index().
		Index(this.index).
		Type(msgType).
		Id(msg.ID).
		Timestamp(fmt.Sprintf("%v", msg.NowTime.Unix())). // 文档创建时间
		BodyJson(msg).
		Do(context.Background())

	if err != nil {
		// 打印日志
		fmt.Println(fmt.Printf("Elastic.save保存数据报错err=%v", err))
		return nil, err
	}

	// 刷新
	_, err = this.client.Flush().Index(this.index).Do(context.Background())
	if err != nil {
		// 打印日志
		fmt.Println(fmt.Printf("Elastic.save Flush数据报错err=%v", err))
		return nil, err
	}

	return indexRes, nil
}

// 查询消息
// 参数：
// query：查询调教
// 结果：
// 1.查询结果
// 2.错误信息
func (this *Elastic) search(query elastic_v5.Query) (*elastic_v5.SearchResult, error) {
	resp, err := this.client.Search().
		Index(this.index).
		Type("doc").
		Query(query).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		// 打印日志
		fmt.Println(fmt.Printf("Elastic.get获取数据报错err=%v", err))
		return nil, err
	}

	return resp, nil
}

// 创建新的NewElastic对象
// 参数：
// url：elastic地址
// index：数据保存的index，类似于mysql的数据库
// 返回值：
// 1.Elastic对象
// 2.错误对象
func NewElastic(url, index string) (*Elastic, error) {
	client, err := elastic_v5.NewClient(elastic_v5.SetURL(url))
	if err != nil {
		return nil, err
	}

	exists, err := client.IndexExists(index).Do(context.Background())
	if err != nil {
		return nil, err
	}

	if !exists {
		// 创建index
		client.CreateIndex(index)
	} else {
		fmt.Println(fmt.Sprintf("已经已经存在了%v", exists))
	}

	return &Elastic{
		client: client,
		url:    url,
		index:  index,
	}, nil
}
