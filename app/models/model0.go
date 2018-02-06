package models

type PageReference struct{
	PageId			int
	LinkTitle		string
	Link 				string
}

type PageIndex struct{
	PageId 			int
	PageTitle 	string
}

type PageBody struct{
	PageId			int
	ErrorCode		string
	ErrorAbst		string
	ErrorDetail string
	Solutions		string
	Importance	int
	Complete		bool
}

type PageSummary struct{
	PageId			int
	SummaryId		int64
	SummaryPage int64
}

type Tags struct{
	PageId			int
	TagName			string
}

type PageRelation struct{
	Page0				PageIndex
	Page1				PageIndex
}

type Page struct{
	Title				PageIndex
	Tags				[]Tags
	Body				PageBody
	Reference		[]Reference
	Relation		[]PageRelation
}
