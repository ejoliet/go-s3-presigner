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
