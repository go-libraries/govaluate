package in_functions_test

import (
	inFunctions "github.com/go-libraries/govaluate/eval/functions/mathx"
	"github.com/go-libraries/govaluate/eval/functions/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("max, min functions", func() {
	Describe("max", func() {
		It("should resolve a dual max", func() {
			c, err := inFunctions.Max(1, 2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(2.0))
		})

		It("should resolve a the max value for 5 values", func() {
			c, err := inFunctions.Max(4, 5, 3, 1, 2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(5.0))
		})

		It("should fail with a param type error (string supplied)", func() {
			c, err := inFunctions.Max(4, 5, 3, 1, 2, "non number")
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should fail with a param type error (boolean supplied)", func() {
			c, err := inFunctions.Max(4, 5, 3, 1, 2, false)
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should fail with a param count error", func() {
			c, err := inFunctions.Max()
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
		})
	})

	Describe("min", func() {
		It("should resolve a dual min", func() {
			c, err := inFunctions.Min(2, 1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.0))
		})

		It("should resolve a the min value for 5 values", func() {
			c, err := inFunctions.Min(4, 5, 3, 1, 2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.0))
		})

		It("should fail with a param type error (string supplied)", func() {
			c, err := inFunctions.Min(4, 5, 3, 1, 2, "non number")
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should fail with a param type error (boolean supplied)", func() {
			c, err := inFunctions.Min(4, 5, 3, 1, 2, false)
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamType(err)).To(BeTrue())
		})

		It("should fail with a param count error", func() {
			c, err := inFunctions.Min()
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
		})
	})
})
