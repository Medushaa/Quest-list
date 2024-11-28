package main

func main() {

    //for store in csv
    todos := Todos{} //new storage
    //storage := NewStorage[Todos]("todos.json") //new todos.json file
    //hard coding json file location for when we call the cmd globally:
    storage := NewStorage[Todos]("D:/Code/GoLang/todo-cli-tool/quests.json")

    storage.Load(&todos) //retrive data from json to todos's address

    //Using commands:
    cmdFlags := NewCmdFlags() //initiate flags and read arg
    cmdFlags.Execute(&todos) //execute flag

    storage.Save(todos) //save the todo slice back into json
}
