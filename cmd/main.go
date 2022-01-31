package cmd

import (
	"github.com/misraak/app/web/app/rest"
)

func Init() error {
	return rest.Register()
}
