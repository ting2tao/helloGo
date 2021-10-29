package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	nowMain := time.Now()
	nums := 500000
	ch := make(chan struct{})
	go goWaitG(nums, ch)
	goNotWaitG(nums)
	<-ch
	fmt.Println("总共过去了：", time.Since(nowMain))
}

func goNotWaitG(nums int) {
	now := time.Now()
	fmt.Println(now)

	houseMap := map[string]string{}
	projectMap := map[string]string{}

	houses, err := GetHouses(nums)
	if err != nil {
		return
	}
	for _, house := range houses {
		houseMap[house.Code] = house.Name
	}

	projects, err := GetProjects(nums)
	if err != nil {
		return
	}
	for _, project := range projects {
		projectMap[project.Code] = project.Name
	}

	fmt.Println("等等他们执行啊Not...")
	fmt.Println(len(houseMap))
	fmt.Println(len(projectMap))
	fmt.Println("过去了多久：", time.Since(now))
}
func goWaitG(nums int, ch chan struct{}) {
	now := time.Now()
	fmt.Println(now)
	wg := sync.WaitGroup{}
	wg.Add(2)
	houseMap := map[string]string{}
	projectMap := map[string]string{}
	go func() {
		defer wg.Done()
		houses, err := GetHouses(nums)
		if err != nil {
			return
		}
		for _, house := range houses {
			houseMap[house.Code] = house.Name
		}
	}()
	go func() {
		defer wg.Done()
		projects, err := GetProjects(nums)
		if err != nil {
			return
		}
		for _, project := range projects {
			projectMap[project.Code] = project.Name
		}
	}()
	fmt.Println("等等他们执行啊Wait...")
	wg.Wait()
	fmt.Println(len(houseMap))
	fmt.Println(len(projectMap))
	fmt.Println("过去了多久：", time.Since(now))
	ch <- struct{}{}
}

type House struct {
	Code string
	Name string
}

func GetHouses(nums int) ([]House, error) {
	houses := make([]House, 0, nums)
	for i := 0; i < nums; i++ {
		houses = append(houses, House{
			Code: fmt.Sprintf("%s%d", "编码", i),
			Name: fmt.Sprintf("%s%d", "房间", i),
		})
	}
	return houses, nil
}

type Project struct {
	Code string
	Name string
}

func GetProjects(nums int) ([]Project, error) {
	projects := make([]Project, 0, nums)
	for i := 0; i < nums; i++ {
		projects = append(projects, Project{
			Code: fmt.Sprintf("%s%d", "编码", i),
			Name: fmt.Sprintf("%s%d", "项目", i),
		})
	}
	return projects, nil
}
