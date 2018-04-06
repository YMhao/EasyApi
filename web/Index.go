package web

import (
	"encoding/base64"
	"fmt"
	"html/template"
)

type IndexInfo struct {
	Name        string
	Description string
	URL         string
	SwaggerJSON string
	SwaggerYAML string
	HTTPS       bool
}

// HTMLIndexInfo 首页信息
type HTMLIndexInfo struct {
	Name           string
	Description    string
	URLBase64      string
	SwaggerJSON    template.HTML
	SwaggerYAML    template.HTML
	SwaggerEditURL string
}

func getSwaggerEditorUURL(https bool) string {
	if https {
		return "https://yuminghao.top:8443"
	}
	return "http://yuminghao.top:8000"
}

func codeHTML(codeType, code string) template.HTML {
	html := fmt.Sprintf(`
		<div class="panel-body">
			<pre class="code %s"">%s</pre>
		</div>
		`, codeType, code)
	return template.HTML(html)
}

func (i *IndexInfo) HTMLIndexInfo() *HTMLIndexInfo {
	return &HTMLIndexInfo{
		Name:           i.Name,
		Description:    i.Description,
		URLBase64:      base64.URLEncoding.EncodeToString([]byte(i.URL)),
		SwaggerJSON:    codeHTML("json", i.SwaggerJSON),
		SwaggerYAML:    codeHTML("yaml", i.SwaggerYAML),
		SwaggerEditURL: getSwaggerEditorUURL(i.HTTPS),
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
	
    <link href="https://cdn.bootcss.com/bootstrap/3.0.3/css/bootstrap.min.css" rel="stylesheet">
	<script src="https://code.jquery.com/jquery-2.2.4.min.js"></script>
	<script src="https://cdn.bootcss.com/bootstrap/3.0.3/js/bootstrap.min.js"></script>

	<link href="https://cdn.bootcss.com/highlight.js/8.0/styles/github.min.css" rel="stylesheet">  
	<script src="https://cdn.bootcss.com/highlight.js/8.0/highlight.min.js"></script>  
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
				<span class="label label-success">Server decription</span>
				<h5>{{.Description}}</h5>
				<hr>
				<p><a href="{{.SwaggerEditURL}}/?yamlUrl={{.URLBase64}}" target="view_window">API Debug Page</a></p>
				<hr>
				<ul id="myTab" class="nav nav-tabs">
					<li class="active"><a href="#yaml" data-toggle="tab">swagger(YAML)</a></li>
					<li><a href="#json" data-toggle="tab">swagger(JSON)</a></li>
				</ul>
				<div id="myTabContent" class="tab-content">
					<div class="tab-pane fade in active" id="yaml">
						{{.SwaggerYAML}}
					</div>
					<div class="tab-pane fade" id="json">
						{{.SwaggerJSON}}
					</div>
				</div>
            </div>
        </div>
    </div>

</body>

</html>

`
