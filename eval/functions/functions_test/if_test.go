package in_functions_test

import (
	inFunctions "github.com/go-libraries/govaluate/eval/functions"
	"github.com/go-libraries/govaluate/eval/functions/mathx"
	"github.com/go-libraries/govaluate/eval/functions/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("if functions", func() {
	It("should evaluate a boolean operation (true)", func() {
		c, err := inFunctions.If(true, 1, 0)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(1))
	})

	It("should evaluate a boolean operation (false)", func() {
		c, err := inFunctions.If(false, 1, 0)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(0))
	})

	It("should get correct value with a float (non-zero)", func() {
		c, err := inFunctions.If(0.000001, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())

		c, err = inFunctions.If(-0.000001, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should get correct value with a float (zero)", func() {
		c, err := inFunctions.If(0.0, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())

		c, err = inFunctions.If(-0.0, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should get correct value with an integer (non-zero)", func() {
		c, err := inFunctions.If(1, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())

		c, err = inFunctions.If(-1, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should get correct value with an integer (zero)", func() {
		c, err := inFunctions.If(0, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())

		c, err = inFunctions.If(-0, true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should get correct value with a string (non-empty)", func() {
		c, err := inFunctions.If("non empty string", true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should get correct value with a string (empty)", func() {
		c, err := inFunctions.If("", true, false)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should get an wrong param count error (too few arguments)", func() {
		c, err := inFunctions.If()
		Expect(c).To(BeNil())
		Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
		Expect(err).NotTo(BeNil())

		c, err = inFunctions.If(0)
		Expect(c).To(BeNil())
		Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
		Expect(err).NotTo(BeNil())

		c, err = inFunctions.If(0, 0)
		Expect(c).To(BeNil())
		Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
		Expect(err).NotTo(BeNil())
	})

	It("should get an wrong param count error (too many arguments)", func() {
		c, err := mathx.Exp(true, true, true, true)
		Expect(c).To(BeNil())
		Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
		Expect(err).NotTo(BeNil())
	})
})
