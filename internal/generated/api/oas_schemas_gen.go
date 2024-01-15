// Code generated by ogen, DO NOT EDIT.

package api

import (
	"io"
	"time"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

type Bearer struct {
	Token string
}

// GetToken returns the value of Token.
func (s *Bearer) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *Bearer) SetToken(val string) {
	s.Token = val
}

type DeleteAccessBadRequest Error

func (*DeleteAccessBadRequest) deleteAccessRes() {}

type DeleteAccessInternalServerError Error

func (*DeleteAccessInternalServerError) deleteAccessRes() {}

// DeleteAccessOK is response for DeleteAccess operation.
type DeleteAccessOK struct{}

func (*DeleteAccessOK) deleteAccessRes() {}

type DeleteAccessUnauthorized Error

func (*DeleteAccessUnauthorized) deleteAccessRes() {}

type DeleteMetricBadRequest Error

func (*DeleteMetricBadRequest) deleteMetricRes() {}

type DeleteMetricForbidden Error

func (*DeleteMetricForbidden) deleteMetricRes() {}

type DeleteMetricInternalServerError Error

func (*DeleteMetricInternalServerError) deleteMetricRes() {}

// DeleteMetricOK is response for DeleteMetric operation.
type DeleteMetricOK struct{}

func (*DeleteMetricOK) deleteMetricRes() {}

type DeleteMetricUnauthorized Error

func (*DeleteMetricUnauthorized) deleteMetricRes() {}

type DeleteUserBadRequest Error

func (*DeleteUserBadRequest) deleteUserRes() {}

type DeleteUserForbidden Error

func (*DeleteUserForbidden) deleteUserRes() {}

type DeleteUserInternalServerError Error

func (*DeleteUserInternalServerError) deleteUserRes() {}

// DeleteUserOK is response for DeleteUser operation.
type DeleteUserOK struct{}

func (*DeleteUserOK) deleteUserRes() {}

type DeleteUserUnauthorized Error

func (*DeleteUserUnauthorized) deleteUserRes() {}

// Ref: #/components/schemas/Error
type Error struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

type GetAccessOwnersBadRequest Error

func (*GetAccessOwnersBadRequest) getAccessOwnersRes() {}

type GetAccessOwnersInternalServerError Error

func (*GetAccessOwnersInternalServerError) getAccessOwnersRes() {}

type GetAccessOwnersOKApplicationJSON []Person

func (*GetAccessOwnersOKApplicationJSON) getAccessOwnersRes() {}

type GetAccessOwnersUnauthorized Error

func (*GetAccessOwnersUnauthorized) getAccessOwnersRes() {}

type GetAccessUsersBadRequest Error

func (*GetAccessUsersBadRequest) getAccessUsersRes() {}

type GetAccessUsersInternalServerError Error

func (*GetAccessUsersInternalServerError) getAccessUsersRes() {}

type GetAccessUsersOKApplicationJSON []Person

func (*GetAccessUsersOKApplicationJSON) getAccessUsersRes() {}

type GetAccessUsersUnauthorized Error

func (*GetAccessUsersUnauthorized) getAccessUsersRes() {}

type GetAvatarBadRequest Error

func (*GetAvatarBadRequest) getAvatarRes() {}

type GetAvatarForbidden Error

func (*GetAvatarForbidden) getAvatarRes() {}

type GetAvatarInternalServerError Error

func (*GetAvatarInternalServerError) getAvatarRes() {}

type GetAvatarOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s GetAvatarOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*GetAvatarOK) getAvatarRes() {}

type GetAvatarUnauthorized Error

func (*GetAvatarUnauthorized) getAvatarRes() {}

type GetMeBadRequest Error

func (*GetMeBadRequest) getMeRes() {}

type GetMeForbidden Error

func (*GetMeForbidden) getMeRes() {}

type GetMeInternalServerError Error

func (*GetMeInternalServerError) getMeRes() {}

type GetMeUnauthorized Error

func (*GetMeUnauthorized) getMeRes() {}

type GetMetricBadRequest Error

func (*GetMetricBadRequest) getMetricRes() {}

type GetMetricForbidden Error

func (*GetMetricForbidden) getMetricRes() {}

type GetMetricInternalServerError Error

func (*GetMetricInternalServerError) getMetricRes() {}

type GetMetricOKApplicationJSON []Metric

func (*GetMetricOKApplicationJSON) getMetricRes() {}

type GetMetricUnauthorized Error

func (*GetMetricUnauthorized) getMetricRes() {}

