package handler

type OfferHandler struct {
	fw *Framework
}

func NewOfferHandler() *OfferHandler {
}

func (h *OfferHandler) Handle(event *sched.Event) {
	for _, offer := range event.Offers.Offers {
		for _, task := range h.pendingTasks {
		}

		h.LaunchTasks(tasks)
	}
}
