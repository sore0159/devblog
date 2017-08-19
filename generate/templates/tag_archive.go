package templates

const TMP_TAG_ARCHIVE = `{{ define "body" }}
{{ template "titlebar" . }}
<h1>{{ .Title }}</h1>
<p class="date">Tag "{{ .Tag }}" last used: {{ .Latest }}</p>
<p class="mytags"><a href="/archives.html">Go to full archive</a></p>

{{ template "linklist" .Posts }}

{{ end }}

{{ define "title" }}Eric Sorell Writes Code (Archive: {{ .Tag }}){{ end }}`