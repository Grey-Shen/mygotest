package fop_test

import (
	"math/big"

	. "gitlab.qiniu.io/pingan/libqiniu/fop"
	"gitlab.qiniu.io/pingan/libqiniu/op"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type StringAlias string
type Int64Alias int64
type Uint32PtrAlias *uint32

type testFopCommand struct {
	Main             string         `fop:"foptest"`
	Bool             bool           `fop:"bool"`
	Float32          float32        `fop:"flt32"`
	Float64          float64        `fop:"flt64"`
	Int              int            `fop:"int"`
	Int16            int16          `fop:"int16"`
	Uint64           uint64         `fop:"int64"`
	_                NoValue        `fop:"novalue"`
	NullValue        interface{}    `fop:"nullvalue"`
	String           string         `fop:"str"`
	Stringer         *big.Float     `fop:"stringer"`
	Entry            op.Entry       `fop:"entry"`
	AnotherString    StringAlias    `fop:"another_str"`
	AnotherInt64     Int64Alias     `fop:"another_int64"`
	AnotherUint32Ptr Uint32PtrAlias `fop:"another_int32_ptr"`
}

var _ = Describe("Fop Struct", func() {
	Context("ConvertStructToFopCommand", func() {
		It("should convert the struct to fop command", func() {
			u32 := uint32(1024)
			t := testFopCommand{
				Main:             "123",
				Bool:             true,
				Float64:          3.14159265354,
				Int:              8,
				Int16:            -3,
				Uint64:           4611686018427387904,
				String:           "hello world",
				Stringer:         big.NewFloat(2.718281828459045),
				Entry:            op.NewEntry("testbucket", "testkey"),
				AnotherString:    StringAlias("this is golang reflect"),
				AnotherInt64:     Int64Alias(-65536),
				AnotherUint32Ptr: &u32,
			}

			cmd, err := NewFopCommand(t)
			Expect(err).NotTo(HaveOccurred())
			Expect(cmd.String()).To(Equal("foptest/123/bool/1/flt64/3.14159265354/int/8/int16/-3/int64/4611686018427387904/novalue/str/hello world/stringer/2.718281828/entry/dGVzdGJ1Y2tldDp0ZXN0a2V5/another_str/this is golang reflect/another_int64/-65536/another_int32_ptr/1024"))

			t2 := testFopCommand{
				Main:    "456",
				Float64: 3.14159265354,
				String:  "hello world",
			}

			cmd2, err := NewFopCommand(&t2)
			Expect(err).NotTo(HaveOccurred())
			Expect(cmd2.String()).To(Equal("foptest/456/flt64/3.14159265354/novalue/str/hello world"))
		})
	})
})
