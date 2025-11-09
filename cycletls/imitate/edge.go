package imitate

import (
	"github.com/ChengHoward/CycleTLS/cycletls"
)

func Edge(options *cycletls.Options) {
	Chrome142(options)

	options.Headers["Sec-Ch-Ua"] = `"Chromium";v="142", "Microsoft Edge";v="142", "Not_A Brand";v="99"`
	options.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/142.0.0.0 Safari/537.36 Edg/142.0.0.0"
	if options.Headers["Accept"] == "" {
		options.Headers["Accept"] = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"
	}

}
