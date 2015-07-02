{{define "manage"}}
<a class="navbar-brand" href="/">首页</a>
<div>
	<ul class="nav navbar-nav">
		<li {{if .isNews}}class="active"{{end}}><a href="/manage">新闻</a></li>
		<li {{if .isMedia}}class="active"{{end}}><a href="/media">媒体</a></li>
	</ul>
</div>

<div class="pull-right">
	<ul class="nav navbar-nav">
		<li><a href="/login?exit=true">退出</a></li>
	</ul>
</div>
{{end}}