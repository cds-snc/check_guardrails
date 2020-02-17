# Check Guardrails with inspec

This tool uses [Inspec](https://www.inspec.io/) to audit an AWS environment for cloud guardrail compliance. 

The easiest way to use the tool is with Docker. Execute the following commands from inside this directory to run it against an AWS account of choice.

```
docker build -t inspec-aws .
docker run -e AWS_ACCESS_KEY_ID=... -e AWS_SECRET_ACCESS_KEY=... inspec-aws
```

## License

MIT


