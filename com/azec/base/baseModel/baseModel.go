package baseModel

type RequestLov struct {
	Insearch string `json:"inSearch"`
}

type BaseResponseModel struct {
	OutError   int64  `json:"outError"`
	OutMessage string `json:"outMessage"`
}

type DatabaseModel struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}
