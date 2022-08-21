package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

type LimiterIface interface {
	//對應的限流器的key名稱
	Key(c *gin.Context) string
	//取得權杖桶
	GetBucket(key string) (*ratelimit.Bucket, bool)
	//新增多個權杖桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	//自訂名稱
	Key string
	//間隔多久放權杖
	FillInterval time.Duration
	//桶的容量
	Capacity int64
	//到間隔時間後所放的實際權杖數量
	Quantum int64
}
