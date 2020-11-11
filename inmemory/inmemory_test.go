package inmemory

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetMetricCountNotExistKeyReturnError(t *testing.T) {
	metrict := NewMetric()
	_, err := metrict.GetMetricCount("testtKey", time.Now())
	assert.Error(t, err)
}
func Test_GetMetricCountExistKey(t *testing.T) {
	key := "testKey"
	metrict := NewMetric()
	metrict.AddMetric(key, time.Now())
	count, err := metrict.GetMetricCount(key, time.Now())
	assert.Nil(t, err)
	assert.Equal(t, 1, count)
}

func Test_AddManyKeyDuringSameHour(t *testing.T) {
	key := "testKey"
	metrict := NewMetric()
	metrict.AddMetric(key, time.Now())
	metrict.AddMetric(key, time.Now().Add(time.Minute*20))
	metrict.AddMetric(key, time.Now().Add(time.Minute*30))
	metrict.AddMetric(key, time.Now().Add(time.Minute*40))
	metrict.AddMetric(key, time.Now().Add(time.Minute*59))
	count, err := metrict.GetMetricCount(key, time.Now())
	assert.Nil(t, err)
	assert.Equal(t, 5, count)
}

func Test_AddManyKeyAfterHour(t *testing.T) {
	key := "testKey"
	metrict := NewMetric()
	metrict.AddMetric(key, time.Now())
	metrict.AddMetric(key, time.Now().Add(time.Minute*59))
	count, err := metrict.GetMetricCount(key, time.Now())
	assert.Nil(t, err)
	assert.Equal(t, 2, count)

	metrict.AddMetric(key, time.Now().Add(time.Minute*60))
	count, err = metrict.GetMetricCount(key, time.Now())
	assert.Nil(t, err)
	assert.Equal(t, 1, count)
}

func Test_AddManyKeyAndGetAfterHour(t *testing.T) {
	key := "testKey"
	metrict := NewMetric()
	metrict.AddMetric(key, time.Now())
	metrict.AddMetric(key, time.Now().Add(time.Minute*59))
	count, err := metrict.GetMetricCount(key, time.Now())
	assert.Nil(t, err)
	assert.Equal(t, 2, count)

	count, err = metrict.GetMetricCount(key, time.Now().Add(time.Minute*60))
	assert.Nil(t, err)
	assert.Equal(t, 0, count)
}
