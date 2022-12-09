package model

type Config struct {
	APP struct {
		NAME string `yaml:"NAME"`
		PORT string `yaml:"PORT"`
	} `yaml:"APP"`
	DATABASE struct {
		HOST     string `yaml:"HOST"`
		PORT     int    `yaml:"PORT"`
		USER     string `yaml:"USER"`
		PASSWORD string `yaml:"PASSWORD"`
	} `yaml:"DATABASE"`
	SECRET struct {
		KEY    string `yaml:"KEY"`
		EXPIRE int    `yaml:"EXPIRE"`
	} `yaml:"SECRET"`
}
