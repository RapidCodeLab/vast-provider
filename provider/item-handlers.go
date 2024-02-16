package provider

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

func (p *Provider) GetItemsHandler(ctx *fasthttp.RequestCtx) {
	data, err := json.Marshal(p.items.data)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}

	ctx.SetBody(data)
}

func (p *Provider) GetItemHandler(ctx *fasthttp.RequestCtx) {
	itemID := ctx.UserValue("id").(string)
	item, ok := p.items.data[itemID]
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	data, err := json.Marshal(item)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}

	ctx.SetBody(data)
}

func (p *Provider) UpsertItemHandler(ctx *fasthttp.RequestCtx) {
	var item item

	err := json.Unmarshal(ctx.Request.Body(), &item)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if item.ID == "" {
		item.ID = uuid.NewString()
	}

	p.items.data[item.ID] = item
}

func (p *Provider) DelItemHandler(ctx *fasthttp.RequestCtx) {
	itemID := ctx.UserValue("id").(string)
	_, ok := p.items.data[itemID]
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	delete(p.items.data, itemID)
}
