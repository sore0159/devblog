### Static site generation for github pages.

### Commands
```go
dv FILENAME [...]
dv g[en[erate]] [DIRECTORY]
```

### Publish
When dv is called with a list of filenames, those files are renamed to prefix a timestamp to their name.

### Generate
* Searches for the given directory, uses "." if none specified
* Attempts to parse all files in this directory
* * Files must be named TIME\_title\_post.md
* * If the first line of the file begins with "TAGS: ", the rest of the line will be parsed as a comma separated list of tags for the file
* * File contents will then be parsed as markdown, using the go package https://github.com/russross/blackfriday

* Generates static pages for each post, and list of posts for each tag (and complete list).
* Generates other pages?  Index/About?  Use some marker tag for non-inclusion?
* Generate will use a resources folder to store template files during production, but these resources will be packed into the binary after sufficient design.




