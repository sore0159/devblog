package templates

const TMP_MAIN_ARCHIVE = `{{ define "body" }}
{{ template "titlebar" . }}
<h1>Archive of all posts</h1>
<p class="date">Last posted: {{ (index . 0).Published }}</p>
{{ template "linklist" . }}

{{ end }}

{{ define "title" }}Eric Sorell Writes Code (Main Archive){{ end }}`
