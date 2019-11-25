package main

type tool struct {
	name string
	use string
	kind struct{
		student string
	}
}

func main()  {
	t :=tool{name:"pencil",use:"write"}
	t.kind.student="jack"
}