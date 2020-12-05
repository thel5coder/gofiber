package twilio

import "github.com/sfreiberg/gotwilio"

type Client struct {
	sid           string
	token         string
	DefaultSender string
	twilio        *gotwilio.Twilio
}

func NewTwilioClient(sid, token, defaultSender string) *Client {
	return &Client{
		sid:           sid,
		token:         token,
		DefaultSender: defaultSender,
		twilio:        gotwilio.NewTwilioClient(sid, token),
	}
}

func (client Client) SendSMS(from, to, message string) (err error) {
	_, _, err = client.twilio.SendSMS(from, to, message, "", "")
	if err != nil {
		return err
	}

	return nil
}
