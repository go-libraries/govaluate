package in_functions_test

import (
	inFunctions "github.com/go-libraries/govaluate/eval/functions"
	"github.com/go-libraries/govaluate/eval/functions/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("is_defined function", func() {
	It("should return that a nil value is not defined", func() {
		c, err := inFunctions.IsDefined(nil)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should return that a non empty string is defined", func() {
		c, err := inFunctions.IsDefined("this is a non empty string")
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return that a non empty string is defined", func() {
		c, err := inFunctions.IsDefined("")
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should return that a zero valued integer is not defined", func() {
		c, err := inFunctions.IsDefined(0)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should return that a non zero valued integer is defined (positive)", func() {
		c, err := inFunctions.IsDefined(1)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return that a non zero valued integer is defined (negative)", func() {
		c, err := inFunctions.IsDefined(-1)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return that a zero valued float is not defined", func() {
		c, err := inFunctions.IsDefined(0.0)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should return that a non zero valued float is defined (positive)", func() {
		c, err := inFunctions.IsDefined(0.000001)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return that a non zero valued float is defined (negative)", func() {
		c, err := inFunctions.IsDefined(-0.000001)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return an param wrong type error", func() {
		_, err := inFunctions.IsDefined(map[string]interface{}{
			"str": "value",
		})
		Expect(err).NotTo(BeNil())
		Expect(utils.IsWrongParamType(err)).To(BeTrue())
	})

	It("should return a param count error (too few arguments)", func() {
		_, err := inFunctions.IsDefined()
		Expect(err).NotTo(BeNil())
		Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
	})

	It("should return a param count error (too many arguments)", func() {
		_, err := inFunctions.IsDefined(1, 2)
		Expect(err).NotTo(BeNil())
		Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
	})
})
