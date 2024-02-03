package config

type Env string

const (
	Prod     Env = "prod"
	Stagging Env = "stagging"
	Dev      Env = "dev"
	Test     Env = "test"
)

type Config interface {
	GetEnv() Env
	GetPort() uint
	GetDSN() string
}
