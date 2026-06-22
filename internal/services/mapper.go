package notes

import (
	"yat-blog-notes/internal/repositories"
)

func ToServiceAuthorNotes(
	repoDto *repositories.GetNotesByAuthorIdDto,
) *GetNotesByAuthorIdDto {
	notes := make([]Note, len(repoDto.Notes))
	for i := range repoDto.Notes {
		notes[i] = Note(repoDto.Notes[i])
	}
	return &GetNotesByAuthorIdDto{Notes: notes}
}

func ToServiceNote(note *repositories.Note) *Note {
	return &Note{
		Id:        note.Id,
		AuthorId:  note.AuthorId,
		Text:      note.Text,
		CreatedAt: note.CreatedAt,
	}
}
