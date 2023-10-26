package ch100_redis

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// redis list示例: 消息聚合推送
type message struct {
	Id      int    // 消息唯一id
	Type    int    // 消息类型
	Content string // 消息内容
	UserId  string // 用户id
}

const (
	msgPushFlagListKey    = "msg_push_flag_list" // 推送消息标志位列表
	msgPushListItemFormat = "%s:%d,%d"           // 推送消息标志位列表项 (userId:msgType:timestamp)
	msgContentKeyFormat   = "msg:%s:%d"          // 推送消息内容 (msg:userId:msgType)
	msgAggDuration        = 3600                 // 消息聚合时长 (秒)
)

// 消息保存
func saveMsg(msg *message) {
	msgContentKey := fmt.Sprintf(msgContentKeyFormat, msg.UserId, msg.Type)

	// 判断该用户该类型消息是否存在
	n, _ := client.Exists(ctx, msgContentKey).Result()
	if exist := n == 1; !exist {
		// 消息不存在
		// 保存消息内容为string类型
		bytes, _ := json.Marshal(msg)
		client.Set(ctx, msgContentKey, string(bytes), 2*msgAggDuration)

		// 添加标志位到推送消息列表
		itemVal := fmt.Sprintf(msgPushListItemFormat, msg.UserId, msg.Type, time.Now().Unix())
		client.RPush(ctx, msgPushFlagListKey, itemVal)
	}
}

// 轮询消息推送
func pollMsg() {
	for {
		itemVal, _ := client.LPop(ctx, msgPushFlagListKey).Result()
		slice := strings.Split(itemVal, ":")
		userId := slice[0]
		msgType, _ := strconv.ParseInt(slice[1], 10, 64)
		ts, _ := strconv.ParseInt(slice[2], 10, 64)

		if time.Now().Unix() >= ts+msgAggDuration {
			// 聚合时间达到推送消息
			msgContentKey := fmt.Sprintf(msgContentKeyFormat, userId, msgType)
			msgVal, _ := client.Get(ctx, msgContentKey).Result()

			var msg message
			json.Unmarshal([]byte(msgVal), &msg)
			pushMsg(msg)

			client.Del(ctx, msgContentKey)
		} else {
			// 聚合时间未到,重新添加到list中
			client.LPush(ctx, msgPushFlagListKey, itemVal)
			time.Sleep(time.Second * 30)
		}
	}
}

// 推送消息
func pushMsg(msg message) {}
