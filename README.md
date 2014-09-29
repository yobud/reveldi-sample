# Reveldi sample - Simple Dependency Injection for Go apps

This Revel project shows how [reveldi](https://github.com/ZeBigDuck/reveldi) can be implemented.

You'll find some messenger Structs in app/modules/message which are just test services.

In the [app/controllers/dic.go](https://github.com/ZeBigDuck/reveldi-sample/blob/master/app/controllers/dic.go)., you can uncomment some lines to see that the registered service can be changed.

The most important point is that your service is registered only once here.

You can instantiate once all Structs you need with custom parameters (and why not others services). And you'll be able to use them in any controller you want.

In your controller, just tell to Container.Get() that your service must be an interfacing whatever you want (here [messengerInterface](https://github.com/ZeBigDuck/reveldi-sample/blob/master/app/modules/message/messengerInterface.go), and you will not have to care if the Struct is a [prefixMessenger](https://github.com/ZeBigDuck/reveldi-sample/blob/master/app/modules/message/prefixMessenger.go), [superMessenger](https://github.com/ZeBigDuck/reveldi-sample/blob/master/app/modules/message/superMessenger.go) or simple [messenger](https://github.com/ZeBigDuck/reveldi-sample/blob/master/app/modules/message/messenger.go).