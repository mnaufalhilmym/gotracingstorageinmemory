package gotracingstorageinmemory

import (
	"sync"

	"github.com/mnaufalhilmym/gotracing"
)

type Storage struct {
	sync.Mutex
	MaxTracesPerLevel uint
	traces            map[gotracing.Level][]gotracing.Traces
}

func (s *Storage) Insert(level gotracing.Level, trace gotracing.Traces) {
	s.Lock()
	if len(s.traces[level]) >= int(s.MaxTracesPerLevel) {
		s.traces[level] = s.traces[level][1:]
	}
	s.traces[level] = append(s.traces[level], trace)
	s.Unlock()
}

func (s *Storage) GetAll(level gotracing.Level) []gotracing.Traces {
	return s.traces[level]
}

func New(maxTracesPerLevel uint) Storage {
	return Storage{
		MaxTracesPerLevel: maxTracesPerLevel,
		traces:            make(map[gotracing.Level][]gotracing.Traces),
	}
}
