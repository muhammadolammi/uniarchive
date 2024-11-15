package main

import "github.com/muhammadolammi/uniarchive/internal/database"

func convertDBUniToMainUni(dbUni database.University) University {
	return University{
		ID:        dbUni.ID,
		CreatedAt: dbUni.CreatedAt,
		UpdatedAt: dbUni.UpdatedAt,
		Name:      dbUni.Name,
		Alias:     dbUni.Alias,
	}
}

func convertDBUnisToMainUnis(dbUnis []database.University) []University {
	unis := []University{}
	for _, dbuni := range dbUnis {
		unis = append(unis, convertDBUniToMainUni(dbuni))
	}
	return unis
}
