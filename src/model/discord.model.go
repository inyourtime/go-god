package model

type DiscordErrorLog struct {
	Level     string `json:"level"`
	Caller    string `json:"caller"`
	Message   string `json:"msg"`
	Timestamp string `json:"timestamp"`
}

type ServerEnvironment struct {
	Hostname string
	Url      string
	Method   string
}
