package main

import (
	"context"
	"github.com/davidgardner11/hellotemporal/workflow"
	"go.temporal.io/sdk/client"
	"log"
)

func main() {
	c, err := client.NewClient(client.Options{}) 
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:	"hello-workflow-video",
		TaskQueue: 	"hello-world"
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, helloworkflow.Workflow, "First Video")
	if err != nil {
		log.Fatalln("Unable to start workflow", err)
	}

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}
	log.Println("Workflow result: " result)

}

