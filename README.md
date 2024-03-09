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
curl "http://localhost:8082/presign?s3path={s3path}

{s3path} as “bucket_name/key” (see ‘/‘ delimiter)

Example:
```
localhost:8082/presign?s3path=ejoliet-dummy/hello.html
```
### encode value

Go to S3 bucket, select the file to get the URL presigned, copy URI, remove the 's3://' part, then go [here](https://www.urlencoder.io) to encode it:

example:
```
roman-airflow-demo/airflow-dev/logs/dag_id=01_hello_world/run_id=manual__2023-05-18T01:08:39.382762+00:00/task_id=hello_world_bash_operator/attempt=1.log
```
becomes
```
roman-airflow-demo%2Fairflow-dev%2Flogs%2Fdag_id%3D01_hello_world%2Frun_id%3Dmanual__2023-05-18T01%3A08%3A39.382762%2B00%3A00%2Ftask_id%3Dhello_world_bash_operator%2Fattempt%3D1.log
```

Use this for the value of `s3path`.

# AWS deployment

See CloudFormation [stack](server-stack.yaml) for deployment.
```bash
aws cloudformation create-stack --stack-name S3PresignedURL --template-body file://./server-stack.yaml
```
