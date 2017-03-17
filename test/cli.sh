#!/bin/sh

# test
STUBHUB_USER_NAME="bruceadowns@gmail.com" STUBHUB_PASSWORD="foobar" STUBHUB_CUSTOMER_KEY="customerkey" STUBHUB_CUSTOMER_SECRET="customersecret" go run main.go

# or via export
export STUBHUB_APPLICATION_TOKEN=toketoke
export STUBHUB_USER_NAME=bruceadowns@gmail.com
export STUBHUB_PASSWORD=foobar
export STUBHUB_CUSTOMER_KEY=customerkey
export STUBHUB_CUSTOMER_SECRET=customersecret
go run main.go

# cli example usages

go run main.go -h
go run main.go -contexts foobar.yaml -cookies .foobar.yaml -context gostubhub -debug events 'Wynton Marsalis'
cat contexts.yaml | go run main.go -contexts -

go run main.go -context gostubhub -debug events 'Wynton Marsalis'
go run main.go -context gostubhub -debug listings 'STATUS:ACTIVE'

cat contexts.yaml | go run main.go -contexts - -context gostubhub -debug events 'Wynton Marsalis'
cat contexts.yaml | go run main.go -contexts - -context gostubhub -debug listings 'STATUS:ACTIVE'

STUBHUB_PASSWORD=foobar cat contexts.yaml | go run main.go -cookies .foo.yaml -contexts - -context gostubhub -debug listings 'STATUS:ACTIVE'
