package actions

func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/events").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}
