<!DOCTYPE html>
<html lang="zh-cn">
{{template "common/header_flat.tpl"}}
<title>图书列表</title>
<body>
{{template "common/navibar.tpl"}}
<div class="container">
    <div class="panel panel-primary">
        <div class="panel-heading">
            <div class="title">
                <a href="{{urlfor "BookController.Edit"}}" class="btn btn-primary">添加</a>
                <a id="check" class="btn btn-primary">检查更新</a>
            </div>
        </div>
        <div class="panel-body">
            <div class="row">
                {{range .page.List}}
                    <div class="col-sm-6 col-md-4">
                        <div class="thumbnail">
                            <div class="caption">
                                <h3>{{.Name}} <div class="news" bid="{{.Id}}" data-url="{{urlfor "ChapterController.HasNewChapter" ":id" .Id}}"></div></h3>
                                <p>
                                    <a role="button" class="btn btn-info btn-block toRead" link="{{urlfor "ChapterController.ListByLog" ":tag" "T_A_G" ":id" .Id}}">继续阅读</a>
                                    <a type="button" class="btn btn-info btn-block"  href="{{urlfor "ChapterController.List" ":id" .Id}}">章节列表</a>
                                    <button type="button" class="btn btn-default  btn-block updateChapter" link="{{urlfor "BookController.TaskUpdate" ":id" .Id}}">更新章节</button>
                                </p>
                            </div>
                        </div>
                    </div>
                {{end}}
            </div>
            <ul id="page"></ul>
        </div>
    </div>

    <!-- Modal -->
    <div class="modal fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
        <div class="modal-dialog">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal"><span aria-hidden="true">&times;</span><span class="sr-only">Close</span></button>
                    <h4 class="modal-title" id="myModalLabel">添加 Tag</h4>
                </div>
                <div class="modal-body">
                    <input id="tag" type="text" value="tag1">
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                    <button id="useTag" type="button" class="btn btn-primary">使用</button>
                </div>
            </div>
        </div>
    </div>
</div>
{{template "common/script.tpl"}}
<!--消息-->
<link href="/static/css/toastr.min.css" rel="stylesheet"/>
<script src="/static/js/toastr.min.js"></script>

    <script>
        $(".updateChapter").bind("click",function () {
            var url=$(this).attr("link")
            console.log(url)
            url && $.getJSON(url,function (data) {
                console.log(data)
                if (data.code==0){
                    alert("success")
                }else{
                    alert(data.msg)
                }
            })
        });

        $("#check").bind("click",function () {
            $(".news").each(function () {
                var url=$(this).data("url");
                var that=$(this);
                $.getJSON(url,function (data) {
                    if(data.code==0){
                        $(that).html(data.data)
                    }
                })
            })
        })

        $(".toRead").bind("click",function () {
            var url = $(this).attr("link");
            console.log(url);
            var tag = Cookies.get("__TAG");
            if(!tag){
                $("#useTag").attr("link",url);
                $('#myModal').modal('show');
                return
            }
            url = url.replace("T_A_G",tag);
            window.location.href=url
        });

        //使用tag
        $("#useTag").bind("click",function (data) {
            var value = $("#tag").val();
            if(value){
                var url = $(this).attr("link");
                url = url.replace("T_A_G",value)
                $('#myModal').modal('hide')
                window.location.href=url
            }
        });
        $(".export").bind("click",function () {
            var link=$(this).attr("link")

            $.getJSON(link,function (data) {
                if(data.code==0){
                    alert("上传成功")
                }else{
                    alert("上传失败="+data.msg)
                }
            })
        })
        $(".localUpdate").bind("click",function () {
            var link=$(this).attr("link")

            $.getJSON(link,function (data) {
                if(data.code==0){
                    alert("远程更新空章节成功")
                }else{
                    alert("远程更新空章节失败="+data.msg)
                }
            })
        })
    </script>
</body>
</html>
