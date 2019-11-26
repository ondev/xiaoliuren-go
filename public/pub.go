//内容主要来自一掌经
package public

import "fmt"

//小六壬 六神
type LiuShen struct {
	大安, 留连, 速喜, 赤口, 小吉, 空亡 string
	LiuDao                 //六神继承了六道的特质
}

// 六道
type LiuDao struct {
	鬼道, 佛道, 仙道, 修罗道, 人道, 畜生道 string
}

//十二宫

//鬼道
type GuiDao struct {
	天驿, 天厄 string
}
//佛道
type FoDao struct {
	天福,天贵 string
}
//仙道
type XianDao struct {
	天文, 天寿 string
}

//修罗道
type XiuLuoDao struct {
	天奸, 天艺 string
}

// 人道
type RenDao struct {
	天权, 天孤 string
}

//畜生道
type ChuShengDao struct {
	天刃, 天破 string
}

//实现结构体的方法
//人道
func (ren *RenDao) ren_tq() string {

	return ren.天权
}
func (ren *RenDao) ren_tg() string {

	return ren.天孤
}
//鬼道
func (gui *GuiDao)gui_te()string{
	return gui.天厄
}
func (gui *GuiDao)gui_ty()string{
	return gui.天驿
}
//仙道
func (xian *XianDao)xian_ts() string  {
	return xian.天寿
}
func (xian *XianDao)xian_tw() string  {
	return xian.天文
}
//修罗道
func (xiuluo *XiuLuoDao)xiuluo_ty()string  {
	return xiuluo.天艺
}
func (xiuluo *XiuLuoDao)xiuluo_tj()string  {
	return xiuluo.天奸
}
//畜生道
func (chusheng *ChuShengDao)chusheng_tr()string  {
	return chusheng.天刃
}
func (chusheng *ChuShengDao)chusheng_tp()string  {
	return chusheng.天破
}
//佛道
func (fodao *FoDao)fodao_tf()string  {
	return fodao.天福
}
func (fodao *FoDao)fodao_tg()string  {
	return fodao.天贵
}

