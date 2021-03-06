package main

import (
	"bytes"
	"github.com/valyala/fasthttp"
	"io"
)

func ServeFavicon(ctx *fasthttp.RequestCtx) {
	b := bytes.NewBuffer(Favicon)
	ctx.Response.Header.Set("Content-Type", "image/x-icon")
	_, e := io.Copy(ctx.Response.BodyWriter(), b)
	if e != nil {
		ctx.Response.SetStatusCode(fasthttp.StatusInternalServerError)
	}
}
