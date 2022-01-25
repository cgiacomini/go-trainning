package math_test

import (
        
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Math", func() {
        Describe("Add", func() {
		It("adds two numbers", func() {
			sum := Add(2, 3)
			Expect(sum).To(Equal(5))
		})
        })
})
