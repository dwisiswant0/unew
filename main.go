package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
)

type skipPaths []string

var (
	combine                  bool
	outfile                  *os.File
	sc                       *bufio.Scanner
	skipPath                 skipPaths
	replace, outtext, fr, fs string
)

func (s *skipPaths) String() string {
	return ""
}

func (s *skipPaths) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func init() {
	flag.StringVar(&replace, "r", "", "Replace parameters value")
	flag.BoolVar(&combine, "combine", false, "Combine parameters")
	flag.Var(&skipPath, "skip-path", "Skip specific paths (regExp pattern)")
	flag.Parse()

	fr := flag.Arg(0)

	if isStdin() {
		sc = bufio.NewScanner(os.Stdin)
	} else if fr != "" {
		r, err := os.Open(fr)
		if err == nil {
			sc = bufio.NewScanner(r)
		}
	} else {
		os.Exit(1)
	}
}

func main() {
	urls := make(map[string]string)

	fs := flag.Arg(1)
	if fs != "" {
		outfile, _ = os.OpenFile(fs, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		defer outfile.Close()
	}

	for sc.Scan() {
		u, err := url.ParseRequestURI(sc.Text())
		if err != nil {
			continue
		}

		if matchPath(skipPath, u) {
			continue
		}

		b := fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, u.Path)

		if _, d := urls[b]; d {
			if combine {
				q := []string{urls[b], u.RawQuery}
				u.RawQuery = remDup(strings.Join(q, "&"))
			} else {
				continue
			}
		} else {
			urls[b] = u.RawQuery
		}

		if replace != "" {
			u.RawQuery = qsReplace(u.Query(), replace)
		}

		if !combine {
			outtext = fmt.Sprintf("%s%s\n", b, qMark(u.RawQuery))
			fmt.Printf("%s", outtext)
			if fs != "" {
				fmt.Fprintf(outfile, "%s", outtext)
			}

			continue
		}

		urls[b] = u.RawQuery
	}

	if !combine {
		return
	}

	for k, v := range urls {
		outtext = fmt.Sprintf("%s%s\n", k, qMark(v))
		fmt.Printf("%s", outtext)
		if fs != "" {
			fmt.Fprintf(outfile, "%s", outtext)
		}
	}
}

func isStdin() bool {
	f, e := os.Stdin.Stat()
	if e != nil {
		return false
	}

	if f.Mode()&os.ModeNamedPipe == 0 {
		return false
	}

	return true
}

func qMark(q string) string {
	if q == "" {
		return ""
	}

	return "?" + q
}

func remDup(q string) string {
	qs := url.Values{}
	ps, _ := url.ParseQuery(q)

	for p, v := range ps {
		for range v {
			v = v[:1:1]
			qs.Set(p, v[0])
		}
	}

	return qs.Encode()
}

func qsReplace(q url.Values, r string) string {
	qs := url.Values{}
	for p := range q {
		qs.Set(p, r)
	}

	return qs.Encode()
}

func matchPath(s []string, u *url.URL) bool {
	for _, p := range s {
		m, e := regexp.MatchString(p, u.Path)
		if e != nil {
			continue
		}

		if m {
			return true
		}
	}

	return false
}
