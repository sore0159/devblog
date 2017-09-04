### Static site generation

### Commands
```bash
dv [c] FILENAME [...]
dv g[en[erate]] [DIRECTORY] [...]
dv t[est] FILENAME

test_server [DIRECTORY]
```

### Publish
* When dv is called with a list of filenames, those files are renamed to append a timestamp prefix to their name.
* If the first argument to dv is 'c', published files will be copied instead of renamed.
* If a file already has a properly formatted timestamp as it's prefix, that timestamp will be replaced rather than appended too.

### Generate
* This will delete ./dv\_generated and all subfolders when run!
* Searches for the given directories, uses "." if none specified
* Attempts to parse all files in these directories
* * Files must be named TIME\_title\_post.md
* * If the first line of the file begins with "TAGS: ", the rest of the line will be parsed as a comma separated list of tags for the file
* * File contents will then be parsed as markdown, using the go package https://github.com/russross/blackfriday

* Generates static pages for each post, and list of posts for each tag (and complete list).
* Must find a directory named "templates" in the current directory with the needed .html files for post generation.
* Files will be created in a directory "dv\_generated", which will be first destroyed if it already exists.
* Tags NODATE and NOTITLE will suppress the inclusion of the date and title on the generated page, as well as suppressing the date's in the filename of the generated file.  This is primarily for any needed one-off files to still be written as adjustable .md files.
* index.html, feed.rss, archives.html, and archives\_TAG.html (for each content tag present) files will be created from templates, using post data.
* Generate will use a resources directory for template files during production, but these resources may be packed into the binary after sufficient design.

### Test
* When dv is called with the t flag and a single file, that file will be processed into a single page, ignoring all other content, in order to quickly examine how the final layout of an entry will look.  DV will attempt to open a browser on the generated file using the mac command "open".

dv t expects a folder named css to be in the current directory with the appropriate css files.

## Test Server
This project includes a simple test server to use while designing the page layouts and site structure.  This server is configured to be run in the parent directory of the generated content file, and a folder containing static content (such as css files) which can be supplied as a commandline argument. 

### Potential Additions
* Dynamic JS navigation
* Sprucing up archive pages with lines/breaks between months/years.
* Configuration files/options
