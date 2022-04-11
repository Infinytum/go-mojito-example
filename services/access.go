package services

import (
	"time"

	"github.com/golobby/container/v3"
	"github.com/infinytum/go-mojito"
)

func init() {
	mojito.Register(func() Access {
		access := Access{}
		container.Fill(&access)
		return access
	}, true)
}

var accessTimeCacheKey = "lastAccessTime"

type Access struct {
	Cache mojito.Cache `container:"type"`
}

func (s *Access) LastVisitedAt() time.Time {
	var accessTime time.Time
	s.Cache.GetOrDefault(accessTimeCacheKey, &accessTime, time.Now())
	return accessTime
}

func (s *Access) RecordVisit() {
	s.Cache.Set(accessTimeCacheKey, time.Now())
}
