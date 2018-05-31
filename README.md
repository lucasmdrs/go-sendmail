# Go-sendmail
  A simple AWS Lambda function written in Go to send emails using Sendgrid

### Requirements
 - Sendgrid account
 - AWS
 - aws-cli

### Up and Running

#### Configure
Configure the lambda.json with your account information
```json
{
    "FunctionName": "go-sendmail",
    "Runtime": "go1.x",
    "Role": "arn:aws:iam::{{ YOUR_AWS_ACCOUNT_ID }}:role/{{ YOUR_ROLE_NAME }}",
    "Handler": "main",
    "Description": "A simple AWS Lambda function written in Go to send emails using Sendgrid",
    "Timeout": 300,
    "MemorySize": 128,
    "Publish": true,
    "Environment": {
        "Variables": {
            "SENDGRID_API_KEY": "{{ YOUR_API_KEY }}"
        }
    },
    "Tags": {
        "Name": "go-sendmail"
    }
}
```

#### Deploy
Deploy the function with aws-cli
```shell
GOOS=linux go build
zip main.zip go-sendmail
aws lambda create-function \
  --region us-east-1 \
  --zip-file fileb://main.zip \
  --cli-input-json file://lambda.json
```

Or deploy using the AWS Console.

#### Invoke
Invoke the function with aws-cli
```shell
aws lambda invoke \
  --region us-east-1 \
  --function-name go-sendmail \
  --log-type Tail \
  --payload file://invoke-payload.json \
  lambda-response
```
