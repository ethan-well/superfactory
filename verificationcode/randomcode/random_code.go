package randomcode

import (
	"fmt"
	"github.com/ItsWewin/superfactory/aerror"
	"github.com/ItsWewin/superfactory/logger"
	"github.com/ItsWewin/superfactory/verificationcode/verificationcode"
	"math/rand"
	"sync"
	"time"
)

type RandomCodeStore struct {
	lock sync.RWMutex
	db   map[string]*RandomCode
}

type RandomCode struct {
	Code string
	// 过期时间，单位秒
	ExpireTime int64
	StartTime  time.Time
}

var randomCodeStore *RandomCodeStore

func RandomCodeStoreInit() *RandomCodeStore {
	logger.Infof("RandomCodeStoreInit RandomCodeStoreInit RandomCodeStoreInit")
	if randomCodeStore != nil {
		return randomCodeStore
	}

	logger.Infof("randomCodeStore is nil")

	randomCodeStore = &RandomCodeStore{
		lock: sync.RWMutex{},
		db:   make(map[string]*RandomCode),
	}

	logger.Infof("RandomCodeStoreInit: %s", logger.ToJson(randomCodeStore))

	return randomCodeStore
}

func (c *RandomCodeStore) Generate() (verificationcode.VerificationCode, aerror.Error) {
	code := generateCode()

	c.lock.Lock()
	defer c.lock.Unlock()

	c.db[code.Code] = code

	logger.Infof("Generate: %s", logger.ToJson(randomCodeStore))

	return code, nil
}

func (c *RandomCodeStore) Delete(code string) {
	delete(c.db, code)
}

func (c *RandomCodeStore) Check(code string) bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	randomCode, ok := c.db[code]

	logger.Infof("%s", logger.ToJson(c.db))

	if !ok {
		return false
	}

	return randomCode.Valid()
}

func (c *RandomCode) Valid() bool {
	return c.StartTime.Add(time.Duration(c.ExpireTime) * time.Second).After(time.Now())
}

func (c *RandomCode) Expire() {
	c.ExpireTime = 0
}

func (c *RandomCode) GetCode() string {
	return c.Code
}

func generateCode() *RandomCode {
	rand.Seed(time.Now().UnixNano())
	min := 2000
	max := 9999
	code := fmt.Sprintf("%d", rand.Intn(max-min+1)+min)

	return &RandomCode{
		Code:       code,
		ExpireTime: 5 * 60,
		StartTime:  time.Now(),
	}
}
