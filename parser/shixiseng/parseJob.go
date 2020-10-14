package shixiseng

import (
	"github.com/PuerkitoBio/goquery"
	"job-hunter/engine"
	"job-hunter/libs"
	"job-hunter/model"
	"regexp"
	"strings"
)

func ParseJob(doc *goquery.Document, url string) engine.ParseResult {

	idRe := regexp.MustCompile(`.*/intern/([^?]+)\?pcm.*`)
	id := libs.ExtractString([]byte(url), idRe)

	name := doc.Find("div.new_job_name>span").First().Text()
	location := doc.Find("span.com_position").First().Text()
	education := doc.Find("span.job_academic").First().Text()
	salary := doc.Find("span.job_money").First().Text()
	low, high := libs.ParseSalary(salary, ConvertNumber)
	description := doc.Find("div.job_detail").First().Find("span,p").Text()
	// replace \n
	description = strings.Replace(description, "\n", "", -1)

	item := engine.Item{
		Url: url,
		Id:  id,
		Data: model.Job{
			Name:        name,
			Location:    location,
			Education:   education,
			Description: description,
			SalaryLow:   low,
			SalaryHigh:  high,
		},
	}

	result := engine.ParseResult{
		Items:    []engine.Item{item},
		Requests: []engine.Request{},
	}

	comUel, ok := doc.Find("a.com-name").First().Attr("href")
	if ok {
		result.Requests = append(result.Requests, engine.Request{
			Url:       comUel,
			ParseFunc: engine.NilParser,
		})
	}
	return result
}

func getParseJob(id string) engine.ParseFunc {
	return ParseJob
}
