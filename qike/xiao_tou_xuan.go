package qike

import "fmt"

//小透玄部分　以日时断

type GongWEi struct {
	DaAn     string
	LiuLian  string
	SuXi     string
	ChiKou   string
	XiaoJi   string
	KongWang string
}

func XiaoTouXuan(yg uint, eg uint, sg uint) {
	/*
		xiao_tou_xuan := GongWEi{
			DaAn:"\n日生大安心慈祥，速喜临时拜佛堂，先克老父有灾殃，就连岳父也说上，" +
				"早见坎坷晚吉昌，遭克必定走外乡，时加空亡财路强，唯犯小吉水潭深，" +
				"流连只时禄学堂，速喜姻缘两分张，大安重照把官当，最忌桃红幻迷汤。",
				LiuLian:"日生流连脚板宽，披霞落马走江川，心直口快不触官，惹犯小人闹不安，" +
					"大安落位不一般，禄马学堂捧金碗，时临小吉相貌端，吉星贵人来把关，" +
					"财地速喜靠嘴积，不需劳刀一生欢，黑道空亡多严官，笔直大路切莫弯。",

					SuXi:"日坐速喜巧舌礼，月同婚姻主分离，各论兄弟皆同理，宴客交朋不满席，" +
					"做事一气不到底，半途心猿旁改题，三月走到空亡地，克妻不死亦各异，" +
					"速喜遭克守法纪，小人多嗔贵人喜，大安落位相貌奇，白虎加时犯官气。",
					ChiKou:"日坐白虎性情怪，临时本位唇舌殃，俊情对意敞心怀，语不达心头不抬，" +
				"月日重照面多白，早运玉在坐中埋，白虎逢克化以财，空见大安父先载，" +
				"是征兵需挂帅，工农商贾离乡外，禄马学堂遭挫败，押身一枝命安排。",
				XiaoJi:"日生小吉人多情，流连桃红泛水盈，身肩背痣做将领，娶妻嫁夫两口井，" +
				"受克小吉福不清，学堂官章巧聪伶，六亲空亡冷如冰，孜身形影犯孤丁，" +
				"多寅速喜不足凭，过房安身劳碌命。",
				KongWang:"日生空亡财禄重，时临大安慢排功，女命刚强性气宏，压夫不起运晚通，" +
				"龙伤子息晚景空，交朋结友忌尾龙，空亡空空到对宫，速喜坐时旺财丰，" +
				"空亡重照牢灾冲，小吉似吉却土崩，时临白虎在梦中，温柔香刮冷雨风。",
		}
	*/
	GongWei := [8]uint{1, 7, 13, 19, 25, 31, 37, 43}
	//日宫（二宫即时辰宫）落大安
	if eg == GongWei[0] ||
		eg == GongWei[1] ||
		eg == GongWei[2] ||
		eg == GongWei[3] ||
		eg == GongWei[4] ||
		eg == GongWei[5] ||
		eg == GongWei[6] {
		//大安:= xiao_tou_xuan.DaAn
		//fmt.Println("小透玄-->大安：",大安)
		fmt.Println("小透玄-->大安:	日生大安心慈祥。")
		//时辰宫落大安解卦
		if sg == GongWei[0] ||
			sg == GongWei[1] ||
			sg == GongWei[2] ||
			sg == GongWei[3] ||
			sg == GongWei[4] ||
			sg == GongWei[5] ||
			sg == GongWei[6] {
			fmt.Println("时辰宫-->大安:	大安重照把官当，最忌桃红幻迷汤。")
		}
		//时辰落留连宫解卦
		if sg == GongWei[0]+1 ||
			sg == GongWei[1]+1 ||
			sg == GongWei[2]+1 ||
			sg == GongWei[3]+1 ||
			sg == GongWei[4]+1 ||
			sg == GongWei[5]+1 ||
			sg == GongWei[6]+1 {
			fmt.Println("时辰宫-->留连:	流连只时禄学堂。")
		}
		//时辰落速喜解卦
		if sg == GongWei[0]+2 ||
			sg == GongWei[1]+2 ||
			sg == GongWei[2]+2 ||
			sg == GongWei[3]+2 ||
			sg == GongWei[4]+2 ||
			sg == GongWei[5]+2 ||
			sg == GongWei[6]+2 {
			fmt.Println("时辰宫-->速喜:	速喜临时拜佛堂，先克老父有灾殃，就连岳父也说上," +
				"早见坎坷晚吉昌，遭克必定走外乡,速喜姻缘两分张。")
		}
		//时辰落宫空亡解卦
		if sg == 0 ||
			sg == GongWei[7]+3 ||
			sg == GongWei[0]+5 ||
			sg == GongWei[1]+5 ||
			sg == GongWei[2]+5 ||
			sg == GongWei[3]+5 ||
			sg == GongWei[4]+5 ||
			sg == GongWei[5]+5 ||
			sg == GongWei[6]+2 {
			fmt.Println("时辰宫-->空亡:	时加空亡财路强，唯犯小吉水潭深。")
			//fmt.Println("作者注解:	心地善良　财路强　谨慎交往异性")
		}
	}
	//日宫（二宫即时辰宫）落留连
	if eg == GongWei[0]+1 ||
		eg == GongWei[1]+1 ||
		eg == GongWei[2]+1 ||
		eg == GongWei[3]+1 ||
		eg == GongWei[4]+1 ||
		eg == GongWei[5]+1 ||
		eg == GongWei[6]+1 {
		//留连:= xiao_tou_xuan.LiuLian
		//fmt.Println("小透玄-->留连：",留连)
		fmt.Println("小透玄-->留连:	日生流连脚板宽，披霞落马走江川，心直口快不触官，惹犯小人闹不安。")
		//时辰宫落大安解卦
		if sg == GongWei[0] ||
			sg == GongWei[1] ||
			sg == GongWei[2] ||
			sg == GongWei[3] ||
			sg == GongWei[4] ||
			sg == GongWei[5] ||
			sg == GongWei[6] {
			fmt.Println("时辰宫-->大安:	大安落位不一般，禄马学堂捧金碗。")
		}
		//时辰落速喜解卦
		if sg == GongWei[0]+2 ||
			sg == GongWei[1]+2 ||
			sg == GongWei[2]+2 ||
			sg == GongWei[3]+2 ||
			sg == GongWei[4]+2 ||
			sg == GongWei[5]+2 ||
			sg == GongWei[6]+2 {
			fmt.Println("时辰宫-->速喜:	财地速喜靠嘴积，不需劳刀一生欢。")
		}
		//时辰宫落小吉解卦
		if sg == GongWei[0]+4 ||
			sg == GongWei[1]+4 ||
			sg == GongWei[2]+4 ||
			sg == GongWei[3]+4 ||
			sg == GongWei[4]+4 ||
			sg == GongWei[5]+4 ||
			sg == GongWei[6]+2 {
			fmt.Println("时辰宫-->小吉:	时临小吉相貌端，吉星贵人来把关。")
		}
		//时辰落宫空亡解卦
		if sg == 0 ||
			sg == GongWei[7]+3 ||
			sg == GongWei[0]+5 ||
			sg == GongWei[1]+5 ||
			sg == GongWei[2]+5 ||
			sg == GongWei[3]+5 ||
			sg == GongWei[4]+5 ||
			sg == GongWei[5]+5 ||
			sg == GongWei[6]+2 {
			fmt.Println("时辰宫-->空亡:	黑道空亡多严官，笔直大路切莫弯。")
		}
	}
	//日宫（二宫即时辰宫）落速喜
	if eg == GongWei[0]+2 ||
		eg == GongWei[1]+2 ||
		eg == GongWei[2]+2 ||
		eg == GongWei[3]+2 ||
		eg == GongWei[4]+2 ||
		eg == GongWei[5]+2 ||
		eg == GongWei[6]+2 {
		//速喜:= xiao_tou_xuan.SuXi
		//fmt.Println("小透玄-->速喜：",速喜)
		fmt.Println("小透玄-->速喜:	日坐速喜巧舌礼，宴客交朋不满席，" +
			"做事一气不到底，半途心猿旁改题，\n三月走到空亡地，克妻不死亦各异，" +
			"速喜遭克守法纪，小人多嗔贵人喜。")
		//月宫落速喜解卦
		if yg == GongWei[0]+2 ||
			yg == GongWei[1]+2 ||
			yg == GongWei[2]+2 ||
			yg == GongWei[3]+2 ||
			yg == GongWei[4]+2 ||
			yg == GongWei[5]+2 ||
			yg == GongWei[6]+2 {
			fmt.Println("月落宫-->速喜:	月同婚姻主分离，各论兄弟皆同理。")
		}
		//时辰宫落大安解卦
		if sg == GongWei[0] ||
			sg == GongWei[1] ||
			sg == GongWei[2] ||
			sg == GongWei[3] ||
			sg == GongWei[4] ||
			sg == GongWei[5] ||
			sg == GongWei[6] {
			fmt.Println("时辰宫-->大安:	大安落位相貌奇。")
		}
		//时辰宫落赤口解卦
		if sg == GongWei[0]+3 ||
			sg == GongWei[1]+3 ||
			sg == GongWei[2]+3 ||
			sg == GongWei[3]+3 ||
			sg == GongWei[4]+3 ||
			sg == GongWei[5]+3 ||
			sg == GongWei[6]+2 {
			fmt.Println("时辰宫-->赤口:	白虎加时犯官气。")
		}
	}
	//日宫（二宫即时辰宫）落赤口
	if eg == GongWei[0]+3 ||
		eg == GongWei[1]+3 ||
		eg == GongWei[2]+3 ||
		eg == GongWei[3]+3 ||
		eg == GongWei[4]+3 ||
		eg == GongWei[5]+3 ||
		eg == GongWei[6]+2 {
		//赤口:= xiao_tou_xuan.ChiKou
		//fmt.Println("小透玄-->赤口：",赤口)
		fmt.Println("小透玄-->赤口:	日坐白虎性情怪，俊情对意敞心怀，语不达心头不抬，" +
			"白虎逢克化以财，\n空见大安父先载，" +
			"若是征兵需挂帅，工农商贾离乡外，禄马学堂遭挫败，押身一枝命安排。")
		//月落宫位　赤口
		if yg == GongWei[0]+3 ||
			yg == GongWei[1]+3 ||
			yg == GongWei[2]+3 ||
			yg == GongWei[3]+3 ||
			yg == GongWei[4]+3 ||
			yg == GongWei[5]+3 ||
			yg == GongWei[6]+2 {
			fmt.Println("月落宫-->赤口:	月日重照面多白，早运玉在坐中埋。")
		}
		//时辰落宫赤口
		if sg == GongWei[0]+3 ||
			sg == GongWei[1]+3 ||
			sg == GongWei[2]+3 ||
			sg == GongWei[3]+3 ||
			sg == GongWei[4]+3 ||
			sg == GongWei[5]+3 ||
			sg == GongWei[6]+2 {
			fmt.Println("时辰宫-->赤口:	临时本位唇舌殃，俊情对意敞心怀，语不达心头不抬。")
		}
	}
	//日宫　落小吉
	if eg == GongWei[0]+4 ||
		eg == GongWei[1]+4 ||
		eg == GongWei[2]+4 ||
		eg == GongWei[3]+4 ||
		eg == GongWei[4]+4 ||
		eg == GongWei[5]+4 ||
		eg == GongWei[6]+2 {
		//小吉:= xiao_tou_xuan.XiaoJi
		//fmt.Println("小透玄-->小吉：",小吉)
		fmt.Println("小透玄-->小吉:	日生小吉人多情，流连桃红泛水盈，身肩背痣做将领，娶妻嫁夫两口井，" +
			"受克小吉福不清，学堂官章巧聪伶，六亲空亡冷如冰，孜身形影犯孤丁，" +
			"多寅速喜不足凭，过房安身劳碌命。")
	}
	//日宫　落空亡
	if eg == 0 ||
		eg == GongWei[0]+5 ||
		eg == GongWei[1]+5 ||
		eg == GongWei[2]+5 ||
		eg == GongWei[3]+5 ||
		eg == GongWei[4]+5 ||
		eg == GongWei[5]+5 ||
		eg == GongWei[6]+2 {
		//空亡:= xiao_tou_xuan.KongWang
		//fmt.Println("小透玄-->空亡：",空亡)
		fmt.Println("小透玄-->空亡:	日生空亡财禄重，女命刚强性气宏，压夫不起运晚通，龙伤子息晚景空，交朋结友忌尾龙。")
		if sg == GongWei[0] ||
			sg == GongWei[1] ||
			sg == GongWei[2] ||
			sg == GongWei[3] ||
			sg == GongWei[4] ||
			sg == GongWei[5] ||
			sg == GongWei[6] {
			fmt.Println("时辰宫-->大安:	时临大安慢排功。")
		}
		if sg == GongWei[0]+2 ||
			sg == GongWei[1]+2 ||
			sg == GongWei[2]+2 ||
			sg == GongWei[3]+2 ||
			sg == GongWei[4]+2 ||
			sg == GongWei[5]+2 ||
			sg == GongWei[6]+2 {
			fmt.Println("时辰宫-->速喜:	空亡空空到对宫，速喜坐时旺财丰。")
		}
		if sg == GongWei[0]+3 ||
			sg == GongWei[1]+3 ||
			sg == GongWei[2]+3 ||
			sg == GongWei[3]+3 ||
			sg == GongWei[4]+3 ||
			sg == GongWei[5]+3 ||
			sg == GongWei[6]+2 {
			fmt.Println("时辰宫-->赤口:	时临白虎在梦中，温柔香刮冷雨风。")
		}
		if sg == GongWei[0]+4 ||
			sg == GongWei[1]+4 ||
			sg == GongWei[2]+4 ||
			sg == GongWei[3]+4 ||
			sg == GongWei[4]+4 ||
			sg == GongWei[5]+4 ||
			sg == GongWei[6]+2 {
			fmt.Println("时辰宫-->小吉:	小吉似吉却土崩。")
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
			fmt.Println("时辰宫-->空亡:	空亡重照牢灾冲。")
		}
	}
}
