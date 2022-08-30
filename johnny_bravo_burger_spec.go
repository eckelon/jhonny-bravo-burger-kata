package jhonny_bravo_burger_kata

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Johnny Bravo's Burger", func() {
	Context("when calculating price for a", func() {
		It("hamburger", func() {
			burger := CalculateBurgerPrice(Burger{name: "plain"})
			Expect(burger.price).To(Equal(5))
		})
		It("cheese burger", func() {
			burger := CalculateBurgerPrice(Burger{name: "cheese burger", ingredients: []string{"cheese"}})
			Expect(burger.price).To(Equal(7))
		})
		It("bacon, cheese burger", func() {
			burger := CalculateBurgerPrice(Burger{name: "bacon cheese burger", ingredients: []string{"cheese", "bacon"}})
			Expect(burger.price).To(Equal(10))
		})
		It("bacon, egg, cheese burger", func() {
			burger := CalculateBurgerPrice(Burger{name: "bacon cheese burger", ingredients: []string{"egg", "cheese", "bacon"}})
			Expect(burger.price).To(Equal(12))
		})
	})
})
