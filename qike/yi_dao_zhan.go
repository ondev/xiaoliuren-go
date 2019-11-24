package qike

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//十二属相宫位
func ShuXiang() string {
	fmt.Println("\n输入你的属相(汉字或者全拼)")
	//fmt.Println("十二属相-->	鼠　牛　虎　兔　龙　蛇　马　羊　猴　鸡　狗　猪")
	input := bufio.NewReader(os.Stdin)
	input_s, err := input.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	s := strings.ToLower(input_s) //转换大写字母为小写
	fmt.Println("你输入的属相是-->:", strings.TrimSpace(s))
	return s
}

//er 二宫定宫　ri 二宫起课数字以此数字定几落几宫　s 属相
func YiDaoZhan(yg uint, eg uint, sg uint, ri uint) {
	//s := strings.TrimSpace(ShuXiang())//去掉尾部换行
	GongWei := [8]uint{1, 7, 13, 19, 25, 31, 37, 43}
	//日　落二宫　以日论 -->大安
	if eg == GongWei[0] || eg == GongWei[1] || eg == GongWei[2] || eg == GongWei[3] || eg == GongWei[4] || eg == GongWei[5] || eg == GongWei[6] {
		fmt.Println("一刀斩-->大安:	大安五行为木数为1 7 金灾之命最克夫　自身流年忌羊猴牛" +
			"日占大安得妻福压祸　见空亡反克妻，妻身体不好心脏病或肝胆有病，岳父娶二妻或独守岳家" +
			"男丁孤单 多个必有早走之人 大安最忌合伙经商")
		//属相相关卦辞
		s := strings.TrimSpace(ShuXiang()) //去掉尾部换行
		//fmt.Println("s de 值 ",s)
		if s == "虎" || s == "hu" || s == "马" || s == "ma" || s == "狗" || s == "gou" {
			fmt.Println("一刀斩-->属相:	", s+" 此属相有车关")
		} else {
			fmt.Println("此属相在此无信息...")
		}
		//大安落不同位置卦辞
		if ri == 10 {
			fmt.Println("一刀斩-->大安落位:  十落大安兄弟两多一个婚必败，兄弟无义不相合，" +
				"学堂必有挫折有技术压身，借妻之福应寅申巳亥")
		}
		if ri == 3 {
			fmt.Println("一刀斩-->大安落位:  三落大安脾气急躁遇事不稳肝火旺，脊柱脑神经病，" +
				"头脚带伤，上葬亡命格不须劳碌财顺可得，一生孤单不免刑克 自身多见伤灾，" +
				"取弟不多扎根不会过两名，压兄运不见抬 互换运可发家，早婚必离")
		}
		if ri == 2 {
			fmt.Println("一刀斩-->大安落位:  二落大安离乡发展 双二落依母亲性格刚烈必见孤单，" +
				"克父缘博克自身体弱，车关占命一落克父母，婚不白头克兄弟运气，克子女婚姻见破坟地必是低坟，" +
				"寻高岗青年灾多，刑克不顺腿脚带伤")
		}
		if ri == 5 {
			fmt.Println("一刀斩-->大安落位:  五落大安妻家姐妹少两个，正位三个克岳母身体不好，" +
				"五落大安灾命最凶，中年克妻在异途 若寅(木虎)相来相合妻花当年必入土")
		}
	}
	//日　落二宫　以日论 -->留连
	if eg == GongWei[0]+1 || eg == GongWei[1]+1 || eg == GongWei[2]+1 || eg == GongWei[3]+1 || eg == GongWei[4]+1 || eg == GongWei[5]+1 || eg == GongWei[6]+1 {
		fmt.Println("一刀斩-->留连:	流连五行为水数为2 8 流连在天为翻天印在地为走马星，" +
			"无论男女定见胸腿有伤，申子辰合来把肾伤，寅午戌未虎口张，同时遇争官司上，" +
			"流连阳宅不出人才，两流连为门户交叠，女占此数三年内婆门必有变更，唯有男丁落地可解，" +
			"男占此数远方必有建树，")
		//属相宫位马
		s := strings.TrimSpace(ShuXiang()) //去掉尾部换行
		if s == "马" || s == "ma" {
			fmt.Println("一刀斩-->属相	", s+" 马星占命走路不稳心思不宁")
		} else {
			fmt.Println("此属相在此无信息...")
		}
		//月宫落留连
		if yg == GongWei[0]+1 || yg == GongWei[1]+1 || yg == GongWei[2]+1 || yg == GongWei[3]+1 || yg == GongWei[4]+1 || yg == GongWei[5]+1 || yg == GongWei[6]+1 {
			fmt.Println("一刀斩-->月宫: 留连	月日流连克父母，父母分离，转别是属鸡属兔的人")
		}
		//时辰落速喜
		if sg == GongWei[0]+2 || sg == GongWei[1]+2 || sg == GongWei[2]+2 || sg == GongWei[3]+2 || sg == GongWei[4]+2 || sg == GongWei[5]+2 || sg == GongWei[6]+2 {
			fmt.Println("一刀斩-->时宫: 速喜	日流连时速喜一生财运多起伏再见空亡伤老父，" +
				"婚姻有再度，刀伤脓血不龟家中出武廉，龙马之年灾还子午卯酉，走动必多 同缘中必有一王姓。")
		}

	}
	//日　落二宫　以日论 -->速喜
	if eg == GongWei[0]+2 || eg == GongWei[1]+2 || eg == GongWei[2]+2 || eg == GongWei[3]+2 || eg == GongWei[4]+2 || eg == GongWei[5]+2 || eg == GongWei[6]+2 {
		fmt.Println("一刀斩-->速喜:	速喜五行为火数为3 9克老人伤兄弟，上见兄弟婚姻有分张，" +
			"十个速喜九个心 脱筰性悚火气大，婚姻最忌火命人必见半路离伤 火逢速喜一生尽犯指卦星，小人随身贵人为土相，" +
			"坟地南临少亡坟东向大路，母有暗痰不免刀伤，命带官司从时起")
		//申子辰属相解卦
		s := strings.TrimSpace(ShuXiang()) //去掉尾部换行
		if s == "猴" || s == "hou" || s == "鼠" || s == "shu" || s == "龙" || s == "long" {
			fmt.Println("一刀斩-->属相:", s+" 申子辰人，早运不强年过3 6得运添财")
		} else {
			fmt.Println("此属相在此卦中无信息...")
		}
		//时辰宫落速喜 空亡
		if sg == GongWei[0]+2 || sg == GongWei[1]+2 || sg == GongWei[2]+2 || sg == GongWei[3]+2 || sg == GongWei[4]+2 || sg == GongWei[5]+2 || sg == GongWei[6]+2 {
			fmt.Println("一刀斩-->时宫: 速喜	时辰落速喜空亡克父母")
		}
		if sg == 0 || sg == GongWei[7]+3 || sg == GongWei[0]+5 || sg == GongWei[1]+5 || sg == GongWei[2]+5 || sg == GongWei[3]+5 || sg == GongWei[4]+5 || sg == GongWei[5]+5 || sg == GongWei[6]+2 {
			fmt.Println("一刀斩-->时宫: 速喜	时辰落速喜空亡克父母")
		}
		if ri == 7 {
			fmt.Println("七落速喜祖基空离，祖成家不靠兄弟自经营，" +
				"吃得前运若力得 后运添娶进小妻不长久　娶得女妻必逃花")
		}
	}
	//日　落二宫　以日论 -->赤口
	if eg == GongWei[0]+3 || eg == GongWei[1]+3 || eg == GongWei[2]+3 || eg == GongWei[3]+3 || eg == GongWei[4]+3 || eg == GongWei[5]+3 || eg == GongWei[6]+2 {
		fmt.Println("一刀斩-->赤口:	白虎五行为金在天为天煞孤星在地为飞血连环，白虎占命不见血星是笑谈" +
			"切忌需防白虎是克双亲，最凶的煞星，父母若不见生离必见死别，" +
			"女占白虎为拔婚煞 未见婚姻一个安婚前，若是身花开 婚姻二度去伤胎，" +
			"婚姻忌走双岁年就是富贵亦旧，连姻缘见出武廉三十大旬必见桃花，出头年" +
			"白虎流年起法不同 只从日上起　不论时上通行年，" +
			"推运最好算同把十二原神串龙虎相斗是祸胎，婚姻破败命中带双虎同山不安泰不见铁骑被人抬。")
		//时辰落白虎
		if sg == GongWei[0]+3 || sg == GongWei[1]+3 || sg == GongWei[2]+3 || sg == GongWei[3]+3 || sg == GongWei[4]+3 || sg == GongWei[5]+3 || sg == GongWei[6]+2 {
			fmt.Println("一刀斩--> 时宫: 赤口	白虎见重为白虎入刑堂格 学堂之上必见灾星对宫之运")
		}
	}
	//日宫落小吉
	if eg == GongWei[0]+4 || eg == GongWei[1]+4 || eg == GongWei[2]+4 || eg == GongWei[3]+4 || eg == GongWei[4]+4 || eg == GongWei[5]+4 || eg == GongWei[6]+2 {
		fmt.Println("一刀斩-->小吉:	小吉五行金水最难辨，在天为灵宝司命，在地为五寿司命 " +
			"与流连遥对为阴阳界限的门户，小吉落命必克双亲，特别大命为火者" +
			"生意场上没有命，带小吉的吃亏人。")
		//属相解卦
		s := strings.TrimSpace(ShuXiang()) //去掉尾部换行
		if s == "羊" || s == "yang" || s == "狗" || s == "gou" {
			fmt.Println("一刀斩-->属相:	", s+" 未戌属相者十五之前克亲")
		} else if s == "龙" || s == "long" || s == "狗" || s == "gou" || s == "牛" || s == "niu" || s == "羊" || s == "yang" {
			fmt.Println("一刀斩-->属相:	", s+" 辰戌丑未克婚姻")
		} else if s == "虎" || s == "hu" || s == "猴" || s == "hou" || s == "蛇" || s == "she" || s == "猪" || s == "zhu" {
			fmt.Println("一刀斩-->属相:	", s+" 寅申巳亥妨自身")
		} else {
			fmt.Println("此属相在此无信息...")
		}
		//二落
		if ri == 2 {
			fmt.Println("一刀斩-->小吉落位:  二落小吉为大命桃花若是女带必不差男" +
				"忌2 5 8上行桃花　此为生死桃花劫煞　必在女人身上丧命" +
				"一生转折都在子午卯酉逢这四年　忌见孝事")
		}
	}
	if eg == 0 || eg == GongWei[0]+5 || eg == GongWei[1]+5 || eg == GongWei[2]+5 || eg == GongWei[3]+5 || eg == GongWei[4]+5 || eg == GongWei[5]+5 || eg == GongWei[6]+2 {
		fmt.Println("一刀斩-->空亡:	空亡五行为土数为6 9 12 为万物生养之数，空亡照命男占厚重，女占孝顺，" +
			"姐妹不挨肩必是重婚人命中必犯刀头煞，四番人逢日占空亡　父母婚姻见波折　" +
			"若是桃花身自占　逃不过晚年孤单，婚姻最忌寅申巳亥，女占空亡无穷命　" +
			"必有财源库 个性倔强不服输　有子女中豪杰土　多厚胸必饱满　" +
			"胸前背后必有一颗克天痣，防40大旬克夫　女命如若有空亡命中必见伤胎疾，" +
			"早年福大晚必穷　早见辛若是真福　辰戌合空最伤姐妹，铁转同血三婚无差，" +
			"空亡在数为六为本根数，推流年一定要加本根数，马前名为六数归天地人三数实为九数。")
	}
}
