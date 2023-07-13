package in_functions_test

import (
	inFunctions "github.com/go-libraries/govaluate/eval/functions/mathx"
	"github.com/go-libraries/govaluate/eval/functions/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

var _ = Describe("cos functions", func() {
	Describe("cos", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Cos(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Cos(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Cos(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Cos(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Cos("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Cos()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Cos(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("acos", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Acos(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Acos(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Acos(1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Acos(1)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Acos("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Acos()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Acos(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("cosh", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Cosh(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Cosh(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Cosh(2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Cosh(2)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Cosh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Cosh()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Cosh(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("acosh", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Acosh(3.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Acosh(3.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Acosh(6)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Acosh(6)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Acosh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Acosh()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Acosh(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
