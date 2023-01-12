# Notes GO

Repository for record my learning in GO.

## Courses and Videos

- [Alura - Go: a linguagem do Google](https://cursos.alura.com.br/course/golang)
- [YouTube Channel HunCoding](https://www.youtube.com/@huncoding)

## Notes

- `:=` : Declaration and attribution  
- `&` : The address of
- Function to return two values

``` go
func getTwoNumbers() (int, int) {
 number1 := 1
 number2 := 2
 return number1, number2
}

number1, number2 := getTwoNumbers()
_, number2 := getTwoNumbers() // Ignore first return
```

- `while` not exist, but you can use `for`

``` go
for {

}
```

- Create a slice (array abstract and dynamic)

``` go
func showNames(){
 names := []string {"Jean", "Alice", "Kal-El"}
 fmt.Println(names)
}
```

- Convert types using strconv

## Project Structure

- `/pkg`: This is where to save all installed dependencies in your projects
- `/src`: This is where all the code is saved.
    - Example: /github.com/jjeanjacques10/crud-golang
- `/bin`: This is where you save the binaries that the "go install" command uses to compile

## GOPATH vs GOROOT

- `$GOROOT` is an environment variable that specifies the location of the Go runtime. It is typically set when Go is installed and should not need to be modified. 
- `$GOPATH`, on the other hand, is an environment variable that specifies the location of your Go workspace, where your Go code and dependencies are stored.

## Documentation

- <https://pkg.go.dev/os>
- <https://pkg.go.dev/time>
