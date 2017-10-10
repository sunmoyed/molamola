package user

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/scrypt"
	"gopkg.in/dgrijalva/jwt-go.v3"

	"github.com/sunmoyed/molamola/server/pkg/util"
)

// XXX Because it's possible to change the state via the CLI independently of
// the server, the in-memory can get out of date. Either make it so that
// the CLI simply talks to the server (preferred). Or force all interactions
// to read and write to the file (current hack).
//
// XXX ^The order of operations is not correct. Memory is adjusted before it's
// written to disk, which is wrong. It's just a pain to do correctly (because
// it would require being able to do a deep copy).

// XXX None of these functions are thread safe.

// UserInfo represents the user. It does not contain the username as that
// is used as the key (in the map) to access this info.
type UserInfo struct {
	UserID   string `json:"userid"`
	Password string `json:"password"`

	// Random data used to generate the token.
	// Used to generate token.
	// Should be refreshed on password change.
	Salt []byte `json:"salt"`
}

// Data is the data
type Data struct {
	// Map of username to user info
	Users map[string]*UserInfo `json:"users"`

	// Map of UID to username
	UIDs map[string]string `json:"uids"`
}

type UserState struct {
	datadir string
	data    *Data
}

const (
	userFileName    string        = "users.json"
	defaultFileMode os.FileMode   = 0600
	pwSaltBytes                   = 32
	pwHashBytes                   = 64
	defaultExpTime  time.Duration = (time.Hour * 24) * 30

	userPerm string = "user"
)

var (
	defaultUserPermissions []string = []string{userPerm}
)

func NewUserState(datadir string) (*UserState, error) {
	us := &UserState{
		datadir: datadir,
	}

	if err := us.loadData(); err != nil {
		return nil, err
	}

	return us, nil
}

func (u *UserState) getUserFilePath() string {
	return path.Join(u.datadir, userFileName)
}

func (u *UserState) loadData() error {
	data, dataErr := dataFromFile(u.getUserFilePath())
	if dataErr != nil {
		return dataErr
	}
	u.data = data
	return nil
}

func dataFromFile(userFilePath string) (*Data, error) {
	data := &Data{}
	if ok, err := util.FileExists(userFilePath); err != nil {
		return nil, err
	} else if !ok {
		return &Data{
			Users: make(map[string]*UserInfo),
			UIDs:  make(map[string]string),
		}, nil
	}

	b, bErr := ioutil.ReadFile(userFilePath)
	if bErr != nil {
		return nil, bErr
	}
	if err := json.Unmarshal(b, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UserState) saveData() error {
	b, bErr := json.Marshal(u.data)
	if bErr != nil {
		return bErr
	}

	return ioutil.WriteFile(u.getUserFilePath(), b, defaultFileMode)
}

func (u *UserState) GetUsers() (map[string]*UserInfo, error) {
	if err := u.loadData(); err != nil {
		return nil, err
	}

	return u.data.Users, nil
}

// LoginUser is called to log in the user. It returns the token that the user
// is to use.
//
// The output of this function should be LOGGED and NOT RETURNED to the client.
func (u *UserState) LoginUser(username, password string) (string, error) {
	userInfo, userInfoOk := u.data.Users[username]
	if !userInfoOk {
		return "", fmt.Errorf("invalid username %s", username)
	}

	if userInfo.Password != password {
		return "", fmt.Errorf("invalid password %s", password)
	}

	hash, hashErr := userInfo.getHash()
	if hashErr != nil {
		return "", hashErr
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// "subject" of the jwt
		"sub": userInfo.UserID,

		// Expiration time in unix time
		"exp": genExpTime(),

		// Comma separated list of permissions
		"permissions": strings.Join(defaultUserPermissions, ","),
	})
	return token.SignedString(hash)
}

