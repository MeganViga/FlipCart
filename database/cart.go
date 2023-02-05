package database

import "errors"

var (
	ErrCantFindProduct = errors.New("cant find the product")
	ErrCantDecodeProducts = errors.New("cant find the produtc")
	ErrUserIdIsNotvalid = errors.New("this users is not valid")
	ErrCantUpdateUser = errors.New("cant add this product to cart")
	ErrCantRemoveItemCart = errors.New("cant remove this item from cart")
	ErrCantGetItem = errors.New("was unable to get item from the cart")
	ErrCantBuyCartItem = errors.New("cant update the purchase")
)

func AddProductTocart(){

}

func RemoveCartItem(){

}

func BuyItemFromCart(){

}

func InstantBuyer(){

}