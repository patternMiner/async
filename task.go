package async

import (
	"log"
)

type Task interface {
	Run()
}

type TaskRunner struct {
	ID int
	TaskChannel chan Task
	TaskRunnerQueue chan chan Task
}

func NewTaskRunner(id int, taskRunnerQueue chan chan Task) TaskRunner {
	return TaskRunner {
		ID: id,
		TaskChannel: make(chan Task),
		TaskRunnerQueue: taskRunnerQueue,
	}
}

func (r *TaskRunner) Start() {
	go func() {
		for {
			// adds this runner's task channel to the task runner queue, to be dequeued by the dispatcher.
			r.TaskRunnerQueue <- r.TaskChannel
			// receives work from the dispatcher after getting its task channel dequeued from task runner queue.
			select {
			case task := <- r.TaskChannel:
				log.Printf("Asynchronous task runner %d is running a task.\n", r.ID)
				// runs the task, and goes back to the the beginning of the loop
				task.Run()
			}
		}
	}()
}

var (
	TaskQueue chan Task
	TaskRunnerQueue chan chan Task
)

func StartTaskDispatcher(runnerCount int) {
	TaskRunnerQueue = make(chan chan Task, runnerCount)
	// creates the specified number of task runners, each of which will add their task channels
	// to the task runner queue.
	for i := 0; i < runnerCount; i++ {
		log.Printf("Starting asynchronous task runner %d", i+1)
		taskRunner := NewTaskRunner(i+1, TaskRunnerQueue)
		taskRunner.Start()
	}
	// receives tasks from the task queue, and dispatches them to available task runners
	go func() {
		for {
			select {
			case task := <- TaskQueue:
				go func() {
					taskChannel := <- TaskRunnerQueue
					taskChannel <- task
				}()
			}
		}
	}()
}

func init() {
	TaskQueue = make(chan Task, 100)
}
