package notification

import "github.com/gen2brain/beeep"

type Notification struct {
	icon    string
	title   string
	message string
}

// Comment
func New(title string, message string, icon string) *Notification {
	return &Notification{
		icon:    icon,
		title:   title,
		message: message,
	}
}

// Comment
func (ctx *Notification) Show() error {
	return beeep.Alert(ctx.title, ctx.message, ctx.icon)
}
