package models

type MessageBoard struct {
	Model
	MessageBoardContent  string `json:"message_board_content"`
	MessageBoardUser     string `json:"message_board_user"`
	MessageBoardContacts string `json:"message_board_contacts"`
}

func GetMessageBoards(pageOffset int, pageSize int, maps interface{}) (messageBoards []MessageBoard) {
	db.Where(maps).Offset(pageOffset).Limit(pageSize).Find(&messageBoards)
	return
}

func DeleteMessageBoard(id int) bool {
	ormDb := db.Where("id = ?", id).Delete(&MessageBoard{})
	return isSuccess(ormDb)
}
