package notifier

var N INotifier

type INotifier interface {
	Send(string) error
	SetChatID(string)
}

type Notifier struct {
	notifier INotifier
}

func (n *Notifier) SetNotification(notifier INotifier) {
	n.notifier = notifier
}
