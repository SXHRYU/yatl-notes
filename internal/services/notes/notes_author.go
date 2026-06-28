package notes

import (
	"context"
	"strings"

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

	for i := range response.Notes {
		shortenNote(&response.Notes[i].Text)
	}
	return notes.ToServiceAuthorNotes(
		response,
		page,
		pagination.TotalPages(count, pagination.DefaultSize),
		count,
	), nil
}

func shortenNote(s *string) {
	const maxTitleSize = 200
	builder := strings.Builder{}
	builder.Grow(maxTitleSize)
	var isBigText bool
	for i, ch := range *s {
		builder.WriteRune(ch)
		// we should count separately 1-byte and 2-, 3-bytes letters
		// space required for "..."
		if i/2+4 >= maxTitleSize {
			isBigText = true
			break
		}
	}

	if isBigText {
		builder.WriteString("...")
	}
	*s = builder.String()
}
