package user
 
import (
    "context"
    "fmt"
	"github.com/hadisetiawan1992/sistem_digitalisasi/config"
    "github.com/hadisetiawan1992/sistem_digitalisasi/models"
    "log"
    "time"
)
 
const (
    table          = "musers"
    layoutDateTime = "2021-01-17 01:08:00"
)
 
// GetAll
func GetAll(ctx context.Context) ([]models.Users, error) {
 
    var users []models.Users
 
    db, err := config.MySQL()
 
    if err != nil {
        log.Fatal("Cant connect to MySQL", err)
    }
 
    queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)
 
    rowQuery, err := db.QueryContext(ctx, queryText)
 
    if err != nil {
        log.Fatal(err)
    }
 
    for rowQuery.Next() {
        var user models.Users
        var createdAt, updatedAt string
 
        if err = rowQuery.Scan(&user.ID,
            &user.Email,
            &user.Nama_lengkap,
            &user.Semester,
            &createdAt,
            &updatedAt); err != nil {
            return nil, err
        }
 
        //  Change format string to datetime for created_at and updated_at
        user.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
 
        if err != nil {
            log.Fatal(err)
        }
 
        user.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
 
        if err != nil {
            log.Fatal(err)
        }
 
        users = append(users, user)
    }
 
    return users, nil
}