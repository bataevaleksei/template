package ping

type Ping struct {
	Message string `json:"message"`
}

func Get() *Ping {
	return &Ping{
		Message: "pong",
	}
}
