Alerter
=======

Simple email sender using SendGrid for server watching script.

## Usage

```
$ go get github.com/macrat/alerter

$ export ALERTER_SENDGRID_APIKEY="your api key"

$ echo hello world | alerter -to to-address@example.com -from from@example.com -subject hello
```

## Command-line Options / Environment Variables

|option          |environment variable     |description                                  |
|----------------|-------------------------|---------------------------------------------|
|                |ALERTER\_SENDGRID\_APIKEY|API Key of SendGrid. This is Required.       |
|-to             |ALERTER\_TO              |To address.                                  |
|-to-name        |ALERTER\_TO\_NAME        |To name.                                     |
|-from           |ALERTER\_FROM            |From address.                                |
|-from-name      |ALERTER\_FROM\_NAME      |From name.                                   |
|-subject        |ALERTER\_SUBJECT         |The subject of email.                        |
|-help           |                         |Show usage.                                  |
|-verbose        |                         |Show verbose debug messages.                 |
|-extra-verbose  |                         |Show verbose debug messages included API key.|
|-dryrun         |                         |Don't send, only parse options.              |
