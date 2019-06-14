package pipeline

//go:generate go get -u github.com/golang/dep/cmd/dep
//go:generate go get -u github.com/rubenv/sql-migrate/...
//go:generate go get -u golang.org/x/tools/cmd/goimports
//go:generate go get -u github.com/xo/xo
//go:generate go get -u github.com/gobuffalo/packr/packr
//go:generate go get -u github.com/mitchellh/gox

//go:generate dep ensure
//go:generate sql-migrate up
//go:generate xo pgsql://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable -o app/models --suffix=.xo.go --template-path templates/
//go:generate packr clean
//go:generate packr
