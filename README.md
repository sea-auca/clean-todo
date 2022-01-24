# clean-todo
An example of go app with clean application

## Functionality

This is a simple apiserver for todo-list application.
Here are basic requirements that we are trying to achieve with this app.
1. It has auth with jwt and oauth, with refresh tokens too
2. We have full CRUD for todos
3. It has endpoints for admins and analytics
4. This application is very error tolerant
5. It is very to configure this application
## How to use

When you clone this repository you have to get all the dependencies with following command in your terminal
```
	go mod tidy
```

After that you can start app with
```
	make run
```

In order to build it use command bellow
```
	make build
```