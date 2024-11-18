package main

import (
	"github.com/muhammadolammi/uniarchive/internal/database"
)

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
func convertDBFacultyToMainFaculty(dbFaculty database.Faculty) Faculty {
	return Faculty{
		ID:           dbFaculty.ID,
		CreatedAt:    dbFaculty.CreatedAt,
		UpdatedAt:    dbFaculty.UpdatedAt,
		Name:         dbFaculty.Name,
		UniversityID: dbFaculty.UniversityID,
	}
}
func convertDBFacultiesToMainFaculties(dbFaculties []database.Faculty) []Faculty {
	faculties := []Faculty{}
	for _, faculty := range dbFaculties {
		faculties = append(faculties, convertDBFacultyToMainFaculty(faculty))
	}
	return faculties
}

func convertDBDepartmentToMainDepartment(dbDepatment database.Department) Department {
	return Department{
		ID:        dbDepatment.ID,
		Name:      dbDepatment.Name,
		CreatedAt: dbDepatment.CreatedAt,
		UpdatedAt: dbDepatment.UpdatedAt,
		FacultyID: dbDepatment.FacultyID,
	}
}
func convertDBDepartmentsToMainDepartments(dbDepatments []database.Department) []Department {
	departments := []Department{}
	for _, department := range dbDepatments {
		departments = append(departments, convertDBDepartmentToMainDepartment(department))
	}
	return departments
}

func convertDBLevelToMainLevel(dbLevel database.Level) Level {
	return Level{
		ID:        dbLevel.ID,
		CreatedAt: dbLevel.CreatedAt,
		UpdatedAt: dbLevel.UpdatedAt,
		Name:      dbLevel.Name,
		Code:      int(dbLevel.Code),
	}
}

func convertDBLevelsToMainLevels(dblevels []database.Level) []Level {
	levels := []Level{}
	for _, dbLevel := range dblevels {
		levels = append(levels, convertDBLevelToMainLevel(dbLevel))
	}
	return levels
}

func convertDBCourseToMainCourse(dbCourse database.Course) Course {
	return Course{
		ID:           dbCourse.ID,
		CreatedAt:    dbCourse.CreatedAt,
		UpdatedAt:    dbCourse.UpdatedAt,
		Name:         dbCourse.Name,
		DepartmentID: dbCourse.DepartmentID,
		CourseCode:   dbCourse.CourseCode,
	}
}

func convertDBCoursesToMainCourses(dbCourses []database.Course) []Course {
	courses := []Course{}
	for _, dbCourse := range dbCourses {
		courses = append(courses, convertDBCourseToMainCourse(dbCourse))
	}
	return courses
}

func convertDBMaterialToMainMaterial(dbMaterial database.Material) Material {
	return Material{
		ID:        dbMaterial.ID,
		CreatedAt: dbMaterial.CreatedAt,
		UpdatedAt: dbMaterial.UpdatedAt,
		Name:      dbMaterial.Name,
		CourseID:  dbMaterial.CourseID,
		CloudUrl:  dbMaterial.CloudUrl,
	}
}

func convertDBMaterialsToMainMaterials(dbMaterials []database.Material) []Material {
	materials := []Material{}
	for _, dbMaterial := range dbMaterials {
		materials = append(materials, convertDBMaterialToMainMaterial(dbMaterial))
	}
	return materials
}

func convertDBUserToMainMaterial(dbUser database.User) User {
	return User{
		ID:           dbUser.ID,
		CreatedAt:    dbUser.CreatedAt,
		UpdatedAt:    dbUser.UpdatedAt,
		Email:        dbUser.Email,
		FirstName:    dbUser.FirstName,
		LastName:     dbUser.LastName,
		OtherName:    dbUser.OtherName,
		MatricNumber: dbUser.MatricNumber,
		LevelID:      dbUser.LevelID,
		UniversityID: dbUser.UniversityID,
		FacultyID:    dbUser.FacultyID,
		DepartmentID: dbUser.DepartmentID,
		Password:     dbUser.Password,
		IsAdmin:      dbUser.IsAdmin,
	}
}

func convertDBUsersToMainUsers(dbUsers []database.User) []User {
	users := []User{}
	for _, dbUser := range dbUsers {
		users = append(users, convertDBUserToMainMaterial(dbUser))
	}
	return users
}
