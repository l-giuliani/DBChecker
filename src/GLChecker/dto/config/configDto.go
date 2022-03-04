package dto

type Config struct {
	Db Db
	Getters []Getter
	Checks []Check
	Reports []Report
}

type Db struct {
    Type		string
    DbName		string
	Ip			string
	Port		uint
	User		string
	Password	string 
}

type Getter struct {
	Source string
    DbName string
}

type Check struct {
	Active		bool
	Type	 	string	`json:"type"`
	Field		Field
}

type Field struct {
	Name 	string
	Step 	uint
	Max		float32
	Min 	float32
}

type Report struct {
	Active 		bool
	Type 		string
	EmailParams EmailParams	
}

type EmailParams struct {
	Sender string
}