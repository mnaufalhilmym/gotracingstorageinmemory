package gotracingstorageinmemory

import (
	"sync"

	"github.com/mnaufalhilmym/gotracing"
)

type storage struct {
	sync.Mutex
	MaxTracesPerLevel uint
	traces            map[gotracing.Level][]gotracing.Stacktraces
}

func (s *storage) Insert(level gotracing.Level, trace gotracing.Stacktraces) {
	s.Lock()
	if len(s.traces[level]) >= int(s.MaxTracesPerLevel) {
		s.traces[level] = s.traces[level][1:]
	}
	s.traces[level] = append(s.traces[level], trace)
	s.Unlock()
}

func (s *storage) GetAll(level gotracing.Level) []gotracing.Stacktraces {
	return s.traces[level]
}

func New(maxTracesPerLevel uint) storage {
	return storage{
		MaxTracesPerLevel: maxTracesPerLevel,
		traces:            make(map[gotracing.Level][]gotracing.Stacktraces),
	}
}
