package entities

var Topics = []string{

	UserCreated{}.Name(),
	UserUpdated{}.Name(),
	UserReaded{}.Name(),
	UserDeleted{}.Name(),
	History{}.Name(),
	Account{}.Name(),
}

type Event interface {
	Name() string
}

type EventHandler interface {
	Handle(topic string, evenBytes []byte)
}

type EventProducer interface {
	Produce(event Event) error
}

type UserCreated struct {
	UserID   string `db:"user_id" json:"user_id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Address  string `db:"address" json:"address"`
}

type UserUpdated struct {
	UserID   string `db:"user_id" json:"user_id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Address  string `db:"address" json:"address"`
}

type UserReaded struct {
	UserID string `db:"user_id" json:"user_id"`
}

type UserDeleted struct {
	UserID string `db:"user_id" json:"user_id"`
}

type Catalog struct {
	ID        string `db:"id" json:"id"`
	Type      string `db:"type" json:"type"`
	Setup     string `db:"setup" json:"setup"`
	Punchline string `db:"punchline" json:"punchline"`
}

type History struct {
	UserID    string `db:"user_id" json:"user_id"`
	ID        string `db:"id" json:"id"`
	Type      string `db:"type" json:"type"`
	Setup     string `db:"setup" json:"setup"`
	Punchline string `db:"punchline" json:"punchline"`
}

type Account struct {
	UserID   string `db:"user_id" json:"user_id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Address  string `db:"address" json:"address"`
}

// Produce / Consume
func (UserCreated) Name() string {
	return "userCreated"
}
func (UserUpdated) Name() string {
	return "userUpdated"
}
func (UserReaded) Name() string {
	return "userReaded"
}
func (UserDeleted) Name() string {
	return "userDeleted"
}
func (History) Name() string {
	return "userHistory"
}
func (Account) Name() string {
	return "userAccount"
}



