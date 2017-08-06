2017 08 05

Overhauling the project after deciding on a completely static site, using github pages for now.  This extremely simplifies a lot of the design considerations.

Publishing now consists only of adding a timestamp, no index database is needed.  Generating the site is now an all or nothing affair, and list/index pages for tags will be statically created on site generation.

Publishing and basic generation is operational: tinkering around with page layout has begun.  The next step is creating inter-page structure, such as "next" and "prev" links, as well as tag listings.


2017 05 23

Began constructing the format library, getting templates into place, and then building a server to test the resulting templated files ended up becoming serious work on the server architecture that will be in the final server.

Project decisions made here are mostly regarding the structure of serving files, which will likely be changed for whatever form the final server takes on external hosting.  Some delves into go's net/http library's source code to see details on how it serves files led to a simple io.Copy approach for now.

The heavy lifting of request routing is not something I want to build myself, having read enough of the subject to respect it's depth.  My routing will be handling UID -> file, or "first/prev/next/last"-> UID -> file.

Current plan is for filter to be stored in a cookie.  No encryption needed since no secure data is involved, and a static javascript file can handle complicated UI for manipulating the filter.


2017 05 22

Parsing now reads user-generated files and creates files and data parsed and ready for the templating step.  Data will be passed directly to a templater for html file generation, but parsed .md files are important to preserve the assignment of UID and submission timestamps while retaining the ability for the user to edit/review his old posts, or re-template them after changing the websites template.

Project vision is now for each post to have it's own page, completely pre-constructed locally and uploaded to a server.  The server will do request processing to construct a filter for next/previous/search requests, using index data to locate and appropriate file to serve.

Main page, tag list, and archive list files may be able to be statically generated or can possibly be dynamically generated.  For a more useful "list" functionality the index may include brief summary data on post content.  Concerns here are the growing size of the index data, which will likely be held in memory.

Plan:

* Create the html templating step using a basic layout.  Pretty designs can come later.
* Create the server request processor filter/generator
* Nail down a UI for controlling view filters
* Work on a minimal, but professional layout (don't take forever)


2017 05 20

First very rough pass complete.  Main things I don't like about it:
1. Poor separation of concerns.  Replacing terrible data persistence with something else will require gutting a lot.
2. Terrible data persistence.  Won't scale well with post count or even tag usage.
3. Initial pass is for multiple posts per-page, but this requires extensive on-demand processing (and file access) to accommodate filters/searches.
4. Posts also have titles.  Summaries might be also good to have somewhere external to the posts themselves, if not in the index.
5. Submit date needs to be added to the unprocessed files for when they are re-processed later.


Plan:

* Ignore 1 & 2 for now.  These are linked and abstracting over arbitrary persistence schemes is not worth my time.  Focus on 3: with a one-post per page scheme, pages for posts can be completely (?) pre-processed, which is sorta the point of all this.  Problems to deal with here:
    1. How to control navigation.  Abstract "next" and "prev" links that somehow take current filter into account (probably via url path)?  Allow for on-serve processing for just the links?
    2. UI for controlling filters/searches.  Keep it simple
* Have post title come from file name.  Have minimal date_time be first part of file field.

