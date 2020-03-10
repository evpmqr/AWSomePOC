# AWSomePOC
This is a POC on a little idea I have.
Takes a yml file and converts it to an AWS Template File with the Serverless Application Model Support.
Adding more resources should pretty much be similar to how Lambda is right now.
The goal is to make a CLI that can generate a template for you, and a github bot that will read PRs with the config.yml file and automatically add it to the branch/PR.

# Requirments
* Go 1.1x

# How To Run
1. Make sure TestTemplate.yaml is deleted
2. Run `go run main.go`
3. TestTemplate.yaml should be generated using testing.yml

# Struct Validations
We need to figure out what rules we want to set for v1 of validation; probably won't be much:
So far I have come up with these validations (will update as I think of more):
* For Lambda Event Properties You cannot Define multiple event property field types. eg: You cannot define `path` and `queue`
* You cannot define any other resources if you are creating an IAM Account Role - Reason Being is that we want IAM Account Roles to be in their own templates

## TODOs
* IAM Role Default Permissions that gets applied without defining any policies
  * The Policy will be attached to the namespace of the IAM Role
* More AWS Resource Definitions
  * IAM Role
  * S3
  * API Gateway
* Better README