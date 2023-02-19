package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Percent int

func (p *Percent) Set(s string) error {
	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*p = Percent(n)
	return nil
}

func (p *Percent) String() string {
	return strconv.Itoa(int(*p))
}

var (
	name    = flag.String("name", "default-name", "名称")
	age     int
	percent Percent
)

func init() {
	flag.IntVar(&age, "age", 10, "年龄")
	flag.Var(&percent, "percent", "百分比")
}

func main() {
	fmt.Println("command-line args:", os.Args)
	fmt.Println("current is parsed:", flag.Parsed())
	flag.Parse()

	fmt.Println("current is parsed:", flag.Parsed())
	fmt.Printf("name:%s, age:%d, percent:%d\n", *name, age, percent)
	fmt.Println("args:", flag.Args())
	fmt.Println("arg 0:", flag.Arg(0))
	fmt.Println("remain arg num after flag:", flag.NArg())
	fmt.Println("flag num:", flag.NFlag())

	// set
	if err := flag.Set("name", "jerry"); err != nil {
		panic(err)
	}
	fmt.Println("after set name:", *name)

	// visit
	flag.Visit(func(f *flag.Flag) {
		fmt.Printf("flag-visit: name:%s, defaultValue:%v, usage:%s, value:%v\n",
			f.Name, f.DefValue, f.Usage, f.Value)
	})
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("flag-visitAll: name:%s, defaultValue:%v, usage:%s, value:%v\n",
			f.Name, f.DefValue, f.Usage, f.Value)
	})

	// go run flag.go -name=tom -age=20 -percent=30 hello world ok
	// output:
	// command-line args: [/var/folders/vh/lks7z1qx6x90j10nwtm3njlw0000gn/T/go-build1964889729/b001/exe/flag -name=tom -age=20 -percent=30 hello world ok]
	// current is parsed: false
	// current is parsed: true
	// name:tom, age:20, percent:30
	// args: [hello world ok]
	// arg 0: hello
	// remain arg num after flag: 3
	// flag num: 3
	// after set name: jerry
	// flag-visit: name:age, defaultValue:10, usage:年龄, value:20
	// flag-visit: name:name, defaultValue:default-name, usage:名称, value:jerry
	// flag-visit: name:percent, defaultValue:0, usage:百分比, value:30
	// flag-visitAll: name:age, defaultValue:10, usage:年龄, value:20
	// flag-visitAll: name:name, defaultValue:default-name, usage:名称, value:jerry
	// flag-visitAll: name:percent, defaultValue:0, usage:百分比, value:30
}
