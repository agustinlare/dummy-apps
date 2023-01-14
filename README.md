# golang

A simple API in `golang` that counts every hit.

## Endpoints
+ `/health`: Status and message dependes on env.HEALTH_CHECK
```
{
  "endpoint": "checklist",
  "ip": "172.17.0.1:57582",
  "counter": 1,
  "status": 200,
  "message": "true"
}
```
+ `/unhealthy`: Status 500
```
{
  "endpoint": "checklist",
  "ip": "172.17.0.1:57582",
  "counter": 1,
  "status": 500,
  "message": "cof...cof"
}
```
+ `/ping`: Status 200
```
{
  "endpoint": "checklist",
  "ip": "172.17.0.1:57582",
  "counter": 1,
  "status": 200,
  "message": "pong"
}
```
+ `/checklist`: Status 200, front 

# Kotlin

S3 Tester

**List all buckets in S3**

GET <APP_URL>/s3tester/buckets

List all files within a bucket in S3
GET <APP_URL>/s3tester/buckets/{bucket_name}/files


S3 Credentials settings
Set the property aws.credentials.provided to true to use local credentials. For example >>

aws.s3.accessKeyId=******************
aws.s3.secretAccessKey=******************

# python

