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
	location := doc.Find("span.job_position").First().Text()
	education := doc.Find("span.job_academic").First().Text()
	salary := doc.Find("span.job_money").First().Text()
	low, high := libs.ParseSalary(salary, NumberMapper)
	description := doc.Find("div.job_detail").First().Find("span,p").Text()
	// replace \n
	description = strings.Replace(description, "\n", "", -1)

	comName := doc.Find("a.com-name").First().Text()
	comDesc := doc.Find(".com-desc").First().Text()

	item := engine.Item{
		Url:    url,
		Id:     id,
		Source: "shixiseng",
		Data: model.Job{
			Name:        name,
			Location:    location,
			Education:   education,
			Description: description,
			SalaryLow:   low,
			SalaryHigh:  high,
			Company: model.Company{
				Name:        comName,
				Description: comDesc,
			},
		},
	}

	result := engine.ParseResult{
		Items:    []engine.Item{item},
		Requests: []engine.Request{},
	}

	return result
}
