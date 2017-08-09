package templates

const TMP_FRAME = `{{ define "frame" -}}
<!doctype html>
<html>
<head>
<meta charset="utf-8">
<title>{{ template "title" . }}</title>
<link rel="stylesheet" href="/css/{{ template "css" . }}.css">
<link rel="shortcut icon" href="/img/yd32.ico">
</head>
<body>
{{ template "body" . }}
</body>
</html>
{{ end }}


{{ define "css" }}basic{{ end }}
{{ define "title" }}Eric Sorell Writes Code{{ end }}
{{ define "body" }}Default Body{{ end }}`
