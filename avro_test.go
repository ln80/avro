package avro_test

import "github.com/ln80/avro/v2"

func ConfigTeardown() {
	// Reset the caches
	avro.DefaultConfig = avro.Config{}.Freeze()
}
