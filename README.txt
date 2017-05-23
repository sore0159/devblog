A project to make a simple, minimalist dev blog for my portfolio projects.

Design Goals:
1.  Don't spend forever on this
2.  Keep content easily extractable for when/if I use something else later.
3.  Look suitably professional for a portfolio.  This means at least some formatting, searchability, and content-tagging.


Project Structure:
* The parse library handles reading raw (markdown) user content and creating markdown files that include metadata on submission date and post UID.
* The format library handles data from the parse library, creating fully-formatted .html files suitable to be directly served.  Creation of organizational files such as pre-generated list pages and index pages is a likely functionality.
* The display library handles routing (may be renamed) of user requests to the appropriate pre-generated file.  It may include the ability to dynamically generate certain kinds of list files.

Plan:
1. Create format, routing libraries.
2. Create prototype http server, run tests on the full processing pipeline.
3. Setup actual public server.  Choose between Google cloud, AWS, Heroku, or something like DreamHost.  This involves an entire learning curve on it's own, so have the prototype server from #2 guide learning.  Get it operational, modifying the server as necessary.
4. Create a system of automated transfer of local parsed content to the server.   

Things Not Present:
1. Viewer Comments
2. RSS Feed (maybe doable later?)
3. Content/date based searching (tag-based searching hopefully enough)
