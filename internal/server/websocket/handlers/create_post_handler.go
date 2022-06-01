package handlers

func (wbsckhandler WebsocketHandler) CreatePost() int {
	postCreated := wbsckhandler.createPostService.Create()
	return postCreated
}
