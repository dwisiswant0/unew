# unew

**u**_**(rl)**_**new** — A tool for append URLs, skipping duplicates & combine parameters. Inspired by [anew](https://github.com/tomnomnom/anew) & [qsreplace](https://github.com/tomnomnom/qsreplace).

## Usage

```bash
▶ cat urls.txt | unew
# or
▶ unew urls.txt
# or, save the results
▶ unew urls.txt output.txt
```

### Flags

Usage of `unew`:
```
  -combine
        Combine parameters
  -r string
        Replace parameters value
  -skip-path value
        Skip specific paths (regExp pattern)
```

## Install

with [Go](https://golang.org/doc/install):

```bash
▶ go get -u github.com/dwisiswant0/unew
```

## Workaround

If you have a `urls.txt` list as

```txt
https://twitter.com/dwisiswant0?href=evilzone.org
https://twitter.com/dwisiswant0
https://twitter.com/dwisiswant0?ref=github&utm_source=github
https://twitter.com/dwisiswant0/status/1305022512590278656
https://www.linkedin.com/in/dwisiswanto/
https://www.linkedin.com/in/dwisiswanto/?originalSubdomain=id
https://www.linkedin.com/in/dwisiswanto/?originalSubdomain=id&utm_medium=github
```

### Regular

Sample workarounds:

```bash
▶ cat urls.txt | unew
https://twitter.com/dwisiswant0?href=evilzone.org
https://www.linkedin.com/in/dwisiswanto/
```

If the list contains multiple URLs with same path, it will save the first one and its parameters.

### Combining parameters

But you can combine parameters if the same path exists by using `-combine` flag.

```bash
▶ cat urls.txt | unew -combine
https://twitter.com/dwisiswant0?href=evilzone.org&ref=github&utm_source=github
https://www.linkedin.com/in/dwisiswanto/?originalSubdomain=id&utm_medium=github
```

### Query replacers

Use the `-r` flag if you want to change the value of all parameters.

```bash
▶ cat urls.txt | unew -combine -r "/etc/passwd"
https://twitter.com/dwisiswant0?href=%2Fetc%2Fpasswd&ref=%2Fetc%2Fpasswd&utm_source=%2Fetc%2Fpasswd
https://www.linkedin.com/in/dwisiswanto/?originalSubdomain=%2Fetc%2Fpasswd&utm_medium=%2Fetc%2Fpasswd
```

### Skipping paths

In case if you want to pass specific/multiple URL paths, you can use `-skip-path` flag for it _(can be set multiple times)_. But, you have to write it with regExp pattern.

```bash
▶ cat urls.txt | unew -skip-path "^/[\w]+/status/[0-9]+" -skip-path "/in/[\w]+"
https://twitter.com/dwisiswant0?href=evilzone.org
```