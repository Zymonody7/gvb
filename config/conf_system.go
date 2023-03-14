package config

import "fmt"

type System struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Env  string `yaml:"env"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%v", s.Host, s.Port)
}
