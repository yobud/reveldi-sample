package controllers

import "github.com/revel/revel"
import "github.com/zebigduck/reveldi"
import "app/app/modules/message"

type Dic struct {
	container *reveldi.Container
}

func (c *Dic) Init() {
	c.container = new(reveldi.Container)

	// Example : Messenger service

		// Register a simple instanciated Struct as a service
		// c.Container.Register("messenger", new(message.Messenger)) // Hi Jérémy!

		// You can easily change the Struct to use
		// c.Container.Register("messenger", new(message.SuperMessenger)) // super Hi Jérémy!

		// Create instance and configure service
		messenger := new(message.PrefixMessenger)
		messenger.SetPrefix("Hi")

		// Or use config.conf parameter
		messenger.SetPrefix(revel.Config.StringDefault("messenger.prefix", "Hey"))
		c.container.Register("messenger", messenger)

	// Now you can use it as follow anywhere in your controllers
	// myMessenger := Container.Get("messenger").(message.MessengerInterface)
}

func (c *Dic) Get(name string) reveldi.Service {
	return c.container.Get(name)
}
