//CRUD services.
package article

import (
	"github.com/goarne/blog/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//The stateless servicefacade for CRUD services on articles.
//It uses the Mongodb wrapper mgo for communicating with Mongdb
type Service interface {
	Create(f *Article) error
	Update(f *Article) error
	Delete(i int32) error
	Find(i int32) (*Article, error)
	FindAll() ([]Article, error)
}

type ServiceImpl struct{}

func (ServiceImpl) Create(f *Article) error {

	query := func(c *mgo.Collection) error {
		i := bson.NewObjectId().Counter()
		f.Id = i
		return c.Insert(f)
	}

	return executeOnCollection(config.COLL_ARTICLE, query)
}

func (ServiceImpl) Update(a *Article) error {

	query := func(c *mgo.Collection) error {
		colQuerier := bson.M{"id": a.Id}

		change := bson.M{"$set": bson.M{"title": a.Title, "abstract": a.Abstract, "text": a.Text, "start": a.Date, "author": a.Author}}
		return c.Update(colQuerier, change)

	}

	return executeOnCollection(config.COLL_ARTICLE, query)
}

func (ServiceImpl) Delete(id int32) error {
	query := func(c *mgo.Collection) error {
		return c.Remove(bson.M{"id": id})
	}

	return executeOnCollection(config.COLL_ARTICLE, query)
}

func (ServiceImpl) Find(id int32) (*Article, error) {
	retrievedArticle := Article{}

	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{"id": id}).One(&retrievedArticle)
	}

	err := executeOnCollection(config.COLL_ARTICLE, query)

	return &retrievedArticle, err
}

func (ServiceImpl) FindAll() ([]Article, error) {
	var articles []Article

	query := func(c *mgo.Collection) error {
		return c.Find(bson.M{}).All(&articles)
	}

	err := executeOnCollection(config.COLL_ARTICLE, query)

	return articles, err
}

//Function pointer created for testing purposes,
//aswell as a nice solution to handle the db session.
var executeOnCollection = func(collName string, q func(*mgo.Collection) error) error {
	session, err := mgo.Dial(config.DOCKER_MONGODB_URL)

	if err != nil {
		panic(err)
	}

	defer session.Close()

	c := session.DB(config.DB_NAME).C(collName)

	return q(c)
}
