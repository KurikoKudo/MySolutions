package models

type Reference struct{
	Link 			string
	LinkTitle	string
}

type Title struct{
	PageId 		int
	PageTitle string
}

type Page struct{
	PageId			int
	PageTitle		string
	ErrorCode		string
	Body				string
	ErrorAbst		string
	ErrorDetail string
	Solutions		string
	Reference		[]Reference
	SummaryId		int
	SummaryPage int
	Importance	int
	Complete		bool
	TagName			[]string
	Relation		[]int
	RelationPage[]Title
}
