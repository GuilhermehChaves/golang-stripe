package config

import "github.com/spf13/viper"

func Setup() {
	viper.SetDefault("STRIPE_URL", "https://api.stripe.com/v1/checkout/sessions")
	viper.SetDefault("STRIPE_AUTHORIZATION", "Bearer stripe_secret_key")
	viper.SetDefault("PORT", "8080")

	viper.BindEnv("STRIPE_URL")
	viper.BindEnv("STRIPE_AUTHORIZATION")
	viper.BindEnv("PORT")
}
