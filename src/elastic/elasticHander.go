package elastic

import (
	"fmt"
	"github.com/pkg/errors"
	elastic_v5 "gopkg.in/olivere/elastic.v5"
	"moqikaka.com/elastic_stack/src/enum"
	"moqikaka.com/elastic_stack/src/model"
	"moqikaka.com/elastic_stack/src/util"
)

const (
	url = "http://10.254.0.162:9200/"
)

var (
	elastic *Elastic
)

func init() {
	tempElastic, err := NewElastic(url, "9504")
	if err != nil {
		panic(errors.New("创建NewElastic报错，请检测"))
	}

	elastic = tempElastic
}

// 打印错误日志
// 参数：
// msg：消息信息
func LogError(msg interface{}) {
	msgObj := model.NewMessageObj(util.GetGuid(), msg)
	_, err := elastic.save(msgObj, enum.Debug)
	if err != nil {
		return
	}
}

// 打印debug日志
// 参数：
// msg：消息信息
func LogDebug(msg interface{}) {
	msgObj := model.NewMessageObj(util.GetGuid(), msg)
	_, err := elastic.save(msgObj, enum.Debug)
	if err != nil {
		return
	}
}

// 打印info日志
// 参数：
// msg：消息信息
func LogInfo(msg interface{}) {
	msgObj := model.NewMessageObj(util.GetGuid(), msg)
	_, err := elastic.save(msgObj, enum.Info)
	if err != nil {
		return
	}
}

// 打印info日志
// 参数：
// msg：消息信息
func LogWarn(msg interface{}) {
	msgObj := model.NewMessageObj(util.GetGuid(), msg)
	_, err := elastic.save(msgObj, enum.Warn)
	if err != nil {
		return
	}
}

// 获取消息信息
func GetMsg() {
	termQuery := elastic_v5.NewTermQuery("Age", "20")
	searchResult, err := elastic.search(termQuery)
	if err != nil {
		return
	}

	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)
		for _, hit := range searchResult.Hits.Hits {
			fmt.Println(fmt.Sprintf("消息数据%v", hit.Source))
		}
	} else {
		fmt.Print("Found no tweets\n")
	}
}
