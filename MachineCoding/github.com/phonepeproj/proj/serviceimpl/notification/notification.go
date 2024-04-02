package notification

// Publisher interface for adding observers and notifying them
type Publisher interface {
	Subscribe(Observer)
	NotifyAllOnDocShare()
}

type NotificationSubjectSvcImpl struct {
	Observers []Observer
}

func NewNotificationSvcImpl() *NotificationSubjectSvcImpl {
	return &NotificationSubjectSvcImpl{Observers: make([]Observer, 0)}
}

// Subscribe adds a new observer to the document subject
func (ns *NotificationSubjectSvcImpl) Subscribe(observer Observer) {
	ns.Observers = append(ns.Observers, observer)
}

func (ns *NotificationSubjectSvcImpl) NotifyAllOnDocShare() {
	for _, observer := range ns.Observers {
		observer.ProcessDocShare()
	}
}

// Observer defines the interface for receiving notifications
type Observer interface {
	ProcessDocShare()
}
