package ecs

type TaskDefinition struct {
}

type Task struct {
	taskDefinition *TaskDefinition
}

type Service struct {
	task *Task
}
