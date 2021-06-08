package configuration

type Configuration struct {
	LokiUrl  string `json:"lokiUrl"`
	User     string `json:"user"`
	Password string `json:"password"`
}
