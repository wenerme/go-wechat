package wechatc

var messageTypes = map[int]string{
	1:         "text",             // 文本
	3:         "image",            // 图像
	34:        "audio",            // 音频 - fileUrl 为 相对 url
	42:        "contact-card",     // 名片 - content 为 xml - 区分个人和公众号
	43:        "video",            // 视频
	47:        "emoji",            // 图片表情,mmemoticon - 表情包,动画表情 - linkUrl 包含 host 但没有 schema
	48:        "position",         // 定位
	50:        "call-end",         // 语音视频聊天结束
	495:       "link",             // 链接 - linkUrl, content 可能为标题
	496:       "file",             // 文件 - 有 fileUrl 为 相对 url , content 为 文件名
	501:       "video-call-start", // 视频通话发起
	502:       "audio-call-start", // 语音通话发起
	4919:      "chat-history",     // 聊天记录 - linkUrl 为链接 content 为标题
	4933:      "mini-program",     // 小程序
	10000:     "system",           // 系统消息
	16777265:  "share-content",    // 分享文本内容
	419430449: "transfer",         // 转账
	436207665: "red-envelope",     // 红包
	855638065: "group-live",       // 群直播
	570425393: "join-group",       // 被人拉进群聊

	// 1048625:   "anime-emoji", // 动画表情
	// 位置共享	开始共享：-1879048186  结束共享：10000
	// 标题和链接一起得消息	16777265
	// 特殊消息类型	以上参数值之外的为特殊消息类型，不用处理
	// 922746929 ?
	// @ 结束后有 4/MSP U+2005 Four-Per-Em Space
	// match(/@(?<name>[^\u2005]+)\u2005/).groups.name
	// 10000 移出群聊也是这个 type
}

func MessageTypeString(t int) string {
	v, ok := messageTypes[t]
	if !ok {
		return ""
	}
	return v
}
