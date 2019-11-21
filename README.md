# PBMM Audit

The purpose of this tool is to check if a cloud vendor account implements the guardrails specified by Shared Services Canada. The tool currently only support AWS, but Azure support will be coming.

Ex:
```
➜  ppbmm_audit aws --aws_key=... --aws_secret=...

Checking AWS root account for MFA ...
 💀  Root MFA is not enabled

Checking AWS users accounts for MFA ...
 ❗  0 out of 3 users have MFA active

Checking AWS for users with admin policies attached ...
 ❗  1 user(s) have admin policies attached

Checking AWS for lambda log export function ...
 ✅  Lambda export function found
```

## AWS implementation

The tool checks the following guardrails:

[X] Protect Root / Global Admins Account with MFA
[x] Management of Administrative Privileges with MFA
[X] Cloud Console Access (Developers/Application Owners) with MFA
[ ] Enterprise Monitoring Accounts
[X] Logging and monitoring with Lambda

You can check your AWS account using the following command:

`pbmm_audit aws --aws_key=YOUR_KEY --aws_secret=YOUR_SECRET`

Refer to `aws.policy.json` to see what account access the tool needs.

## Azure implementation

Coming soon.

## Long term objectives

The long term objective is to build a tool that ensures continous compliance with the guardrails.

## License

MIT