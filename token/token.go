package token

import (
	"fmt"
	"net/url"
	"bytes"
	"encoding/json"
	"github.com/codegangsta/cli"
	"net/http"
	"io/ioutil"
	"log"
)

func CmdToken(c *cli.Context) {

	username, password, clientId, clientSecret := checkFlags(c)

	request_url := fmt.Sprintf(IDP_ENDPOINT, clientId, clientSecret)

	form := url.Values{
		"username":   {username},
		"password":   {password},
		"grant_type": {"password"},
	}

	body := bytes.NewBufferString(form.Encode())
	rsp, err := http.Post(request_url, "application/x-www-form-urlencoded", body)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()
	body_byte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	tokenResponse := TokeResponse{}
	err = json.Unmarshal(body_byte, &tokenResponse)
	if err != nil {
		fmt.Println("Unable to get access token.")
	}

	fmt.Println("Put this in the Authorization header in you browser:\n\n")
	fmt.Printf("Bearer %\n\n", tokenResponse.AccsessToken)
}

func checkFlags(c *cli.Context) (string, string, string, string) {

	var username, password, clientId, clientSecret string

	if len(c.String(FLAG_CLIENT_ID)) > 0 {
		clientId = c.String(FLAG_CLIENT_ID)
	} else {
		log.Fatal(fmt.Sprintf("%s is missing", FLAG_CLIENT_ID))
	}

	if len(c.String(FLAG_CLIENT_SECRET)) > 0 {
		clientSecret = c.String(FLAG_CLIENT_SECRET)
	} else {
		log.Fatal(fmt.Sprintf("%s is missing", FLAG_CLIENT_SECRET))
	}
	if len(c.String(FLAG_USERNAME)) > 0 {
		username = c.String(FLAG_USERNAME)
	} else {
		log.Fatal(fmt.Sprintf("%s is missing", FLAG_USERNAME))
	}
	if len(c.String(FLAG_PASSWORD)) > 0 {
		password = c.String(FLAG_PASSWORD)
	} else {
		log.Fatal(fmt.Sprintf("%s is missing", FLAG_PASSWORD))
	}

	return username, password, clientId, clientSecret
}
