// Package downloader is the main module of GO_SPIDER for download page.
package downloader

import (
	"visual_spider_go/spider/core/common/page"
	"visual_spider_go/spider/core/common/request"
)

// The Downloader interface.
// You can implement the interface by implement function Download.
// Function Download need to return Page instance pointer that has request result downloaded from Request.
type Downloader interface {
	Download(req *request.Request) *page.Page
}