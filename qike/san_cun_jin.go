package qike

import "fmt"

//三寸金 观月论

func SanCunJin(yg uint, yue uint, eg uint) {
	GongWei := [8]uint{1, 7, 13, 19, 25, 31, 37, 43}
	//大安
	if yg == GongWei[0] || yg == GongWei[1] || yg == GongWei[2] || yg == GongWei[3] || yg == GongWei[4] || yg == GongWei[5] || yg == GongWei[6] {
		if yue == 1 {
			fmt.Println("三寸金-->大安:	正月大安性情直 见喜前奉遇恶耻 女命流连旺夫志 三九起家运不迟")
		} else {
			fmt.Println("三寸金-->大安:	大安七月同父换 初年马星运渐宽 门出军匪少亡汉 四九方能将运缓")
		}
	}
	//留连
	if yg == GongWei[0]+1 || yg == GongWei[1]+1 || yg == GongWei[2]+1 || yg == GongWei[3]+1 || yg == GongWei[4]+1 || yg == GongWei[5]+1 || yg == GongWei[6]+1 {
		if yue == 2 {
			fmt.Println("三寸金-->留连:	二月流连运不稳 做官总将桃色混 运低受触多穷困 重照需挪外乾坤")
		} else {
			fmt.Println("三寸金-->留连:	流连八月共潮生  信从军戎备官成 天命安排福自增")
			if eg == GongWei[0]+1 || eg == GongWei[1]+1 || eg == GongWei[2]+1 || eg == GongWei[3]+1 || eg == GongWei[4]+1 || eg == GongWei[5]+1 || eg == GongWei[6]+1 {
				fmt.Println("三寸金-->时留连:	重连必见双井绳")
			}
		}
	}
	//速喜
	if yg == GongWei[0]+2 || yg == GongWei[1]+2 || yg == GongWei[2]+2 || yg == GongWei[3]+2 || yg == GongWei[4]+2 || yg == GongWei[5]+2 || yg == GongWei[6]+2 {
		if yue == 3 {
			fmt.Println("三寸金-->速喜:	三月速喜闹东堂 七一九三两分张 子午破财合中伤 小人指背号无常")
		} else {
			fmt.Println("三寸金-->速喜:	速喜九月官不和 双亲堂前多见克 唯有车马并刀戈 唯怕木命来过河")
		}
	}
	//赤口
	if yg == GongWei[0]+3 || yg == GongWei[1]+3 || yg == GongWei[2]+3 || yg == GongWei[3]+3 || yg == GongWei[4]+3 || yg == GongWei[5]+3 || yg == GongWei[6]+2 {
		if yue == 4 {
			fmt.Println("三寸金-->赤口:	四月白虎有桃情 七一九三必多井 若是二三福不清 姻缘早动冲马星")
		} else {
			fmt.Println("三寸金-->赤口:	十月白虎卧寒松 二八逢年运不通 男迎女送映桃红 牵手姻缘走西东")
		}
	}
	//小吉
	if yg == GongWei[0]+4 || yg == GongWei[1]+4 || yg == GongWei[2]+4 || yg == GongWei[3]+4 || yg == GongWei[4]+4 || yg == GongWei[5]+4 || yg == GongWei[6]+2 {
		if yue == 5 {
			fmt.Println("三寸金-->小吉	五月小吉力多能 贵人面前去点灯 运见二七多调孥 大安合局为上芍")
		} else {
			fmt.Println("三寸金-->小吉	小吉冬月性情倔 克伐六亲有伤缺 脐带缠处斜痠雀 兄弟多来独寒月")
		}
	}
	//空亡
	if yg == 0 || yg == GongWei[0]+5 || yg == GongWei[1]+5 || yg == GongWei[2]+5 || yg == GongWei[3]+5 || yg == GongWei[4]+5 || yg == GongWei[5]+5 || yg == GongWei[6]+2 {
		if yue == 6 {
			fmt.Println("三寸金-->空亡	六月空亡心聪颖 亲朋好有念善情 男子回身戴官翎 女子身弱念佛经")
		} else {
			fmt.Println("三寸金-->空亡	空亡腊月心刚强 自力更生挑家梁 争强好胜把话讲 唯有婚姻大不详")
		}
	}
}
