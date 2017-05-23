package parse

import "mule/devblog"

type ParsedFile struct {
	devblog.IndexData
	Content []byte
}
