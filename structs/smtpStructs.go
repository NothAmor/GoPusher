package pusherStruct

type SmtpRequestStruct struct {
	Host     string
	Account  string
	Password string
	Port     int    `default: 0`
	Nickname string `default: "GoPusher"`
	MailType string
	Sender   string
	SendTo   []string
	Title    string
	Content  string
}
