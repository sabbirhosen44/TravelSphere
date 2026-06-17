package controllers

// DashboardController serve the SSR dashboard page.
type DashboardController struct {
	BaseController
}

// Get renders the dashboard template populated with statistics.
func (c *DashboardController) Get() {
	c.Data["ActivePage"] = "dashboard"
	c.TplName = "dashboard.tpl"

	total, planned, visited, err := WishlistSvc.GetSummary()
	if err != nil {
		c.Data["TotalCount"] = 0
		c.Data["PlannedCount"] = 0
		c.Data["VisitedCount"] = 0
		c.Data["Error"] = "Failed to compute wishlist dashboard metrics."
		return
	}

	c.Data["TotalCount"] = total
	c.Data["PlannedCount"] = planned
	c.Data["VisitedCount"] = visited
}
