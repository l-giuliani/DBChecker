package dto

type Config struct {
	Db Db
	Checks []Check
	Report []Report
}

type Db struct {
    Type		string
    DbName		string
	Ip			string
	Port		uint
	User		string
	Password	string 
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