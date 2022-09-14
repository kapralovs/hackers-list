package hackers

type Usecase interface {
	// CreateBookmark(ctx context.Context, user *models.Hacker, url, title string) error
	// GetBookmarks(ctx context.Context, user *models.Hacker) ([]*models.Hacker, error)
	// DeleteBookmark(ctx context.Context, user *models.Hacker, id string) error
}

type Repository interface {
	// CreateBookmark(ctx context.Context, user *models.Hacker, bm *models.Hacker) error
	// GetBookmarks(ctx context.Context, user *models.Hacker) ([]*models.Hacker, error)
	// DeleteBookmark(ctx context.Context, user *models.Hacker, id string) error
}
