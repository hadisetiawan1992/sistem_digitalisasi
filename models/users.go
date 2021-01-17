package models
 
import (
    "time"
)
 
type (
    // User
    Users struct {
        ID        		int       `json:"id"`
        Email       	int       `json:"email"`
        Nama_lengkap   	string    `name:"nama_lengkap"`
        Semester  		int       `json:"semester"`
        CreatedAt time.Time `json:"created_at"`
        UpdatedAt time.Time `json:"updated_at"`
    }
)
 