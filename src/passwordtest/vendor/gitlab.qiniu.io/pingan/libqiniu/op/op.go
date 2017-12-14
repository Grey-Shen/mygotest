package op

import (
	"encoding/base64"
	"strconv"
)

type Stat struct {
	Entry Entry
}

func (stat Stat) Encode() string {
	return "/stat/" + stat.Entry.Encode()
}

type Chgm struct {
	Entry Entry
	Mime  string
}

func (chgm Chgm) Encode() string {
	return "/chgm/" + chgm.Entry.Encode() + "/mime/" + base64.URLEncoding.EncodeToString([]byte(chgm.Mime))
}

type Chtype struct {
	Entry Entry
	Type  Type
}

type Type uint8

const (
	TypeNormal     = Type(0)
	TypeInfrequent = Type(1)
)

func (t Type) Encode() string {
	return strconv.FormatUint(uint64(t), 10)
}

func (chtype Chtype) Encode() string {
	return "/chtype/" + chtype.Entry.Encode() + "/type/" + chtype.Type.Encode()
}

type Move struct {
	Src   Entry
	Dest  Entry
	Force bool
}

func (move Move) Encode() string {
	return "/move/" + move.Src.Encode() + "/" + move.Dest.Encode() + "/force/" + strconv.FormatBool(move.Force)
}

type Copy struct {
	Src   Entry
	Dest  Entry
	Force bool
}

func (cpy Copy) Encode() string {
	return "/copy/" + cpy.Src.Encode() + "/" + cpy.Dest.Encode() + "/force/" + strconv.FormatBool(cpy.Force)
}

type Delete struct {
	Entry Entry
}

func (del Delete) Encode() string {
	return "/delete/" + del.Entry.Encode()
}

type DeleteAfterDays struct {
	Entry     Entry
	AfterDays uint64
}

func (policy DeleteAfterDays) Encode() string {
	return "/deleteAfterDays/" + policy.Entry.Encode() + "/" + strconv.FormatUint(policy.AfterDays, 10)
}
