package impl_utils

import (
	"account-service/utils"
	"math/rand"
	"strconv"
	"time"
)

type registerUtils struct {
	numChar           int
	timeExistSaveInfo int
}

func (r *registerUtils) CreateCode() (code string) {
	createCode := ""
	for i := 0; i < r.numChar; i++ {
		charNumRandom := strconv.Itoa(rand.Intn(10))
		createCode += charNumRandom
	}
	return createCode
}

func (r *registerUtils) CreateTimeExist() (startAt time.Time, finishAt time.Time) {
	timeStartAt := time.Now()
	timeFinishAt := timeStartAt.Add(time.Second * time.Duration(r.timeExistSaveInfo))

	return timeStartAt, timeFinishAt
}

func RegisterUtilsInit() utils.RegisterUtils {
	return &registerUtils{
		numChar:           6,
		timeExistSaveInfo: 30,
	}
}
