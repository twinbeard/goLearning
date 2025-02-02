package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"database"`
}

func main() {
	viper := viper.New()            // Create a new instance of viper
	viper.AddConfigPath("./config") // Path to the config folder.
	viper.SetConfigName("local")    // Name of the config file
	viper.SetConfigType("yaml")     // Type of the config file

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fail to read configuration file %w \n", err)) //if error in reading the config file, STOP APP
	}
	// Read configuration -> simple config
	fmt.Println("Server Port: ", viper.GetInt("server.host"))             // server.port = 8080 (trong file local.yaml)
	fmt.Println("Server security: ", viper.GetString("security.jwt.key")) // server.port = 8080 (trong file local.yaml)

	// Read onfiguration -> complex config
	var config Config              // Create a new instance of the Config struct
	err = viper.Unmarshal(&config) // Unmarshal the config into the struct
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))

	}
	fmt.Println("Config port: ", config.Server.Port) // database[0].user = root (trong file local.yaml)
	for _, db := range config.Database {
		fmt.Printf("Database user: %s, password: %s, host: %s \n", db.User, db.Password, db.Host) // database
	}
}
