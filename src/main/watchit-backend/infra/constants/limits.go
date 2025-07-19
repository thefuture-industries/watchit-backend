package constants

const (
	LimitFree = 101010
	LimitPay  = 111111
)

const (
	FreeMaxQueryLengthUsage   = 255  // Длина поискового запроса (макс)
	PayMaxQueryLengthUsage    = 1000 // Длина поискового запроса (макс)
	FreeDailySearchLimitUsage = 10   // Кол-во поисков в день
	PayDailySearchLimitUsage  = 100  // Кол-во поисков в день
	FreeSearchPriorityUsage   = 0    // Приоритет в очереди (0 — низкий)
	PaySearchPriorityUsage    = 1    // Приоритет в очереди (0 — низкий)
)
