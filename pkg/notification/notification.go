package notification

var N INotification

type INotification interface {
	send(string) bool
}

type Notifier struct {
	notification INotification
}

func (n *Notifier) SetNotification(notification INotification) {
	n.notification = notification
}

func (n *Notifier) Send(msg string) bool {
	return n.notification.send(msg)
}
