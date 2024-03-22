package structs
type SignUpStruct struct {
	Id       string `bson:"_id"`
	Name     string
	Surname  string
	Login    string
	Password string
}
type SearchMusic struct{
	Id       string `bson:"_id"`
	Name string
	Author string
	Date string
}