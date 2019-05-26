package cmd

import (
	"github.com/markthub/apis/api/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	addressFlag      = "address-host"
	dbHostFlag       = "db-host"
	dbPortFlag       = "db-port"
	dbUserFlag       = "db-user"
	dbPasswordFlag   = "db-password"
	dbNameFlag       = "db-name"
	mollieAPIKeyFlag = "mollie-api-key"
	mollieTestFlag   = "mollie-test"
)

var cmdAPI = &cobra.Command{
	Use:   "server",
	Short: "Run the API server",
	Long: `
MarktHub APIs
This command will run the APIs for this toy project.
`,
	SilenceUsage: true,
	Run:          runAPI,
}

func runAPI(cmd *cobra.Command, args []string) {

	addr := viper.GetString(addressFlag)
	dbHost := viper.GetString(dbHostFlag)
	dbPort := viper.GetString(dbPortFlag)
	dbUser := viper.GetString(dbUserFlag)
	dbPassword := viper.GetString(dbPasswordFlag)
	dbName := viper.GetString(dbNameFlag)
	mollieAPIKey := viper.GetString(mollieAPIKeyFlag)
	mollieTest := viper.GetBool(mollieTestFlag)

	err := api.Serve(addr, dbHost, dbPort, dbUser, dbPassword, dbName, mollieAPIKey, mollieTest)
	if err != nil {
		log.Panic().Err(err).Msg("could not serve the application")
	}
}

func init() {
	f := cmdAPI.Flags()

	f.String(addressFlag, ":8000", "server address")
	f.String(dbHostFlag, "127.0.0.1", "database host")
	f.String(dbPortFlag, "3306", "database port")
	f.String(dbUserFlag, "markthub", "database username")
	f.String(dbPasswordFlag, "markthub", "database password")
	f.String(dbNameFlag, "markthub", "database name")
	f.String(mollieAPIKeyFlag, "test_5UKptB8fj2Jje2Um2FCU6Dba5Fxmnw", "mollie API key")
	f.Bool(mollieTestFlag, true, "mollie debug flag")

	viper.BindEnv(addressFlag, "ADDRESS_HOST")
	viper.BindEnv(dbHostFlag, "DB_HOST")
	viper.BindEnv(dbPortFlag, "DB_PORT")
	viper.BindEnv(dbUserFlag, "DB_USER")
	viper.BindEnv(dbPasswordFlag, "DB_PASSOWRD")
	viper.BindEnv(dbNameFlag, "DB_NAME")
	viper.BindEnv(mollieAPIKeyFlag, "MOLLIE_API_KEY")
	viper.BindEnv(mollieTestFlag, "MOLLIE_TEST")

	viper.BindPFlags(f)
}

// initConfig sets AutomaticEnv in viper to true.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}

// Execute will start the application
func Execute() {
	cobra.OnInitialize(initConfig)
	if err := cmdAPI.Execute(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
