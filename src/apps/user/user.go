package user

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"libs/abstract"
	"settings"
	"utils"

	"github.com/mongodb/mongo-go-driver/bson"
	reflections "gopkg.in/oleiade/reflections.v1"
)

const collectionName = "user"

type ExportUserVersion struct {
	*User
	Password string `json:"-"`
}

type User struct {
	abstract.AbstractModel `bson:",inline"`
	FirstName              string `json:"firstName" bson:"firstName"`
	LastName               string `json:"lastName" bson:"lastName"`
	FullName               string `json:"fullName" bson:"fullName"`
	UserName               string `json:"username" bson:"username"`
	Password               string `json:"password,omitempty"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User

	safeUser := struct {
		Password string `json:"password,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	return json.Marshal(safeUser)
}

func (u *User) Bind(r *http.Request) error {
	if err := u.Clean(); err != nil {
		return err
	}
	return nil
}

func (u *User) Clean() error {
	if u.FirstName == "" || u.LastName == "" {
		return errors.New("first name or last name cannot be blank")
	}

	if u.Password == "" {
		return errors.New("password cannot be blank")
	}
	SetPassword(u, u.Password)

	if u.FullName == "" {
		u.FullName = fmt.Sprintf("%v %v", u.FirstName, u.LastName)
	}
	return nil
}

func (u *User) GetAbstractModel() *abstract.AbstractModel {
	return &u.AbstractModel
}

func (u User) GetCollectionName() string {
	return collectionName
}

func (u User) toBSON() bson.M {
	var data = bson.M{}
	fieldList, err := reflections.Fields(u)
	utils.CheckError(err)

	for _, fieldName := range fieldList {
		value, err := reflections.GetField(u, fieldName)
		utils.CheckError(err)
		data[fieldName] = value
	}
	return data
}

func SetPassword(u *User, password string) {
	hash := hmac.New(sha256.New, []byte(settings.SecretKey))
	fmt.Println(hash, password)
}
