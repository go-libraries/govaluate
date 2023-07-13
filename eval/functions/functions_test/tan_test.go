package in_functions_test

import (
	inFunctions "github.com/go-libraries/govaluate/eval/functions/mathx"
	"github.com/go-libraries/govaluate/eval/functions/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

var _ = Describe("tan functions", func() {
	Describe("tan", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Tan(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Tan(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Tan(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Tan(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Tan("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Tan()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Tan(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("atan", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Atan(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Atan(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Atan(1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Atan(1)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Atan("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Atan()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Atan(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("atan2", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Atan2(0.1363122, 0.7362629)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Atan2(0.1363122, 0.7362629)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Atan2(3, 1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Atan2(3, 1)))
		})

		It("should return an param wrong type error (first param)", func() {
			_, err := inFunctions.Atan2("invalid param value", 0.73)
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should return an param wrong type error (second param)", func() {
			_, err := inFunctions.Atan2(0.13, "invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Atan2()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too few arguments 2)", func() {
			c, err := inFunctions.Atan2(1)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Atan2(1, 2, 3)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("tanh", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Tanh(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Tanh(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Tanh(2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Tanh(2)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Tanh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Tanh()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Tanh(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("atanh", func() {
		It("should get correct value with a float param", func() {
			c, err := inFunctions.Atanh(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Atanh(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := inFunctions.Atanh(0.12345)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Atanh(0.12345)))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Atanh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Atanh()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Atanh(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
