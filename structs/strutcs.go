package structs

type User struct {
	Username string `json:"username"`
	Password string `json:"pass"`
}

type Student struct {
	ID    int64
	Name  string `json:"name"`
	Major string `json:"major"`
}

type Admin struct {
	ID         int64
	Name       string `json:"name"`
	Password   string `json:"pass"`
	Username   string `json:"username"`
	Birth_date string `json:"birth_date"`
}

type Book struct {
	ID           int64
	Title        string `json:"title"`
	Author       string `json:"author"`
	Release_year string `json:"release_year"`
	Category     string `json:"category"`
	IsAvailable  bool
}

type LendingHistory struct {
	ID         int64
	Student_id int64 `json:"student_id"`
	Admin_id   int64 `json:"admin_id"`
	Book_id    int64 `json:"book_id"`
	Duration   int64 `json:"duration"`
	Created_at string
}

type ReturnHistory struct {
	ID         int64
	Student_id int64 `json:"student_id"`
	Admin_id   int64 `json:"admin_id"`
	Book_id    int64 `json:"book_id"`
	Created_at string
}

type History struct {
	ID         int64
	Student    string
	Book       string
	Admin      string
	Duration   int64
	Created_at string
}

type ReturnHist struct {
	ID         int64
	Student    string
	Book       string
	Admin      string
	Created_at string
}
