package main

import (
	"fmt"
	"log"
	"os"
)

func setEnvMany(keyValuePairs ...string) {
	if (len(keyValuePairs) % 2) != 0 {
		log.Fatalln("odd number of values provided")
	}
	for i := 1; i < len(keyValuePairs); i += 2 {
		key := keyValuePairs[i-1]
		value := keyValuePairs[i] // default value

		if override := os.Getenv(key); override != "" {
			value = override
		}

		if err := os.Setenv(key, value); err != nil {
			log.Fatalln(err, key, value)
		}
		fmt.Println("key:", key, "Value: ", value)
	}
}
func main() {
	setEnvMany(
		"WORKER", "gstreamer-mp",
		"GST_DEBUG_DIR", "logDir",
		"GST_PLUGIN_PATH", "/opt/firmware/current/ronin/lib/gstreamer-1.0",
		"LOG_LEVEL", "debug",
		"LOG_FILE", "/tmp/ronin/pipeline_testcase",
		"LOG_BUFFER_LEVELS", "1",
		"STREAM_INFO", "streamInfo",
	)

}
