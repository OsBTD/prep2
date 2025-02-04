package main

import (
	"bytes"
	"fmt"
	"io"
	"math/big"
	"os"
	"os/exec"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

func Format(a ...interface{}) string {
	ss := make([]string, len(a))
	for i, v := range a {
		switch v.(type) {
		case nil:
			ss[i] = "nil" // instead of "<nil>"
		case
			string,
			byte, // uint8
			rune: // int32

			// string     : a double-quoted string safely escaped with Go syntax
			// byte, rune : a single-quoted character literal safely escaped with Go syntax
			ss[i] = fmt.Sprintf("%q", v)
		default:
			if reflect.TypeOf(v).Kind() == reflect.Func {
				// Function passed as parameter should output the name of the function
				ss[i] = runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
			} else {
				// a Go-syntax representation of the value
				ss[i] = fmt.Sprintf("%#v", v)
			}
		}
	}
	return strings.Join(ss, ", ")
}

func Call(fn interface{}, args []interface{}) []interface{} {
	// Convert args from []interface{} to []reflect.Value
	vals := make([]reflect.Value, len(args))
	for i, v := range args {
		if v != nil {
			vals[i] = reflect.ValueOf(v)
		} else {
			vals[i] = reflect.Zero(reflect.TypeOf((*interface{})(nil)).Elem())
		}
	}

	vals = reflect.ValueOf(fn).Call(vals)

	// Convert the return values from []reflect.Value to []interface{}
	result := make([]interface{}, len(vals))
	for i, v := range vals {
		result[i] = v.Interface()
	}
	return result
}

type Output struct {
	Results []interface{}
	Stdout  string
}

func Monitor(fn interface{}, args []interface{}) (out Output) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w
	out.Results = Call(fn, args)
	outC := make(chan string)
	var buf strings.Builder
	go func() {
		_, _ = io.Copy(&buf, r)
		outC <- buf.String()
	}()
	os.Stdout = old
	_ = w.Close()
	out.Stdout = <-outC
	return out
}

func Function(name string, submittedFunction, expectedFunction interface{}, args ...interface{}) {
	submitted := Monitor(submittedFunction, args)
	expected := Monitor(expectedFunction, args)
	if !reflect.DeepEqual(submitted.Results, expected.Results) {
		Fatalf("%s(%s) == %s instead of %s\n",
			name,
			Format(args...),
			Format(submitted.Results...),
			Format(expected.Results...),
		)
	}
	if !reflect.DeepEqual(submitted.Stdout, expected.Stdout) {
		Fatalf("%s(%s) prints:\n%s\ninstead of:\n%s\n",
			name,
			Format(args...),
			Format(submitted.Stdout),
			Format(expected.Stdout),
		)
	}
}

func Fatal(a ...interface{}) {
	_, _ = fmt.Fprint(os.Stderr, a...)
	os.Exit(1)
}

func Fatalln(a ...interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}

func Fatalf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func ProgramStdin(exercise, input string, args ...string) {
	run := func(pkg string) (string, bool) {
		binaryPath := path.Join(os.TempDir(), "binaries", path.Base(path.Dir(pkg)), path.Base(pkg))
		if _, err := os.Stat(binaryPath); os.IsNotExist(err) {
			if b, err := exec.Command("go", "build", "-o", binaryPath, pkg).CombinedOutput(); err != nil {
				return string(b), false
			}
		}
		if fi, err := os.Stat(binaryPath); err != nil || fi.Mode()&0111 == 0 {
			return "go run: cannot run non-main package\n", false
		}
		cmd := exec.Command(binaryPath, args...)
		if input != "" {
			cmd.Stdin = bytes.NewBufferString(input)
		}
		b, err := cmd.CombinedOutput()
		if err != nil {
			if _, ok := err.(*exec.ExitError); !ok {
				return err.Error(), false
			}
			return string(b) + err.Error(), false
		}
		return string(b), true
	}
	console := func(out string) string {
		var quotedArgs []string
		for _, arg := range args {
			quotedArgs = append(quotedArgs, strconv.Quote(arg))
		}
		s := "\n$ "
		if input != "" {
			s += "echo -ne " + strconv.Quote(input) + " | "
		}
		return fmt.Sprintf(s+"go run . %s\n%s$", strings.Join(quotedArgs, " "), out)
	}
	student, studentOK := run(path.Join("student", exercise))
	solution, solutionOK := run(exercise)
	if solutionOK {
		if !studentOK {
			Fatalln("Your program fails (non-zero exit status) when it should not :\n" +
				console(student) +
				"\n\nExpected :\n" +
				console(solution))
		}
	} else {
		if studentOK {
			Fatalln("Your program does not fail when it should (with a non-zero exit status) :\n" +
				console(student) +
				"\n\nExpected :\n" +
				console(solution))
		}
	}
	if student != solution {
		Fatalln("Your program output is not correct :\n" +
			console(student) +
			"\n\nExpected :\n" +
			console(solution))
	}
}

