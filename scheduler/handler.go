package scheduler

func (s *Scheduler) ResourceOffer(offers []*mesos.Offer) {
	select {
		case job := <- s.pendingJobs:
                	s.LaunchJob(job, offers)   
                default:
               		s.DeclineOffers(offers) 
	}
}
