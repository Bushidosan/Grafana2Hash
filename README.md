# Grafana2Hash
Simple script which converts the secret + salt of Grafana passwords to a hashcat usable hash.
Made this script for a [VulnLab](https://vulndev.io/lab/) machine called Data. Writeup of this machine is available at [Bushidosan.com](https://bushidosan.com/posts/vl-data/)


## Usage
```
go run .\Grafana2Hash.go <password> <hash>
```

## Credits
[VulnCheck](https://www.vulncheck.com/blog/grafana-cve-2021-43798)
