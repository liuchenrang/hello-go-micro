package config

type Host struct {
	Address string `json:"address"`
	Port int `json:"port"`
}

type Config struct{
	Hosts map[string]Host `json:"hosts"`
}
