//后期用户交互部分
package public

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//获取输入信息

func Get_In_Put() string {
	input := bufio.NewReader(os.Stdin)
	input_s, err := input.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	s := strings.ToLower(strings.TrimSpace(input_s))
	return s
}

func Choise() {
	fmt.Println("查看六神属性按p")
	if Get_In_Put() == "p" {
		Liu_Shen() //六神属性
	}
}
