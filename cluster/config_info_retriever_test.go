package cluster

import (
	"testing"

	"github.com/felippeduran/pitaya/v2/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfigInfoRetrieverRegion(t *testing.T) {
	t.Parallel()

	c := viper.New()
	c.Set("pitaya.cluster.info.region", "us")
	conf := config.NewConfig(c)

	infoRetriever := NewConfigInfoRetriever(config.NewConfigInfoRetrieverConfig(conf))

	assert.Equal(t, "us", infoRetriever.Region())
}
