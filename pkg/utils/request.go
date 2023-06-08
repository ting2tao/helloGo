package utils

import (
	"github.com/parnurzeal/gorequest"
	"time"
)

func Get(targetUrl string) *gorequest.SuperAgent {
	agent := gorequest.New().Get(targetUrl)
	agent.Timeout(time.Second * 30)
	return agent
}

func Post(targetUrl string) *gorequest.SuperAgent {
	agent := gorequest.New().Post(targetUrl)
	agent.Timeout(time.Second * 30)
	return agent
}

func Put(targetUrl string) *gorequest.SuperAgent {
	agent := gorequest.New().Put(targetUrl)
	agent.Timeout(time.Second * 30)
	return agent
}

func Patch(targetUrl string) *gorequest.SuperAgent {
	agent := gorequest.New().Patch(targetUrl)
	agent.Timeout(time.Second * 30)
	return agent
}

func Delete(targetUrl string) *gorequest.SuperAgent {
	agent := gorequest.New().Delete(targetUrl)
	agent.Timeout(time.Second * 30)
	return agent
}
