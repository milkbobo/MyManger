package account

type Account struct {
	AccountId  int `xorm:"autoincr"`
	UserId     int
	Name       string
	Money      int
	Remark     string
	CategoryId int
	CardId     int
	Type       int
	CreateTime string `xorm:"created"`
	ModifyTime string `xorm:"updated"`
}

type Accounts struct {
	Count int
	Data  []Account
}

type WeekTypeStatistic struct {
	CardId     int
	CardName   string
	Money      int
	Name       string
	Type       int
	TypeName   string
	Week       int
	Year       int
	CreateTime string
	ModifyTime string
}