# go-s3-presigner
Presigned URL microservice prototype

Effectively returns same as
“aws s3 presign bucket_name/key --expires-in 60*15”


URL:
curl "http://localhost:8080/presign?s3path={s3path}

{s3path} as “bucket_name/key” (see ‘/‘ delimiter)
go run main.go

or:
go build main.go

./main &


See stack for deployment.