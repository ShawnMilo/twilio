/* package Twilio allows easy use of the Twilio restful API for sending SMS messages.

https://www.twilio.com/docs/api/rest/sending-messages

*/

package twilio

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

var URL, SID, TOKEN, FROM string

// init reads environment variables for configuration.
func init() {
	SID = os.Getenv("TWILIO_SID")
	TOKEN = os.Getenv("TWILIO_TOKEN")
	URL = "https://api.twilio.com/2010-04-01/Accounts/" + SID + "/Messages"
	FROM = os.Getenv("TWILIO_NUMBER")
}

// Send sends a text SMS via Twilio.
func Send(toNum, msg string) error {
	// HTTP POST values
	values := url.Values{}
	values.Set("From", FROM)
	values.Set("To", toNum)
	values.Set("Body", msg)

	// Create and configure a request instance.
	req, err := http.NewRequest("POST", URL, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	req.SetBasicAuth(SID, TOKEN)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send the request.
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()

	return nil
}
