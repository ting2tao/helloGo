package channel

import "fmt"

/*   把channel用在数据流动的地方：
消息传递、消息过滤
信号广播
事件订阅与广播
请求、响应转发
任务分发
结果汇总
并发控制
同步与异步
*/

type Job struct {
	Name string
	ID   int
}
type Handler struct {
	JobCh  chan *Job
	StopCh chan int
}

// 分配job时，如果收到关闭的通知则退出，不分配job
func (h *Handler) Handle(job *Job) {
	select {
	case h.JobCh <- job:
		ch := <-h.JobCh
		id := ch.ID
		name := ch.Name
		fmt.Println(id)
		fmt.Println(name)
		return
	case <-h.StopCh:
		fmt.Println("stop")
		return
	}
}
