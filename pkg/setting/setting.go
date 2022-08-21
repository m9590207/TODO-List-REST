package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

//系統設定初始化
func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	//檔名
	vp.SetConfigName("config")
	for _, config := range configs {
		if config != "" {
			//viper允許多次呼叫設定多個config的相對路徑 config1/ config2/....
			vp.AddConfigPath(config)
		}
	}
	//副檔名
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	s := &Setting{vp}
	s.WatchSettingChange()
	return s, nil
}

//使用viper提供的熱更新功能,對config進行監聽
func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}
