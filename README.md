# Todoapps

```text
create a folder named todoapps

go mod init todoapps

gogen init todocore

gogen usecase RunTodoCreate

gogen usecase RunTodoCheck

gogen usecase GetAllTodo

gogen repository SaveTodo Todo RunTodoCreate

gogen service GenerateRandomID RunTodoCreate

gogen error MessageMustNotEmpty

gogen repository FindOneTodo Todo RunTodoCheck 

gogen valuestring TodoID

gogen error Random6CharLengthNotSatisfied

gogen repository FindAllTodo Todo GetAllTodo

gogen gateway mysqldb

gogen gateway mongodb

gogen gateway inmemory

gogen application apptodo restapi inmemory

gogen error CannotCheckTheDoneTodoItem
```

