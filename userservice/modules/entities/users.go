package entities

type UserService interface {
	UserCreated(command *UserCreated) (UserCreated, error)
	UserUpdated(command *UserUpdated) (UserUpdated, error)
	UserDeleted(command *UserDeleted) (UserDeleted, error)
}

type Account_userRepository interface {
	CreateAccount(*Account_user) (*Account_user, error)
	UpdateAccount(*Account_user) (*Account_user, error)
	DeleteAccount(string) (string, error)
	CheckAccount(string) (string, error)
}

type HistoryRepository interface {
	GetHistory(string) ([]History, error)
	CreateHistory(History) (*History, error)
}

type HistoryService interface {
	GetHistory(string) ([]History, error)
}

// fix table name
func (a Account_user) TableName() string {
	return "account_user"
}
func (h History) TableName() string {
	return "history"
}

type Account_user struct {
	UserID   string `db:"user_id" json:"user_id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Address  string `db:"address" json:"address"`
}

type UserRead struct {
	UserID string `db:"user_id" json:"user_id"`
	ID     string `db:"id" json:"id"`
}

type UserGetCatalogs struct {
	UserID string `db:"user_id" json:"user_id"`
}

type UserFavorite struct {
	UserID string `db:"user_id" json:"user_id"`
	ID     string `db:"id" json:"id"`
}

type UserGetFavorite struct {
	UserID string `db:"user_id" json:"user_id"`
}