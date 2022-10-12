# Simple CLI app in Go

Creating Conference booking App

### Some codes related to Go

- If you have multiple "**.go**" files(e.g., main.go and helper.go) interrelated to each other, then run the command:
  - `go run main.go helper.go`
- As it is not possible to add multiple files one by one in a single command, we can use an alternative command:
  - `go run .`
  - This command will run all the files in the current folder("**.**" signifies the current folder)
- If we have multiple packages and we want to use the each others' code then we can do it by following:
  - Different package should have different folders.
  - E.g., In our case, we have **helper.go** in **helper** folder, making it a different package
  - Making the code available/access by other packages we have to make it public/export it from there
  - In **Go** we make this happen by capitalizing the first letter of the code block
  - E.g., In our case, we achieved it by making **HandleUserName()**, the **H** is capital
  - Then, we need to use it in another package, **main.go** in our case
  - We did it by importing **booking-app/helper**
    - **booking-app** is the module name tht we used while starting the project, present in the **go.mod**
    - **helper** is the another package where the code block is present that we want to access in **main.go**
    - We simply accessed using **helper.HandleUserName()**
