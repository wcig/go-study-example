package url

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"testing"
)

func TestPathEscape(t *testing.T) {
	fmt.Println(url.PathEscape("ok"))
	fmt.Println(url.PathEscape("<>"))
	fmt.Println(url.PathEscape("好的"))
	fmt.Println(url.PathEscape("1+2"))
	// output:
	// ok
	// %3C%3E
	// %E5%A5%BD%E7%9A%84
	// 1+2
}

func TestPathUnescape(t *testing.T) {
	fmt.Println(url.PathUnescape("ok"))
	fmt.Println(url.PathUnescape("%3C%3E"))
	fmt.Println(url.PathUnescape("%E5%A5%BD%E7%9A%84"))
	fmt.Println(url.PathUnescape("1+2"))
	// output:
	// ok <nil>
	// <> <nil>
	// 好的 <nil>
	// 1+2 <nil>
}

func TestQueryEscape(t *testing.T) {
	fmt.Println(url.QueryEscape("ok"))
	fmt.Println(url.QueryEscape("<>"))
	fmt.Println(url.QueryEscape("好的"))
	fmt.Println(url.QueryEscape("1+2"))
	// output:
	// ok
	// %3C%3E
	// %E5%A5%BD%E7%9A%84
	// 1%2B2
}

func TestQueryUnescape(t *testing.T) {
	fmt.Println(url.QueryUnescape("ok"))
	fmt.Println(url.QueryUnescape("%3C%3E"))
	fmt.Println(url.QueryUnescape("%E5%A5%BD%E7%9A%84"))
	fmt.Println(url.QueryUnescape("1%2B2"))
	fmt.Println(url.QueryUnescape("1+2"))
	// output:
	// ok <nil>
	// <> <nil>
	// 好的 <nil>
	// 1+2 <nil>
	// 1 2 <nil>
}

func TestTypeUrl(t *testing.T) {
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(u)
	fmt.Println(string(b))

	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Println(u)
	// output:
	// {"Scheme":"http","Opaque":"","User":null,"Host":"bing.com","Path":"/search","RawPath":"","ForceQuery":false,"RawQuery":"q=dotnet","Fragment":"","RawFragment":""}
	// https://google.com/search?q=golang
}

func TestUrlEscapeFragment(t *testing.T) {
	u, err := url.Parse("http://example.com/#x/y%2Fz")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fragment:", u.Fragment)
	fmt.Println("RawFragment:", u.RawFragment)
	fmt.Println("EscapedFragment:", u.EscapedFragment())
	// output:
	// Fragment: x/y/z
	// RawFragment: x/y%2Fz
	// EscapedFragment: x/y%2Fz
}

func TestUrlEscapePath(t *testing.T) {
	u, err := url.Parse("http://example.com/x/y%2Fz")
	// u, err := url.Parse("http://example.com/x/y/z?a=b")
	// u, err := url.Parse("http://localhost:6060/pkg/net/url/#PathEscape")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Path:", u.Path)
	fmt.Println("RawPath:", u.RawPath)
	fmt.Println("EscapedPath:", u.EscapedPath())
	// output:
	// Path: /x/y/z
	// RawPath: /x/y%2Fz
	// EscapedPath: /x/y%2Fz
}

func TestUrlHostname(t *testing.T) {
	u, err := url.Parse("https://example.org:8000/path")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Hostname())

	u, err = url.Parse("https://[2001:0db8:85a3:0000:0000:8a2e:0370:7334]:17000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Hostname())
	// output:
	// example.org
	// 2001:0db8:85a3:0000:0000:8a2e:0370:7334
}

func TestUrlIsAbs(t *testing.T) {
	u := url.URL{Host: "example.com", Path: "foo"}
	fmt.Println(u.String(), u.IsAbs())

	u.Scheme = "http"
	fmt.Println(u.String(), u.IsAbs())
	// output:
	// //example.com/foo false
	// http://example.com/foo true
}

func TestUrlMarshalBinary(t *testing.T) {
	u, _ := url.Parse("https://example.org?q=ok")
	b, err := u.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b) // https://example.org?q=ok
}

func TestUrlPare(t *testing.T) {
	u, err := url.Parse("https://example.org")
	if err != nil {
		log.Fatal(err)
	}

	rel, err := u.Parse("/foo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rel) // https://example.org/foo

	_, err = u.Parse(":foo")
	if _, ok := err.(*url.Error); !ok {
		log.Fatal(err)
	}
}

func TestUrlPort(t *testing.T) {
	u, err := url.Parse("https://example.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Port())

	u, err = url.Parse("https://example.org:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Port())
	// output:
	//
	// 8080
}

func TestUrlQuery(t *testing.T) {
	u, err := url.Parse("https://example.org/?a=1&a=2&b=&=3&&&&")
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	fmt.Println(q["a"])
	fmt.Println(q.Get("b"))
	fmt.Println(q.Get(""))
	// output:
	// [1 2]
	//
	// 3
}

func TestUrlRedacted(t *testing.T) {
	u := &url.URL{
		Scheme: "https",
		User:   url.UserPassword("user", "password"),
		Host:   "example.com",
		Path:   "foo/bar",
	}
	fmt.Println(u.Redacted())

	u.User = url.UserPassword("me", "newerPassword")
	fmt.Println(u.Redacted())
	// output:
	// https://user:xxxxx@example.com/foo/bar
	// https://me:xxxxx@example.com/foo/bar
}

func TestUrlRequestURI(t *testing.T) {
	u, err := url.Parse("https://example.org/path?foo=bar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.RequestURI()) // /path?foo=bar
}

func TestUrlResolveReference(t *testing.T) {
	u, err := url.Parse("../../..//search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}
	base, err := url.Parse("http://example.com/directory/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(base.ResolveReference(u)) // http://example.com/search?q=dotnet
}

func TestUrlString(t *testing.T) {
	u := &url.URL{
		Scheme:   "https",
		User:     url.UserPassword("me", "pass"),
		Host:     "example.com",
		Path:     "foo/bar",
		RawQuery: "x=1&y=2",
		Fragment: "anchor",
	}
	fmt.Println(u.String())
	u.Opaque = "opaque"
	fmt.Println(u.String())
	// output:
	// https://me:pass@example.com/foo/bar?x=1&y=2#anchor
	// https:opaque?x=1&y=2#anchor
}

func TestUrlUnmarshalBinary(t *testing.T) {
	u := &url.URL{}
	err := u.UnmarshalBinary([]byte("https://example.org/foo?a=b"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", u) // https://example.org/foo?a=b
}

func TestTypeValues(t *testing.T) {
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	fmt.Println(v.Encode())
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])
	// output:
	// friend=Jess&friend=Sarah&friend=Zoe&name=Ava
	// Ava
	// Jess
	// [Jess Sarah Zoe]
}

func TestValueParseQuery(t *testing.T) {
	m, err := url.ParseQuery(`x=1&y=2&y=3;z`)
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(m)
	fmt.Println(string(b)) // {"x":["1"],"y":["2","3"],"z":[""]}
}
