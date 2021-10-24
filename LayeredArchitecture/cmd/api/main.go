package main

import interfaces "github.com/xxarupakaxx/go-todo-api/LayeredArchitecture/interfaces/handler"

func main() {
	newsHandler := InjectNewsHandler()
	topicHandler := InjectTopicHandler()
	interfaces.InitRouting(topicHandler,newsHandler)
}