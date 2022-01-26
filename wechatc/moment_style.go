package wechatc

var momentStyleType = map[int]string{
	1:  "image",
	2:  "images",
	3:  "link",  // contentUrl, title
	15: "video", // mediaUrl
	// 28 - unknown
}

func MomentStyleTypeString(t int) string {
	v, ok := momentStyleType[t]
	if !ok {
		return ""
	}
	return v
}
