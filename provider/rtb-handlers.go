package provider

import (
	"encoding/json"
	"strings"

	"github.com/google/uuid"
	"github.com/prebid/openrtb/v19/openrtb2"
	"github.com/valyala/fasthttp"
)

func (p *Provider) BidRequestHandler(ctx *fasthttp.RequestCtx) {
	itemID := ctx.UserValue("id").(string)
	item, ok := p.items.data[itemID]
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	var bidRequest openrtb2.BidRequest

	err := json.Unmarshal(ctx.Request.Body(), &bidRequest)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if len(bidRequest.Imp) < 1 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	bidResponse := openrtb2.BidResponse{
		ID:    bidRequest.ID,
		BidID: uuid.NewString(),
		SeatBid: []openrtb2.SeatBid{
			{
				Bid: []openrtb2.Bid{
					{
						ID:    uuid.NewString(),
						ImpID: bidRequest.Imp[0].ID,
						AdID:  item.ID,
						Price: item.Bid,
						CID:   item.ID,
						CrID:  item.ID,
						NURL:  p.baseURL + "/rtb/notify/" + item.ID,
						AdM: p.baseURL +
							"/rtb/vast/" +
							item.ID +
							"?source_id=" +
							bidRequest.Site.ID,
					},
				},
			},
		},
	}

	data, err := json.Marshal(bidResponse)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadGateway)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetBody(data)
}

func (p *Provider) NotifyHandler(ctx *fasthttp.RequestCtx) {
	itemID := ctx.UserValue("id").(string)
	item, ok := p.items.data[itemID]
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}
	p.stats.IncrementValue(EventTypeNotify, item.ID)
}

func (p *Provider) VASTHandler(ctx *fasthttp.RequestCtx) {
	itemID := ctx.UserValue("id").(string)
	item, ok := p.items.data[itemID]
	if !ok {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	sourceID := ctx.QueryArgs().Peek("source_id")

	destURL := strings.Replace(item.DestinationUrl, "{{ad_id}}", item.ID, -1)
	destURL = strings.Replace(destURL, "{{source_id}}", string(sourceID), -1)

	data := strings.Replace(vastTemplate, "{{ad_id}}", item.ID, -1)
	data = strings.Replace(data, "{{video_url}}", item.VideoURL, -1)
	data = strings.Replace(data, "{{title}}", item.Title, -1)
	data = strings.Replace(data, "{{video_duration}}", item.VideoDuration, -1)
	data = strings.Replace(data, "{{destination_url}}", destURL, -1)
	data = strings.Replace(data, "{{base_url}}", p.baseURL, -1)

	ctx.SetContentType("text/xml")
	ctx.SetBody([]byte(data))
}
