# WeCarry API Lambdas

A collection of Lambda functions for the WeCarry App

## Dev setup
1. Copy `.env.example` to `.env` and fill in appropriate values.

## Staging Deploy
1. Create a local `aws.credentials` file
2. Run `make deploy` to build and deploy lambda service

## AWS Credentials
AWS IAM Credentials are needed to deploy the Lambda function using Serverless.
Terraform config files for creating the IAM resources can be found in
[wecarry-terraform](https://github.com/silinternational/wecarry-terraform)
