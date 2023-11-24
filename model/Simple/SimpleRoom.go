package simple

// SimpleRoom โครงสร้างข้อมูลที่แสดงเฉพาะบางฟิลด์
type SimpleRoom struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
}
