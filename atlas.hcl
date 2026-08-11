# Atlas configuration for Auth Service schema management.
# Desired schema is derived from GORM models via an external schema loader.
# Migrations are versioned under ./migrations — never AutoMigrate in production.

env "local" {
  url = getenv("DATABASE_URL")
  dev = getenv("DEV_DATABASE_URL")

  migration {
    dir = "file://migrations"
  }
}

lint {
  destructive {
    error = true
  }
}
