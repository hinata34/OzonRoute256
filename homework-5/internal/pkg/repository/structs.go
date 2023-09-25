package repository

type User struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	Age  int64  `db:"age"`
}

type Address struct {
	ID          int64  `db:"id"`
	HouseNumber int64  `db:"house_number"`
	StreetName  string `db:"street_name"`
	UserID      int64  `db:"user_id"`
}
