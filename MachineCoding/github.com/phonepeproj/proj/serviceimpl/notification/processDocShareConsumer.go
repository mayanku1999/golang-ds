package notification

import (
	"github.com/phonepeproj/proj/dao/model"
	"log"
)

type DocShareConsumerImpl struct {
	User   *model.User
	Doc    *model.Document
	UserId string
	DocId  string
}

func (d *DocShareConsumerImpl) ProcessDocShare() {
	log.Println("called ProcessDocShare", d.UserId, d.DocId)
}
