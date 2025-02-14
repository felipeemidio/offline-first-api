package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var note models.Note
	if err = json.Unmarshal(reqBody, &note); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = note.Check(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.CreateNotesRepository(db)
	noteId, err := repository.Create(note)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	note.ID = noteId

	responses.JSON(w, http.StatusCreated, note)
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.CreateNotesRepository(db)
	notes, err := repository.GetAll()
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusCreated, notes)
}

func GetNoteById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	noteId, err := strconv.ParseUint(params["noteId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.CreateNotesRepository(db)
	note, err := repository.GetById(noteId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, note)
}

func EditNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	noteId, err := strconv.ParseUint(params["noteId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var newNote models.Note
	if err = json.Unmarshal(reqBody, &newNote); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.CreateNotesRepository(db)
	err = repository.EditNote(noteId, newNote)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	noteId, err := strconv.ParseUint(params["noteId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.CreateNotesRepository(db)
	err = repository.DeleteNote(noteId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
