package controllers

import (
	"KnowledgeManage/commands"
	"KnowledgeManage/models"
	"github.com/astaxie/beego"
	"github.com/huichen/wukong/types"
	"strings"
	"time"
)

type SearchController struct {
	BaseController
}

type JsonResponse struct {
	Docs []*models.Knowledgedata `json:"docs"`
}

var (
	searcher       = &commands.Searcher
	knowledgedatas = map[uint64]models.Knowledgedata{}
)

// Get 搜索首页
func (c *SearchController) Get() {
	c.Prepare()
	var result []map[string]interface{}
	var ejfl interface{}
	yjflList, _, err := models.NewClassify().FindByClassify("", -1)
	if err == nil {
		for _, v := range yjflList {
			first := make(map[string]interface{})
			second := make(map[string]interface{})
			ejflList, _, err := models.NewClassify().FindByClassify(v.(string), 0)
			if err == nil {
				for _, y := range ejflList {
					sjflList, _, err := models.NewClassify().FindByClassify(y.(string), 1)
					if err == nil {
						second[y.(string)] = sjflList
					} else {
						beego.Error(err.Error())
						c.Abort("500")
					}
					ejfl = second
				}
				first[v.(string)] = ejfl
				result = append(result, first)
			} else {
				beego.Error(err.Error())
				c.Abort("500")
			}

		}
	} else {
		beego.Error(err.Error())
		c.Abort("500")
	}

	c.Data["classify"] = result
	c.Data["title"] = "检索"
	c.Layout = "index.tpl"
	c.TplName = "search.html"
}

// Post 搜索首页
func (c *SearchController) Post() {
	c.Prepare()
	query := c.GetString("query")
	t1 := time.Now()
	docs := SearchResult(query)

	var result []map[string]interface{}
	var ejfl interface{}
	yjflList, _, err := models.NewClassify().FindByClassify("", -1)
	if err == nil {
		for _, v := range yjflList {
			first := make(map[string]interface{})
			second := make(map[string]interface{})
			ejflList, _, err := models.NewClassify().FindByClassify(v.(string), 0)
			if err == nil {
				for _, y := range ejflList {
					sjflList, _, err := models.NewClassify().FindByClassify(y.(string), 1)
					if err == nil {
						second[y.(string)] = sjflList
					} else {
						beego.Error(err.Error())
						c.Abort("500")
					}
					ejfl = second
				}
				first[v.(string)] = ejfl
				result = append(result, first)
			} else {
				beego.Error(err.Error())
				c.Abort("500")
			}

		}
	} else {
		beego.Error(err.Error())
		c.Abort("500")
	}
	elapsed := time.Since(t1)
	c.Data["elapsed"] = strings.Replace(elapsed.String(), "ms", "", -1)
	c.Data["docs"] = docs
	c.Data["count"] = len(docs)
	c.Data["classify"] = result
	c.Data["query"] = query
	c.Data["title"] = "检索结果"
	c.Layout = "index.tpl"
	c.TplName = "result.html"
}

// SearchResult 搜索引擎实现
func SearchResult(query string) []*models.Knowledgedata {
	output := searcher.Search(types.SearchRequest{
		Text: query,
		RankOptions: &types.RankOptions{
			ScoringCriteria: &commands.KnowledgeScoringCriteria{},
			OutputOffset:    0,
			MaxOutputs:      100,
		},
	})

	docs := []*models.Knowledgedata{}
	knowResutlt, err := models.NewKnowledge().GetAllKnowledgeData()
	if err == nil {
		for _, knowledge := range knowResutlt {
			knowledgedatas[uint64(knowledge.Id)] = knowledge
		}
	}

	for _, doc := range output.Docs {
		kd := knowledgedatas[doc.DocId]
		for _, t := range output.Tokens {
			kd.Title = strings.Replace(kd.Title, t, "<font color=red>"+t+"</font>", -1)
			kd.Content = strings.Replace(kd.Content, t, "<font color=red>"+t+"</font>", -1)
		}
		start := strings.Index(kd.Content, "<font color=red>") - 200
		end := strings.Index(kd.Content, "</font>") + 200
		if end > len(kd.Content) {
			end = len(kd.Content) - 1
		}
		if start < 0 {
			start = 0
		}

		kd.Content = kd.Content[start:end]

		docs = append(docs, &kd)
	}
	return docs
}