//仙道
func Display_XianDao()  {
	xian:= XianDao{
		天文:"命遇天文秀氣清，聰明智慧意惺惺，男才女秀身清吉，滿腹文章錦繡成。" +
			"此星照命，主人聰明伶俐，學識過人，作事和美，若逢天貴天福天藝星相助，" +
			"定主鰲頭獨占，虎榜登名，金階玉陛之人也，若遇天權天刃星者，" +
			"文武多才，乃為上命，如遇破厄孤驛及犯重者，乃多學少成，" +
			"不為書算文墨之輩，必為雲遊湖海之人，乃手藝術士之下命也。",
		天寿:"夫妻生時命最長，上恭下敬性溫良，一聞千悟心慈善，喜怒中間有主張。" +
			"此星照命，主人長壽康健，智慧聰明，作事溫良有救人之心，無傷人之意，" +
			"恩中招怨，作事樸實，眾人欽敬，平生安穩，有始有終，喜怒不形，" +
			"若得天權福貴刃星相助，必主寬洪大量，福壽綿綿，犯重者有壽無福，犯孤破厄星乃中命也。",
	}
	fmt.Println(xian.xian_tw())
	fmt.Println(xian.xian_ts())
}
//修罗道
func Displsy_XiuLuo()  {
	x:= XiuLuoDao{
		天奸:"大如滄海細如毛，佛口蛇心兩面刀，奸狡狠謀藏毒性，意多翻覆最難調。" +
			"此星照命，主人一生勞碌，啾唧奔波，指東說西，機變難測，" +
			"若得天貴福星相助，財帛豐盈，亦為上命，若逢權刃星者，" +
			"必為奸權殘忍之人，言清行濁，執性凶謀，有善人之心，無容人之量，" +
			"貪嗔太重，非善人也，若逢孤厄破驛，定為慳貪嫉妒之小人，乃下命也。",
		天艺:"天藝生人性最靈，將南作北逞多能，諱為見靈機關巧，到處和同作事勤。" +
			"此星照命，主人多智多能，機巧近貴，若犯重者，主資質昏鈍，懶惰愚頑，" +
			"多學少成，匠作用力之輩，若得天權貴福文壽俱全，剛柔相濟，雖為藝術，" +
			"亦可成立，若為天孤可為僧道之出類者，乃中命也，若逢破厄，則藝舉無成終為下命。",
	}
	fmt.Println(x.xiuluo_tj())
	fmt.Println(x.xiuluo_ty())
}
//佛道
func Display_FoDao()  {
	f:=FoDao{
		天福:"命逢天福是生時，定然倉庫有盈餘，寬洪大量根基穩，財帛光華百福齊。" +
			"此星照命，主人受福清閑，性情自在，度量寬洪，根基穩實，又得權刃相扶，" +
			"衣帛充足，倉庫盈餘，堆金積玉之命，若犯重者，衣祿不多，" +
			"若驛孤奸破星者，必主貪慳嫉妒，衣祿艱難之下命也。",
		天贵:"",
	}
	fmt.Println(f.fodao_tf())
	fmt.Println(f.fodao_tg())
}
//畜生道
func Display_ChuSheng()  {
	c:=ChuShengDao{
		天刃:"天刃為人性大剛，是非終日要爭強，持刀弄斧刑心重，好似將軍入戰場。" +
			"此星照命，主人一生剛很性格，躁暴自做，自是不受人觸，受不得閑氣，風火性過端然無事，" +
			"若得權貴，福星為人，不俗禮義，足以化強暴乃上命也，" +
			"若逢孤破奸厄，膽大心粗，形體殘疾，不免斷髮身死乃下命也，" +
			"惡星少而吉星多者，亦為中命犯重必主殘疾。",
		天破:"時辰落在天破宮，堆金積玉也成空，夜眠算計圖家富，鈔袋誰知有蛀蟲 " +
			"此星主財帛空虛，祖業耗散，若得權貴福星相助，亦為中命，如遇驛刃孤厄犯重者，" +
			"作事艱難，重重破敗，浮浪東西之下命也。",
	}
	fmt.Println(c.chusheng_tp())
	fmt.Println(c.chusheng_tr())
}
//人道
func Display_RenDao() {
	r := RenDao{
		天权: "時辰落在天權星，性格操持志氣雄，作事差遲人也喜，一呼百喏有威風。" +
		"時辰落在天權星，性格操持志氣雄，作事差遲人也喜，一呼百喏有威風。" +
		"此星在命，主人聰明，俊秀灑落，襟懷有權有勢，多智多能，若逢貴福文壽星相助者，人人欽敬，" +
		"權而無權乃中命也，若逢厄破孤驛在命者，作事勞力，財帛不聚，未能先能，未會先會，浮浪中命也。",
		天孤:"時辰若逢此天孤，六親兄弟有如無，空作空門清靜客，總有妻兒情分疏。" +
		"此星照命，主一生孤獨，男人得之，六親無分，女人得之，剋子妨夫，" +
		"孤星犯重者，反不為孤，必為半僧半俗，若得權福貴壽星相助，乃上命也，" +
		"亦不免少年刑剋，若逢破驛奸厄刃星，必為雲水漂流下命也，凡選故出之命，要看孤星為主。"}
	fmt.Println(r.ren_tq())
	fmt.Println(r.ren_tg())
}
//鬼道
func Display_GuiDao()  {
	g:= GuiDao{
		天驿:"人道若逢天驛星，搬移離祖不曾停，身心不得片時靜，走遍天涯是未甯" +
			"此星照命，主人離鄉別井，骨肉情多勞碌，身心自成自立之命，" +
			"若逢福權貴刃壽五星者，必主官供，給車馬相隨乃顯榮之命，" +
			"若逢孤破厄星，猶如風吹樹葉水上浮萍，心猿意馬，奔馳不定，" +
			"方外雲遊江湖上之下命也，若犯重而刃厄相沖者必為流徒之類。",
		天厄:"",
	}
	fmt.Println(g.gui_te())
	fmt.Println(g.gui_ty())
}

