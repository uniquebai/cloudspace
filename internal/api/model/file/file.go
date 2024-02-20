package file

type RenameReq struct {
	ID      string `json:"ID,omitempty" path:"id"`
	NewName string `json:"newName,omitempty"`
}
type ModifyFileContentReq struct {
	ID      string `json:"ID,omitempty" path:"id"`
	Content string `json:"content,omitempty"`
}
type CreateTextFileReq struct {
	ParentID string `json:"parentID,omitempty" path:"id"`
	FileName string `json:"fileName,omitempty"`
	Content  string `json:"content,omitempty"`
}
