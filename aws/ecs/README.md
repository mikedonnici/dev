# AWS ECS and ECR

Notes on using Elastic Container Service and Elastic Container Registry. 


## Authenticating with ECR using Docker

On Linux docker generally requires `sudo` so, for convenience, add your user 
to the `docker` group. 

See: https://docs.docker.com/engine/install/linux-postinstall/

Steps to authenticate with ECR:

- [Install AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html)
- [Configure AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html)
- [Private Registry Auth](https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry_auth.html)

In short, command below fetches an auth token and pipes that to 
[docker login](https://docs.docker.com/engine/reference/commandline/login/#login-to-a-self-hosted-registry):

```shell
$ aws ecr get-login-password --region ap-southeast-2 / 
  | docker login --username AWS --password-stdin / 
  896154582544.dkr.ecr.ap-southeast-2.amazonaws.com
```

Region must be specified and last arg is: `[account id].dkr.ecr.[region].amazonaws.com`

This will store credentials in `~/.docker/config.json` and the aws token should 
remain valid for 12 hours.

