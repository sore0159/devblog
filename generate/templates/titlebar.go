package templates

const TMP_TITLEBAR = `{{ define "titlebar" }}<div id="titlebar">
<a id="titlelink" class="toplink" href="/">Eric Sorell Writes Code</a>
</div>
<div id="topright">
        <a class="toplink" href="/index.html">Home</a>
        &bull; <a class="toplink" href="/about.html">About Me</a>
        &bull; <a class="toplink" href="/archives.html">Archives</a>
 </div>{{ end }}`