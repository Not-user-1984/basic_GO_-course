package json
import "gorm.io/gorm"



type Person struct {
	gorm.Model
	Name  string `json:"name"`
	Age   string `json:"age"`
	Email string `json:"email"`
	JobID int    `json:"-"`
	Job   Job    `json:"job" gorm:"foreignKey:JobID"`
}
type Job struct {
	ID      int  
	Name    string `json:"name"`
	Company string `json:"company"`
	Pay     float32 `json:"pay"`
}