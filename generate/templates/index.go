package templates

const TMP_INDEX = `{{ define "body" }}
{{ template "titlebar" . }}
<h3>Latest Post:</h3>
<p class="index">{{ template "fulllink" . }}</p>
<h3>About This Site</h3>
<p class="index">This is the index page!  We'll have some stuff here eventually.</p>

{{ end }}

{{ define "title" }}Eric Sorell Writes Code{{ end }}`