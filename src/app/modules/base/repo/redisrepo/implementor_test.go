package redisrepo

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redismock/v9"
	"github.com/oktopriima/marvel/pkg/cache"
	"github.com/oktopriima/marvel/src/app/modules/base/models"
	"github.com/oktopriima/marvel/src/test"
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

var redisInstance, redisMock = test.RedisInstance()

type S struct {
	instance  cache.RedisInstance
	redisMock redismock.ClientMock
	ctx       context.Context
}

func Test(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&S{
	instance:  redisInstance,
	redisMock: redisMock,
	ctx:       context.Background(),
})

type RedisTestExample struct {
	models.BaseModel
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var res = RedisTestExample{}

var example = RedisTestExample{
	BaseModel: models.BaseModel{
		Id:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	Name: "john doe",
	Age:  20,
}

var key = "example"
var ttl = 60 * time.Second

func (s *S) TestStoreCache(c *C) {
	var err error
	// init repo with mocking instance
	repo := NewBaseRedisRepo(s.instance)
	c.Assert(repo, NotNil)

	// create bytes from an example object
	bytesExample, err := json.Marshal(&example)
	c.Assert(bytesExample, NotNil)
	c.Assert(err, IsNil)

	// create the expectation
	s.redisMock.ExpectSet(key, string(bytesExample), ttl).SetVal("OK")

	// store test
	err = repo.StoreCache(s.ctx, key, ttl, &example)
	c.Assert(err, IsNil)

	err = s.redisMock.ExpectationsWereMet()
	c.Assert(err, IsNil)
}

func (s *S) TestFindCache(c *C) {
	var err error
	// init repo with mocking instance
	repo := NewBaseRedisRepo(s.instance)
	c.Assert(repo, NotNil)

	// create bytes from an example object
	bytesExample, err := json.Marshal(&example)
	c.Assert(bytesExample, NotNil)
	c.Assert(err, IsNil)

	// create the expectation
	s.redisMock.ExpectGet(key).SetVal(string(bytesExample))
	err = repo.FindCache(s.ctx, &res, key)

	c.Assert(err, IsNil)
	c.Assert(res, NotNil)
	c.Assert(res.Name, Equals, example.Name)
	c.Assert(res.Age, Equals, example.Age)
	c.Assert(res.Id, Equals, example.Id)

	err = s.redisMock.ExpectationsWereMet()
	c.Assert(err, IsNil)
}

func (s *S) TestFindRawCache(c *C) {
	var err error
	// init repo with mocking instance
	repo := NewBaseRedisRepo(s.instance)
	c.Assert(repo, NotNil)

	// create bytes from an example object
	bytesExample, err := json.Marshal(&example)
	c.Assert(bytesExample, NotNil)
	c.Assert(err, IsNil)

	// create the expectation
	s.redisMock.ExpectGet(key).SetVal(string(bytesExample))

	res, err := repo.FindRawCache(s.ctx, key)
	c.Assert(err, IsNil)
	c.Assert(res, NotNil)
	c.Assert(res, DeepEquals, bytesExample)

	err = s.redisMock.ExpectationsWereMet()
	c.Assert(err, IsNil)
}

func (s *S) TestStoreRawCache(c *C) {
	var err error
	// init repo with mocking instance
	repo := NewBaseRedisRepo(s.instance)
	c.Assert(repo, NotNil)

	// create bytes from an example object
	bytesExample, err := json.Marshal(&example)
	c.Assert(bytesExample, NotNil)
	c.Assert(err, IsNil)

	// create the expectation
	s.redisMock.ExpectSet(key, string(bytesExample), ttl).SetVal("OK")

	// store test
	err = repo.StoreObjectCache(s.ctx, key, ttl, bytesExample)
	c.Assert(err, IsNil)

	err = s.redisMock.ExpectationsWereMet()
	c.Assert(err, IsNil)
}

func (s *S) TestRemoveCache(c *C) {
	var err error
	// init repo with mocking instance
	repo := NewBaseRedisRepo(s.instance)
	c.Assert(repo, NotNil)

	// create bytes from an example object
	bytesExample, err := json.Marshal(&example)
	c.Assert(bytesExample, NotNil)
	c.Assert(err, IsNil)

	// create the expectation
	s.redisMock.ExpectDel(key).SetVal(int64(1))

	// store test
	err = repo.RemoveCache(s.ctx, key)
	c.Assert(err, IsNil)

	err = s.redisMock.ExpectationsWereMet()
	c.Assert(err, IsNil)
}
