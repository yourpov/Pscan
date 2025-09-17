# Tools

A collection of lightweight networking utilities.

| Tool       | Description                                | Example Usage                           | Output                  |
|------------|--------------------------------------------|-----------------------------------------|-------------------------|
| isPrivate  | Check if an IP is in a private (RFC 1918) range | `isPrivate 192.168.1.10`                | `true`                  |
| lookup     | Forward DNS lookup (A/AAAA records)        | `lookup example.com`                     | `93.184.216.34`         |
| pscan      | Simple port scanner                        | `pscan 192.168.1.5 20-100`               | `Open ports: 22, 80`    |
| rdns       | Reverse DNS lookup (PTR records)           | `rdns 8.8.8.8`                           | `dns.google`            |
| validate   | Validate IPv4/IPv6 addresses               | `validate 10.0.0.1`                      | `valid IPv4`            |


## Usages:

```bash
isPrivate.go <ip>
    <ip>: IPv4 or IPv6 address to check


lookup.go <domain>
    <domain>: Hostname to resolve into an IP address


pscan.go <host> <maxPort>
    <host>: Target IP or hostname
    <maxPort>: Highest port number to scan (starts at 1)


rdns.go <ip>
    <ip>: IPv4 or IPv6 address to resolve into a domain


validate.go <ip>
    <ip>: IPv4 or IPv6 address to validate
```