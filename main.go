package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sirupsen/logrus"
)

var (
	defaultFromAddr = os.Getenv("ALERTER_FROM")
	defaultFromName = os.Getenv("ALERTER_FROM_NAME")
	defaultToAddr = os.Getenv("ALERTER_TO")
	defaultToName = os.Getenv("ALERTER_TO_NAME")
	defaultSubject = os.Getenv("ALERTER_SUBJECT")
	apiKey = os.Getenv("ALERTER_SENDGRID_APIKEY")
)

func main() {
	toAddr := flag.String("to", defaultToAddr, "The To address. Read environment variable ALERTER_TO in default.")
	toName := flag.String("to-name", defaultToName, "The To name. Read environment variable ALERTER_TO_NAME in default.")
	fromAddr := flag.String("from", defaultFromAddr, "The From address. Read environment variable ALERTER_FROM in default.")
	fromName := flag.String("from-name", defaultFromName, "The From name. Read environment variable ALERTER_FROM_NAME in default.")

	subject := flag.String("subject", defaultSubject, "A subject of email. Read environment variable ALERTER_SUBJECT in default.")

	verbose := flag.Bool("verbose", false, "Verbose output.")
	exVerbose := flag.Bool("extra-verbose", false, "Verbose output. included API KEY.")
	dryrun := flag.Bool("dryrun", false, "Don't sent, only parse.")

	flag.Parse()

	if *verbose {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.ErrorLevel)
	}
	if *exVerbose {
		logrus.Info("api key: " + apiKey)
	}

	if *fromAddr == "" {
		logrus.Fatal("From address is not set. Please see -help.")
	}

	if *toAddr == "" {
		logrus.Fatal("To address is not set. Please see -help.")
	}

	if apiKey == "" {
		logrus.Fatal("Envronment variable ALERTER_SENDGRID_APIKEY was not set. Please set API KEY of SendGrid.")
	}

	logrus.Info("loading message...")

	message, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		logrus.Fatal("Failed to read message from stdin: " + err.Error())
	}

	logrus.Info("## message ##")
	logrus.Info("from address default: " + defaultFromAddr)
	logrus.Info("from name default: " + defaultFromName)
	logrus.Info("to address default: " + defaultToAddr)
	logrus.Info("to name default: " + defaultToName)
	logrus.Info("from address: " + *fromAddr)
	logrus.Info("from name: " + *fromName)
	logrus.Info("to address: " + *toAddr)
	logrus.Info("to name: " + *toName)
	logrus.Info("default subject: " + defaultSubject)
	logrus.Info("subject: " + *subject)
	logrus.Info("payload:")
	for _, line := range strings.Split(string(message), "\n") {
		logrus.Info(line)
	}
	logrus.Info("## message ##")

	from := mail.NewEmail(*fromName, *fromAddr)
	to := mail.NewEmail(*toName, *toAddr)
	content := mail.NewContent("text/plain", string(message))
	email := mail.NewV3MailInit(from, *subject, to, content)

	logrus.Info("construct...")

	if *dryrun {
		logrus.Info("dryrun")
	} else {
		if response, err := sendgrid.NewSendClient(apiKey).Send(email); err != nil {
			logrus.Fatal("Failed to send message: " + err.Error())
		} else if response.StatusCode != 202 {
			logrus.Fatalf("Failed to send message: code=%d: %s", response.StatusCode, response.Body)
		}

		logrus.Info("message sent")
	}
}
