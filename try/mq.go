package main

import (
	"fmt"
	"github.com/gogap/errors"
	"strings"
	"time"

	"github.com/aliyunmq/mq-http-go-sdk"
)

func main() {

	// 设置HTTP协议客户端接入点，进入消息队列RocketMQ版控制台实例详情页面的接入点区域查看。
	endpoint := ""
	// AccessKey ID阿里云身份验证，在阿里云RAM控制台创建。
	accessKey := ""
	// AccessKey Secret阿里云身份验证，在阿里云RAM控制台创建。
	secretKey := ""
	// 消息所属的Topic，在消息队列RocketMQ版控制台创建。
	//不同消息类型的Topic不能混用，例如普通消息的Topic只能用于收发普通消息，不能用于收发其他类型的消息。
	topic := ""
	// Topic所属的实例ID，在消息队列RocketMQ版控制台创建。
	// 若实例有命名空间，则实例ID必须传入；若实例无命名空间，则实例ID传入null空值或字符串空值。实例的命名空间可以在消息队列RocketMQ版控制台的实例详情页面查看。
	instanceId := ""
	// 您在控制台创建的Group ID。
	groupId := ""

	client := mq_http_sdk.NewAliyunMQClient(endpoint, accessKey, secretKey, "")

	mqConsumer := client.GetConsumer(instanceId, topic, groupId, "")

	for {
		endChan := make(chan int)
		respChan := make(chan mq_http_sdk.ConsumeMessageResponse)
		errChan := make(chan error)
		go func() {
			select {
			case resp := <-respChan:
				{
					// 处理业务逻辑。
					var handles []string
					fmt.Printf("Consume %d messages---->\n", len(resp.Messages))
					for _, v := range resp.Messages {
						handles = append(handles, v.ReceiptHandle)
						fmt.Printf("\tMessageID: %s, PublishTime: %d, MessageTag: %s\n"+
							"\tConsumedTimes: %d, FirstConsumeTime: %d, NextConsumeTime: %d\n"+
							"\tBody: %s\n"+
							"\tProps: %s\n",
							v.MessageId, v.PublishTime, v.MessageTag, v.ConsumedTimes,
							v.FirstConsumeTime, v.NextConsumeTime, v.MessageBody, v.Properties)
					}

					// NextConsumeTime前若不确认消息消费成功，则消息会被重复消费。
					// 消息句柄有时间戳，同一条消息每次消费拿到的都不一样。
					ackerr := mqConsumer.AckMessage(handles)
					if ackerr != nil {
						// 某些消息的句柄可能超时，会导致消息消费状态确认不成功。
						fmt.Println(ackerr)
						if errAckItems, ok := ackerr.(errors.ErrCode).Context()["Detail"].([]mq_http_sdk.ErrAckItem); ok {
							for _, errAckItem := range errAckItems {
								fmt.Printf("\tErrorHandle:%s, ErrorCode:%s, ErrorMsg:%s\n",
									errAckItem.ErrorHandle, errAckItem.ErrorCode, errAckItem.ErrorMsg)
							}
						} else {
							fmt.Println("ack err =", ackerr)
						}
						time.Sleep(time.Duration(3) * time.Second)
					} else {
						fmt.Printf("Ack ---->\n\t%s\n", handles)
					}

					endChan <- 1
				}
			case err := <-errChan:
				{
					// Topic中没有消息可消费。
					if strings.Contains(err.(errors.ErrCode).Error(), "MessageNotExist") {
						fmt.Println("\nNo new message, continue!")
					} else {
						fmt.Println(err)
						time.Sleep(time.Duration(3) * time.Second)
					}
					endChan <- 1
				}
			case <-time.After(35 * time.Second):
				{
					fmt.Println("Timeout of consumer message ??")
					endChan <- 1
				}
			}
		}()

		// 长轮询消费消息，网络超时时间默认为35s。
		// 长轮询表示如果Topic没有消息，则客户端请求会在服务端挂起3s，3s内如果有消息可以消费则立即返回响应。
		mqConsumer.ConsumeMessage(respChan, errChan,
			3, // 一次最多消费3条（最多可设置为16条）。
			3, // 长轮询时间3s（最多可设置为30s）。
		)
		<-endChan
	}
}
