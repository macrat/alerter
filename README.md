Alerter
=======

Simple email sender using SendGrid for server watching script.

```
$ go get github.com/macrat/alerter

$ export ALERTER_SENDGRID_APIKEY="your api key"

$ echo hello world | alerter -to to-address@example.com -from from@example.com -subject hello
```
