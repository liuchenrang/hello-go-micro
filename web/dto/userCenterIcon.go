package dto

type UserCenterIcon struct {
	Ico string `json:"icon"`
	Title string `json:"title"`
	TitleClr string `json:"titleClr"`
	SubTitle string `json:"subtitle"`
	SubTitleClr string `json:"subtitleClr"`
	Link string `json:"link"`
	IsJumpAli string `json:"isJumpAli"`
	IsOauth string `json:"isOauth"`
	IsUpgrade string `json:"isUpgrade"`
	IsLogin string `json:"isLogin"`
}
