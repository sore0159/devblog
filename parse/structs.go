package parse

import "mule/devblog/route"

type ParsedFile struct {
	route.IndexData
	Content []byte
}
