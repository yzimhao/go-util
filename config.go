package utilgo

import "github.com/spf13/viper"

func ViperInit(configFile string) *viper.Viper {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return viper.GetViper()
}
