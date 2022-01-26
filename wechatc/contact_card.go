package wechatc

import "encoding/xml"

type ContactCardContent struct {
	XMLName                 xml.Name `xml:"msg"`
	Text                    string   `xml:",chardata"`
	Head                    string   `xml:"head,attr"`
	BigHeadImgURL           string   `xml:"bigheadimgurl,attr"`
	SmallHeadImgURL         string   `xml:"smallheadimgurl,attr"`
	Username                string   `xml:"username,attr"`
	Nickname                string   `xml:"nickname,attr"`
	FullPinyin              string   `xml:"fullpy,attr"`
	ShortPinyin             string   `xml:"shortpy,attr"`
	Alias                   string   `xml:"alias,attr"`
	ImageStatus             string   `xml:"imagestatus,attr"`
	Scene                   string   `xml:"scene,attr"`
	Province                string   `xml:"province,attr"`
	City                    string   `xml:"city,attr"`
	Sign                    string   `xml:"sign,attr"`
	Sex                     string   `xml:"sex,attr"`
	CertFlag                string   `xml:"certflag,attr"`
	CertInfo                string   `xml:"certinfo,attr"`
	BrandIconURL            string   `xml:"brandIconUrl,attr"`
	BrandHomeURL            string   `xml:"brandHomeUrl,attr"`
	BrandSubscriptConfigURL string   `xml:"brandSubscriptConfigUrl,attr"`
	BrandFlags              string   `xml:"brandFlags,attr"`
	RegionCode              string   `xml:"regionCode,attr"`
	AntiSpamTicket          string   `xml:"antispamticket,attr"`
}

func ParseContactCard(s string) (ContactCardContent, error) {
	var c ContactCardContent
	return c, xml.Unmarshal([]byte(s), &c)
}
