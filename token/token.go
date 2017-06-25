package token

import (
	"fmt"
	"net/url"
	"bytes"
	"encoding/json"
	"github.com/codegangsta/cli"
	"net/http"
	"io/ioutil"
)

func CmdToken(c *cli.Context) {

	clientId := c.String("client_id")
	clientSecret := c.String("client_secret")
	username := c.String("username")
	password := c.String("password")

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
	fmt.Println("Bearer " + tokenResponse.AccsessToken)
}
