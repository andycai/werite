package handler

type AdminHandler struct{}

var Admin = new(AdminHandler)

//#region private methods

//

//#endregion

//#region handler

// Login login to admin
func (ah AdminHandler) Login(c *Ctx) error {
	return nil
}

//#endregion