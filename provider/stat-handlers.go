package provider

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func (p *Provider) StatsHandler(ctx *fasthttp.RequestCtx) {
	data, err := json.Marshal(p.stats.data)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(data)
}

func (p *Provider) ItemStatsHandler(ctx *fasthttp.RequestCtx) {
	itemID := ctx.UserValue("id").(string)
	item, ok := p.stats.data[itemID]
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	data, err := json.Marshal(item)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(data)
}
