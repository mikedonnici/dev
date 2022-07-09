AWS Certifications

Suggested Path for Cloud Architect gig:

1. [Cloud Practitioner](./cloud-practitioner/)
1. [Solutions Architect - Associate](./solutions-architect-associate/)
1. [Developer - Associate](./developer-associate/)
1. SysOps Administrator - Associate
1. Security - Specialty
1. Big Data - Specialty (opt)
1. Advanced Networking - Specialty (opt)
1. DevOps Engineer - Professional
1. Solutions Architect - Professional

Resources:

- [How I Passed the AWS Certified Developer Associate Exam](https://www.freecodecamp.org/news/how-i-passed-the-aws-certified-developer-associate-exam/)
- https://thefreeccp.com (free practitioner)
- https://www.exampro.co (paid - cheaper that cloud guru)

## Disable MFA 

- If MFA stops working, but still have cli access, can disable MFA:

```bash
$ aws iam list-mfa-devices --user-name SomeUser
# copy SerialNumber in response
$ aws iam deactivate-mfa-device --user-name SomeUser --serial-number [SerialNumber]
```

- Can then log into console and re-activate MFA
