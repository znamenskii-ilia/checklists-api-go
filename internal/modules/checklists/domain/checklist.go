package domain

import "time"

type Checklist struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tasks     []Task    `json:"tasks"`
}

func (c *Checklist) Rename(title string) {
	c.Title = title
	c.UpdatedAt = time.Now()
}

func (c *Checklist) AddTask(task Task) {
	c.Tasks = append(c.Tasks, task)
}

func (c *Checklist) RemoveTask(taskIndex int) {
	c.Tasks = append(c.Tasks[:taskIndex], c.Tasks[taskIndex+1:]...)
}

func (c *Checklist) UpdateTask(taskIndex int, task Task) {
	c.Tasks[taskIndex] = task
}
