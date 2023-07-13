package in_functions_test

import (
	inFunctions "github.com/go-libraries/govaluate/eval/functions/mathx"
	"github.com/go-libraries/govaluate/eval/functions/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

var _ = Describe("sin functions", func() {
	Describe("sin", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Sin(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Sin(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Sin(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Sin(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Sin("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Sin()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Sin(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("asin", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Asin(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Asin(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Asin(1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Asin(1)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Asin("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Asin()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Asin(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("sinh", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Sinh(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Sinh(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Sinh(2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Sinh(2)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Sinh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Sinh()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Sinh(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("asinh", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Asinh(3.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Asinh(3.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Asinh(6)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Asinh(6)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Asinh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Asinh()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Asinh(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
