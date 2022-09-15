package redis

type config struct {
	port string
}

func NewConfig() *config {
	return &config{
		port: ":6379",
	}
}
