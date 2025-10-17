package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
)

func rbg(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

func rainbow(i int)(int,int,int){
	switch i%3{
	case 0:
		return 255,0,0 
	case 1:
		return 0,255,0
	default:
		return 0,0,255


	}}

func yellow(i int)(int,int,int){
	if(i==0){
	return 255,255,0
}
return 255,0,0
}



func color_printer(output []rune) {
	for i := 0; i < len(output); i++ {
		r, g, b := rbg(i)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[i])
	}
	for i := 0; i < len(output); i++ {
		r, g, b := rainbow(i)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[i])
	}
	for i := 0; i < len(output); i++ {
		r, g, b := yellow(i)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[i])
	}


	fmt.Println()
}



func main() {
	var output []rune
	var cmd1 *exec.Cmd

	cli_arg := os.Args
	cli_args_etc := cli_arg[2:]
	if len(cli_args_etc) == 0 {
		cmd1 = exec.Command(cli_arg[1])
	} else {
		cmd1 = exec.Command(cli_arg[1], cli_args_etc...)
	}

	cmd1_op, _ := cmd1.StdoutPipe()

	if err := cmd1.Start(); err != nil {
		return
	}

	reader := bufio.NewReader(cmd1_op)
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}
	// println("gotokekll")

	cmd1.Wait()

	color_printer(output)
}
