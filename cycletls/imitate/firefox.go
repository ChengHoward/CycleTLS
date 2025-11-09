package imitate

import (
	"github.com/ChengHoward/CycleTLS/cycletls"
)

var firefoxH2Settings = &cycletls.H2Settings{
	Settings: map[string]int{
		"HEADER_TABLE_SIZE":   65536,
		"ENABLE_PUSH":         0,
		"INITIAL_WINDOW_SIZE": 131072,
		"MAX_FRAME_SIZE":      16384,
	},
	SettingsOrder: []string{
		"HEADER_TABLE_SIZE",
		"ENABLE_PUSH",
		"INITIAL_WINDOW_SIZE",
		"MAX_FRAME_SIZE",
	},
	ConnectionFlow: 12517377,
	HeaderPriority: map[string]interface{}{
		"weight":    42,
		"streamDep": 0,
		"exclusive": false,
	},
	PriorityFrames: []map[string]interface{}{
		/*{
			"streamID": 0,
			"priorityParam": map[string]interface{}{
				"weight":    255,
				"streamDep": 0,
				"exclusive": true,
			},
		},*/
		/*{
			"streamID": 3,
			"priorityParam": map[string]interface{}{
				"weight":    0,
				"streamDep": 0,
				"exclusive": false,
			},
		},
		{
			"streamID": 5,
			"priorityParam": map[string]interface{}{
				"weight":    101,
				"streamDep": 0,
				"exclusive": false,
			},
		},
		{
			"streamID": 7,
			"priorityParam": map[string]interface{}{
				"weight":    1,
				"streamDep": 0,
				"exclusive": false,
			},
		},
		{
			"streamID": 9,
			"priorityParam": map[string]interface{}{
				"weight":    1,
				"streamDep": 7,
				"exclusive": false,
			},
		},
		{
			"streamID": 11,
			"priorityParam": map[string]interface{}{
				"weight":    1,
				"streamDep": 3,
				"exclusive": false,
			},
		},
		{
			"streamID": 13,
			"priorityParam": map[string]interface{}{
				"weight":    241,
				"streamDep": 0,
				"exclusive": false,
			},
		},*/
	},
}
var firefoxHttp2Setting = cycletls.ToHTTP2Settings(firefoxH2Settings)

func Firefox(options *cycletls.Options) {
	//"0-23-65281-10-11-35-16-5-51-43-13-45-28-21"
	//"0-23-65281-10-11-35-16-5-34-51-43-13-45-28-21"
	options.Ja3 = "771,4865-4867-4866-49195-49199-52393-52392-49196-49200-49162-49161-49171-49172-156-157-47-53,0-23-65281-10-11-35-16-5-34-18-51-43-13-45-28-27-65037,4588-29-23-24-25-256-257,0"
	options.HTTP2Settings = firefoxHttp2Setting
	options.PHeaderOrderKeys = []string{
		":method",
		":path",
		":authority",
		":scheme",
	}

	//options.Headers["Sec-Ch-Ua"] = `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`
	options.Headers["upgrade-insecure-requests"] = "1"
	options.Headers["Sec-Fetch-Dest"] = "document"
	options.Headers["Sec-Fetch-Mode"] = "navigate"
	options.Headers["Sec-Fetch-Site"] = "none"
	options.Headers["Sec-Fetch-User"] = "?1"
	options.Headers["Accept-Encoding"] = "gzip, deflate, br, zstd"
	options.Headers["Priority"] = "u=0, i"
	options.Headers["te"] = "trailers"
	if options.Headers["Accept"] == "" {
		options.Headers["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"
	}
	if options.Headers["Accept-Language"] == "" {
		options.Headers["Accept-Language"] = "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2"
	}
	options.HeaderOrderKeys = []string{
		"host",
		"connection",
		"cache-control",
		"device-memory",
		"viewport-width",
		"rtt",
		"downlink",
		"ect",
		"sec-ch-ua",
		"sec-ch-ua-mobile",
		"sec-ch-ua-full-version",
		"sec-ch-ua-arch",
		"sec-ch-ua-platform",
		"sec-ch-ua-platform-version",
		"sec-ch-ua-model",
		"user-agent",
		"accept",
		"accept-language",
		"accept-encoding",
		"upgrade-insecure-requests",
		"sec-fetch-dest",
		"sec-fetch-mode",
		"sec-fetch-site",
		"sec-fetch-user",
		"cookie",
		"referer",
		"priority",
		"te",
	}
	options.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:144.0) Gecko/20100101 Firefox/144.0"

}
