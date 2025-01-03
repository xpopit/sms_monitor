package cls

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadENV() {
	// log.SetFlags(log.Lshortfile | log.LstdFlags)
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Fprintf(os.Stderr, "Unable to identify current directory (needed to load .env.test)")
		os.Exit(1)
	}
	basepath := filepath.Dir(file)
	err := godotenv.Load(filepath.Join(basepath, "../.env"))
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Error loading .env file")
	}
	if os.Getenv("ENV") == "development" {
		// log.SetFlags(log.L/stdFlags | log.Llongfile)
		log.SetFlags(log.Ltime | log.Lshortfile)
		// log.SetFlags(0)
		// log.SetOutput(new(logWriter))
	}
	// zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
	// 	_str := ""
	// 	for i, v := range strings.Split(file, "/") {
	// 		if v == "agent-api-golang" {
	// 			_str = strings.Join(strings.Split(file, "/")[i+1:], "/")
	// 		}
	// 	}
	// 	return _str + ":" + strconv.Itoa(line)
	// }

	// multi := zerolog.MultiLevelWriter(os.Stdout)
	// zerolog.New(multi).With().Timestamp().Logger()

	fmt.Println("Init cls")
}