func Program(exercise string, args ...string) {
	ProgramStdin(exercise, "", args...)
}

// TODO: check unhandled errors on all solutions (it should contains "ERROR" on the first line to prove we correctly handle the error)
// TODO: remove the number of rand functions, refactor test cases (aka "table")
func PrintMemorys(a [10]byte) {
	str := ""
	for i, nbr := range a {
		fmt.Printf("%.2x", nbr)

		if ((i+1)%4 == 0 && i != 0) || i == len(a)-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}

		if nbr >= 33 && nbr <= 126 {
			str += string(rune(nbr))
		} else {
			str += "."
		}
	}
	fmt.Println(str + strings.Repeat(".", 10-len(a)))
}

const (
	MinInt = ^MaxInt
	MaxInt = 1<<63 - 1
)

var (
	nsSince1970 = time.Now().UnixNano()
	bigRand     = rand.New(rand.NewSource(nsSince1970))
)

func init() {
	rand.Seed(nsSince1970)
}

func makeIntFunc(f func() int) (s []int) {
	for i := 0; i < 8; i++ {
		s = append(s, f())
	}
	return
}

// IntBetween returns a random int between a and b included.
func IntBetween(a, b int) int {
	if a > b {
		a, b = b, a
	}
	n := big.NewInt(int64(b))      // b
	n.Sub(n, big.NewInt(int64(a))) // b-a
	n.Add(n, big.NewInt(1))        // b-a+1
	n.Rand(bigRand, n)             // 0 <= n <= b-a
	n.Add(n, big.NewInt(int64(a))) // a <= n <= b
	return int(n.Int64())
}

// Int returns a random int between MinInt and MaxInt included.
func Int() int {
	return IntBetween(MinInt, MaxInt)
}

// IntSlice returns a slice of 8 random ints between MinInt and MaxInt included.
func IntSlice() []int {
	return makeIntFunc(Int)
}

// IntSliceBetween returns a slice of 8 random ints between a and b included.
func IntSliceBetween(a, b int) []int {
	return makeIntFunc(func() int {
		return IntBetween(a, b)
	})
}

// Str returns a string with l random characters taken from chars.
func Str(chars string, length int) (dst string) {
	if length <= 0 {
		return ""
	}
	if chars == "" {
		panic("No charset provided")
	}
	for ; length > 0; length-- {
		r := rand.Intn(len(chars))
		dst += string(chars[r])
	}
	return string(dst)
}

// Str returns a slice of 8 strings with 13 random characters taken from chars.
func StrSlice(chars string) (s []string) {
	for i := 0; i < 8; i++ {
		s = append(s, Str(chars, 13))
	}
	return
}

func main() {
	table := [10]byte{}

	for j := 0; j < 5; j++ {
		for i := 0; i < IntBetween(7, 10); i++ {
			table[i] = byte(IntBetween(13, 126))
		}
		Function("PrintMemory", PrintMemory, PrintMemory, table)
	}
	table2 := [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'}
	Function("PrintMemory", PrintMemory, PrintMemorys, table2)
}

func PrintMemory(l [10]byte) {
	// PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})

	var s []string
	for i := 0; i < len(l); i++ {
		s = append(s, Convert(l[i]))
	}

	// fmt.Print(s)

	for j := 0; j < len(s); j++ {
		for k := len(s[j]) - 1; k >= 0; k-- {
			fmt.Print(string(s[j][k]))
		}
		if s[j] == "" {
			fmt.Print("00")
		}
		if j != 0 && (j == len(s)-1 || (j+1)%4 == 0) {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}

	}

	for m := 0; m < len(l); m++ {
		if string(l[m]) < " " || string(l[m]) > "~" {
			fmt.Print(".")
		}
		fmt.Print(string(l[m]))

	}
}

// func PrintMemory(l [10]byte) string {

// }
func Convert(n byte) string {
	Hex := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	var x int
	var s string
	for n > 0 {
		x = int(n) % 16
		s += Hex[x]
		n = n / 16
	}
	return s
}
