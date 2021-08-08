package ch33_regexp

import (
	"fmt"
	"regexp"
	"testing"
)

func TestMatch(t *testing.T) {
	matched, err := regexp.Match(`foo.*`, []byte(`seafood`))
	fmt.Println(matched, err)
	matched, err = regexp.Match(`bar.*`, []byte(`seafood`))
	fmt.Println(matched, err)
	matched, err = regexp.Match(`a(b`, []byte(`seafood`))
	fmt.Println(matched, err)
	// output:
	// true <nil>
	// false <nil>
	// false error parsing regexp: missing closing ): `a(b`
}

func TestMatchString(t *testing.T) {
	matched, err := regexp.MatchString(`foo.*`, `seafood`)
	fmt.Println(matched, err)
	matched, err = regexp.MatchString(`bar.*`, `seafood`)
	fmt.Println(matched, err)
	matched, err = regexp.MatchString(`a(b`, `seafood`)
	fmt.Println(matched, err)
	// output:
	// true <nil>
	// false <nil>
	// false error parsing regexp: missing closing ): `a(b`
}

func TestRegexpExpand(t *testing.T) {
	content := []byte(`
    # comment line
    option1: value1
    option2: value2

    # another comment line
    option3: value3
`)

	// Regex pattern captures "key: value" pair from the content.
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	// Template to convert "key: value" to "key=value" by
	// referencing the values captured by the regex pattern.
	template := []byte("$key=$value\n")

	result := []byte{}

	// For each match of the regex in the content.
	for _, submatches := range pattern.FindAllSubmatchIndex(content, -1) {
		// Apply the captured submatches to the template and append the output
		// to the result.
		result = pattern.Expand(result, template, content, submatches)
	}
	fmt.Println(string(result))
}

func TestRegexpFind(t *testing.T) {
	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.Find([]byte(`seafood fool`))) // "food"
}

func TestRegexpFindAll(t *testing.T) {
	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.FindAll([]byte(`seafood fool`), -1)) // ["food" "fool"]
}

func TestRegexpFindAllIndex(t *testing.T) {
	content := []byte("London")
	re := regexp.MustCompile(`o.`)
	fmt.Println(re.FindAllIndex(content, 1))
	fmt.Println(re.FindAllIndex(content, -1))
	// output:
	// [[1 3]]
	// [[1 3] [4 6]]
}

func TestRegexpFindAllString(t *testing.T) {
	re := regexp.MustCompile(`a.`)
	fmt.Println(re.FindAllString("paranormal", -1))
	fmt.Println(re.FindAllString("paranormal", 2))
	fmt.Println(re.FindAllString("graal", -1))
	fmt.Println(re.FindAllString("none", -1))
	// output:
	// [ar an al]
	// [ar an]
	// [aa]
	// []
}

func TestRegexpFindAllStringIndex(t *testing.T) {
	re := regexp.MustCompile(`a.`)
	fmt.Println(re.FindAllStringIndex("paranormal", -1))
	fmt.Println(re.FindAllStringIndex("paranormal", 2))
	fmt.Println(re.FindAllStringIndex("graal", -1))
	fmt.Println(re.FindAllStringIndex("none", -1))
	// output:
	// [[1 3] [3 5] [8 10]]
	// [[1 3] [3 5]]
	// [[2 4]]
	// []
}

func TestRegexpFindAllStringSubmatch(t *testing.T) {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-", -1))
	fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))
	// output:
	// [["ab" ""]]
	// [["axxb" "xx"]]
	// [["ab" ""] ["axb" "x"]]
	// [["axxb" "xx"] ["ab" ""]]
}

func TestRegexpFindAllStringSubmatchIndex(t *testing.T) {
	re := regexp.MustCompile(`a(x*)b`)
	// Indices:
	//    01234567   012345678
	//    -ab-axb-   -axxb-ab-
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-ab-axb-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-axxb-ab-", -1))
	fmt.Println(re.FindAllStringSubmatchIndex("-foo-", -1))
	// output:
	// [[1 3 2 2]]
	// [[1 5 2 4]]
	// [[1 3 2 2] [4 7 5 6]]
	// [[1 5 2 4] [6 8 7 7]]
	// []
}

func TestRegexpFindAllSubmatch(t *testing.T) {
	re := regexp.MustCompile(`foo(.?)`)
	fmt.Printf("%q\n", re.FindAllSubmatch([]byte(`seafood fool`), -1)) // [["food" "d"] ["fool" "l"]]
}

func TestRegexpFindAllSubmatchIndex(t *testing.T) {
	content := []byte(`
    # comment line
    option1: value1
    option2: value2
`)
	// Regex pattern captures "key: value" pair from the content.
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
	allIndexes := pattern.FindAllSubmatchIndex(content, -1)
	for _, loc := range allIndexes {
		fmt.Println(loc)
		fmt.Println(string(content[loc[0]:loc[1]]))
		fmt.Println(string(content[loc[2]:loc[3]]))
		fmt.Println(string(content[loc[4]:loc[5]]))
	}
	// output:
	// [24 39 24 31 33 39]
	// option1: value1
	// option1
	// value1
	// [44 59 44 51 53 59]
	// option2: value2
	// option2
	// value2
}

func TestRegexpFindIndex(t *testing.T) {
	content := []byte(`
    # comment line
    option1: value1
    option2: value2
`)
	// Regex pattern captures "key: value" pair from the content.
	pattern := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)

	loc := pattern.FindIndex(content)
	fmt.Println(loc)
	fmt.Println(string(content[loc[0]:loc[1]]))
	// output:
	// [24 39]
	// option1: value1
}

