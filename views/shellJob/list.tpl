<!DOCTYPE html>
<html lang="zh-cn">
{{template "common/header_flat.tpl"}}
<title>shell任务列表</title>
<body>
{{template "common/navibar.tpl"}}
<div class="container">
    <div class="panel panel-primary">
        <div class="panel-heading">
            <div class="title">
                shell任务列表
            </div>
        </div>
        <div class="panel-body">
            {{range .jobs}}
            <div class="row">
                <div class="col-md-5 col-sm-5">{{.Name}}:{{.Cron}}</div>
            </div>
            {{end}}
            <ul id="page"></ul>
        </div>
    </div>
</div>
{{template "common/script.tpl"}}
<!--消息-->
<link href="/static/css/toastr.min.css" rel="stylesheet"/>
<script src="/static/js/toastr.min.js"></script>
</body>
</html>
