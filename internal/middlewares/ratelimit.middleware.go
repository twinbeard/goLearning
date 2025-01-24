package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/twinbeard/goLearning/global"
	"github.com/ulule/limiter/v3"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
)

type RateLimiter struct {
	globalRateLimiter         *limiter.Limiter // Global API
	publicAPIRateLimiter      *limiter.Limiter // Public API
	userPrivateAPIRateLimiter *limiter.Limiter // Private API
}

/*
You can also use the simplified format "<limit>-<period>"", with the given periods:

  - "S": second

  - "M": minute

  - "H": hour

  - "D": day

    Examples:

  - 5 reqs/second: "5-S"

  - 10 reqs/minute: "10-M"

  - 1000 reqs/hour: "1000-H"

  - 2000 reqs/day: "2000-D"
*/

func NewRateLimitter() *RateLimiter {
	rate := &RateLimiter{
		globalRateLimiter:         rateLimiter("100-S"), // => 100 req/1s
		publicAPIRateLimiter:      rateLimiter("80-S"),
		userPrivateAPIRateLimiter: rateLimiter("50-S"),
	}
	return rate
}

func rateLimiter(interval string) *limiter.Limiter {
	store, err := redisStore.NewStoreWithOptions(global.Rdb, limiter.StoreOptions{
		Prefix:          "rate-limiter", // Prefix for keys in redis such rate-limitter:xxxx
		MaxRetry:        3,              // if error occurs when access redis, redis will retry 3 times
		CleanUpInterval: time.Hour,      // remove expired keys in redis every hour => Good for performance
	})
	if err != nil {
		return nil
	}

	rate, err := limiter.NewRateFromFormatted(interval) // 10-M => 10 requests per minute
	if err != nil {
		panic(err)
	}

	instance := limiter.New(store, rate)
	return instance
}

// GLOBAL LIMITER

func (rl *RateLimiter) GlobalRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "global" // unit

		log.Println("global--->")

		limitContext, err := rl.globalRateLimiter.Get(c, key) // Get the rate limit threshold setup and current limit (count) => This limit will automatically increment limit when user access
		if err != nil {
			fmt.Println("Failed to check rate limit GLOBAL", err)
			c.Next()
			return
		}
		// If the current limit is beyond the limit threshold, return error below
		if limitContext.Reached {
			log.Printf("Rate limit breached GLOBAL %s", key)
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached GLOBAL, try later"})
			return
		}

		c.Next()
	}
}

// PUBLIC API LIMITER

func (rl *RateLimiter) PublicAPIRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path // Get the URL path: /v1/2024/user/login

		rateLimitPath := rl.filterLimitUrlPath(urlPath)

		if rateLimitPath != nil {
			log.Println("Client Ip--->", c.ClientIP())

			key := fmt.Sprintf("%s-%s", "111-222-333-44")

			limitContext, err := rateLimitPath.Get(c, key)

			if err != nil {
				fmt.Println("Failed to check rate limit", err)
				c.Next()
				return
			}

			if limitContext.Reached {
				log.Printf("Rate limit breached %s", key)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached, try later"})
				return
			}
		}

		c.Next()
	}
}

// PRIVATE API LIMITER

func (rl *RateLimiter) UserAndPrivateRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		rateLimitPath := rl.filterLimitUrlPath(urlPath)

		if rateLimitPath != nil {
			userId := 1001 // Replace with actual logic to get user ID => context.GetSubjectUUID()

			key := fmt.Sprintf("%d-%s", userId, urlPath) // User ID + Path

			limitContext, err := rateLimitPath.Get(c, key)

			if err != nil {
				fmt.Println("Failed to check rate limit", err)
				c.Next()
				return
			}

			if limitContext.Reached {
				log.Printf("Rate limit breached %s", key)
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit breached, try later"})
				return
			}
		}

		c.Next()
	}
}
func (rl *RateLimiter) filterLimitUrlPath(urlPath string) *limiter.Limiter {
	// we can add more path here but we have to figure out which path is for public and which is for private
	if urlPath == "/v1/2024/user/login" || urlPath == "/ping/80" {
		return rl.publicAPIRateLimiter
	} else if urlPath == "/v1/2024/user/info" || urlPath == "/ping/50" {
		return rl.userPrivateAPIRateLimiter
	} else {
		return rl.globalRateLimiter
	}
}
