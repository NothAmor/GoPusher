package pusherStruct

type ServerChanRequestStruct struct {
	Key     string
	Title   string
	Desp    string `default: ""`
	Short   string `default: ""`
	Channel string `default: ""`
	Openid  string `default: ""`
}

type ServerChanPushSuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Pushid  string `json:"pushid"`
		Readkey string `json:"readkey"`
		Error   string `json:"error"`
		Errno   int    `json:"errno"`
	} `json:"data"`
}
