### Static site generation

### Commands
```bash
dv [m] FILENAME [...]
dv g[en[erate]] [DIRECTORY] [...]
```

### Publish
* When dv is called with a list of filenames, those files are copied to append a timestamp prefix to their name.
* If the first argument to dv is 'm', published files will be moved instead of copied.
* If a file already has a properly formatted timestamp as it's prefix, that timestamp will be replaced rather than appended too.

### Generate
* Searches for the given directories, uses "." if none specified
* Attempts to parse all files in these directories
* * Files must be named TIME\_title\_post.md
* * If the first line of the file begins with "TAGS: ", the rest of the line will be parsed as a comma separated list of tags for the file
* * File contents will then be parsed as markdown, using the go package https://github.com/russross/blackfriday

* Generates static pages for each post, and list of posts for each tag (and complete list).
* Must find a directory named "templates" in the current directory with the needed .html files for post generation.
* Files will be created in a directory "generated", which will be made if it does not exist.
* Tags NODATE and NOTITLE will suppress the inclusion of the date and title on the generated page, as well as suppressing the date's in the filename of the generated file.  This is primarily for any needed one-off files to still be written as adjustable .md files.
* index.html, archives.html, and archives\_TAG.html (for each content tag present) files will be created from templates, using post data.
* Generate will use a resources directory for template files during production, but these resources may be packed into the binary after sufficient design.

### Test Server
This project includes a simple test server to use while designing the page layouts and site structure

### Potential Additions
* Dynamic JS navigation
* Sprucing up archive pages with lines/breaks between months/years.
* RSS feed generation
* Configuration files/options
