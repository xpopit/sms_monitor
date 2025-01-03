package models

type RegexConfig struct {
	Regex struct {
		Banks []Bank `yaml:"banks"`
	} `yaml:"regex"`
}

type Bank struct {
	Name    string `yaml:"name"`
	TestStr string `yaml:"test_str"`
	Regex   string `yaml:"regex"`
}
