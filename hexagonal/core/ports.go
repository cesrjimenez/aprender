package core

type Inventory interface {
	Reserve(itemID string) error
}

type Notifier interface {
	SendConfirmation(email, message string) error
}
