package dto

type SiginResponse struct {
	Id          string  `json:"id"`
	Firstname   string  `json:"firstname"`
	Lastname    string  `json:"lastname"`
	Email       string  `json:"email"`
	Mobile      string  `json:"mobile"`
	Username    string  `json:"username"`
	Rolename    string  `json:"rolename"`
	Isactivated int     `json:"isactivated"`
	Isblocked   int     `json:"isblocked"`
	Userpicture string  `json:"userpicture"`
	Qrcodeurl   *string `json:"qrcodeurl"`
	Token       string  `json:"token"`
}