// ValidateUserToken is used to check that the token from the client is valid.
// This is called on every request from the client.
//
// The userid is on who's behalf this request was made. This must match the
// userid in the token. This is so someone can't make a valid request with
// someone else's token.
//
// The output of this function should be LOGGED and NOT RETURNED to the client.
func (u *UserState) ValidateUserToken(userid, tokenStr string) error {
	username, usernameOk := u.data.UIDs[userid]
	if !usernameOk {
		return fmt.Errorf("invalid userid %s", userid)
	}

	userInfo, userInfoOk := u.data.Users[username]
	if !userInfoOk {
		return fmt.Errorf("invalid username %s", username)
	}

	// The error here is purposefully delayed in checking, that's how the
	// library example uses it.
	token, tokenErr := jwt.Parse(tokenStr, userInfo.ValidateUserTokenFn)

	claims, claimsOk := token.Claims.(jwt.MapClaims)
	// The following two error checks are done in a very specific order,
	// according to the library example.
	if !claimsOk {
		return fmt.Errorf("bad claims: %s", tokenErr)
	}
	if !token.Valid {
		return fmt.Errorf("token invalid: %s", tokenErr)
	}

	if sub, ok := claims["sub"]; !ok {
		return fmt.Errorf("token missing sub")
	} else {
		if userid != sub.(string) {
			return fmt.Errorf("userid %s doesn't match sub %s", userid, sub)
		}
	}

	if expStr, ok := claims["exp"]; !ok {
		return fmt.Errorf("token missing exp")
	} else {
		exp, expErr := strconv.ParseInt(expStr.(string), 10, 64)
		if expErr != nil {
			return expErr
		}
		expTime := time.Unix(exp, 0)
		nowTime := time.Now()
		if expTime.Before(nowTime) {
			return fmt.Errorf("token expired: exp %s now %s", expTime, nowTime)
		}
	}

	if permissions, ok := claims["permissions"]; !ok {
		return fmt.Errorf("token missing permissions")
	} else {
		perms := strings.Split(permissions.(string), ",")
		foundPerm := false
		for _, p := range perms {
			if p == userPerm {
				foundPerm = true
			}
		}
		if !foundPerm {
			return fmt.Errorf("could not find perm: %s", userPerm)
		}
	}
	return nil
}

func (u *UserInfo) ValidateUserTokenFn(outerToken *jwt.Token) (interface{}, error) {
	// Since we are using the HMAC encryption on the JWT, the interface{} is a
	// []byte(). The library doesn't enforce it but we will.

	return func(token *jwt.Token) ([]byte, error) {
		// Purposefully ignore the "alg" in the token header. We shouldn't trust
		// it because we hardcode the algorithm in here anyways.

		return u.getHash()
	}(outerToken)
}

func (u *UserState) AddUser(username, password string) error {
	if err := u.loadData(); err != nil {
		return err
	}

	if _, ok := u.data.Users[username]; ok {
		return fmt.Errorf("user already exists: %s", username)
	}

	if password == "" {
		return fmt.Errorf("invalid password: must supply password")
	}

	userUUID := genUUID()

	salt, saltErr := genSalt()
	if saltErr != nil {
		return saltErr
	}

	u.data.Users[username] = &UserInfo{
		UserID:   userUUID,
		Password: password,
		Salt:     salt,
	}
	u.data.UIDs[userUUID] = username
	return u.saveData()
}

func (u *UserInfo) getHash() ([]byte, error) {
	return computeHash([]byte(u.Password), u.Salt)
}

func genExpTime() int64 {
	return time.Now().Add(defaultExpTime).Unix()
}

func genSalt() ([]byte, error) {
	salt := make([]byte, pwSaltBytes)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func computeHash(value, salt []byte) ([]byte, error) {
	return scrypt.Key(value, salt, 1<<14, 8, 1, pwHashBytes)
}

func genUUID() string {
	return uuid.NewV4().String()
}

func (u *UserState) DelUser(username string) error {
	if err := u.loadData(); err != nil {
		return err
	}

	if _, ok := u.data.Users[username]; !ok {
		return fmt.Errorf("user doesn't exists: %s", username)
	}

	uid := u.data.Users[username].UserID
	delete(u.data.Users, username)
	delete(u.data.UIDs, uid)
	return u.saveData()
}
