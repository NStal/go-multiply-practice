package main

import "fmt"
import "math/rand"
import "time"
import "strconv"
import "bufio"
import "os"
import "strings"

func main() {
	sdate := time.Now()
	success, fail := 0, 0
	ia,ib := 0,0
	var a,b float32;
	for true {
		rand.Seed(time.Now().Unix())
		a = rand.Float32()
		b = rand.Float32()
		ia = int(a * 100)
		ib = int(b * 100)
		fmt.Printf("%v x %v\n", ia, ib)
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, " \n")
		if strings.Compare(text, "end") == 0 {
			fmt.Printf("Done\n")
			break
		}
		number, err := strconv.Atoi(text)
		if err != nil {
			continue
		}
		if number == ia*ib {
			success = success + 1
		} else {
			fmt.Printf("Result is:%d \n", int(ia*ib))
			fail = fail + 1
		}
	}
	if success == 0 {
		fmt.Printf("No result\n")
		return
	}
	edate := time.Now()
	fmt.Printf("Success %d, Total %d, SuccessRate %f%\n", success, fail+success, float32(success)*100/(float32(success)+float32(fail)))
	msec := edate.Unix() - sdate.Unix()
	sec := float32(msec)
	speed := sec / float32(success)
	fmt.Printf("%f sec/Op", speed)
}
