package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	defaultFromAddr = os.Getenv("ALERTER_FROM")
	defaultFromName = os.Getenv("ALERTER_FROM_NAME")
	defaultToAddr = os.Getenv("ALERTER_TO")
	defaultToName = os.Getenv("ALERTER_TO_NAME")
	apiKey = os.Getenv("ALERTER_SENDGRID_APIKEY")
)

func main() {
	toAddr := flag.String("to", defaultToAddr, "The To address. Read environment variable ALERTER_TO in default.")
	toName := flag.String("to-name", defaultToName, "The To name. Read environment variable ALERTER_TO_NAME in default.")
	fromAddr := flag.String("from", defaultFromAddr, "The From address. Read environment variable ALERTER_FROM in default.")
	fromName := flag.String("from-name", defaultFromName, "The From name. Read environment variable ALERTER_FROM_NAME in default.")

	subject := flag.String("subject", "alerter", "A subject of email.")

	verbose := flag.Bool("verbose", false, "Verbose output.")
	exVerbose := flag.Bool("extra-verbose", false, "Verbose output. included API KEY.")


	flag.Parse()

	if *fromAddr == "" {
		log.Fatalln("From address is not set. Please see -help.")
	}

	if *toAddr == "" {
		log.Fatalln("To address is not set. Please see -help.")
	}

	if apiKey == "" {
		log.Fatalln("Envronment variable ALERTER_SENDGRID_APIKEY was not set. Please set API KEY of SendGrid.")
	}


	if *verbose || *exVerbose {
		log.Println("loading message...")
	}


	message, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln("Failed to read message from stdin: " + err.Error())
	}


	if *verbose || *exVerbose {
		log.Println("message:")

		if *exVerbose {
			log.Println("api key: " + apiKey)
		}

		log.Println("from address default: " + defaultFromAddr)
		log.Println("from name default: " + defaultFromName)
		log.Println("to address default: " + defaultToAddr)
		log.Println("to name default: " + defaultToName)
		log.Println("from address: " + *fromAddr)
		log.Println("from name: " + *fromName)
		log.Println("to address: " + *toAddr)
		log.Println("to name: " + *toName)
		log.Println("subject: " + *subject)
		log.Println("payload:")
		for _, line := range strings.Split(string(message), "\n") {
			log.Println(line)
		}
	}


	from := mail.NewEmail(*fromName, *fromAddr)
	to := mail.NewEmail(*toName, *toAddr)
	content := mail.NewContent("text/plain", string(message))
	email := mail.NewV3MailInit(from, *subject, to, content)


	if *verbose || *exVerbose {
		log.Println("construct...")
	}


	if response, err := sendgrid.NewSendClient(apiKey).Send(email); err != nil {
		log.Fatalln("Failed to send message: " + err.Error())
	} else if response.StatusCode != 202 {
		log.Fatalf("Failed to send message: code=%d: %s", response.StatusCode, response.Body)
	}

	if *verbose || *exVerbose {
		log.Println("message sent")
	}
}
