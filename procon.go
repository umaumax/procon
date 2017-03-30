package procon

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	//	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func SolveWithArgs(solve func(in io.Reader, out io.Writer) (err error)) {
	flag.Parse()
	nArg := flag.NArg()
	if nArg%2 == 1 || nArg == 0 {
		log.Fatalln("worng args")
	}
	nTest := nArg / 2
	inputs := flag.Args()[:nTest]
	outputs := flag.Args()[nTest:]

	//	in := strings.NewReader("hel lo")
	//	ans := strings.NewReader("hello2")
	for i := 0; i < nTest; i++ {
		var err error
		fmt.Println("----++++----++++----")
		fmt.Printf("[%03d]:'%s-%s'\n", i, inputs[i], outputs[i])

		in, err := os.Open(inputs[i])
		if err != nil {
			log.Fatalln(i, "file open err:", err)
		}

		ans, err := os.Open(outputs[i])
		if err != nil {
			log.Fatalln(i, "file open err:", err)
		}

		w := bytes.Buffer{}
		out := &w
		err = solve(in, out)
		if err != nil {
			log.Fatalln(err)
		}

		ansTmp, _ := ioutil.ReadAll(ans)
		ansText := string(ansTmp)
		outText := out.String()
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(ansText, outText, false)

		fmt.Print(dmp.DiffPrettyText(diffs))
		fmt.Println("----++++----++++----")
	}
}
