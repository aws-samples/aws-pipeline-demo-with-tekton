# aws-pipeline-demo-with-tekton

This repository contains the demo resources discussed in the blog post [Cloud Native CI/CD with Tekton and ArgoCD on AWS](https://aws.amazon.com/blogs/containers/cloud-native-ci-cd-with-tekton-and-argocd-on-aws/).

The code provided is for demo purposes only and not ready for production.

## Prerequisites
This demo requires multiple tools to be installed on your machine.

Please make sure that the following tools are installed and ready to use:

- [AWS CLI v2](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)
- [eksctl](https://eksctl.io/introduction/#installation)
- [kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Helm](https://helm.sh/docs/intro/install/)
- [aws-iam-authenticator](https://docs.aws.amazon.com/eks/latest/userguide/install-aws-iam-authenticator.html)
- [jq](https://stedolan.github.io/jq/download/)
- [Golang](https://go.dev/dl/)
- [Docker](https://www.docker.com/products/docker-desktop)
- [envsubst](https://formulae.brew.sh/formula/gettext)

Further we suggest to use a dedicated AWS account.
The install script should be executed with the credentials of an Admin user.

The following articles provide guidance to setup an AWS Account and configure the required Admin user:

- [Create an AWS Account](https://aws.amazon.com/premiumsupport/knowledge-center/create-and-activate-aws-account/)
- [Create an Admin User](https://docs.aws.amazon.com/IAM/latest/UserGuide/getting-started_create-admin-group.html)
- [Setup your CLI credentials](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html)

Please note that the tests were made within the eu-central-1 (Frankfurt) region and a VPC that had subnets deployed into 3 availability zones. Further the installation took place from a computer with the MacOS operating system installed

Next a Kubernetes cluster is required in order to deploy Tekton and other related resources. In order to work with the demo script please use the official eksctl command line tool to create the cluster.
Please find below the required steps in order to create a cluster (cluster config file can be found in the root folder of the repository):

```console
$ eksctl create cluster -f eks-cluster-template.yaml
```

Please wait until the cluster has been provisioned successfully and you obtained the kubeconfig file.
You can test the successful installation by running:

```console
$ eksctl get clusters
$ kubectl get nodes
```

If both of the above commands completed successfully please continue with the installation steps.

## Install demo environment

Clone the repository and run the installation script:

```console
$ git clone https://github.com/aws-samples/aws-pipeline-demo-with-tekton.git
$ cd aws-pipeline-demo-with-tekton
$ chmod u+x install.sh
$ ./install.sh
```
Please note that the script requires your public ip address to continue. This ip address will be used to restrict the access to the resources deployed through the script.

The script installs the environment and takes approximately 10 minutes to complete (depends on your internet connectivity). Please keep your Terminal open until everything is installed and the output section is displayed.

## Uninstall

To uninstall all resources, please switch back into the root folder:

```console
$ cd aws-tekton-pipeline-demo
$ chmod u+x uninstall.sh
$ ./uninstall.sh
```

Wait until all resources have been removed. We suggest to double check your AWS account for not cleaned up resources which needs to be removed manually.
Especially the cluster needs to be removed manually.

## Security

See [CONTRIBUTING](CONTRIBUTING.md#security-issue-notifications) for more information.

## License

This library is licensed under the MIT-0 License. See the LICENSE file.

