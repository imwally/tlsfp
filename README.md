# tlsfp

Get TLS certificate fingerprints.

```
usage: tlsfp [-a algorithm] host
  -a int
    	algorithm: 1, 256 (default 1)
```
## Examples

```
$ tlsfp -a 256 google.com
A7 1C B5 68 76 0A CF 4F 21 52 C9 35 98 B0 87 48 C2 14 5C 61 46 8D E3 9C DB 85 04 28 FA 13 54 97
```