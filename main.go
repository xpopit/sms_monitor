package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"sms_server/cls"
	"sms_server/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cls.LoadENV()
	rdb := db.Init()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/:service/:id", func(c *fiber.Ctx) error {
		body := c.Body()
		values, err := url.ParseQuery(string(body))
		if err != nil {
			fmt.Println("Error parsing query string:", err)
		}
		// log.Println(values)
		result := make(map[string]string)
		for key, value := range values {
			if len(value) > 0 {
				result[key] = value[0] // take the first value if there are multiple
			}
		}
		// if v, ok := result["title"]; ok {
		// 	log.Println(v)
		// }
		if _, ok := result["text"]; ok {
			// _bank := strings.ToUpper(c.Params("service"))
			// fmt.Println(cls.Call(_bank, v))
			// log.Println(v)
			log.Println(c.Params("service"))
			log.Println(result)
			log.Println(c.Params("id"))
			_m := map[string]interface{}{
				"service": c.Params("service"),
				"id":      c.Params("id"),
				"result":  result,
			}
			b, err := json.Marshal(_m)
			if err != nil {
				fmt.Println("error:", err)
			}
			base64Str := base64.StdEncoding.EncodeToString(b)
			// log.Println("base64Str:", base64Str)
			rdb.RDB_Pub("sms", base64Str)
		}
		return c.SendString("ok")
	})

	app.Listen(":13000")
	// data := `name=Messages&pkg=com.google.android.apps.messaging&title=Krungsri&text=%E0%B8%A1%E0%B8%B5%E0%B9%80%E0%B8%87%E0%B8%B4%E0%B8%99%E0%B8%AD%E0%B8%AD%E0%B8%81xxx105469x%20%E0%B8%9C%E0%B9%88%E0%B8%B2%E0%B8%99%20Online%201.00%20%E0%B9%80%E0%B8%AB%E0%B8%A5%E0%B8%B7%E0%B8%AD%20%2051.63%20%2827%2F6%2F67%2C13%3A39%29&subtext=my&bigtext=&infotext=`
	// data2 := `name=KMA&pkg=com.krungsri.kma&title=krungsri%20app&text=%E0%B8%A1%E0%B8%B5%E0%B9%80%E0%B8%87%E0%B8%B4%E0%B8%99%E0%B8%AD%E0%B8%AD%E0%B8%81%20XXX-1-05469-X%20%E0%B8%88%E0%B8%B3%E0%B8%99%E0%B8%A7%E0%B8%99%201.00%20%E0%B8%9A%E0%B8%B2%E0%B8%97&subtext=&bigtext=%E0%B8%A1%E0%B8%B5%E0%B9%80%E0%B8%87%E0%B8%B4%E0%B8%99%E0%B8%AD%E0%B8%AD%E0%B8%81%20XXX-1-05469-X%20%E0%B8%88%E0%B8%B3%E0%B8%99%E0%B8%A7%E0%B8%99%201.00%20%E0%B8%9A%E0%B8%B2%E0%B8%97&infotext=`
	// datas := []string{data, data2}``
	// for _, v := range datas {
	// 	decodedString, err := url.QueryUnescape(v)
	// 	if err != nil {
	// 		fmt.Println("Error decoding:", err)
	// 		return
	// 	}
	// 	// fmt.Println(decodedString)
	// 	values, err := url.ParseQuery(decodedString)
	// 	if err != nil {
	// 		fmt.Println("Error parsing query string:", err)
	// 		return
	// 	}

	// 	// วนลูปและพิมพ์ค่าของคีย์และค่าต่าง ๆ
	// 	for key, value := range values {
	// 		fmt.Printf("%s = %s\n", key, value[0])
	// 		if key == "text" {

	// 		}
	// 	}
	// }

}
