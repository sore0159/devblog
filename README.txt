A project to make a simple, minimalist dev blog for my portfolio projects.

Design Goals:
1.  Don't spend forever on this
2.  Keep content easily extractable for when/if I use something else later.
3.  Look suitably professional for a portfolio.  This means at least some formatting, searchability, and content-tagging.

Plan:
1. Create local 'parse' tool that reads files and generates parsed-content and index files.
2. Create prototype http server-program that formats, displays, and enables filtering of parsed-content.
3. Setup actual public server.  Choose between Google cloud, AWS, Heroku, or something like DreamHost.  This involves an entire learning curve on it's own, so have the prototype server from #2 guide learning.  Get it operational, modifying the server as necessary.
4. Create a system of automated transfer of local parsed content to the server.   

Things Not Present:
1. Viewer Comments
2. RSS Feed (maybe doable later?)
3. Content/date based searching 
