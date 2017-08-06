##Pure static site generation for github pages.

File Name: post_title.md
File First Line (optional): TAGS: Comma separated, Case Sensitive
File Contents: Markdown content 

##Commands:
dv publish post_title.md [...]
dv generate [directory]

## Publish
Renames files to prepend TIME\_ to the filename.

## Generate
* Finds files properly named in directory (default dir "."), parses tags/titles/content for each file.
* Generates static pages for each post, and list of posts for each tag (and complete list).
* Generates other pages?  Index/About?  Use some marker tag for non-inclusion?
* Generate will use a resources folder to store template files during production, but these resources will be packed into the binary after sufficient design.




