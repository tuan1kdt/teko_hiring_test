package main

import "time"

func newRateLimitChecker(maxRequestPerHour int) rateLimitChecker {
	return rateLimitChecker{
		maxRequestPerHour: maxRequestPerHour,
		requestLogs:       map[time.Time]bool{},
	}
}

type rateLimitChecker struct {
	maxRequestPerHour int
	total             int
	requestLogs       map[time.Time]bool
}

func (r *rateLimitChecker) removeExpiredLog(requestTime time.Time) {
	lastAcceptedTime := requestTime.Add(-time.Hour)
	for t, _ := range r.requestLogs {
		if t.Before(lastAcceptedTime) {
			//TODO: handle race condition
			delete(r.requestLogs, t)
			r.total--
		}
	}
}

func (r *rateLimitChecker) addLog(requestTime time.Time) {
	//TODO: handle race condition
	r.requestLogs[requestTime] = true
	r.total++
}

func (r *rateLimitChecker) isRequestOverflow() bool {
	if r.total == r.maxRequestPerHour {
		return true
	}
	return false

}
