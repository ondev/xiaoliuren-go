package qike

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//起课后形成的三个落宫
var yi_gong = make(chan string, 70)
var er_gong = make(chan string, 70)
var san_gong = make(chan string, 70)

//童子命
var tong_zi = make(chan string, 70)

//起课部分

func GetInPut() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	input.Text()
	return input.Text()
}

func QiKe() {
	fmt.Println("\n道家小六壬-->命法卷　\n按时辰起课　输入农历月日时\n")
	fmt.Println("1->子时(23:00~1:00)　2->丑时(1:00~3:00) 3->寅时(3:00~5:00)\n" +
		"4->卯时(5:00~7:00) 5->辰时(7:00~9:00) 6->巳时(9:00~11:00)\n" +
		"7->午时(11:00~13:00) 8->未时(13:00~15:00) 9->申时(15:00~17:00) \n" +
		"10->酉时(17:00~19:00) 11->戌时(19:00~21:00) 12亥时(21:00~23:00)\n")
	fmt.Println("例: 4月26辰时 输入数字: 4  26  5")

	var one = GetInPut()
	var two = GetInPut()
	var three = GetInPut()

	var yue, err = strconv.Atoi(one)
	if err != nil {
		fmt.Println(err)
	}
	ri, _ := strconv.Atoi(two)
	shi, _ := strconv.Atoi(three)

	//fmt.Println("yue =", yue)
	//yue ri shi 对应解卦部分的几落　比如二宫十落大安　指的是起课数字two转换而来的ri
	// 在二宫落定的基础上走数十　所落的大安宫位　例如４ 10 2 赤口　十落大安
	var YiGong = yue % 6
	//	fmt.Println("yigong = ", YiGong)
	if ri > 6 {
		ErGong := YiGong + ri - 7
		//fmt.Println("ri >6 =",ri)
		if shi > 6 {
			SanGong := ErGong + shi - 7
			//fmt.Println("sangong >6(shi >6) ", SanGong)
			JieGua(uint(YiGong), uint(ErGong), uint(SanGong), uint(ri), uint(yue))
		}
		if shi <= 6 {
			ErGong := YiGong + ri - 7
			//fmt.Println("ri>6 =",ri)
			SanGong := ErGong + shi - 1
			//fmt.Println("sangong >6(shi<=6) =", SanGong)
			JieGua(uint(YiGong), uint(ErGong), uint(SanGong), uint(ri), uint(yue))
		}
	} else if ri <= 6 {
		var ErGong = YiGong + ri - 1
		var SanGong = ErGong + shi - 1
		//fmt.Println("ri<6: =", ri)
		//fmt.Println("sangong<=6: =", SanGong)
		JieGua(uint(YiGong), uint(ErGong), uint(SanGong), uint(ri), uint(yue))
	}

}

