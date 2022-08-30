package jhonny_bravo_burger_kata

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Johnny Bravo's Burger", func() {

	Context("when calculating price for a ", func() {
		It("hamburger", func() {
			aBurger := MakeBurger([]Ingredient{})
			anotherBurger := MakeBurger([]Ingredient{burger})
			Expect(aBurger.price).To(Equal(5))
			Expect(anotherBurger.price).To(Equal(5))
		})
		It("cheese burger", func() {
			burger := MakeBurger([]Ingredient{cheese})
			Expect(burger.price).To(Equal(7))
		})
		It("bacon, cheese burger", func() {
			burger := MakeBurger([]Ingredient{cheese, bacon})
			Expect(burger.price).To(Equal(10))
		})
		It("bacon, egg, cheese burger", func() {
			burger := MakeBurger([]Ingredient{cheese, bacon, egg})
			Expect(burger.price).To(Equal(12))
		})
	})
})
