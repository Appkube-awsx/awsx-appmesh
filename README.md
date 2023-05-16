- [What is awsx-appmesh](#awsx-appmesh)
- [How to write plugin subcommand](#how-to-write-plugin-subcommand)
- [How to build / Test](#how-to-build--test)
- [what it does ](#what-it-does)
- [command input](#command-input)
- [command output](#command-output)
- [How to run ](#how-to-run)

# awsx-appmesh

This is a plugin subcommand for awsx cli ( https://github.com/Appkube-awsx/awsx#awsx ) cli.

For details about awsx commands and how its used in Appkube platform , please refer to the diagram below:

![alt text](https://raw.githubusercontent.com/AppkubeCloud/appkube-architectures/main/LayeredArchitecture-phase2.svg)

This plugin subcommand will implement the Apis' related to AppMesh services , primarily the following API's:

- getConfigData

This cli collect data from metric / logs / traces of the AppMesh services and produce the data in a form that Appkube Platform expects.

This CLI , interacts with other Appkube services like Appkube vault , Appkube cloud CMDB so that it can talk with cloud services as
well as filter and sort the information in terms of product/env/ services, so that Appkube platform gets the data that it expects from the cli.

# How to write plugin subcommand

Please refer to the instruction -
https://github.com/Appkube-awsx/awsx#how-to-write-a-plugin-subcommand

It has detailed instruction on how to write a subcommand plugin , build/test/debug/publish and integrate into the main commmand.

# How to build / Test

            go run main.go
                - Program will print Calling aws-cloudelements on console

            Another way of testing is by running go install command
            go install
            - go install command creates an exe with the name of the module (e.g. awsx-appmesh) and save it in the GOPATH
            - Now we can execute this command on command prompt as below
           awsx-appmesh getConfigData --zone=us-east-1 --accessKey=xxxxxxxxxx --secretKey=xxxxxxxxxx --crossAccountRoleArn=xxxxxxxxxx  --externalId=xxxxxxxxxx

# what it does

This subcommand implement the following functionalities -
getConfigData - It will get the resource count summary for a given AWS account id and region.

# command input

1. --valutURL = specifies the URL of the AWS Key Management Service (KMS) customer master key (CMK) that you want to use to encrypt a table.
2. --acountId = specifies the AWS account id.
3. --zone = specifies the AWS region where the meshes is located.
4. --accessKey = specifies the AWS access key to use for authentication.
5. --secretKey = specifies the AWS secret key to use for authentication.t
6. --crossAccountRoleArn = specifies the Amazon Resource Name (ARN) of the role that allows access to a meshes in another account.
7. --external Id = specifies the AWS External id.
8. --meshName = Insert your meshName which you get in meshes list in aws account.
9. --meshOwner = Insert your meshOwner which you get in meshes list in aws account.

# command output

Meshes: [{
Arn: "arn:aws:appmesh:us-east-1:657907747545:mesh/abdul-test-1",
CreatedAt: 2023-02-27 09:53:19.527 +0000 UTC,
LastUpdatedAt: 2023-02-27 09:53:19.527 +0000 UTC,
MeshName: "abdul-test-1",
MeshOwner: "657907747545",
ResourceOwner: "657907747545",
Version: 1
},{
Arn: "arn:aws:appmesh:us-east-1:657907747545:mesh/ab-test-2",
CreatedAt: 2023-02-27 14:31:29.78 +0000 UTC,
LastUpdatedAt: 2023-02-27 14:31:29.78 +0000 UTC,
MeshName: "ab-test-2",
MeshOwner: "657907747545",
ResourceOwner: "657907747545",
Version: 1
}]

# How to run

From main awsx command , it is called as follows:

```bash
awsx-appmesh  --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<>
```

If you build it locally , you can simply run it as standalone command as:

```bash
go run main.go  --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<> --externalId=<>
```

# awsx-appmesh

appmesh extension

# AWSX Commands for AWSX-AppMesh Cli's :

1. CMD used to get list of appmesh instance's :

```bash
./awsx-appmesh --zone=us-east-1 --accessKey=<6f> --secretKey=<> --crossAccountRoleArn=<> --externalId=<>
```

2. CMD used to get Config data (metadata) of AWS AppMesh instances :

```bash
./awsx-appmesh --zone=us-east-1 --accessKey=<#6f> --secretKey=<> --crossAccountRoleArn=<> --externalId=<> getConfigData --meshName=<> --meshOwner=<>
```
