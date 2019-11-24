package qike

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

//童男童女
//十二时辰
type Time struct {
	子午时 string
	丑未时 string
	寅申时 string
	卯酉时 string
	辰戌时 string
	巳亥时 string
}
//get number
func GetNumber() int{
	fmt.Println("\n如果你看到了这里的内容说明落宫应童子命 输入出生时辰对应数字查看更多内容")
	fmt.Println("十二时辰数字参考上文提示")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	input.Text()
	number,err := strconv.Atoi(input.Text())
	if err!=nil{
		fmt.Println(err)
	}
	return number
}
func TongZiMing(sg uint) {

	GongWei:= [8]uint{1,7,13,19,25,31,37,43}
	if (sg==GongWei[0]+1 || sg== GongWei[1]+1 || sg== GongWei[2]+1	|| sg== GongWei[3]+1|| sg== GongWei[4]+1 || sg== GongWei[5]+1 || sg== GongWei[6]+1)	{
		fmt.Println("三宫-->留连")

		number := GetNumber()
		if number != 1 && number != 2 && number != 3 && number != 4 && number != 5 && number != 6 && number != 7 && number != 8 && number != 9 && number != 10 && number != 11 && number != 12 {
			fmt.Println("时辰输入错误　退出")
			os.Exit(0)
		}

		time := Time{
			子午时: "子午佛道扫墙根",
			丑未时: "丑未关帝牵马蹲",
			寅申时: "寅申三清同道尊",
			卯酉时: "卯酉娘娘草木春",
			辰戌时: "辰戌老爷掌公文",
			巳亥时: "巳亥太子文武身",
		}
		if number == 1 || number == 7 {
			fmt.Println(time.子午时)
		}
		if number == 2 || number == 8 {
			fmt.Println(time.丑未时)
		}
		if number == 3 || number == 9 {
			fmt.Println(time.寅申时)
		}
		if number == 4 || number == 10 {
			fmt.Println(time.卯酉时)
		}
		if number == 5 || number == 11 {
			fmt.Println(time.辰戌时)
		}
		if number == 6 || number == 12 {
			fmt.Println(time.巳亥时)
		}
		fmt.Println("######################################")
		fmt.Println("以下信息自行选择")
		fmt.Println("还送替身敬天恩，三尺三寸童子身，穿着颜色要细分，甲己年来浓黄坤，" +
			"乙庚岁头绿草春，丙辛年首盯蓝深，丁壬青云透花纹，戊癸浅红必为真，" +
			"阴月逢八还替身，八卦对冲向巽辰，舅舅提半酉黄昏")
	}
}

func About()  {
	time.Sleep(time.Second*9)
	fmt.Println("					六壬通神暄未然\n"+
		"					因缘造化系自身\n"+
		"					阴阳共判劝君明\n"+
		"					虚实明辨壬中求\n")
	fmt.Println("						/by liangzi")
	fmt.Println("						mail: bGlhbmd6aTEyMTZAb3V0bG9vay5jb20K")
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")

}
