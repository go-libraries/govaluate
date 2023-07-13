package in_functions_test

import (
	inFunctions "github.com/go-libraries/govaluate/eval/functions/mathx"
	"github.com/go-libraries/govaluate/eval/functions/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("round functions", func() {
	Describe("round", func() {
		It("should round to zero", func() {
			c, err := inFunctions.Round(0.49999999)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(0.0))
		})

		It("should round to one", func() {
			c, err := inFunctions.Round(0.5)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.0))
		})

		It("should round to zero", func() {
			c, err := inFunctions.Round(-0.4999999)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(0.0))
		})

		It("should round to minus one", func() {
			c, err := inFunctions.Round(-0.5)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(-1.0))
		})

		It("should round taking a decimal point in count", func() {
			c, err := inFunctions.Round(1.2345678, 2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.23))
		})

		It("should round taking a decimal point in count (rounding up)", func() {
			c, err := inFunctions.Round(1.2385674, 2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.24))
		})

		It("should return an param wrong type error", func() {
			_, err := inFunctions.Round("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Round()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Round(1, 2, 3)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("trunc", func() {
		It("should trunc down the value", func() {
			c, err := inFunctions.Trunc(0.49999999)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(0.0))
		})

		It("should round to one", func() {
			c, err := inFunctions.Trunc(0.9)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(0.0))
		})

		It("should return a param type error", func() {
			_, err := inFunctions.Trunc("no valid value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Trunc()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Trunc(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("ceil", func() {
		It("should ceil a low float to one", func() {
			c, err := inFunctions.Ceil(0.009)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.0))
		})

		It("should ceil a high float to one", func() {
			c, err := inFunctions.Ceil(0.9)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.0))
		})

		It("should ceil a negative value", func() {
			c, err := inFunctions.Ceil(-1.9)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(-1.0))
		})

		It("should return a param type error", func() {
			_, err := inFunctions.Ceil("no valid value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Ceil()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Ceil(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("floor", func() {
		It("should floor a low float to zero", func() {
			c, err := inFunctions.Floor(0.009)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(0.0))
		})

		It("should floor a high float to one", func() {
			c, err := inFunctions.Floor(0.9)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(0.0))
		})

		It("should floor a negative value", func() {
			c, err := inFunctions.Floor(-1.4)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(-2.0))
		})

		It("should return a param type error", func() {
			_, err := inFunctions.Floor("no valid value")
			Expect(err).NotTo(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := inFunctions.Floor()
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := inFunctions.Floor(1, 2)
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
