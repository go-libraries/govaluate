package in_functions_test

import (
	inFunctions "github.com/go-libraries/govaluate/eval/functions"
	"github.com/go-libraries/govaluate/eval/functions/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ifnull function", func() {
	It("should return the first param", func() {
		c, err := inFunctions.IfNull(1, 2)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(1))
	})

	It("should return the second param", func() {
		c, err := inFunctions.IfNull(nil, 2)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(2))
	})

	It("should return the thrid param", func() {
		c, err := inFunctions.IfNull(nil, nil, 3)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(3))
	})

	It("should return nil", func() {
		c, err := inFunctions.IfNull(nil, nil, nil, nil)
		Expect(err).To(BeNil())
		Expect(c).To(BeNil())
	})

	It("should return a param count error (too few params)", func() {
		c, err := inFunctions.IfNull()
		Expect(err).NotTo(BeNil())
		Expect(utils.IsWrongParamsCount(err)).To(BeTrue())
		Expect(c).To(BeNil())
	})
})
