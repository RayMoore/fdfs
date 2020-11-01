package fdfs

import (
	"bufio"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Config struct {
	TrackerAddr []string
	MaxConnections    int
}

func newConfig(configName string) (*Config, error) {
	config := &Config{}
	f, err := os.Open(configName)
	if err != nil {
		return nil, err
	}
	splitFlag := "\n"
	if runtime.GOOS == "windows" {
		splitFlag = "\r\n"
	}
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, splitFlag)
		str := strings.SplitN(line, "=", 2)
		switch str[0] {
		case "tracker_server":
			config.TrackerAddr = append(config.TrackerAddr, str[1])
		case "max_connections":
			config.MaxConnections, err = strconv.Atoi(str[1])
			if err != nil {
				return nil, err
			}
		}
		if err != nil {
			if err == io.EOF {
				return config, nil
			}
			return nil, err
		}
	}
	//return config, nil
}
