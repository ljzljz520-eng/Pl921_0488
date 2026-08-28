package storage

func (s *Store) Healthy() bool    { return s != nil && s.db != nil }
func (s *Store) BucketCount() int { return len(buckets) }
