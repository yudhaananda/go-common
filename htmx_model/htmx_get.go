package htmxmodel

import "html/template"

type MemberStruct struct {
	Member template.HTML
}

type HTMXGet struct {
	Header       []MemberStruct
	Column       []Column
	SideBar      []SideBar
	Link         template.HTML
	SectionName  template.HTML
	DateJQuery   []DateJQuery
	Filter       []HTMXFilter
	Pagination   []HTMXPagination
	IsFirst      bool
	IsLast       bool
	PreviousPage template.HTML
	NextPage     template.HTML
	LastPage     template.HTML
	Take         template.HTML
	QueryPage    template.HTML
	QueryTake    template.HTML
}

type DateJQuery struct {
	Value string
}

type HTMXPagination struct {
	Active    template.HTML
	Link      template.HTML
	Page      template.HTML
	QueryPage template.HTML
}

type HTMXFilter struct {
	Type  template.HTML
	Id    template.HTML
	Label template.HTML
	Value template.HTML
}

type SideBar struct {
	Active template.HTML
	Name   template.HTML
	Link   template.HTML
}

type Column struct {
	Row  []MemberStruct
	Id   template.HTML
	Name template.HTML
}

type Modal struct {
	Name    template.HTML
	Link    template.HTML
	Id      template.HTML
	Method  template.HTML
	Members []ModalMember
}

type ModalMember struct {
	Type        template.HTML
	Id          template.HTML
	Name        template.HTML
	Value       template.HTML
	Placeholder template.HTML
}
