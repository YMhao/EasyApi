package web

import (
	"fmt"
	"html/template"
)

type IndexInfo struct {
	Name        string
	Description string
	SwaggerJSON string
}

// IndexInfo 首页信息
type HtmlIndexInfo struct {
	Name        string
	Description string
	SwaggerJSON template.HTML
}

func (i *IndexInfo) HtmlIndexInfo() *HtmlIndexInfo {
	html := fmt.Sprintf(`
		<div class="panel-body">
			<pre class="code json"">%s</pre>
		</div>
		`, i.SwaggerJSON)
	return &HtmlIndexInfo{
		Name:        i.Name,
		Description: i.Description,
		SwaggerJSON: template.HTML(html),
	}
}

const IndexPage = `
<!DOCTYPE html>
<html lang="zh-CN">

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
	<title>{{.Name}}</title>
	
    <link href="http://libs.baidu.com/bootstrap/3.0.3/css/bootstrap.min.css" rel="stylesheet">
	<script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
	<script src="http://libs.baidu.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>

	<link href="http://cdn.bootcss.com/highlight.js/8.0/styles/github.min.css" rel="stylesheet">  
	<script src="http://cdn.bootcss.com/highlight.js/8.0/highlight.min.js"></script>  
	<script type="text/javascript">
		$(document).ready(function(){
			$('pre.code').each(function(i, block) {
				hljs.highlightBlock(block);
			});
		});
	</script> 
</head>

<body>
	<div class="container">
        <div class="panel panel-primary">
            <div class="panel-heading">
                <h3 class="panel-title">{{.Name}}</h3>
			</div>
			<div class="panel-body">
				<span class="label label-success">服务描述</span>
				<h5>{{.Description}}</h5>
				<hr>
				<span class="label label-success">swagger Edit</span>
				<p>把swagger协议文档复制之后，点击下方链接，在打开的网页中粘贴可进行Api调试。（后面计划修改editor，修改后无需粘贴，直接点击即可打开当前api调试页）</p>
				<p><a href="http://yuminghao.top:8000/" target="view_window">http://yuminghao.top:8000/</a></p>
				<p> Generate Server 选项可生成多种语言服务端代码，Generate Client 选项可生成多种语言服务端代码 </p>
				<p> 转grpc 正在开发中 </p>
				<hr>
				<span class="label label-success">Swagger协议文档:</span>
				{{.SwaggerJSON}}
            </div>
        </div>
    </div>

</body>

</html>

`
