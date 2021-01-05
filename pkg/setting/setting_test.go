package setting

import "testing"

func TestSetup(t *testing.T) {
	Setup()
	t.Logf("AppSetting: %v", AppSetting)
	t.Logf("ServerSetting: %v", ServerSetting)
	t.Logf("RedisSetting: %v", RedisSetting)
}
