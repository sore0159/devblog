### Static site generation

### Commands
```bash
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
* Tags NODATE and NOTITLE will suppress the inclusion of the date and title on the generated page, as well as suppressing the date's in the filename of the generated file.  This is primarily for index and archive files to still be written as adjustable .md files.
* Generate will use a resources directory for template files during production, but these resources will be packed into the binary after sufficient design.

### Test Server
This project includes a simple test server to use while designing the page layouts and site structure


