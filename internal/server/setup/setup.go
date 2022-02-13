package setup

import "github.com/arpit006/go_txns/internal/config"

func initialize() error {
	// load config
	if err := config.Load(); err != nil {
		// TODO: log
		return err
	}
	// instrumentation setup
	return nil
}
