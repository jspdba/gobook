<!DOCTYPE html>
<html lang="zh-cn">
<style>
    body {
        padding-top: 70px;
        padding-bottom: 50px;
    }
</style>
{{template "common/header_flat.tpl"}}
<title>str解码</title>
<body>
{{template "common/navibar.tpl"}}

<div class="container">
    <h4>Base64解码</h4>
    <row>
        <div class="table-responsive">
            <table class="table table-bordered">
                <caption>Base64解码</caption>
                <tr>
                    <td>
                        <textarea autocomplete="off" data-provide="typeahead" name="codeEncoded" id="codeEncoded" placeholder="编码字符" class="form-control"></textarea>
                    </td>
                    <td>
                        <div class="btn-group pull-right">
                            <button id="decode" type="button" class="btn btn-primary">解码</button>
                        </div>
                    </td>
                </tr>
                <tr>
                    <td colspan="2">
                        <div id="code"></div>
                    </td>
                </tr>
            </table>
        </div>
    </row>
</div>
{{template "common/script.tpl"}}
<!--消息-->
<link href="/static/css/toastr.min.css" rel="stylesheet"/>
<script src="/static/js/toastr.min.js"></script>

<script type="text/javascript" src="/static/paste/paste.js"></script>
<script>
    $(function () {
        function setData(data) {
            var result="";
            try {
                result = atob(jQuery.trim(data))
            }catch (e){
                result="解析失败"
            }
            $("#code").html(result)
        }

        $("#decode").bind("click",function () {
            var data = $("#codeEncoded").val();
            setData(data)
        });
        //黏贴加密字符
        $('#codeEncoded').pastableTextarea().on('pasteText', function(ev, data){
            setData(data.text)
        });
    })
</script>
</body>
</html>
