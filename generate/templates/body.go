package templates

const TMP_BODY = `{{ define "body" }}
{{ template "titlebar" . }}
{{ if not .NoTitle }}
<h1>{{ .Title }}</h1> 
{{ end }}
<div id="postheader">
{{ if not .NoDate }}<p class="date">Published {{ .Published }}</p>{{ end }}
{{ if .ContentTags }}<p class="mytags">Tagged: {{ range .ContentTags }}<a class="taglink" href="/{{ index . 1 }}">{{ index . 0 }}</a> {{ end }}</p>{{ end }}
{{ if not .NoTitle }}
<div class="topnav">
<b>Go to:</b> {{ template "navbar" (index .TagNavs 0) }}
</div>
{{ end }}
</div>
{{ .Content }}
<hr class=navbar />
<div id="navbar">
{{ range .TagNavs }}
<div class="navbar">
{{ template "navbar" . }}
</div>
{{ end }}
</div>
{{ end }}

{{ define "title" }}Eric Sorell Writes Code{{ if .Title }}: ({{ .Title }}){{ end }}{{ end }}`