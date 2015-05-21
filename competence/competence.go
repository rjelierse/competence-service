package competence

// Competence model defines the available data
type Competence struct {
	Name           string `bson:"name" json:"name"`
	RequiredPoints int    `bson:"requiredPoints" json:"required_points"`
	AssignedPoints int    `bson:"assignedPoints" json:"assigned_points"`
	UsedPoints     int    `bson:"usedPoints" json:"used_points"`
}
