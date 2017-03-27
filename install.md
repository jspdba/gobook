# gobook
 
 https://github.com/jspdba/gobook.git
 
 ## create a new repository on the command line
     echo "# gobook" >> README.md
     git init
     git add README.md
     git commit -m "first commit"
     git remote add origin https://github.com/jspdba/gobook.git
     git push -u origin master
 ## or push an existing repository from the command line
     git remote add origin https://github.com/jspdba/gobook.git
     git push -u origin master
 ## 安装包
     Dependency list (8):
     -> github.com/PuerkitoBio/goquery
     -> github.com/astaxie/beego
     -> github.com/fatih/color
     -> github.com/go-sql-driver/mysql
     -> github.com/henrylee2cn/mahonia
     -> github.com/jakecoffman/cron
     -> github.com/juju/persistent-cookiejar
     -> github.com/sluu99/uuid

## 使用代理
        代理下载地址  https://github.com/shadowsocks/shadowsocks-windows/releases
        在终端设置： 
        git config --global http.proxy http://127.0.0.1:1080 
        git config --global https.proxy https://127.0.0.1:1080
        
        默认不设置代理： 
        git config --global --unset http.proxy 
        git config --global --unset https.proxy
        
        查看已经设置的值：
        git config http.proxy
## gopm 部分包需要用gopm下载 
    go get -u github.com/gpmgo/gopm
1. gopm 安装：
>这个十分简单只需一go条命令就可以了：
    go get -u github.com/gpmgo/gopm
    
2. 使用 gopm安装需要的包
>gopm 具有丰富的包管理功能，具体的管理命令可以参考官方文档（官方文档有中文版　各位爽不爽）链接
    这里只需要一条命令就可以搞定了：
    gopm bin -d $GOPATH/bin  PackageName
    gopm bin -d %GOPATH%/bin  PackageName
## gopm 使用
    https://github.com/gpmgo/docs/blob/master/zh-CN/README.md
    https://github.com/gpmgo/docs/blob/master/zh-CN/Quickstart.md
## mysql免安装版配置
    http://blog.csdn.net/q98842674/article/details/12094777/