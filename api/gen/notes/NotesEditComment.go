// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package requests

import (
	"github.com/defany/govk/api"
	"github.com/defany/govk/api/gen/models"
)

// NotesEditComment Edits a comment on a note.
type NotesEditCommentRequest api.Params

func NewNotesEditCommentRequest() NotesEditCommentRequest {
	params := make(NotesEditCommentRequest, 4)
	return params
}

func (n NotesEditCommentRequest) WithCommentId(n_comment_id int) NotesEditCommentRequest{
	n["comment_id"] = n_comment_id
	return n
}

func (n NotesEditCommentRequest) WithOwnerId(n_owner_id int) NotesEditCommentRequest{
	n["owner_id"] = n_owner_id
	return n
}

func (n NotesEditCommentRequest) WithMessage(n_message string) NotesEditCommentRequest{
	n["message"] = n_message
	return n
}

func (n NotesEditCommentRequest) Params() api.Params {
	return api.Params(n)
}

// May execute with listed access token types:
//    [ user ]
// When executing method, may return one of global or with listed codes API errors:
//    [ Error_AccessComment ]
//
// https://dev.vk.com/method/notes.editComment
func (n *Notes) NotesEditComment(params ...api.MethodParams) (resp models.BaseOkResponse, err error) {
	req := api.NewRequest[models.BaseOkResponse](n.api)

	res, err := req.Execute("notes.editComment", api.ParamsOrNil(params))
	if err != nil {
		return res, err
	}

	return res, nil
}

