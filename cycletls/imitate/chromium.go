package imitate

import (
	"github.com/ChengHoward/CycleTLS/cycletls"
)

// 在chromium中 21 是时有时无的
const chromiumExtension = "0-5-10-11-13-16-18-21-23-27-35-43-45-51-17513-65037-65281"

var chromiumH2Settings = &cycletls.H2Settings{
	Settings: map[string]int{
		"HEADER_TABLE_SIZE":      65536,
		"ENABLE_PUSH":            0,
		"MAX_CONCURRENT_STREAMS": 1000,
		"INITIAL_WINDOW_SIZE":    6291456,
		"MAX_HEADER_LIST_SIZE":   262144,
		// 在chromium中是随机的
		//"UNKNOWN_SETTING_15082 ": 0,
	},
	SettingsOrder: []string{
		"HEADER_TABLE_SIZE",
		"ENABLE_PUSH",
		"MAX_CONCURRENT_STREAMS",
		"INITIAL_WINDOW_SIZE",
		"MAX_HEADER_LIST_SIZE",
		"UNKNOWN_SETTING_15082 ",
	},
	ConnectionFlow: 15663105,
	HeaderPriority: map[string]interface{}{
		"weight":    256,
		"streamDep": 0,
		"exclusive": true,
	},
	PriorityFrames: []map[string]interface{}{
		/*{
			"streamID": 0,
			"priorityParam": map[string]interface{}{
				"weight":    0,
				"streamDep": 0,
				"exclusive": true,
			},
		},*/
		/*{
			"streamID": 3,
			"priorityParam": map[string]interface{}{
				"weight":    201,
				"streamDep": 0,
				"exclusive": false,
			},
		},*/
		/*{
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
var chromiumHttp2Setting = cycletls.ToHTTP2Settings(chromiumH2Settings)

func Chromium(options *cycletls.Options) {
	options.Ja3 = "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53" + "," + shuffleExtension(chromiumExtension, 7) + "-41,29-23-24,0"
	options.HTTP2Settings = chromiumHttp2Setting
	options.PHeaderOrderKeys = []string{
		":method",
		":authority",
		":scheme",
		":path",
	}
	if options.Headers == nil {
		options.Headers = make(map[string]string)
	}

	options.Headers["Sec-Ch-Ua"] = `"Chromium";v="115", "Not/A)Brand";v="99"`
	options.Headers["Sec-Ch-Ua-Mobile"] = "?0"
	options.Headers["Sec-Ch-Ua-Platform"] = `"Windows"`
	options.Headers["Sec-Fetch-Dest"] = "document"
	options.Headers["Sec-Fetch-Mode"] = "navigate"
	options.Headers["Sec-Fetch-Site"] = "none"
	options.Headers["Sec-Fetch-User"] = "?1"
	if options.Headers["Accept"] == "" {
		options.Headers["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"
	}

	//options.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36"
	options.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

}
