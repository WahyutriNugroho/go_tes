//Models/UserModel.go
package Models

type Predict struct {
	// Id      uint   `json:"id"`
	// Name    string `json:"name"`
	// Email   string `json:"email"`
	// Phone   string `json:"phone"`
	// Address string `json:"address"`
	Id          int     `json:id`
	Class_name  string  `json:"class_name"`
	ImgUrl      string  `json:"imgUrl"`
	Date        string  `json:"date"`
	Latitude    string  `json:"latitude"`
	Longtitude  string  `json:"longtitude"`
	Probability float64 `json:"probability"`
}

func (b *Predict) TableName() string {
	return "predict"
}
