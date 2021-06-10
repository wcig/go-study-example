package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	name = flag.String("name", "default-name", "名称")
	age  int
)

func init() {
	flag.IntVar(&age, "age", 10, "年龄")
}

func main() {
	fmt.Println("command-line args:", os.Args)
	fmt.Println("current is parsed:", flag.Parsed())
	flag.Parse()

	fmt.Println("current is parsed:", flag.Parsed())
	fmt.Printf("name:%s, age:%d\n", *name, age)
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
		fmt.Printf("flag-vskit: name:%s, defaultValue:%v, usage:%s, value:%v\n",
			f.Name, f.DefValue, f.Usage, f.Value)
	})
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("flag-vskitAll: name:%s, defaultValue:%v, usage:%s, value:%v\n",
			f.Name, f.DefValue, f.Usage, f.Value)
	})

	// go run flag.go -name=tom -age=20 hello world ok
	// output:
	// command-line args: [/tmp/go-build708948192/b001/exe/flag -name=tom -age=20 hello world ok]
	// current is parsed: false
	// current is parsed: true
	//name:tom, age:20
	// args: [hello world ok]
	// arg 0: hello
	// remain arg num after flag: 3
	// flag num: 2
	// after set name: jerry
	// flag-vskit: name:age, defaultValue:10, usage:年龄, value:20
	// flag-vskit: name:name, defaultValue:default-name, usage:名称, value:jerry
	// flag-vskitAll: name:age, defaultValue:10, usage:年龄, value:20
	// flag-vskitAll: name:name, defaultValue:default-name, usage:名称, value:jerry
}
