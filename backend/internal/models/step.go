package models

type Step struct {
	Condition       string            `json:"condition,omitempty" bson:"condition"`
	Function        string            `json:"function" bson:"function"`
	Input           map[string]string `json:"input" bson:"input"`
	Integration     string            `json:"integration" bson:"integration"`
	IntegrationType string            `json:"integrationType" bson:"integrationType"`
	Name            string            `json:"name" bson:"name"`
	DisplayName     string            `json:"displayName" bson:"displayName"`
	Output          map[string]string `json:"output,omitempty" bson:"output"`
}