// Ref: #/components/schemas/Metric
type Metric struct {
	UUID       uuid.UUID  `json:"uuid"`
	User       uuid.UUID  `json:"user"`
	MetricType MetricType `json:"metricType"`
	Value      float64    `json:"value"`
	Date       time.Time  `json:"date"`
}

// GetUUID returns the value of UUID.
func (s *Metric) GetUUID() uuid.UUID {
	return s.UUID
}

// GetUser returns the value of User.
func (s *Metric) GetUser() uuid.UUID {
	return s.User
}

// GetMetricType returns the value of MetricType.
func (s *Metric) GetMetricType() MetricType {
	return s.MetricType
}

// GetValue returns the value of Value.
func (s *Metric) GetValue() float64 {
	return s.Value
}

// GetDate returns the value of Date.
func (s *Metric) GetDate() time.Time {
	return s.Date
}

// SetUUID sets the value of UUID.
func (s *Metric) SetUUID(val uuid.UUID) {
	s.UUID = val
}

// SetUser sets the value of User.
func (s *Metric) SetUser(val uuid.UUID) {
	s.User = val
}

// SetMetricType sets the value of MetricType.
func (s *Metric) SetMetricType(val MetricType) {
	s.MetricType = val
}

// SetValue sets the value of Value.
func (s *Metric) SetValue(val float64) {
	s.Value = val
}

// SetDate sets the value of Date.
func (s *Metric) SetDate(val time.Time) {
	s.Date = val
}

func (*Metric) postMetricRes() {}
func (*Metric) putMetricRes()  {}

// Ref: #/components/schemas/MetricType
type MetricType string

const (
	MetricTypeCalories      MetricType = "calories"
	MetricTypeProteins      MetricType = "proteins"
	MetricTypeFats          MetricType = "fats"
	MetricTypeCarbohydrates MetricType = "carbohydrates"
	MetricTypeWeight        MetricType = "weight"
)

// AllValues returns all MetricType values.
func (MetricType) AllValues() []MetricType {
	return []MetricType{
		MetricTypeCalories,
		MetricTypeProteins,
		MetricTypeFats,
		MetricTypeCarbohydrates,
		MetricTypeWeight,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s MetricType) MarshalText() ([]byte, error) {
	switch s {
	case MetricTypeCalories:
		return []byte(s), nil
	case MetricTypeProteins:
		return []byte(s), nil
	case MetricTypeFats:
		return []byte(s), nil
	case MetricTypeCarbohydrates:
		return []byte(s), nil
	case MetricTypeWeight:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *MetricType) UnmarshalText(data []byte) error {
	switch MetricType(data) {
	case MetricTypeCalories:
		*s = MetricTypeCalories
		return nil
	case MetricTypeProteins:
		*s = MetricTypeProteins
		return nil
	case MetricTypeFats:
		*s = MetricTypeFats
		return nil
	case MetricTypeCarbohydrates:
		*s = MetricTypeCarbohydrates
		return nil
	case MetricTypeWeight:
		*s = MetricTypeWeight
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// NewOptDate returns new OptDate with value set to v.
func NewOptDate(v time.Time) OptDate {
	return OptDate{
		Value: v,
		Set:   true,
	}
}

// OptDate is optional time.Time.
type OptDate struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDate was set.
func (o OptDate) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDate) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDate) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDate) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDate) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFloat64 returns new OptFloat64 with value set to v.
func NewOptFloat64(v float64) OptFloat64 {
	return OptFloat64{
		Value: v,
		Set:   true,
	}
}

// OptFloat64 is optional float64.
type OptFloat64 struct {
	Value float64
	Set   bool
}

// IsSet returns true if OptFloat64 was set.
func (o OptFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFloat64) Or(d float64) float64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptMetricType returns new OptMetricType with value set to v.
func NewOptMetricType(v MetricType) OptMetricType {
	return OptMetricType{
		Value: v,
		Set:   true,
	}
}

// OptMetricType is optional MetricType.
type OptMetricType struct {
	Value MetricType
	Set   bool
}

// IsSet returns true if OptMetricType was set.
func (o OptMetricType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptMetricType) Reset() {
	var v MetricType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptMetricType) SetTo(v MetricType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptMetricType) Get() (v MetricType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptMetricType) Or(d MetricType) MetricType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Person
type Person struct {
	UUID  uuid.UUID `json:"uuid"`
	Email string    `json:"email"`
	Login string    `json:"login"`
}

