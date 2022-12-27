package Email

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//var str1 string = "suryagarimella@beautifulcode.in"

func SendEmail(data string) {
	fmt.Println(data)
	url := "https://rapidprod-sendgrid-v1.p.rapidapi.com/mail/send"

	str := fmt.Sprintf(`{
		"personalizations": [
			{
				"to": [
					{
						"email":  "%s"
					}
				],
				"subject": "Linkedin"
			}
		],
		"from": {
			"email": "hemanth.kakumanu@beautifulcode.in"
		},
		"content": [
			{
				"type": "text/plain",
				"value": "Hey you got a new connection "
			}
		]
	}`, data)

	payload := strings.NewReader(str)
	req, _ := http.NewRequest("POST", url, payload)
	rapid_key := os.Getenv("X-RapidAPI-Key")

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", rapid_key)
	req.Header.Add("X-RapidAPI-Host", "rapidprod-sendgrid-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(str)
	fmt.Println(string(body))

}