//解卦
func JieGua(yg uint, eg uint, sg uint, ri uint, yue uint) {

	GongWei := [8]uint{1, 7, 13, 19, 25, 31, 37, 43}
	//一宫 月落宫
	if yg == GongWei[0] ||
		yg == GongWei[1] ||
		yg == GongWei[2] ||
		yg == GongWei[3] ||
		yg == GongWei[4] ||
		yg == GongWei[5] ||
		yg == GongWei[6] {
		fmt.Println("一宫:大安")
		yi_gong <- "大安"
		defer close(yi_gong)
	}
	if yg == GongWei[0]+1 ||
		yg == GongWei[1]+1 ||
		yg == GongWei[2]+1 ||
		yg == GongWei[3]+1 ||
		yg == GongWei[4]+1 ||
		yg == GongWei[5]+1 ||
		yg == GongWei[6]+1 {
		fmt.Println("一宫:留连")
		yi_gong <- "留连"
		defer close(yi_gong)
	}
	if yg == GongWei[0]+2 ||
		yg == GongWei[1]+2 ||
		yg == GongWei[2]+2 ||
		yg == GongWei[3]+2 ||
		yg == GongWei[4]+2 ||
		yg == GongWei[5]+2 ||
		yg == GongWei[6]+2 {
		fmt.Println("一宫:速喜")
		yi_gong <- "速喜"
		defer close(yi_gong)
	}
	if yg == GongWei[0]+3 ||
		yg == GongWei[1]+3 ||
		yg == GongWei[2]+3 ||
		yg == GongWei[3]+3 ||
		yg == GongWei[4]+3 ||
		yg == GongWei[5]+3 ||
		yg == GongWei[6]+2 {
		fmt.Println("一宫:赤口")
		yi_gong <- "赤口"
		defer close(yi_gong)
	}
	if yg == GongWei[0]+4 ||
		yg == GongWei[1]+4 ||
		yg == GongWei[2]+4 ||
		yg == GongWei[3]+4 ||
		yg == GongWei[4]+4 ||
		yg == GongWei[5]+4 ||
		yg == GongWei[6]+2 {
		fmt.Println("一宫:小吉")
		yi_gong <- "小吉"
		defer close(yi_gong)
	}
	if yg == 0 ||
		yg == GongWei[0]+5 ||
		yg == GongWei[1]+5 ||
		yg == GongWei[2]+5 ||
		yg == GongWei[3]+5 ||
		yg == GongWei[4]+5 ||
		yg == GongWei[5]+5 ||
		yg == GongWei[6]+2 {
		fmt.Println("一宫:空亡")
		yi_gong <- "空亡"
		defer close(yi_gong)
	}

	//二宫 以日轮
	if eg == GongWei[0] ||
		eg == GongWei[1] ||
		eg == GongWei[2] ||
		eg == GongWei[3] ||
		eg == GongWei[4] ||
		eg == GongWei[5] ||
		eg == GongWei[6] {
		fmt.Println("二宫:大安")
		er_gong <- "大安"
		defer close(er_gong)
	}
	if eg == GongWei[0]+1 ||
		eg == GongWei[1]+1 ||
		eg == GongWei[2]+1 ||
		eg == GongWei[3]+1 ||
		eg == GongWei[4]+1 ||
		eg == GongWei[5]+1 ||
		eg == GongWei[6]+1 {
		fmt.Println("二宫:留连")
		er_gong <- "留连"
		defer close(er_gong)
	}
	if eg == GongWei[0]+2 ||
		eg == GongWei[1]+2 ||
		eg == GongWei[2]+2 ||
		eg == GongWei[3]+2 ||
		eg == GongWei[4]+2 ||
		eg == GongWei[5]+2 ||
		eg == GongWei[6]+2 {
		fmt.Println("二宫:速喜")
		er_gong <- "速喜"
		defer close(er_gong)
	}
	if eg == GongWei[0]+3 ||
		eg == GongWei[1]+3 ||
		eg == GongWei[2]+3 ||
		eg == GongWei[3]+3 ||
		eg == GongWei[4]+3 ||
		eg == GongWei[5]+3 ||
		eg == GongWei[6]+2 {
		fmt.Println("二宫:赤口")
		er_gong <- "赤口"
		defer close(er_gong)
	}
	if eg == GongWei[0]+4 ||
		eg == GongWei[1]+4 ||
		eg == GongWei[2]+4 ||
		eg == GongWei[3]+4 ||
		eg == GongWei[4]+4 ||
		eg == GongWei[5]+4 ||
		eg == GongWei[6]+2 {
		fmt.Println("二宫:小吉")
		er_gong <- "小吉"
		defer close(er_gong)
	}
	if eg == 0 ||
		eg == GongWei[0]+5 ||
		eg == GongWei[1]+5 ||
		eg == GongWei[2]+5 ||
		eg == GongWei[3]+5 ||
		eg == GongWei[4]+5 ||
		eg == GongWei[5]+5 ||
		eg == GongWei[6]+2 {
		fmt.Println("二宫:空亡")
		er_gong <- "空亡"
		defer close(er_gong)
	}

	//三宫
	if sg == GongWei[0] ||
		sg == GongWei[1] ||
		sg == GongWei[2] ||
		sg == GongWei[3] ||
		sg == GongWei[4] ||
		sg == GongWei[5] ||
		sg == GongWei[6] {
		fmt.Println("三宫:大安")
		san_gong <- "大安"
		defer close(san_gong)
	}
	if sg == GongWei[0]+1 ||
		sg == GongWei[1]+1 ||
		sg == GongWei[2]+1 ||
		sg == GongWei[3]+1 ||
		sg == GongWei[4]+1 ||
		sg == GongWei[5]+1 ||
		sg == GongWei[6]+1 {
		fmt.Println("三宫:留连")
		san_gong <- "留连"
		defer close(san_gong)
		//童男童女
		//TongZiMing()
		/*		tong_zi <- "此为童子命"
				defer close(tong_zi)*/ //死锁
	}
	if sg == GongWei[0]+2 ||
		sg == GongWei[1]+2 ||
		sg == GongWei[2]+2 ||
		sg == GongWei[3]+2 ||
		sg == GongWei[4]+2 ||
		sg == GongWei[5]+2 ||
		sg == GongWei[6]+2 {
		fmt.Println("三宫:速喜")
		san_gong <- "速喜"
		defer close(san_gong)
	}
	if sg == GongWei[0]+3 ||
		sg == GongWei[1]+3 ||
		sg == GongWei[2]+3 ||
		sg == GongWei[3]+3 ||
		sg == GongWei[4]+3 ||
		sg == GongWei[5]+3 ||
		sg == GongWei[6]+2 {
		fmt.Println("三宫:赤口")
		san_gong <- "赤口"
		defer close(san_gong)
	}
	if sg == GongWei[0]+4 ||
		sg == GongWei[1]+4 ||
		sg == GongWei[2]+4 ||
		sg == GongWei[3]+4 ||
		sg == GongWei[4]+4 ||
		sg == GongWei[5]+4 ||
		sg == GongWei[6]+2 {
		fmt.Println("三宫:小吉")
		san_gong <- "小吉"
		defer close(san_gong)
	}
	if sg == 0 ||
		sg == GongWei[7]+3 ||
		sg == GongWei[0]+5 ||
		sg == GongWei[1]+5 ||
		sg == GongWei[2]+5 ||
		sg == GongWei[3]+5 ||
		sg == GongWei[4]+5 ||
		sg == GongWei[5]+5 ||
		sg == GongWei[6]+2 {
		fmt.Println("三宫:空亡")
		san_gong <- "空亡"
		defer close(san_gong)
	}
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	fmt.Println("金口定-->马前论命分百般，六壬推卦说一番")
	jin_kou_ding() //金口定
	fmt.Println("百句金口暂活用，人生自古无十全，月缺尚须等月圆，请君续观小透玄")
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	XiaoTouXuan(yg, eg, sg) //小透玄
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	SanCunJin(yg, yue, eg) //三寸金
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	YiDaoZhan(yg, eg, sg, ri) //一刀斩
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")
	TongZiMing(sg)
	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::")

}
