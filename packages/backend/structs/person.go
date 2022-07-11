package structs

type Person struct {
	ID string `uri:"id" binding:"required,uuid"`
}
