package controllers

import "github.com/revel/revel"
import "app/app/modules/message"

type App struct {
	*revel.Controller
}

// App
func (c App) Index() revel.Result {
	myMessenger := Container.Get("messenger").(message.MessengerInterface)
	myMessenger.SetMessage("Jérémy!")

	return c.Render(c, myMessenger)
}
