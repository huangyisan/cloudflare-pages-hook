package notification

type INotification interface {
	send(string) bool
}

type notifier struct {
	notification INotification
}

func (n *notifier) setNotification(notification INotification) {
	n.notification = notification
}

func (n *notifier) send(msg string) bool {
	return n.notification.send(msg)
}
