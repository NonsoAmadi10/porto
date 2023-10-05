# Porto
Born out of the frustration of having to upload files to s3 buckets via the UI, I want to encourage my laziness by having a cli tool do it for me

## Prerequisite
- aws cli
- AWS SECRET KEY
- AWS ACCESS ID

### Installation
- Configure your aws credentials by running:
```bash
aws configure
```
and following the prompt.

Download `porto` by running:
```bash
wget https://github.com/NonsoAmadi10/porto/releases/download/1.0.0/porto
```

- Make the binary executable:

```bash
chmod +x ./porto
```
- Move the kubectl binary to a file location on your system PATH.

```bash
sudo mv ./porto /usr/local/bin/porto
sudo chown root: /usr/local/bin/porto
```
Test to ensure the version you installed is up-to-date:

porto --version

### Usage
```bash
porto upload --region=<bucket-region> --bucket=<bucket-name> --file=<file>
```
You can also upload multiple file by simply adding another file extension:

```bash
porto upload --region=<bucket-region> --bucket=<bucket-name> --file=<file1> --file=<file2> --file=<file3>
```

## Contributions
This is tool is currently in beta mode and hopefully it gets to alpha and a final release. If you feel you have ideas on how we can improve this, shoot me a mail at nonsoamadi@aol.com and raise a PR or open an issue. I respond fast!

