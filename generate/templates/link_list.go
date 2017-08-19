package templates

const TMP_LINKLIST = `{{ define "linklist" }}
{{ if . }}
<ul>
{{ range . }}
{{ if not .NoDate }}
<li> {{ template "fulllink" . }} </li>

{{ end }}
{{ end }}
</ul>
{{ end }}
{{ end }}

{{ define "fulllink" }}
        <a href="{{ .FileName }}">{{ .Title }}</a> {{ .Published }}
        {{ if .ContentTags }}<b>Tags:</b> 
        {{ range .ContentTags }}<a class="taglink" href="/{{ index . 1 }}">{{ index . 0 }}</a>{{ end }}
        {{ end }}
{{ end }}`