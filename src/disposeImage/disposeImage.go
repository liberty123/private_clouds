package disposeImage

import (
	"github.com/spf13/viper"
	"os"
	"fmt"
)


func GetConfig() (string, string) {
	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		os.Exit(1)
	}
	original_image_path := viper.Get("imagePath.original_image_path")
	thumbnail_path := viper.Get("imagePath.thumbnail_path")
	a, ok_1 := original_image_path.(string)
	b, ok_2 := thumbnail_path.(string)
	if ok_1 && ok_2{
		return a, b
	} else {
		return "It's not ok for type string", "It's not ok for type string"
	}
}
