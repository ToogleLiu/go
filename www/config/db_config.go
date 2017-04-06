package config

type DBConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

func DB_order() DBConfig {
    var config DBConfig
    config.Host = "127.0.0.1"
    config.Port = "3306"
    config.User = "N5lux"
    config.Password = "idbT3W3n0"
    config.DBName = "v30_db_orders"
    return config
}

func DB_user() DBConfig {
    var config DBConfig
    config.Host = "127.0.0.1"
    config.Port = "3306"
    config.User = "N5lux"
    config.Password = "idbT3W3n0"
    config.DBName = "v30_db_users"
    return config
}
