package provider

import (
	"context"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
	"golang.org/x/exp/slog"
)

const timeFormat = "2006-01-02 15:04:05"

type (
	Provider struct {
		httpServer *fasthttp.Server
		items      *itemStore
		stats      *statStore
		baseURL    string
	}
)

func New() *Provider {
	return &Provider{
		httpServer: &fasthttp.Server{
			Name: "VAST Provider",
		},
		items: NewItemStore(),
		stats: NewStatStore(),
	}
}

func (p *Provider) Start(
	ctx context.Context,
	listenNetwork,
	listenAddr,
	baseURL string,
) error {
	ln, err := reuseport.Listen(
		listenNetwork,
		listenAddr,
	)
	if err != nil {
		return err
	}
	defer ln.Close()

	p.baseURL = baseURL

	r := fasthttprouter.New()

	// rtb
	r.POST("/rtb/:id", p.BidRequestHandler)   // OpenRTB requests
	r.GET("/rtb/notify/:id", p.NotifyHandler) // OpenRTB Notify Event (nurl)
	r.GET("/rtb/vast/:id", p.VASTHandler)     // OpenRTB VAST tag (adm)

	// VAST Events
	r.GET("/vast/:event/:id", p.VASTEventHandler) // VAST Events tracking

	// items crud
	r.GET("/data/items", p.GetItemsHandler)
	r.GET("/data/item/:id", p.GetItemHandler)
	r.PUT("/data/item", p.UpsertItemHandler)
	r.DELETE("/data/item/:id", p.DelItemHandler)

	// stats
	r.GET("/stats", p.StatsHandler)
	r.GET("/stats/:id", p.ItemStatsHandler)

	p.httpServer.Handler = r.Handler

	go func() {
		<-ctx.Done()
		err := p.httpServer.Shutdown()
		if err != nil {
			slog.Error(
				"Parsing Config Error",
				"datetime", time.Now().Format(timeFormat),
				"error", err.Error(),
			)
		}
		// stop all connections to db, etc
	}()

	go func() {
		for range time.Tick(5 * time.Second) {
			err := writToFile(itemStorePath, p.items.data)
			if err != nil{
				slog.Error("items snapshot", "error", err.Error())
			}
			err = writToFile(statStorePath, p.stats.data)
			if err != nil{
				slog.Error("stats snapshot", "error", err.Error())
			}
		}
	}()
	return p.httpServer.Serve(ln)
}
