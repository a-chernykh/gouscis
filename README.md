gouscis is a little daemon which I've created to get familiar with [go](http://golang.org/). It will periodically check if [USCIS case status](https://egov.uscis.gov/cris/Dashboard/CaseStatus.do) for given receipt number have changed. If it was changed, it will send e-mail to given address.

# Install

```sh
make install
```

# Run

```sh
export EMAIL=my@email.com
export SMTP_SERVER=127.0.0.1:25
export CASE_NUMBER=MSC0000000000

gouscis
```
