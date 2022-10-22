package main

import (
	"fmt"
	"math/rand"
	"time"
)

//定义了一个全局的map来存储问题的id
type Ids map[int64]bool
var IDS Ids

func (ids Ids)Check(id int64)bool {
	//检查id是否已经存在
	return ids[id]
}

func (ids Ids)Add(id int64) {
	//将id添加到表中
	if ids == nil {
		ids = make(map[int64]bool)
	}
	ids[id]=true
}

func GenId()int64 {
	//生成一个id，这个id不会与已经存在的id重名
	var id int64
	rand.Int63()
	for IDS.Check(id) {
		id = rand.Int63()
	}
	IDS.Add(id)
	return id
}

type  Question struct {
	Id        int64
	title     string
	content   string
	like 	  int         //点赞数
	username string
	comments []string
	commentCount int
	CreatedAt time.Time
	DeletedAt time.Time
	UpdateAt  time.Time
}

func NewQuestion()Question {
	q := Question{
		Id:        GenId(),
		CreatedAt: time.Now(),
		like: 0,
		commentCount: 0,
	}
	return q
}

var Questions map[string]Question
func Publish() {
	//发布一个问题,将这个问题记录在Questions中
	q := NewQuestion()
	title := ""
	content:=""
	fmt.Println("请输入问题的题目：")
	fmt.Scanln(&title)
	fmt.Println("请输入你要发布的问题的具体内容: ")
	fmt.Scanln(&content)
	q.content=content
	q.title=title
	Questions[title]=q
}

func (q *Question)DianZan() {
	q.like++
}

func (q *Question)Comment() {
	tmp:=""
	fmt.Println("请输入评论: ")
	fmt.Scanln(&tmp)
	q.comments=append(q.comments,tmp)
	q.commentCount++
}

func FindQuestion(title string)(Question,bool) {
	//按照问题内容查找问题，返回问题结构体，map的值不能取址，因此不能返回指针
	//由于返回的是结构体，在对这个问题结构体进行操作后不会改变questions表中的问题，因此每次操作完了之后要更新一下。
	q,ok:=Questions[title]
	return q,ok
}

func Init() {
	//初始化id表
	IDS = make(map[int64]bool)
	//初始化Questions表
	Questions = make(map[string]Question)
}

func main() {
	//用户可以选择发布问题，也可以选择搜索想要查询的问题
	//用户查询问题的时候可以评论，点赞
	//妈的数电还没开始学，下周考试了，我就熟悉一下语法，随便写的，大概率是有不合理的地方的，见谅
	Init()
	for {
		var flag int
		fmt.Println("1、发布问题")
		fmt.Println("2、搜索问题")
		fmt.Scanln(&flag)
		switch flag{
		case 1:Publish()
		case 2:
			var title string
			fmt.Println("请输入要查找到问题: ")
			fmt.Scanln(&title)
			q,ok := FindQuestion(title)
			if !ok {
				fmt.Println("没有此问题")
				break
			}
			fmt.Println(q.title)
			fmt.Println("问题描述:  ",q.content)
			fmt.Println("点赞：",q.like,"评论：",q.commentCount)
			fmt.Println("当前评论:")
			for key,comment:=range q.comments {
				fmt.Println(fmt.Sprintf("%d、%s",key+1,comment))
			}
			fmt.Println("-------------------------------------")
			fmt.Println("1、点赞")
			fmt.Println("2、评论")
			var tmp int
			fmt.Scanln(&tmp)
			switch tmp {
			case 1:q.DianZan()
			case 2:q.Comment()
			}
			Questions[title]=q//更新问题
		}
	}
}