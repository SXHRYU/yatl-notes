package notes

import (
	"context"

	"golang.org/x/sync/errgroup"
	"yat-blog-notes/internal/pagination"
	"yat-blog-notes/internal/repositories"
	notes "yat-blog-notes/internal/services"
)

func (nc *NotesService) GetAuthorNotes(
	ctx context.Context,
	authorId, page, limit int,
) (*notes.GetNotesByAuthorIdDto, error) {
	var (
		response *repositories.GetNotesByAuthorIdDto
		count    int
	)
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		response, err = nc.notesRepo.GetNotesByAuthorId(
			ctx,
			authorId,
			limit,
			(page-1)*limit,
		)
		// todo: process service errors
		return err
	})

	g.Go(func() error {
		var err error
		count, err = nc.notesRepo.CountNotesByAuthorId(ctx, authorId)
		// todo: process service errors
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return notes.ToServiceAuthorNotes(
		response,
		page,
		pagination.TotalPages(count, pagination.DefaultSize),
		count,
	), nil
}
