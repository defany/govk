// Code generated by https://github.com/defany/govk. DO NOT EDIT.

package tests

import (
	"github.com/defany/govk/api/gen/models"
	"github.com/defany/govk/api/gen/notes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/defany/govk/vk"
	"testing"
	"encoding/json"
)

func fillRandomlyNotesAddRequest(r *requests.NotesAddRequest) {
	r.WithTitle(randString())
	r.WithText(randString())
	lPrivacyView := randIntn(maxArrayLength + 1)
	r.WithPrivacyView(randStringArr(lPrivacyView))
	lPrivacyComment := randIntn(maxArrayLength + 1)
	r.WithPrivacyComment(randStringArr(lPrivacyComment))
}

func TestVKNotesAddSuccess(t *testing.T) {
	params := requests.NewNotesAddRequest()
	fillRandomlyNotesAddRequest(&params)
	var expected models.NotesAddResponse
	fillRandomlyNotesAddResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "notes.add", params.Params(), expectedJSON))
	resp, err := vk.Api.Notes.NotesAdd(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNotesCreateCommentRequest(r *requests.NotesCreateCommentRequest) {
	r.WithNoteId(randInt())
	r.WithOwnerId(randInt())
	r.WithReplyTo(randInt())
	r.WithMessage(randString())
	r.WithGuid(randString())
}

func TestVKNotesCreateCommentSuccess(t *testing.T) {
	params := requests.NewNotesCreateCommentRequest()
	fillRandomlyNotesCreateCommentRequest(&params)
	var expected models.NotesCreateCommentResponse
	fillRandomlyNotesCreateCommentResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "notes.createComment", params.Params(), expectedJSON))
	resp, err := vk.Api.Notes.NotesCreateComment(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNotesDeleteRequest(r *requests.NotesDeleteRequest) {
	r.WithNoteId(randInt())
}

func TestVKNotesDeleteSuccess(t *testing.T) {
	params := requests.NewNotesDeleteRequest()
	fillRandomlyNotesDeleteRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "notes.delete", params.Params(), expectedJSON))
	resp, err := vk.Api.Notes.NotesDelete(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNotesDeleteCommentRequest(r *requests.NotesDeleteCommentRequest) {
	r.WithCommentId(randInt())
	r.WithOwnerId(randInt())
}

func TestVKNotesDeleteCommentSuccess(t *testing.T) {
	params := requests.NewNotesDeleteCommentRequest()
	fillRandomlyNotesDeleteCommentRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "notes.deleteComment", params.Params(), expectedJSON))
	resp, err := vk.Api.Notes.NotesDeleteComment(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNotesEditRequest(r *requests.NotesEditRequest) {
	r.WithNoteId(randInt())
	r.WithTitle(randString())
	r.WithText(randString())
	lPrivacyView := randIntn(maxArrayLength + 1)
	r.WithPrivacyView(randStringArr(lPrivacyView))
	lPrivacyComment := randIntn(maxArrayLength + 1)
	r.WithPrivacyComment(randStringArr(lPrivacyComment))
}

func TestVKNotesEditSuccess(t *testing.T) {
	params := requests.NewNotesEditRequest()
	fillRandomlyNotesEditRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "notes.edit", params.Params(), expectedJSON))
	resp, err := vk.Api.Notes.NotesEdit(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNotesEditCommentRequest(r *requests.NotesEditCommentRequest) {
	r.WithCommentId(randInt())
	r.WithOwnerId(randInt())
	r.WithMessage(randString())
}

func TestVKNotesEditCommentSuccess(t *testing.T) {
	params := requests.NewNotesEditCommentRequest()
	fillRandomlyNotesEditCommentRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "notes.editComment", params.Params(), expectedJSON))
	resp, err := vk.Api.Notes.NotesEditComment(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNotesGetRequest(r *requests.NotesGetRequest) {
	lNoteIds := randIntn(maxArrayLength + 1)
	r.WithNoteIds(randIntArr(lNoteIds))
	r.WithUserId(randInt())
	r.WithOffset(randInt())
	r.WithCount(randInt())
	r.WithSort(randInt())
}

func TestVKNotesGetSuccess(t *testing.T) {
	params := requests.NewNotesGetRequest()
	fillRandomlyNotesGetRequest(&params)
	var expected models.NotesGetResponse
	fillRandomlyNotesGetResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "notes.get", params.Params(), expectedJSON))
	resp, err := vk.Api.Notes.NotesGet(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNotesGetByIdRequest(r *requests.NotesGetByIdRequest) {
	r.WithNoteId(randInt())
	r.WithOwnerId(randInt())
	r.WithNeedWiki(randBool())
}

func TestVKNotesGetByIdSuccess(t *testing.T) {
	params := requests.NewNotesGetByIdRequest()
	fillRandomlyNotesGetByIdRequest(&params)
	var expected models.NotesGetByIdResponse
	fillRandomlyNotesGetByIdResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "notes.getById", params.Params(), expectedJSON))
	resp, err := vk.Api.Notes.NotesGetById(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNotesGetCommentsRequest(r *requests.NotesGetCommentsRequest) {
	r.WithNoteId(randInt())
	r.WithOwnerId(randInt())
	r.WithSort(randInt())
	r.WithOffset(randInt())
	r.WithCount(randInt())
}

func TestVKNotesGetCommentsSuccess(t *testing.T) {
	params := requests.NewNotesGetCommentsRequest()
	fillRandomlyNotesGetCommentsRequest(&params)
	var expected models.NotesGetCommentsResponse
	fillRandomlyNotesGetCommentsResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "notes.getComments", params.Params(), expectedJSON))
	resp, err := vk.Api.Notes.NotesGetComments(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

func fillRandomlyNotesRestoreCommentRequest(r *requests.NotesRestoreCommentRequest) {
	r.WithCommentId(randInt())
	r.WithOwnerId(randInt())
}

func TestVKNotesRestoreCommentSuccess(t *testing.T) {
	params := requests.NewNotesRestoreCommentRequest()
	fillRandomlyNotesRestoreCommentRequest(&params)
	var expected models.BaseOkResponse
	fillRandomlyBaseOkResponse(&expected)
	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err)
	token := randString()
	vk, err := govk.NewVK(token)
	assert.NoError(t, err)
	vk.Api.WithHTTP(NewTestClient(t, "notes.restoreComment", params.Params(), expectedJSON))
	resp, err := vk.Api.Notes.NotesRestoreComment(params)
	assert.EqualValues(t, expected, resp)
	assert.NoError(t, err)
}