// GetUUID returns the value of UUID.
func (s *Person) GetUUID() uuid.UUID {
	return s.UUID
}

// GetEmail returns the value of Email.
func (s *Person) GetEmail() string {
	return s.Email
}

// GetLogin returns the value of Login.
func (s *Person) GetLogin() string {
	return s.Login
}

// SetUUID sets the value of UUID.
func (s *Person) SetUUID(val uuid.UUID) {
	s.UUID = val
}

// SetEmail sets the value of Email.
func (s *Person) SetEmail(val string) {
	s.Email = val
}

// SetLogin sets the value of Login.
func (s *Person) SetLogin(val string) {
	s.Login = val
}

func (*Person) getMeRes()    {}
func (*Person) postUserRes() {}
func (*Person) putUserRes()  {}

// Ref: #/components/schemas/PostAccess
type PostAccess struct {
	Email OptString `json:"email"`
	Login OptString `json:"login"`
}

// GetEmail returns the value of Email.
func (s *PostAccess) GetEmail() OptString {
	return s.Email
}

// GetLogin returns the value of Login.
func (s *PostAccess) GetLogin() OptString {
	return s.Login
}

// SetEmail sets the value of Email.
func (s *PostAccess) SetEmail(val OptString) {
	s.Email = val
}

// SetLogin sets the value of Login.
func (s *PostAccess) SetLogin(val OptString) {
	s.Login = val
}

type PostAccessBadRequest Error

func (*PostAccessBadRequest) postAccessRes() {}

type PostAccessInternalServerError Error

func (*PostAccessInternalServerError) postAccessRes() {}

// PostAccessOK is response for PostAccess operation.
type PostAccessOK struct{}

func (*PostAccessOK) postAccessRes() {}

type PostAccessUnauthorized Error

func (*PostAccessUnauthorized) postAccessRes() {}

type PostLoginBadRequest Error

func (*PostLoginBadRequest) postLoginRes() {}

type PostLoginForbidden Error

func (*PostLoginForbidden) postLoginRes() {}

type PostLoginInternalServerError Error

func (*PostLoginInternalServerError) postLoginRes() {}

// Ref: #/components/schemas/PostMetric
type PostMetric struct {
	MetricType MetricType `json:"metricType"`
	Value      float64    `json:"value"`
	Date       time.Time  `json:"date"`
}

// GetMetricType returns the value of MetricType.
func (s *PostMetric) GetMetricType() MetricType {
	return s.MetricType
}

// GetValue returns the value of Value.
func (s *PostMetric) GetValue() float64 {
	return s.Value
}

// GetDate returns the value of Date.
func (s *PostMetric) GetDate() time.Time {
	return s.Date
}

// SetMetricType sets the value of MetricType.
func (s *PostMetric) SetMetricType(val MetricType) {
	s.MetricType = val
}

// SetValue sets the value of Value.
func (s *PostMetric) SetValue(val float64) {
	s.Value = val
}

// SetDate sets the value of Date.
func (s *PostMetric) SetDate(val time.Time) {
	s.Date = val
}

type PostMetricBadRequest Error

func (*PostMetricBadRequest) postMetricRes() {}

type PostMetricForbidden Error

func (*PostMetricForbidden) postMetricRes() {}

type PostMetricInternalServerError Error

func (*PostMetricInternalServerError) postMetricRes() {}

type PostMetricUnauthorized Error

func (*PostMetricUnauthorized) postMetricRes() {}

// Ref: #/components/schemas/PostUser
type PostUser struct {
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

// GetEmail returns the value of Email.
func (s *PostUser) GetEmail() string {
	return s.Email
}

// GetLogin returns the value of Login.
func (s *PostUser) GetLogin() string {
	return s.Login
}

// GetPassword returns the value of Password.
func (s *PostUser) GetPassword() string {
	return s.Password
}

// SetEmail sets the value of Email.
func (s *PostUser) SetEmail(val string) {
	s.Email = val
}

// SetLogin sets the value of Login.
func (s *PostUser) SetLogin(val string) {
	s.Login = val
}

// SetPassword sets the value of Password.
func (s *PostUser) SetPassword(val string) {
	s.Password = val
}

type PostUserBadRequest Error

func (*PostUserBadRequest) postUserRes() {}

type PostUserInternalServerError Error

func (*PostUserInternalServerError) postUserRes() {}

type PutAvatarBadRequest Error

func (*PutAvatarBadRequest) putAvatarRes() {}

type PutAvatarForbidden Error

func (*PutAvatarForbidden) putAvatarRes() {}

type PutAvatarInternalServerError Error

func (*PutAvatarInternalServerError) putAvatarRes() {}

// PutAvatarOK is response for PutAvatar operation.
type PutAvatarOK struct{}

func (*PutAvatarOK) putAvatarRes() {}

type PutAvatarReq struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s PutAvatarReq) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

