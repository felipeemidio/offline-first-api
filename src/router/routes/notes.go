package routes

import (
	"api/src/controllers"
	"net/http"
)

var notesRoutes = []Route{
	{
		URI:       "/notes",
		Method:    http.MethodPost,
		Callback:  controllers.CreateNote,
		IsPrivate: false,
	},
	{
		URI:       "/notes",
		Method:    http.MethodGet,
		Callback:  controllers.GetNotes,
		IsPrivate: false,
	},
	{
		URI:       "/notes/{noteId}",
		Method:    http.MethodGet,
		Callback:  controllers.GetNoteById,
		IsPrivate: false,
	},
	{
		URI:       "/notes/{noteId}",
		Method:    http.MethodPut,
		Callback:  controllers.EditNote,
		IsPrivate: false,
	},
	{
		URI:       "/notes/{noteId}",
		Method:    http.MethodDelete,
		Callback:  controllers.DeleteNote,
		IsPrivate: false,
	},
}
