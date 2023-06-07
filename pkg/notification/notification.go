package notification

var N INotification

type INotification interface {
	send(string) error
}

type Notifier struct {
	notification INotification
}

func (n *Notifier) SetNotification(notification INotification) {
	n.notification = notification
}

func (n *Notifier) Send(msg string) error {
	return n.notification.send(msg)
}
