# Check Guardrails

The purpose of this tool is to check if a cloud vendor account implements the guardrails specified by the Canadian Government. The tool currently only support AWS, but Azure support will be coming.

Ex:
```
➜  check_guardrails aws --aws_key=... --aws_secret=...

Checking AWS root account for MFA ...
 ❌  Root MFA is not enabled
 
 Checking AWS root account for programmatic keys ...
 ✅  Root MFA has no programmatic keys
 
Checking AWS console users accounts for MFA ...
 ✅  All user accounts use MFA (taking into account 2 breakglass accounts)
 
Checking AWS for users with admin policies attached ...
 ❌  3 user(s) have admin policies attached (2 expected)
 
Checking AWS for lambda log export function ...
 ✅  Lambda export function found
 
Checking AWS password policy ...
 ✅  Password must be 15 characters or longer
 
Checking AWS GuardDuty ...
 ✅  GuardDuty found with master account enabled
 
Checking AWS EC2 data residency ...
 ❌  EC2 instances found outside ca-central-1
 
Checking AWS S3 bucket encryption settings ...
 ✅  No unexpected S3 bucket found without encryption
 
Checking AWS RDS encryption settings ...
 ❌  RDS instance found without encryption
 
Checking AWS EC2 security groups for port 80 ingress ...
 ❌  Security group with port 80 found
```

## AWS implementation

This tool should only be used to check application level accounts, not organisation accounts.

The tool checks the following guardrails:

| Guardrail  |  Verfication method |
|---|---|
| Protect Root / Global Admins Account  | Validates that there is MFA active on root account  |
| Protect Root / Global Admins Account  | Validates that root account does not have programmatic keys |
| Protect Root / Global Admins Account  | Validates that password policy requires 15 characters  |
| Protect Root / Global Admins Account  | Validates that break glass accounts exist  |
| Cloud Console Access (Developers/Application Owners)  | Validates that console users have MFA active  |
| Cloud Console Access (Developers/Application Owners)  | Validates that non-console users do not have an admin policy attached  |
| Enterprise Monitoring Accounts  | Validates that GuardDuty is active with a master account enabled  |
| Data location in Canada | Validates that no EC2 instances exist outside of `CA-CENTRAL-1`  |
| Protection of data-at-rest | Validates that all S3 buckets are encrypted unless they are on the safelist  |
| Protection of data-at-rest | Validates that all RDS instances are encrypted |
| Protection of data-in-transit | Validates that no security groups allow traffic on TCP port 80 |
| Logging and monitoring | Validates that the lambda export function exists |

You can check your AWS account using the following command:

`check_guardrails aws --aws_key=YOUR_KEY --aws_secret=YOUR_SECRET`

You can also define these and other variables in a `yaml` file. Review `.check_guardrails.yaml.example` for more information.

Refer to `aws.policy.json` to see what account access the tool needs.

## Azure implementation

Coming soon.

## Long term objectives

The long term objective is to build a tool that ensures continous compliance with the guardrails.

## License

MIT
