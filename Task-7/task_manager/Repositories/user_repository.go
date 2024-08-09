type UserRepository interface {
	Save(user models.User) (models.UserRes, error)
	Update(user models.User) (models.UserRes, error)
	FindUser(id string) (models.User, error)
	FindByUsername(username string) (models.User, bool)
	FindAll() []models.UserRes
	Delete(username string) error
	DeleteAll() error
}

type MongoUserRepository struct {
	collection *mongo.Collection
}
