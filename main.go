package main

import (
    "fmt"
    "strings"
    "log"
    "net/http"
    "time"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    //"github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/s3"
)

func main() {
    http.HandleFunc("/presign", presignHandler)
    log.Println("Server is running on port 8082...")
    log.Println("Try http://localhost:8082/presign?s3path=...")
    log.Fatal(http.ListenAndServe(":8082", nil))
}

func presignHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Unsupported method. Please use GET.", http.StatusMethodNotAllowed)
        return
    }

    s3Path := r.URL.Query().Get("s3path")
    if s3Path == "" {
        http.Error(w, "s3path query parameter is required", http.StatusBadRequest)
        return
    }

    presignedURL, err := generatePresignedURL(s3Path)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprint(w, presignedURL)
}

func generatePresignedURL(s3Path string) (string, error) {
    // Assume the S3 path is in the format "bucket/key"
    parts := strings.SplitN(s3Path, "/", 2)
    if len(parts) != 2 {
        return "", fmt.Errorf("invalid S3 path; expected format 'bucket/key'")
    }
    bucket, key := parts[0], parts[1]
    log.Println(s3Path)
    //if running in EC2, Go will take credentials from IAM role directly, n need to setup env variable
    // Initialize a session using the default credential provider chain
    //sess, err := session.NewSession()
    
    // awsAccessKeyID := ""
    // awsSecretAccessKey := ""
    // awsToken := ""
    //for testing locally or using sets of access key and secrets, use the expanded configuration to set individual keys

    sess, err := session.NewSession(&aws.Config{
    	Region:      aws.String("us-east-1"),
    //     //Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, awsToken),
    //     LogLevel:    aws.LogLevel(aws.LogDebugWithHTTPBody),
    })
    //aws.String(os.Getenv("AWS_REGION")),
        // Endpoint: aws.String("https://s3.us-east-1.amazonaws.com"),
        // Credentials: credentials.NewStaticCredentials(conf.AWS_ACCESS_KEY_ID, conf.AWS_SECRET_ACCESS_KEY, "")
    if err != nil {
        return "", err
    }
    svc := s3.New(sess)

    req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
    })

    urlStr, err := req.Presign(15 * time.Minute)

    if err != nil {
        log.Println("Failed to sign request", err)
    }
//    log.Println("The URL is", urlStr)
    return urlStr, err
}
