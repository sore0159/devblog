package templates

const TMP_NAVBAR = `{{ define "navbar" }}
{{ if index .Has 0 }}
<a class="navbar" href="{{ .FirstFN }}" title="{{.FirstT }} ({{ .FirstPub }})">First</a>
{{ else }}
<span class="inv">First</span>
{{ end }}
&bull;
{{ if index .Has 0 }}
<a class="navbar" href="{{ .PrevFN }}" title="{{.PrevT }} ({{ .PrevPub }})">Previous</a>
{{ else }}
<span class="inv">Previous</span>
{{ end }}
<div class="archnav">
        &bull;
{{ if index .Tag 1 }}
<a class="navbar" href="{{ index .Tag 1 }}">Posts tagged "{{ index .Tag 0 }}"</a>
{{ else }}
<a class="navbar" href="archives.html">All posts</a>
{{ end }}
        &bull;
</div>
{{ if index .Has 1 }}
<a class="navbar" href="{{ .NextFN }}" title="{{.NextT }} ({{ .NextPub }})">Next</a>
{{ else }}
<span class="inv">Next</span>
{{ end }}
&bull;
{{ if index .Has 1 }}
<a class="navbar" href="{{ .LastFN }}" title="{{.LastT }} ({{ .LastPub }})">Latest</a>
{{ else }}
<span class="inv">Latest</span>
{{ end }}
{{ end }}`