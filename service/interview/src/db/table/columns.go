package table

// Define schemas for tables
// Due to clustering order, the order of columns in maps are necessary

type columnDetails struct {
	columnType            string
	isPartitionKey        bool
	isClusteringKey       bool
	isClusteringOrderDesc bool
}

var InterviewTemplatesColumns map[string]columnDetails = map[string]columnDetails{
	"interview_template_id": {columnType: "TIMEUUID", isPartitionKey: true, isClusteringKey: false, isClusteringOrderDesc: false},
	"average_score":         {columnType: "INT", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
	"average_rating":        {columnType: "INT", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
	"amount_conducted":      {columnType: "INT", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
	"company":               {columnType: "TEXT", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
	"role":                  {columnType: "TEXT", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
	"skills":                {columnType: "SET<TEXT>", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
	"description":           {columnType: "TEXT", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
	"user_id":               {columnType: "INT", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
	"questions":             {columnType: "LIST<TEXT>", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
}

var InterviewTemplatesByCompanyColumns = map[string]columnDetails{
	"company":               {columnType: "TEXT", isPartitionKey: true, isClusteringKey: false, isClusteringOrderDesc: false},
	"interview_template_id": {columnType: "TIMEUUID", isPartitionKey: false, isClusteringKey: true, isClusteringOrderDesc: true},
}

var AverageScoresByRoleAndCompanyColumns = map[string]columnDetails{
	"role":                  {columnType: "TEXT", isPartitionKey: true, isClusteringKey: false, isClusteringOrderDesc: false},
	"company":               {columnType: "TEXT", isPartitionKey: true, isClusteringKey: false, isClusteringOrderDesc: false},
	"average_score":         {columnType: "INT", isPartitionKey: false, isClusteringKey: true, isClusteringOrderDesc: true},
	"interview_template_id": {columnType: "TIMEUUID", isPartitionKey: false, isClusteringKey: true, isClusteringOrderDesc: true},
}

var AverageRatingsByRoleAndCompanyColumns = map[string]columnDetails{
	"role":                  {columnType: "TEXT", isPartitionKey: true, isClusteringKey: false, isClusteringOrderDesc: false},
	"company":               {columnType: "TEXT", isPartitionKey: true, isClusteringKey: false, isClusteringOrderDesc: false},
	"average_rating":        {columnType: "INT", isPartitionKey: false, isClusteringKey: true, isClusteringOrderDesc: true},
	"interview_template_id": {columnType: "TIMEUUID", isPartitionKey: false, isClusteringKey: true, isClusteringOrderDesc: true},
}

var AmountConductedByRoleAndCompanyColumns = map[string]columnDetails{
	"role":                  {columnType: "TEXT", isPartitionKey: true, isClusteringKey: false, isClusteringOrderDesc: false},
	"company":               {columnType: "TEXT", isPartitionKey: true, isClusteringKey: false, isClusteringOrderDesc: false},
	"amount_conducted":      {columnType: "INT", isPartitionKey: false, isClusteringKey: true, isClusteringOrderDesc: false},
	"interview_template_id": {columnType: "TIMEUUID", isPartitionKey: false, isClusteringKey: true, isClusteringOrderDesc: true},
}

var InterviewTemplatesByUserColumns = map[string]columnDetails{
	"user_id":               {columnType: "INT", isPartitionKey: true, isClusteringKey: false, isClusteringOrderDesc: false},
	"interview_template_id": {columnType: "TIMEUUID", isPartitionKey: false, isClusteringKey: true, isClusteringOrderDesc: true},
}

var ConductedInterviewsColumns = map[string]columnDetails{
	"conducted_interview_id": {columnType: "TIMEUUID", isPartitionKey: true, isClusteringKey: false, isClusteringOrderDesc: false},
	"interview_template_id":  {columnType: "TIMEUUID", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
	"score":                  {columnType: "INT", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
	"user_id":                {columnType: "INT", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
	"role":                   {columnType: "TEXT", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
	"rating":                 {columnType: "INT", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
	"responses":              {columnType: "FROZEN <response_type>", isPartitionKey: false, isClusteringKey: false, isClusteringOrderDesc: false},
}

var ConductedInterviewsByUserColumns = map[string]columnDetails{
	"user_id":                {columnType: "INT", isPartitionKey: true, isClusteringKey: false, isClusteringOrderDesc: false},
	"conducted_interview_id": {columnType: "TIMEUUID", isPartitionKey: false, isClusteringKey: true, isClusteringOrderDesc: true},
}