func TestRegexpFindString(t *testing.T) {
	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.FindString("seafood fool"))
	fmt.Printf("%q\n", re.FindString("meat"))
	// output:
	// "food"
	// ""
}

func TestRegexpFindStringIndex(t *testing.T) {
	re := regexp.MustCompile(`foo.?`)
	fmt.Println(re.FindStringIndex("seafood fool"))
	fmt.Println(re.FindStringIndex("meat"))
	// output:
	// [3 7]
	// []
}

func TestRegexpFindStringSubmatch(t *testing.T) {
	re := regexp.MustCompile(`a(x*)b(y|z)c`)
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
	// output:
	// ["axxxbyc" "xxx" "y"]
	// ["abzc" "" "z"]
}

func TestRegexpFindStringSubmatchIndex(t *testing.T) {
	re := regexp.MustCompile(`a(x*)b(y|z)c`)
	fmt.Println(re.FindStringSubmatchIndex("-axxxbyc-"))
	fmt.Println(re.FindStringSubmatchIndex("-abzc-"))
	// output:
	// [1 8 2 5 6 7]
	// [1 5 2 2 3 4]
}

func TestRegexpFindSubmatch(t *testing.T) {
	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.FindSubmatch([]byte("seafood fool")))
	fmt.Printf("%q\n", re.FindSubmatch([]byte("meat")))
	// output:
	// "food"
	// ""
}

func TestRegexpFindSubmatchIndex(t *testing.T) {
	re := regexp.MustCompile(`a(x*)b`)
	// Indices:
	//    01234567   012345678
	//    -ab-axb-   -axxb-ab-
	fmt.Println(re.FindSubmatchIndex([]byte("-ab-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-axxb-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-ab-axb-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-axxb-ab-")))
	fmt.Println(re.FindSubmatchIndex([]byte("-foo-")))
	// output:
	// [1 3 2 2]
	// [1 5 2 4]
	// [1 3 2 2]
	// [1 5 2 4]
	// []
}

func TestRegexpLiteralPrefix(t *testing.T) {
	re := regexp.MustCompile(`foo.?`)
	fmt.Println(re.LiteralPrefix()) // foo false
}

func TestRegexpLongest(t *testing.T) {
	re := regexp.MustCompile(`a(|b)`)
	fmt.Println(re.FindString("ab"))
	re.Longest()
	fmt.Println(re.FindString("ab"))
	// output:
	// a
	// ab
}

func TestRegexpMatch(t *testing.T) {
	re := regexp.MustCompile(`foo.?`)
	fmt.Println(re.Match([]byte(`seafood fool`)))
	fmt.Println(re.Match([]byte(`something else`)))
	// output:
	// true
	// false
}

func TestRegexpMatchString(t *testing.T) {
	re := regexp.MustCompile(`foo.?`)
	fmt.Println(re.MatchString(`seafood fool`))
	fmt.Println(re.MatchString(`something else`))
	// output:
	// true
	// false
}

func TestRegexpNumSubexp(t *testing.T) {
	re0 := regexp.MustCompile(`a.`)
	fmt.Printf("%d\n", re0.NumSubexp())

	re := regexp.MustCompile(`(.*)((a)b)(.*)a`)
	fmt.Println(re.NumSubexp())
	// output:
	// 0
	// 4
}

func TestRegexpReplaceAll(t *testing.T) {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("T")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("$1W")))
	fmt.Printf("%s\n", re.ReplaceAll([]byte("-ab-axxb-"), []byte("${1}W")))
	// output:
	// -T-T-
	// --xx-
	// ---
	// -W-xxW-
}

func TestRegexpReplaceAllLiteral(t *testing.T) {
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Printf("%s\n", re.ReplaceAllLiteral([]byte("-ab-axxb-"), []byte("T")))
	fmt.Printf("%s\n", re.ReplaceAllLiteral([]byte("-ab-axxb-"), []byte("$1")))
	fmt.Printf("%s\n", re.ReplaceAllLiteral([]byte("-ab-axxb-"), []byte("$1W")))
	fmt.Printf("%s\n", re.ReplaceAllLiteral([]byte("-ab-axxb-"), []byte("${1}W")))
	// output:
	// -T-T-
	// -$1-$1-
	// -$1W-$1W-
	// -${1}W-${1}W-
}

func TestRegexpSplit(t *testing.T) {
	re := regexp.MustCompile("a*")
	s := re.Split("abaabaccadaaae", 5)
	for _, v := range s {
		fmt.Printf("%q, ", v)
	}
	// output:
	// "", "b", "b", "c", "cadaaae"
}

func TestRegexpString(t *testing.T) {
	re := regexp.MustCompile("a*")
	fmt.Println(re.String()) // a*
}

func TestRegexpSubexpIndex(t *testing.T) {
	re := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
	fmt.Println(re.MatchString("Alan Turing"))
	matches := re.FindStringSubmatch("Alan Turing")
	lastIndex := re.SubexpIndex("last")
	fmt.Printf("last => %d\n", lastIndex)
	fmt.Println(matches[lastIndex])
	// output:
	// true
	// last => 2
	// Turing
}

func TestRegexpSubexpNames(t *testing.T) {
	re := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
	fmt.Println(re.MatchString("Alan Turing"))
	fmt.Printf("%q\n", re.SubexpNames())
	reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
	fmt.Println(reversed)
	fmt.Println(re.ReplaceAllString("Alan Turing", reversed))
	// output:
	// true
	// ["" "first" "last"]
	// ${last} ${first}
	// Turing Alan
}
