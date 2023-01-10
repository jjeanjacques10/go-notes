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
