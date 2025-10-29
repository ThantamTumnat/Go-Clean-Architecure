package entities

type UserService interface {
	UserCreated(command *UserCreated) (UserCreated, error)
	UserUpdated(command *UserUpdated) (UserUpdated, error)
	UserDeleted(command *UserDeleted) (UserDeleted, error)
}


type AccountRepository interface {
	CreateAccount(*Account) (*Account, error)
	UpdateAccount(*Account) (*Account, error)
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

type CatalogRepository interface {
	GetCatalogs() ([]CatalogDB, error)
	GetCatalog(string) (*CatalogDB, error)
	GetCatalogId(string) (string, error)
}

type FavoriteRepository interface {
	GetFavorite(string) ([]Favorite, error)
	CreateFavorite(Favorite) (*Favorite, error)
}

type FavoriteService interface {
	GetFavorite(string) ([]Favorite, error)
	CreatedFavorite(string, string) (*Favorite, error)
}

type CatalogService interface {
	GetCatalogs() ([]Catalog, error)
	GetCatalog(string, string) (History, error)
}


// fix table name
//
//	func (a Account_user) TableName() string {
//		return "account_user"
//	}
func (h History) TableName() string {
	return "history"
}
func (a Account) TableName() string {
	return "account"
}
func (c CatalogDB) TableName() string {
	return "catalog"
}
func (f Favorite) TableName() string {
	return "favorite"
}

// type Account_user struct {
// 	UserID   string `db:"user_id" json:"user_id"`
// 	Username string `db:"username" json:"username"`
// 	Password string `db:"password" json:"password"`
// 	Address  string `db:"address" json:"address"`
// }

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

type Favorite struct {
	UserID    string `db:"user_id" json:"user_id"`
	ID        string `db:"id" json:"id"`
	Type      string `db:"type" json:"type"`
	Setup     string `db:"setup" json:"setup"`
	Punchline string `db:"punchline" json:"punchline"`
}

type CatalogDB struct {
	ID        int `db:"id" json:"id"`
	Type      string `db:"type" json:"type"`
	Setup     string `db:"setup" json:"setup"`
	Punchline string `db:"punchline" json:"punchline"`
}


