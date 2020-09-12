# unew

**u**_**(rl)**_**new** — A tool for append URLs, skipping duplicates & combine parameters. Inspired by [anew](https://github.com/tomnomnom/anew) & [qsreplace](https://github.com/tomnomnom/qsreplace).

# Usage

```bash
▶ cat urls.txt | unew
# or
▶ unew urls.txt
# or, save the results
▶ unew urls.txt output.txt
```

## Flags

Usage of `unew`:
```
  -combine
      Combine parameters
  -r string
      Replace parameters value
```

# Install

with [Go](https://golang.org/doc/install):

```bash
▶ go get -u github.com/dwisiswant0/unew
```

# Workaround

If you have a list as

```txt
https://twitter.com/dwisiswant0?href=evilzone.org
https://twitter.com/dwisiswant0
https://twitter.com/dwisiswant0?ref=github&utm_source=github
https://www.linkedin.com/in/dwisiswanto/
https://www.linkedin.com/in/dwisiswanto/?originalSubdomain=id
https://www.linkedin.com/in/dwisiswanto/?originalSubdomain=id&utm_medium=github
```

Sample workaround
```bash
▶ cat urls.txt | unew
https://twitter.com/dwisiswant0?href=evilzone.org
https://www.linkedin.com/in/dwisiswanto/
```

If the list contains multiple URLs with same path, it will save the first one and its parameters.

But you can combine parameters if the same path exists by using `-combine` flag.

```bash
▶ cat urls.txt | unew -combine
https://twitter.com/dwisiswant0?href=evilzone.org&ref=github&utm_source=github
https://www.linkedin.com/in/dwisiswanto/?originalSubdomain=id&utm_medium=github
```

Use the `-r` flag if you want to change the value of all parameters.

```bash
▶ cat urls.txt | unew -combine -r "/etc/passwd"
https://twitter.com/dwisiswant0?href=%2Fetc%2Fpasswd&ref=%2Fetc%2Fpasswd&utm_source=%2Fetc%2Fpasswd
https://www.linkedin.com/in/dwisiswanto/?originalSubdomain=%2Fetc%2Fpasswd&utm_medium=%2Fetc%2Fpasswd
```