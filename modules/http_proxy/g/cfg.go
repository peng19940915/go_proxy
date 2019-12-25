package g

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/toolkits/file"
	"sync"
	"time"
)

type Log struct {
	LogLevel     string        `json:"logLevel"`
	RotationTime time.Duration `json:"rotationTime"`
	LogPath      string        `json:"logPath"`
	MaxAge       time.Duration `json:"maxAge"`
	FileName     string        `json:"fileName"`
}

type GlobalConfig struct {
	Log *Log `json:"log"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func ParseConfig(cfg string) error {
	if cfg == "" {
		return fmt.Errorf("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		return fmt.Errorf("config file: %s is not existent. maybe you need `mv cfg.example.json cfg.json`", cfg)
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		return fmt.Errorf("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		return fmt.Errorf("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c

	logrus.Infof("read config file: %s successfully", cfg)
	return nil
}
