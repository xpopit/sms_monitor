package cls

import (
	"io/ioutil"
	"log"
	"os"
	"sms_server/models"

	"gopkg.in/yaml.v2"
)

func LoadRegexBank(BankName string) (regex *models.Bank) {
	// log.Println()
	data, err := ioutil.ReadFile(os.Getenv("REGEX_YAML_PATH"))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	var regexConfig models.RegexConfig
	err = yaml.Unmarshal(data, &regexConfig)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for _, v := range regexConfig.Regex.Banks {
		if v.Name == BankName {
			regex = &v
		}
	}
	return
}
