package service

import (
	"sync"
	"gobook/models"
	"strings"
	"regexp"
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
)

func GetChapterContent(selector string,chapters  []*models.Chapter,threadsCount int)  {
	lenChapters:=len(chapters)

	itemCount:=lenChapters/threadsCount
	var wg sync.WaitGroup

	if (threadsCount>=lenChapters){
		itemCount=1
		for i:=0;i<lenChapters;i++{
			wg.Add(1)
			go DownloadOne(&wg,chapters[i],selector)
		}
	}else{
		for start,i:=0,0;i<threadsCount;i++{
			end:=start+itemCount
			wg.Add(1)
			if i==threadsCount-1{
				beego.Info(start,":-",lenChapters)
				go Download(&wg,chapters[start:],selector)
				break
			}else{
				beego.Info(start,":",end,lenChapters)
				go Download(&wg,chapters[start:end],selector)
			}
			start=end
		}
	}
	wg.Wait()
}

func Download(wg *sync.WaitGroup,chapters []*models.Chapter,selector string){
	defer wg.Done()
	for _,chapter:=range chapters{
		if chapter.Url!=""{
			beego.Info("DownLoad="+chapter.Url)
			content:=GetContent(chapter.Url,selector)
			chapter.Content=content
		}
	}
}
func DownloadOne(wg *sync.WaitGroup,chapter *models.Chapter,selector string){
	defer wg.Done()
	if chapter.Url!=""{
		content:=GetContent(chapter.Url,selector)
		chapter.Content=content
		beego.Info("DownLoad="+chapter.Url)
	}
}

func GetContent(url string, selector string) string{
	doc, err := goquery.NewDocument(url)
	if err == nil {
		if doc != nil {
			content:= doc.Find(selector).Text()
			doc=nil
			return AddLine(content)
		}
	}else{
		beego.Error(err)
	}
	return ""
}
//http://www.biquge.tw/81_81260/
func GetBookInfo(url string) *models.Book{
	doc, err := goquery.NewDocument(url)
	book:=&models.Book{}
	if err == nil {
		if doc != nil {
			name:= doc.Find("#info > h1").Text()
			maker:= doc.Find("#info > p:nth-child(2)").Text()
			content:= doc.Find("#intro").Text()
			book.Name=name
			book.Maker=maker
			book.Content=content
			book.ChapterRules="#list > dl > dd"
			book.ContentRules="#content"
			doc=nil
			return book
		}
	}else{
		beego.Error(err)
	}
	return book
}

func Search(key string) *models.Book{
	key = strings.NewReplacer("\n", "", "\r", "", ";","", " ","").Replace(key)
	/*if key!=""{
		u,_:=url.Parse(key)
		key=u.EscapedPath()
	}*/
	//#results > div.result-list > div:nth-child(1) > div.result-game-item-detail > p
	//*[@id="results"]/div[2]/div[1]/div[2]/p
	url:="http://zhannei.baidu.com/cse/search?ie=utf-8&q="+key+"&entry=1&s=16829369641378287696&nsid="
	doc, err := goquery.NewDocument(url)
	book:=&models.Book{}
	if err == nil {
		if doc != nil {
			selector:="#results > div.result-list > div:nth-child(1) > div.result-game-item-detail"
			replacer := strings.NewReplacer("\n", "", "\r", "", ";","", " ","")
			doc.Find(selector).EachWithBreak(func(i int, contentSelection *goquery.Selection) bool{
				name,_:= contentSelection.Find("h3 > a").Attr("title")
				Url,_:= contentSelection.Find("h3 > a").Attr("href")
				maker:= contentSelection.Find("div > p:nth-child(1) > span:nth-child(2)").Text()
				maker=replacer.Replace(maker)
				//#results > div.result-list > div:nth-child(1) > div.result-game-item-detail > p
				content:= contentSelection.Find("p").First().Text()
				content=replacer.Replace(content)
				book.Name=name
				book.Maker=maker
				book.Content=content
				book.ChapterRules="#list > dl > dd"
				book.ContentRules="#content"
				book.Url=Url
				doc=nil
				return false;
			})
		}
	}else{
		beego.Error(err)
	}
	return book
}

func AddLine(str string) string{
	str=strings.Replace(str,"&nbsp;&nbsp;&nbsp;&nbsp;","\r\n&nbsp;&nbsp;&nbsp;&nbsp;",-1)
	str=strings.Replace(str,"　　　　","\r\n　　　　",-1)
	return strings.Replace(str,"\r\n","<br/>",-1)
}

func GetUrlInfo(url string,selector string,limit int) []*models.Chapter{
	doc, err := goquery.NewDocument(url)

	chapters := []*models.Chapter{}
	if err == nil && doc!=nil {
		doc.Find(selector).EachWithBreak(func(i int, contentSelection *goquery.Selection) bool{
			title := contentSelection.Find("a").Text()
			//log.Println("第", i + 1, "个帖子的标题：", title)
			//index:=getIndex(utils.Convert2Digit(utils.FindDigit(title)))

			href, _ := contentSelection.Find("a").Attr("href")
			if !strings.HasPrefix(href, "http") {
				href = getHost(url) + href
			}
			chapter:=&models.Chapter{
				Title:title,
				//Index:index,
				Url:href,
			}
			chapters=append(chapters,chapter)

			if limit>0 && len(chapters)==limit{
				return false;
			}
			return true
		})
	}
	return chapters
}

func getHost(url string) string{
	url=strings.Replace(url,"http://","", -1)
	url=url[0:strings.Index(url,"/")]
	return "http://"+url
}

func getIndex(s string) int{

	valid ,err:= regexp.Compile("[0-9]{1,5}")
	if !valid.MatchString(s){
		return -1
	}
	if err==nil{
		data := valid.FindAllStringSubmatch(s, -1)
		if len(data)>=0{
			if num,err :=strconv.Atoi(strings.Join(data[0],"")); err==nil{
				return num
			}
		}
	}
	return -1
}

func ResetJob(){
	models.ResetJob();
}