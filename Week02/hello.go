package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	gerrors "github.com/pkg/errors"
)

type resultMsg struct {
	code   int
	msg    string
	result string
}

func dao() (int, error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	a := r.Intn(10)
	if a == 1 {
		var sqlError = errors.New("sqlRowError")
		return 1, gerrors.Wrap(sqlError, "no data found")
	}
	return a, nil

}

func service() (int, error) {
	n, err := dao()
	if err != nil {
		return n, err
	}
	return n, nil
}
func web() resultMsg {

	fmt.Println("\r\n call service")
	n, err := service()
	if err != nil {
		fmt.Printf("error %+v", err)
		return resultMsg{500, "fail", "there is something wrong"}
	}
	return resultMsg{200, "success", fmt.Sprint(n)}
}

func main() {
	code := 200
	resultMsg := resultMsg{}
	for code == 200 {
		resultMsg = web()
		code = resultMsg.code
		time.Sleep(time.Duration(1) * time.Second)
	}
	fmt.Println("\r\nfail to call")
	fmt.Printf("%v \r\n", resultMsg)

}
