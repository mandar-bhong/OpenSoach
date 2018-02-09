package models

type ConfigSettings struct {
	ServerConfig   ServerSettings   `json:"serverconfig"`
	EndpointConfig EndpointSettings `json:"endpointconfig"`
	WebConfig      WebSettings      `json:"webconfig"`
	LoggerConfig   LoggerSettings   `json:"logger"`
}

type ServerSettings struct {
}

type EndpointSettings struct {
	WebSocketPort int `json:"websocketport"`
}

type WebSettings struct {
	ServicePort    int `json:"webport"`
	SessionTimeOut int `json:"sessiontimeout"`
}

type LoggerSettings struct {
	Filename   string `json:"filename"`
	MaxSize    int    `json:"maxfilesize"`
	MaxBackups int    `json:"maxbackups"`
	MaxAge     int    `json:"maxageindays"`
	Level      string `json:"loglevel"`
}
