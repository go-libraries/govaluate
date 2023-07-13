package in_functions_test

import (
	inFunctions "github.com/go-libraries/govaluate/eval/functions/mathx"
	"github.com/go-libraries/govaluate/eval/functions/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

var _ = Describe("log functions", func() {
	Describe("log", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Log(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Log(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Log(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Log(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Log("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Log()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Log(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("log10", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Log10(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Log10(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Log10(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Log10(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Log10("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Log10()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Log10(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
