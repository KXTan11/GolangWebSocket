package company

import (
	"sort"
)

// Machine struct
type Machine struct {
	MachineName string
	MachineType byte

	// Owner
	Company *Company

	// Owning objects
	Tasks     []*Task
	FirstTask *Task
	LastTask  *Task
}

// CreateTask method
func (machine *Machine) CreateTask(TaskType byte, Duration int) *Task {
	task := &Task{
		TaskType:     TaskType,
		Duration:     Duration,
		Machine:      machine,
		PreviousTask: nil,
		NextTask:     nil,
	}

	// Run declarative functions here
	task.SetStartDateTime() // omit SetEndDateTime

	// Add task to this Machine list
	machine.Tasks = append(machine.Tasks, task)

	// Re-sort the tasks for machine

	return task
}

// RelationUpdateTasksSorting xaxa
func (machine *Machine) RelationUpdateTasksSorting() {
	// Sort tasks based on StartDateTime
	sort.SliceStable(machine.Tasks, func(i, j int) bool {
		return machine.Tasks[i].StartDateTime < machine.Tasks[j].StartDateTime
	})

	// Set machine first and last task, and every task's previous and next task
	for k, t := range machine.Tasks {
		if k == 0 {
			machine.FirstTask = t
		} else {
			value := machine.Tasks[k-1]
			t.PreviousTask = CalcFuncRelation(t.PreviousTask, value, t.SetStartDateTime).(*Task)
		}

		if k == len(machine.Tasks)-1 {
			machine.LastTask = t
		} else {
			value := machine.Tasks[k+1]
			t.NextTask = CalcFuncRelation(t.NextTask, value).(*Task)
		}
	}
}
