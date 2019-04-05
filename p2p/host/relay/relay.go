package relay

import (
	"context"
	"time"

	discovery "github.com/libp2p/go-libp2p-discovery"
)

var (
	AdvertiseBootDelay = 30 * time.Second
)

func Advertise(ctx context.Context, advertise discovery.Advertiser) {
	go func() {
		select {
		case <-time.After(AdvertiseBootDelay):
			discovery.Advertise(ctx, advertise, RelayRendezvous)
		case <-ctx.Done():
		}
	}()
}