type PutAvatarUnauthorized Error

func (*PutAvatarUnauthorized) putAvatarRes() {}

// Ref: #/components/schemas/PutMetric
type PutMetric struct {
	UUID       uuid.UUID     `json:"uuid"`
	MetricType OptMetricType `json:"metricType"`
	Value      OptFloat64    `json:"value"`
	Date       OptDate       `json:"date"`
}

// GetUUID returns the value of UUID.
func (s *PutMetric) GetUUID() uuid.UUID {
	return s.UUID
}

// GetMetricType returns the value of MetricType.
func (s *PutMetric) GetMetricType() OptMetricType {
	return s.MetricType
}

// GetValue returns the value of Value.
func (s *PutMetric) GetValue() OptFloat64 {
	return s.Value
}

// GetDate returns the value of Date.
func (s *PutMetric) GetDate() OptDate {
	return s.Date
}

// SetUUID sets the value of UUID.
func (s *PutMetric) SetUUID(val uuid.UUID) {
	s.UUID = val
}

// SetMetricType sets the value of MetricType.
func (s *PutMetric) SetMetricType(val OptMetricType) {
	s.MetricType = val
}

// SetValue sets the value of Value.
func (s *PutMetric) SetValue(val OptFloat64) {
	s.Value = val
}

// SetDate sets the value of Date.
func (s *PutMetric) SetDate(val OptDate) {
	s.Date = val
}

type PutMetricBadRequest Error

func (*PutMetricBadRequest) putMetricRes() {}

type PutMetricForbidden Error

func (*PutMetricForbidden) putMetricRes() {}

type PutMetricInternalServerError Error

func (*PutMetricInternalServerError) putMetricRes() {}

type PutMetricUnauthorized Error

func (*PutMetricUnauthorized) putMetricRes() {}

// Ref: #/components/schemas/PutUser
type PutUser struct {
	UUID        uuid.UUID `json:"uuid"`
	Email       OptString `json:"email"`
	Login       OptString `json:"login"`
	OldPassword OptString `json:"oldPassword"`
	NewPassword OptString `json:"newPassword"`
}

// GetUUID returns the value of UUID.
func (s *PutUser) GetUUID() uuid.UUID {
	return s.UUID
}

// GetEmail returns the value of Email.
func (s *PutUser) GetEmail() OptString {
	return s.Email
}

// GetLogin returns the value of Login.
func (s *PutUser) GetLogin() OptString {
	return s.Login
}

// GetOldPassword returns the value of OldPassword.
func (s *PutUser) GetOldPassword() OptString {
	return s.OldPassword
}

// GetNewPassword returns the value of NewPassword.
func (s *PutUser) GetNewPassword() OptString {
	return s.NewPassword
}

// SetUUID sets the value of UUID.
func (s *PutUser) SetUUID(val uuid.UUID) {
	s.UUID = val
}

// SetEmail sets the value of Email.
func (s *PutUser) SetEmail(val OptString) {
	s.Email = val
}

// SetLogin sets the value of Login.
func (s *PutUser) SetLogin(val OptString) {
	s.Login = val
}

// SetOldPassword sets the value of OldPassword.
func (s *PutUser) SetOldPassword(val OptString) {
	s.OldPassword = val
}

// SetNewPassword sets the value of NewPassword.
func (s *PutUser) SetNewPassword(val OptString) {
	s.NewPassword = val
}

type PutUserBadRequest Error

func (*PutUserBadRequest) putUserRes() {}

type PutUserForbidden Error

func (*PutUserForbidden) putUserRes() {}

type PutUserInternalServerError Error

func (*PutUserInternalServerError) putUserRes() {}

type PutUserUnauthorized Error

func (*PutUserUnauthorized) putUserRes() {}

// Ref: #/components/schemas/TokenInfo
type TokenInfo struct {
	Token OptString `json:"token"`
}

// GetToken returns the value of Token.
func (s *TokenInfo) GetToken() OptString {
	return s.Token
}

// SetToken sets the value of Token.
func (s *TokenInfo) SetToken(val OptString) {
	s.Token = val
}

func (*TokenInfo) postLoginRes() {}