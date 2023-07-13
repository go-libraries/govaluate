package in_functions_test

import (
	inFunctions "github.com/go-libraries/govaluate/eval/functions/mathx"
	"github.com/go-libraries/govaluate/eval/functions/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

var _ = Describe("exp functions", func() {
	Describe("exp", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Exp(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Exp(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Exp(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Exp(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Exp("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Exp()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Exp(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("exp2", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Exp2(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Exp2(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Exp2(1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Exp2(1)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Exp2("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Exp2()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Exp2(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
