# NTUOSS Go Lambda Workshop Repository

## Overview

This workshop will guide you through the process of creating a simple FX backend service using Go and AWS Lambda.

## Prerequisites

- Any code editor/IDE, e.g. VSCode, GoLand, Vim (including necessary extensions)
- Install Go ([documentation](https://go.dev/doc/install))
- Install Postman
- Create an AWS account ([homepage](https://aws.amazon.com/console/))

> Do note that a valid payment method is required to sign-up (for AWS), even though they provide generous free tier that is enough for this workshop

## AWS Setup

### Create Lambda function

1. Select "Author from scratch"
1. Select "Provide your own bootstrap on Amazon Linux 2" as the runtime
1. Select "arm64" as the architecture
1. Click "Create function"

### Create API Gateway

1. Select "HTTP API"
1. Add the created Lambda function as an integration
1. Add "ANY /{proxy+}" as a route for proxy integration
1. Ensure "Auto-deploy" is enabled
1. Create the API
