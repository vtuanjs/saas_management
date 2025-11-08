module github.com/vtuanjs/saas_management/apps/saas_mngt_service

go 1.24.5

replace github.com/vtuanjs/saas_management/packages/go => ../../packages/go

require (
	github.com/bsm/redislock v0.9.4
	github.com/dgraph-io/ristretto v0.2.0
	github.com/google/wire v0.6.0
	github.com/vtuanjs/saas_management/packages/go v0.0.0-00010101000000-000000000000
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sagikazarmark/locafero v0.10.0 // indirect
	github.com/sourcegraph/conc v0.3.1-0.20240121214520-5f936abd7ae8 // indirect
	github.com/spf13/afero v1.14.0 // indirect
	github.com/spf13/cast v1.9.2 // indirect
	github.com/spf13/pflag v1.0.7 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/redis/go-redis/v9 v9.11.0
	github.com/spf13/viper v1.20.1
	go.uber.org/zap v1.27.0
	golang.org/x/net v0.43.0
	golang.org/x/text v0.28.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)
