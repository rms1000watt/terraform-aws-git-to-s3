# Terraform AWS Git to S3

## Introduction

Terraform module in AWS to clone Git repo and place in S3


## Contents

- [Usage](#usage)
- [Dependencies](#dependencies)

## Usage

```hcl
module "git_to_s3" {
  source = "rms1000watt/git-to-s3/aws"

  git_user = "username"
  git_pass = "password"
  git_url  = "https://github.com/rms1000watt/serverless-tf.git"
}
```

See `examples` folder for working examples

## Dependencies

Golang environment with `github.com/kardianos/govendor`
