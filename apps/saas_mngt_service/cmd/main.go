package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/infrastructures/viper"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// var (
// 	// For this example, we're using just a simple logger implementation,
// 	// You probably want to ship your own implementation of `watermill.LoggerAdapter`.
// 	logger = watermill.NewStdLogger(false, false)
// )

func main() {
	mux := http.NewServeMux()

	config := viper.NewConfig()
	// logger := zap.NewLogger(config)
	// redis := redis.NewRedisCache(config)

	// userCommand := wire.InitUserCommand()
	// userQuery := wire.InitUserQuery()
	// userAPI := wire.InitUserAPIHandler(userCommand, userQuery)

	fmt.Print("Config loaded::: ", config.Server.Port, "\n")

	// mux.Handle(saasv1connect.NewUserAPIHandler(userAPI))

	// config := wire.InitViper()
	// logger := wire.InitZapLogger(config)
	// redisCache := wire.InitRedisCache(config)
	// wire.InitUserRouterHandler(redisCache, logger)

	////////////////

	// router, err := message.NewRouter(message.RouterConfig{}, logger)
	// if err != nil {
	// 	panic(err)
	// }

	// SignalsHandler will gracefully shutdown Router when SIGTERM is received.
	// You can also close the router by just calling `r.Close()`.
	// router.AddPlugin(plugin.SignalsHandler)

	// Router level middleware are executed for every message sent to the router
	// router.AddMiddleware(
	// 	// CorrelationID will copy the correlation id from the incoming message's metadata to the produced messages
	// 	middleware.CorrelationID,

	// 	// The handler function is retried if it returns an error.
	// 	// After MaxRetries, the message is Nacked and it's up to the PubSub to resend it.
	// 	middleware.Retry{
	// 		MaxRetries:      3,
	// 		InitialInterval: time.Millisecond * 100,
	// 		Logger:          logger,
	// 	}.Middleware,

	// 	// Recoverer handles panics from handlers.
	// 	// In this case, it passes them as errors to the Retry middleware.
	// 	middleware.Recoverer,
	// )

	// For simplicity, we are using the gochannel Pub/Sub here,
	// You can replace it with any Pub/Sub implementation, it will work the same.
	// pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)

	// Producing some incoming messages in background
	// go publishMessages(pubSub)

	// AddHandler returns a handler which can be used to add handler level middleware
	// or to stop handler.
	// handler := router.AddHandler(
	// 	"struct_handler",          // handler name, must be unique
	// 	"incoming_messages_topic", // topic from which we will read events
	// 	pubSub,
	// 	"outgoing_messages_topic", // topic to which we will publish events
	// 	pubSub,
	// 	userRouter.Handler2,
	// )

	// Handler level middleware is only executed for a specific handler
	// Such middleware can be added the same way the router level ones
	// handler.AddMiddleware(func(h message.HandlerFunc) message.HandlerFunc {
	// 	return func(message *message.Message) ([]*message.Message, error) {
	// 		log.Println("executing handler specific middleware for ", message.UUID)
	// 		return h(message)
	// 	}
	// })

	// just for debug, we are printing all messages received on `incoming_messages_topic`
	// router.AddNoPublisherHandler(
	// 	"print_incoming_messages",
	// 	"incoming_messages_topic",
	// 	pubSub,
	// 	userRouter.PrintMessages,
	// )

	// just for debug, we are printing all events sent to `outgoing_messages_topic`
	// router.AddNoPublisherHandler(
	// 	"print_outgoing_messages",
	// 	"outgoing_messages_topic",
	// 	pubSub,
	// 	printMessages,
	// )

	// Now that all handlers are registered, we're running the Router.
	// Run is blocking while the router is running.
	// ctx := context.Background()
	// if err := router.Run(ctx); err != nil {
	// 	panic(err)
	// }

	//////////////////
	fmt.Print("Starting server at ", config.Server.Port, "\n")

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", 4000),
		Handler:      h2c.NewHandler(mux, &http2.Server{}),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	err := server.ListenAndServe()

	log.Fatalf("listen failed: %v", err)
}

// func publishMessages(publisher message.Publisher) {
// 	for {
// 		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))
// 		middleware.SetCorrelationID(watermill.NewUUID(), msg)

// 		log.Printf("sending message %s, correlation id: %s\n", msg.UUID, middleware.MessageCorrelationID(msg))

// 		if err := publisher.Publish("incoming_messages_topic", msg); err != nil {
// 			panic(err)
// 		}

// 		time.Sleep(time.Second)
// 	}
// }

// func printMessages(msg *message.Message) error {
// 	fmt.Printf(
// 		"\n> Received message: %s\n> %s\n> metadata: %v\n\n",
// 		msg.UUID, string(msg.Payload), msg.Metadata,
// 	)
// 	return nil
// }

// type structHandler struct {
// 	// we can add some dependencies here
// }

// func (s structHandler) Handler2(msg *message.Message) ([]*message.Message, error) {
// 	log.Println("structHandler received message", msg.UUID)

// 	msg = message.NewMessage(watermill.NewUUID(), []byte("message produced by structHandler"))
// 	return message.Messages{msg}, nil
// }
