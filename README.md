# go-aws-examples

This repository contains a collection of Go code examples for interacting with Amazon Web Services (AWS) services. These examples demonstrate various use cases and best practices for working with AWS services in Go.

## Prerequisites

Before you begin, ensure you have the following prerequisites:

1.  **Go**: You need to have Go (Golang) installed on your system. You can download and install Go from [the official Go website](https://golang.org/dl/).
2.  **AWS Account**: You must have an AWS account to use the AWS services. If you don't have one, you can create an account at [AWS Sign-Up](https://aws.amazon.com/).
3.  **AWS Credentials**: To run these examples, you need to configure your AWS credentials. The recommended way to do this is by using the AWS Command Line Interface (AWS CLI) and the `aws configure` command. Make sure your credentials have the necessary permissions for the AWS services you intend to use in these examples.

## Getting Started

1\. Clone the Repository:

```
git clone https://github.com/vuduongtp/go-aws-examples.git
```

2\. Change Directory to example folder:

```
cd go-aws-examples/example-folder
```

3\. Install Dependencies:

```
make depends
```

4\. Deploy

```
make deploy
```

5\. Check resource on AWS

6\. Remove resource on AWS

```
make remove
```

7\. Double check resource on AWS to make sure all resources are deleted

## Examples

The repository is organized into directories, each containing a set of Go examples for a specific AWS service or use case. Below is a list of the available examples:

1.  [Lambda + API Gateway + DynamoDB](https://github.com/vuduongtp/go-aws-examples/tree/main/lambda-rest-api-with-dynamodb): Build simple rest api.
2.  [Lambda  + SNS + SQS](https://github.com/vuduongtp/go-aws-examples/tree/main/lambda-sns-sqs): Send message to SNS topic, SNS send message to SQS, Lambda will receive message from SQS and process.
3.  [Lambda  + Step Functions](https://github.com/vuduongtp/go-aws-examples/tree/main/lambda-step-functions): Build processing stream have many steps.

Each example directory will have its own README file with specific usage instructions and explanations.

## Contributing

We welcome contributions from the community. If you have an example or improvement to an existing one, please open a pull request.

## License

This repository is licensed under the MIT License. See the LICENSE file for details.

**Please note that while these examples are provided as a resource, you should be cautious when working with AWS services to avoid incurring unexpected costs or security vulnerabilities.**
