package provider

import "github.com/valyala/fasthttp"

func (p *Provider) VASTEventHandler(ctx *fasthttp.RequestCtx) {
	itemID := ctx.UserValue("id").(string)
	item, ok := p.items.data[itemID]
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	event := ctx.UserValue("event").(string)
	p.stats.IncrementValue(event, item.ID)
}
