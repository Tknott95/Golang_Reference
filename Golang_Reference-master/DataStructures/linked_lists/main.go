package main

type Student struct {
	age    int
	weight int
	name   string
	Next   *Student
}

type Teacher struct {
	age    int
	weight int
	name   string
	Next   *Student
}

func main() {
	tex := Student{22, 200, "Tex", nil}
	ollie := Student{20, 150, "Ollie", &tex}

	msStinky := Teacher{55, 133, "msStinky", &ollie}

	println("The teacher is: ", msStinky.name)
	println("The name of the student directly behind the teacher dealing with her smell is: ", msStinky.Next.name)
	println("The Student laughing at ", msStinky.Next.name, " is named ", msStinky.Next.Next.name, " ... ", ollie.Next.name, " IS RUDE MOST LIKELY BECAUSE HE IS ", ollie.Next.age, " YEARS OLD!")
}
