package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/phonepeproj/proj/config"
	"github.com/phonepeproj/proj/enums"
	"github.com/phonepeproj/proj/serviceimpl/document"
	"github.com/phonepeproj/proj/serviceimpl/user"
)

func main() {
	singletonConfigInstance := config.GetSingleConfigInstance()
	spew.Dump("singletonConfigInstance", singletonConfigInstance)
	userSvcClient := user.NewUserSvcImpl(singletonConfigInstance)
	docSvcClient := document.NewDocSvcImpl(singletonConfigInstance, userSvcClient)
	u1, err := userSvcClient.Signup("User1", "abc", enums.UserRole_L1)
	if err != nil {
		fmt.Println("error creating user", err)
		return
	}
	spew.Dump("user1 : ", u1)

	loggedInUser, err := userSvcClient.Login(u1.GetId(), "User1", "abc")
	if err != nil {
		fmt.Println("error in logging user", err)
		return
	}
	spew.Dump("loggedInUser : ", loggedInUser)

	u2, err := userSvcClient.Signup("User2", "def", enums.UserRole_L2)
	if err != nil {
		fmt.Println("error creating user", err)
		return
	}
	spew.Dump("user2 : ", u2)

	loggedInUser2, err := userSvcClient.Login(u2.GetId(), "User2", "def")
	if err != nil {
		fmt.Println("error in logging user", err)
		return
	}
	spew.Dump("loggedInUser2 : ", loggedInUser2)

	doc1, err := docSvcClient.Create(u1.GetId(), "title1", "content1", enums.PublishedMode_PUBLISHED)
	if err != nil {
		fmt.Println("error creating doc", err)
		return
	}
	spew.Dump("doc1 : ", doc1)

	updatedDoc1, err := docSvcClient.Update(u1.GetId(), doc1.GetId(), "updated content")
	if err != nil {
		fmt.Println("error creating doc", err)
		return
	}
	spew.Dump("updatedDoc1 : ", updatedDoc1)

	//updatedDoc2, err := docSvcClient.Update(u2.GetId(), doc1.GetId(), "updated content")
	//if err != nil {
	//	fmt.Println("error updating doc", err)
	//	return
	//}
	//spew.Dump("updatedDoc2 : ", updatedDoc2)

	getDoc1, err := docSvcClient.GetDocById(doc1.GetId())
	if err != nil {
		fmt.Println("error reading doc", err)
		return
	}
	spew.Dump("getDoc1 : ", getDoc1)

	//isFeatEnabled, err := userSvcClient.CheckUserEligibilityForFeature(u1.GetId(), "FEATURE_1")
	//if err != nil {
	//	panic(err)
	//}
	//spew.Dump("isFeatEnabled : ", isFeatEnabled)
	//
	//err = docSvcClient.ShareDoc([]*document.DocShareSubject{{DocId: "docId1", UserId: u1.GetId()}, {DocId: "docId2", UserId: u2.GetId()}})
	//if err != nil {
	//	panic(err)
	//}
}
