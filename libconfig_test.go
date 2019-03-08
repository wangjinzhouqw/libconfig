package libconfig

import "testing"

func TestParseConfigFile(t *testing.T) {
	ParseConfigFile("/Users/wangjz/Desktop/work/project/opensource/janus-gateway/conf/janus.jcfg.sample.in")
	//ParseJsonConfigFile("/Users/wangjz/Desktop/janus.jcfg.sample.in")
}