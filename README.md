# Overview
S3 Presigned URL microservice prototype written in Go.
Using AWS SDK in Go.

## Intro
Returns the presigned URL for downalod with 15 mins expiration.
Service listen to 8082 (changeable)

Effectively returns same as
“aws s3 presign bucket_name/key --expires-in 60*15”

# Installation

go run main.go

or:
go build main.go

./main &

## URL pattern
curl "http://localhost:8080/presign?s3path={s3path}

{s3path} as “bucket_name/key” (see ‘/‘ delimiter)

# AWS deployment

See CloudFormation [stack](server-stack.yaml) for deployment.

# Extra - ICS file generated

Added bonus: service to create ICS file and QR to add calendar event on the fly, just with title and time (optional with link).
See [ics.go](ics.go)
