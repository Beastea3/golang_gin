package services
//
//import (
//	"github.com/beastea/golang-gin/models"
//)
//
//type Joke = models.Joke
//
////var jokes = []Joke{
////	Joke{1, 0,"停车场，我看到一个女孩用湿纸巾擦了擦车屁股，然后一口亲了上去。我走过去：“你需要帮忙吗？她不好意思地松开口，撩了一下头发：“我倒车把车 屁 股撞了一个小小的瘪，我看能不能把瘪给吸回来！”"},
////	Joke{2, 0,"爸爸，为什么每次我一犯错，你和妈妈就要打我？”“宝宝，你知道吗？玉不琢，不成器！你妈打你，是因为她对你有很高的期望。”“哦，那爸爸你呢？”“我不一样，就是单纯地想打你！” "},
////	Joke{3, 0,"舍友跟宿舍里放了个臭屁，还在那边狡辩：这不叫放屁！这是菊花在呼吸！我：你这光呼了，没吸啊！舍友：你们不是吸了么！"},
////	Joke{4, 0,"我已经48小时没敢回家了！！！老婆搞卫生、不小心摔碎了三岁半女儿的存钱罐，摔出来只有满地的小石头。现在家里的俩女人依然处于嗜 血 狂 暴状态！ "},
////	Joke{5, 0,"同事聚会，饭桌上，经理：“给每人先来一杯果汁。”服务员端来几杯果汁放到女同事面前，经理：“怎么只给女的果汁？”服务员：“你说的，给美人先来一杯果汁。。。”经理。。。"},
////	Joke{6, 0,"猫A：“你为什么要把抓到的老鼠撕成碎片？你不觉得这样做很残忍吗？” 猫B：“可是鼠片真的很好吃。”"},
////	Joke{7, 0,"早上，一个流浪汉对我说：“哥你有零钱么？我饿坏了。” 我摇着头就直接走了。。。哪有人饿了吃硬币的啊？！"},
////	Joke{8, 0,"我问室友：“你知道什么是父债子还吗？” 室友说：“你欠我的钱不会是想等你儿子长大以后才赚钱来还我吧。” 我笑着说：“当然不是，我是说父亲欠的钱用儿子来抵债。”然后递给他一张卫生纸。"},
////}
//
////func JokeHandler(db) []Joke{
////	jokes := []Joke{}
////	err := db.C(models.CollectionJoke).Find(nil).Sort("-update_on").ALL(&jokes)
////	return jokes
////}
//
//func LikeJoke(jokeid int) *[]Joke{
//
//		for i := 0; i< len(jokes); i++ {
//			if jokes[i].ID == jokeid {
//				jokes[i].Likes += 1
//			}
//		}
//		return &jokes
//
//}