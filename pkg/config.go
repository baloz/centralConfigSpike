package pkg

type Config struct {
	AppName string
}

func NewConfig() {
	// readUsingConsule()
	readUsingSpringConfig()
}
