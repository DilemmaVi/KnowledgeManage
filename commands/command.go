package commands

import (
	"KnowledgeManage/models"
	"encoding/gob"
	"flag"
	"log"
	"reflect"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Searcher      = engine.Engine{}
	dictFile      = flag.String("dict_file", "static/data/dictionary.txt", "词典文件")
	stopTokenFile = flag.String("stop_token_file", "static/data/stop_tokens.txt", "停用词文件")
)

const (
	SecondsInADay     = 86400
	MaxTokenProximity = 2
)

func Register() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RegisterModel(new(models.Knowledgedata))
	orm.RegisterModel(new(models.Member))
	orm.RegisterModel(new(models.Classifydata))
	createTable()
	searchInit()
}

func createTable() {
	o := orm.NewOrm()
	o.Using("default")
	o.Raw("CREATE TABLE Knowledgedata (Id integer not null primary key, Yjfl text,Ejfl text,Sjfl text,Title text,Content text,ContentHTML text,Keyword text,CreateTime text,ModifyTime text,Creator text,Reviser text);").Exec()
	o.Raw("CREATE TABLE members (Id integer not null primary key, account text,password text,name text,email text,phone text,role text,role_name text,status text,create_time text,last_login_time text);").Exec()
	o.Raw("CREATE TABLE Classifydata (Id integer not null primary key, Yjfl text,Ejfl text,Sjfl text,Status,CreateTime text,ModifyTime text,Creator text,Reviser text);").Exec()

}

func searchInit() {
	flag.Parse()

	// 初始化搜索引擎
	gob.Register(knowledgeScoringFields{})
	log.Print("引擎开始初始化")
	Searcher.Init(types.EngineInitOptions{
		SegmenterDictionaries: *dictFile,
		StopTokenFile:         *stopTokenFile,
		IndexerInitOptions: &types.IndexerInitOptions{
			IndexType: types.LocationsIndex,
		},
	})
	log.Print("引擎初始化完毕")
	// 索引
	log.Print("建索引开始")
	go IndexKnowledage()
	log.Print("建索引完毕")
}

/*******************************************************************************
    索引
*******************************************************************************/
func IndexKnowledage() {
	// 读入知识库数据
	knowResutlt, _ := models.NewKnowledge().GetAllKnowledgeData()
	log.Print("添加索引")
	for _, knowledge := range knowResutlt {
		stamp, _ := time.Parse("2006-01-02 15:04:05", knowledge.Createtime)
		content := knowledge.Title + knowledge.Keyword + knowledge.Content
		Searcher.IndexDocument(uint64(knowledge.Id), types.DocumentIndexData{
			Content: content,
			Fields: knowledgeScoringFields{
				Timestamp: stamp.Unix(),
			},
		}, false)
	}

	Searcher.FlushIndex()
	log.Printf("索引了%d条知识\n", len(knowResutlt))
}

/*******************************************************************************
    评分
*******************************************************************************/
type knowledgeScoringFields struct {
	Timestamp int64
}

type KnowledgeScoringCriteria struct {
}

func (criteria KnowledgeScoringCriteria) Score(
	doc types.IndexedDocument, fields interface{}) []float32 {
	if reflect.TypeOf(fields) != reflect.TypeOf(knowledgeScoringFields{}) {
		return []float32{}
	}
	wsf := fields.(knowledgeScoringFields)
	output := make([]float32, 2)
	if doc.TokenProximity > MaxTokenProximity {
		output[0] = 1.0 / float32(doc.TokenProximity)
	} else {
		output[0] = 1.0
	}
	output[1] = float32(wsf.Timestamp / (SecondsInADay * 3))
	return output
}
