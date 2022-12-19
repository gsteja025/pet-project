package Email

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "004470f396msh422de6c41ab61eap1a3894jsn9d0e034003a3")
	req.Header.Add("X-RapidAPI-Host", "rapidprod-sendgrid-v1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(str)
	fmt.Println(string(body))

}
