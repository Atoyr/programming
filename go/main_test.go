package main

import (
	"os"
	"io"
	"strings"
	"strconv"
	"testing"
)

type testValue struct {
	arg []string
	ans string
}

func getTests() []testValue{
	testValues := []testValue{
		testValue{ 
			[]string{ 
				"4 5", 
				"4 2",
				"5 2",
				"2 1",
				"8 3",
			} ,
		"13" },
		testValue{ 
			[]string{ 
				"2 20",
				"5 9",
				"4 10",
			} ,
		"9"},
	}

	return testValues
}


func stubio(inbuf string, f func()) (string,string) {
	inr, inw, _ := os.Pipe()
	outr, outw, _ := os.Pipe()
	errr, errw, _ := os.Pipe()

	orgStdin := os.Stdin
	orgStdout := os.Stdout
	orgStderr := os.Stderr

	inw.Write([]byte(inbuf))
	inw.Close()

	os.Stdin = inr
	os.Stdout = outw
	os.Stderr = errw

	outC := make(chan string)
	errC := make(chan string)
	defer close(outC)
	defer close(errC)
	go func() {
		var buf strings.Builder
		io.Copy(&buf,outr)
		outr.Close()
		outC <- buf.String()
	}()

	go func() {
		var buf strings.Builder
		io.Copy(&buf,errr)
		errr.Close()
		errC <- buf.String()
	}()

	f()
	
	os.Stdin = orgStdin
	os.Stdout = orgStdout
	os.Stderr = orgStderr
	outw.Close()
	errw.Close()

	return <-outC,<-errC
}

func Test_main(t *testing.T) {
	tests := getTests()
	for i, tt := range tests {
		si := strconv.Itoa(i)
		t.Run("Case "+si, func(t *testing.T) {
			ret,err := stubio(strings.Join(tt.arg," "),main)
			if err != "" {
				t.Errorf("error func: %s ", err)
			}
			if ret != tt.ans {
				t.Errorf("Unexpected output: '%s' Need: '%s'", ret, tt.ans)
			}
		})
	}
}
