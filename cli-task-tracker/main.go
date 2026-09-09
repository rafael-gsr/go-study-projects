package main

func main() {
	persistence := NewPersistence("tasks.txt", "")

	if persistence != nil {
		panic("Error on creating persistence package")
	}

	storedContent := persistence.Read()
	tasks := NewTaskList(storedContent)
	flags := NewFlags(persistence, tasks)
	flags.Setup()
}
