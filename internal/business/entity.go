type BusinessEntity struct {
    ID          string    `db:"id"`
    Name        string    `db:"name"`
    Password    string    `db:"password_hash"`
    Rating      float64   `db:"rating"`
    CreatedAt   time.Time `db:"created_at"`
    UpdatedAt   time.Time `db:"updated_at"`
}