package extractor

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

var (
	blankReg *regexp.Regexp
)

func init()  {
	var err error
	blankReg,err = regexp.Compile("\\s+")
	if err != nil {
		panic(err)
	}
}

type FeatureExtractor []*common.FeatureConfig

func (m *FeatureExtractor) Init(confFile string) error{
	if fin,err := os.Open(confFile); err != nil {
		return err
	} else {
		defer fin.Close()
		const MAX_CONFIG_SIZE = 1 << 20		//配置文件大小不能超过1M
		bs := make([]byte,MAX_CONFIG_SIZE)
		if n,err := fin.Read(bs); err != nil {
			return err
		} else {
			if n >= MAX_CONFIG_SIZE {
				return fmt.Errorf("config file size more than %dB",MAX_CONFIG_SIZE)
			}
			var confList Common.FeatureConfigList
			if err = json.Unmarshal(bs[:n],&confList); err != nil {
				return err
			} else {
				productType := reflect.TypeOf(common.Product{})
				for _, conf := range confList {
					paths := splitString(conf.Path,".")
					
				}
			}
		}
	}
	return nil
}



