package services

import (
	model "crud-hex/internals/core/domain"
	"crud-hex/internals/core/ports"
	"time"
)
   
   type ProfilingService struct {
	profilingRepo ports.IProfilingRepository
   }
   
   func NewProfilingService(profilingRepo ports.IProfilingRepository) *ProfilingService {
	return &ProfilingService{
	 profilingRepo: profilingRepo,
	}
   }
   
   func (s *ProfilingService) Log(profiling model.Profiling) error {
	profiling.Timestamp = time.Now()
	return s.profilingRepo.Create(profiling)
   }