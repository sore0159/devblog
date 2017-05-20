package main

// First Google result for 'golang markdown parser'
// Selling points:
//  * Simple API
//  * Been around a while, still worked on
//  * Handles more features than _I_ need
//  * Depends only on golang std-lib
//
// func([]byte) []byte
// parsed := blackfriday.MarkdownCommon(unparsed)
import "github.com/russross/blackfriday"

// Possible alternative:
// Why?
//  * Blackfriday parsing didn't seem to catch ## as h2
// Downside:
//  * More dependancies
//  * Requires specific css
//  * Also, didn't fix the problem -_-
//
// parsed := gfm.Markdown(unparsed)
//import gfm "github.com/shurcooL/github_flavored_markdown"

func ParseMarkdown(in []byte) (out []byte) {
	return blackfriday.MarkdownCommon(in)
}
