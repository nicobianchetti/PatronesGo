package main

func main() {
	shirtItem := newItem("Nike Shirt")

	observerFirst := &customer{id: "email1@gmail.com"}
	observerSecond := &customer{id: "email2@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
