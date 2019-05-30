# get_ips
Go APP for bulk get IPs for domain names

Build:

```
cd get_ips
go build
```

Input file:

```
domain1.tld
domain2.tld
domain3.tld
```

Usage:

```
./get_ips inputfile.txt > outputfile.json
```

Output file:

```
{"domain":"domain1.tld","ip":"1.2.3.4","datum":"2019-05-29"}
{"domain":"domain2.tld","ip":"5.6.7.8","datum":"2019-05-29"}
{"domain":"domain3.tld","ip":"0.0.0.0","datum":"2019-05-29"}
```
