package main

import (
	"C"
	"log"
	"os"
	"plugin"

	"github.com/urfave/cli"
)
import "fmt"

//Linker interface plugin library
type Linker interface {
	Add(a, b int) int
	Cosine(x float64) float64
	Sort(vals []int)
	Log(msg string) int
}

var Link Linker

func findPlgugin(name string) (Linker, error) {
	//conecting plugin
	p, err := plugin.Open("plugin.so")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//find
	greetSymbol, err := p.Lookup("Link")
	if err != nil {
		panic(err)
	}

	Link, ok := greetSymbol.(Linker)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	return Link, nil
}

func main() {

	//cli ap
	Ap := cli.NewApp()
	Ap.Version = "v0.0.1"
	Ap.Name = "LinkerGo"
	Ap.Author = "Ivan Cherkasov"
	Ap.HideHelp = true
	Ap.HideVersion = true

	Ap.Description = "Library building modules"
	Ap.Usage = "LinkerGo call func in building go programms.."
	//commands cli
	sliceComands := []cli.Command{}
	com := AddComandCli("h", "Description on module commands")
	com.Action = func(c *cli.Context) {
		fmt.Println(discription())
	}

	sliceComands = append(sliceComands, com)
	Ap.Commands = sliceComands

	//call in cli plug for terminal
	com = AddComandCli("sum", `"A" + "B"`)
	com.Action = func(c *cli.Context) {
		plug, err := findPlgugin("Add")

		if err != nil {
			fmt.Printf("not found!!")
		}

		var a int
		var b int
		fmt.Printf(`please enter a number "А"`)
		fmt.Scanf("%d", &a)
		fmt.Printf(`please enter a number "B"`)
		fmt.Scanf("%d", &b)
		f := plug.Add(a, b)
		fmt.Println(f)
	}

	sliceComands = append(sliceComands, com)
	Ap.Commands = sliceComands
	err := Ap.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func add(a, b int) {
	// 4. use the module
	f := Link.Add(a, b)
	fmt.Println(f)
}

//Init commands
func AddComandCli(name string, usage string) cli.Command {

	commanda := cli.Command{}
	commanda.Name = name
	commanda.Usage = usage
	commanda.SkipFlagParsing = false
	commanda.HideHelp = false
	commanda.Hidden = false
	return commanda
}

func discription() string {

	txt := `one dir main.go + plugin.go 
$ dep init
$ dep ensure			
$ go build -buildmode=plugin plugin.go
$ rm plugin.go
$ go build ;)
			
			import "plugin"

			//Linker interface plugin library
			type Linker interface {
				Add(a, b int) int
				Cosine(x float64) float64
				Sort(vals []int)
				Log(msg string) int
			}
			
			var Link Linker
			
			//conecting plugin
			p, err := plugin.Open("plugin.so")
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			//find
			greetSymbol, err := p.Lookup("Link")
			if err != nil {
				panic(err)
			}

			Link, ok := greetSymbol.(Linker)
			if !ok {
				fmt.Println("unexpected type from module symbol")
				os.Exit(1)
			}

			return Link, nil			
			f := Link.Add(a, b)
			fmt.Println(f)
			`
	return txt
}
