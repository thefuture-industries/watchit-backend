package action

func CreateFavourite(uuid string, genreID uint16) error {
	logger := lib.NewLogger()
	db := database.GetDB()

	var genres schema.Genres
	if err := db.Where("genre_id = ?", genreID).First(&genres).Error; err != nil {
		logger.Error(fmt.Sprintf("Database -> 'Failed to create recommendation: %v'", err))
		return fmt.Errorf("failed to create recommendation.")
	}

	recomSchemaPrepare := schema.Recommendations{
		UUID:    uuid,
		GenreID: genres.ID,
		Count:   1,
	}

	if err := db.Create(&recomSchemaPrepare).Error; err != nil {
		logger.Error(fmt.Sprintf("Database -> 'Failed to create recommendation: %v'", err))
		return fmt.Errorf("failed to create recommendation.")
	}

	return nil
}
