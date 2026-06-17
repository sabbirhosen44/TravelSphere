package controllers

// WishlistController serve the SSR travel wishlist page.
type WishlistController struct {
	BaseController
}

// Get renders the wishlist template populated with current entries.
func (c *WishlistController) Get() {
	c.Data["ActivePage"] = "wishlist"
	c.TplName = "wishlist.tpl"

	list, err := WishlistSvc.List()
	if err != nil {
		c.Data["WishlistItems"] = []interface{}{}
		c.Data["Error"] = "Failed to load wishlist items from local storage."
		return
	}

	c.Data["WishlistItems"] = list
}
