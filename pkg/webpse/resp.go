package webpse

// generic response

type Response[T any] struct {
	Data    T    `json:"data"`
	Success bool `json:"success"`
}

type UserInfo struct {
	UUID            string `json:"user_uuid"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	AvatarUrl       string `json:"avatar_url"`
	DailyQuota      int64  `json:"daily_quota"`
	DailyQuotaLimit int64  `json:"daily_quota_limit"`
	PermanentQuota  int64  `json:"permanent_quota"`
	UserPlan        string `json:"user_plan"`
}

type UserStats struct {
	UUID           string `json:"user_uuid"`
	TotalBytesSent int64  `json:"user_total_bytes_sent"`
}

type ProxyStatus struct {
	UUID      string `json:"proxy_uuid"`
	Name      string `json:"proxy_name"`
	OriginUrl string `json:"proxy_origin_url"`
	ProxyUrl  string `json:"proxy_proxy_url"`
	UserAgent string `json:"proxy_ua"`

	Quality        int64 `json:"proxy_quality"`
	CacheExpire    int64 `json:"proxy_cache_expire"`
	CacheSize      int64 `json:"proxy_cache_size"`
	CacheSizeLimit int64 `json:"proxy_cache_size_limit"`
}

type ProxyList = []ProxyStatus
