package format

import (
	"github.com/ddwapple/joy4/format/mp4"
	"github.com/ddwapple/joy4/format/ts"
	"github.com/ddwapple/joy4/format/rtmp"
	"github.com/ddwapple/joy4/format/rtsp"
	"github.com/ddwapple/joy4/format/flv"
	"github.com/ddwapple/joy4/format/aac"
	"github.com/ddwapple/joy4/av/avutil"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}

