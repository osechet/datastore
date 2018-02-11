package example

import (
	"reflect"
	"testing"

	test "github.com/osechet/go-datastore/_proto/osechet/test"
	datastore "google.golang.org/genproto/googleapis/datastore/v1"
)

func TestAutoQuery(t *testing.T) {
	dbBooks := []*test.Book{
		{
			Isbn:   60929871,
			Title:  "Brave New World",
			Author: "Aldous Huxley",
		},
		{
			Isbn:   140009728,
			Title:  "Nineteen Eighty-Four",
			Author: "George Orwell",
		},
		{
			Isbn:   9780140301694,
			Title:  "Alice's Adventures in Wonderland",
			Author: "Lewis Carroll",
		},
		{
			Isbn:   140008381,
			Title:  "Animal Farm",
			Author: "George Orwell",
		},
	}
	type args struct {
		query   datastore.Query
		dbBooks []*test.Book
	}
	tests := []struct {
		name string
		args args
		want []*test.Book
	}{
		{
			"sort on title - ascending",
			args{
				datastore.Query{
					Order: []*datastore.PropertyOrder{
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "title"},
							Direction: datastore.PropertyOrder_ASCENDING,
						},
					},
				},
				dbBooks,
			},
			[]*test.Book{
				{
					Isbn:   9780140301694,
					Title:  "Alice's Adventures in Wonderland",
					Author: "Lewis Carroll",
				},
				{
					Isbn:   140008381,
					Title:  "Animal Farm",
					Author: "George Orwell",
				},
				{
					Isbn:   60929871,
					Title:  "Brave New World",
					Author: "Aldous Huxley",
				},
				{
					Isbn:   140009728,
					Title:  "Nineteen Eighty-Four",
					Author: "George Orwell",
				},
			},
		},
		{
			"sort on title - descending",
			args{
				datastore.Query{
					Order: []*datastore.PropertyOrder{
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "title"},
							Direction: datastore.PropertyOrder_DESCENDING,
						},
					},
				},
				dbBooks,
			},
			[]*test.Book{
				{
					Isbn:   140009728,
					Title:  "Nineteen Eighty-Four",
					Author: "George Orwell",
				},
				{
					Isbn:   60929871,
					Title:  "Brave New World",
					Author: "Aldous Huxley",
				},
				{
					Isbn:   140008381,
					Title:  "Animal Farm",
					Author: "George Orwell",
				},
				{
					Isbn:   9780140301694,
					Title:  "Alice's Adventures in Wonderland",
					Author: "Lewis Carroll",
				},
			},
		},
		{
			"sort on author then title",
			args{
				datastore.Query{
					Order: []*datastore.PropertyOrder{
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "author"},
							Direction: datastore.PropertyOrder_ASCENDING,
						},
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "title"},
							Direction: datastore.PropertyOrder_ASCENDING,
						},
					},
				},
				dbBooks,
			},
			[]*test.Book{
				{
					Isbn:   60929871,
					Title:  "Brave New World",
					Author: "Aldous Huxley",
				},
				{
					Isbn:   140008381,
					Title:  "Animal Farm",
					Author: "George Orwell",
				},
				{
					Isbn:   140009728,
					Title:  "Nineteen Eighty-Four",
					Author: "George Orwell",
				},
				{
					Isbn:   9780140301694,
					Title:  "Alice's Adventures in Wonderland",
					Author: "Lewis Carroll",
				},
			},
		},
		{
			"sort on author then descending title",
			args{
				datastore.Query{
					Order: []*datastore.PropertyOrder{
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "author"},
							Direction: datastore.PropertyOrder_ASCENDING,
						},
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "title"},
							Direction: datastore.PropertyOrder_DESCENDING,
						},
					},
				},
				dbBooks,
			},
			[]*test.Book{
				{
					Isbn:   60929871,
					Title:  "Brave New World",
					Author: "Aldous Huxley",
				},
				{
					Isbn:   140009728,
					Title:  "Nineteen Eighty-Four",
					Author: "George Orwell",
				},
				{
					Isbn:   140008381,
					Title:  "Animal Farm",
					Author: "George Orwell",
				},
				{
					Isbn:   9780140301694,
					Title:  "Alice's Adventures in Wonderland",
					Author: "Lewis Carroll",
				},
			},
		},
		{
			"filter on author, sort on author then title",
			args{
				datastore.Query{
					Filter: &datastore.Filter{
						FilterType: &datastore.Filter_CompositeFilter{
							CompositeFilter: &datastore.CompositeFilter{
								Filters: []*datastore.Filter{
									&datastore.Filter{
										FilterType: &datastore.Filter_PropertyFilter{
											PropertyFilter: &datastore.PropertyFilter{
												Property: &datastore.PropertyReference{Name: "author"},
												Op:       datastore.PropertyFilter_GREATER_THAN_OR_EQUAL,
												Value: &datastore.Value{
													ValueType: &datastore.Value_StringValue{StringValue: "Geor"},
												},
											},
										},
									},
									&datastore.Filter{
										FilterType: &datastore.Filter_PropertyFilter{
											PropertyFilter: &datastore.PropertyFilter{
												Property: &datastore.PropertyReference{Name: "author"},
												Op:       datastore.PropertyFilter_LESS_THAN,
												Value: &datastore.Value{
													ValueType: &datastore.Value_StringValue{StringValue: "Geor" + "\uFFFF"},
												},
											},
										},
									},
								},
							},
						},
					},
					Order: []*datastore.PropertyOrder{
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "author"},
							Direction: datastore.PropertyOrder_ASCENDING,
						},
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "title"},
							Direction: datastore.PropertyOrder_ASCENDING,
						},
					},
				},
				dbBooks,
			},
			[]*test.Book{
				{
					Isbn:   140008381,
					Title:  "Animal Farm",
					Author: "George Orwell",
				},
				{
					Isbn:   140009728,
					Title:  "Nineteen Eighty-Four",
					Author: "George Orwell",
				},
			},
		},
		{
			"filter on author, no sort",
			args{
				datastore.Query{
					Filter: &datastore.Filter{
						FilterType: &datastore.Filter_CompositeFilter{
							CompositeFilter: &datastore.CompositeFilter{
								Filters: []*datastore.Filter{
									&datastore.Filter{
										FilterType: &datastore.Filter_PropertyFilter{
											PropertyFilter: &datastore.PropertyFilter{
												Property: &datastore.PropertyReference{Name: "author"},
												Op:       datastore.PropertyFilter_GREATER_THAN_OR_EQUAL,
												Value: &datastore.Value{
													ValueType: &datastore.Value_StringValue{StringValue: "Geor"},
												},
											},
										},
									},
									&datastore.Filter{
										FilterType: &datastore.Filter_PropertyFilter{
											PropertyFilter: &datastore.PropertyFilter{
												Property: &datastore.PropertyReference{Name: "author"},
												Op:       datastore.PropertyFilter_LESS_THAN,
												Value: &datastore.Value{
													ValueType: &datastore.Value_StringValue{StringValue: "Geor" + "\uFFFF"},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				dbBooks,
			},
			[]*test.Book{
				{
					Isbn:   140009728,
					Title:  "Nineteen Eighty-Four",
					Author: "George Orwell",
				},
				{
					Isbn:   140008381,
					Title:  "Animal Farm",
					Author: "George Orwell",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := AutoQuery(tt.args.query, tt.args.dbBooks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuery(t *testing.T) {
	dbBooks := []*test.Book{
		{
			Isbn:   60929871,
			Title:  "Brave New World",
			Author: "Aldous Huxley",
		},
		{
			Isbn:   140009728,
			Title:  "Nineteen Eighty-Four",
			Author: "George Orwell",
		},
		{
			Isbn:   9780140301694,
			Title:  "Alice's Adventures in Wonderland",
			Author: "Lewis Carroll",
		},
		{
			Isbn:   140008381,
			Title:  "Animal Farm",
			Author: "George Orwell",
		},
	}
	type args struct {
		query   datastore.Query
		dbBooks []*test.Book
	}
	tests := []struct {
		name string
		args args
		want []*test.Book
	}{
		{
			"sort on title - ascending",
			args{
				datastore.Query{
					Order: []*datastore.PropertyOrder{
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "title"},
							Direction: datastore.PropertyOrder_ASCENDING,
						},
					},
				},
				dbBooks,
			},
			[]*test.Book{
				{
					Isbn:   9780140301694,
					Title:  "Alice's Adventures in Wonderland",
					Author: "Lewis Carroll",
				},
				{
					Isbn:   140008381,
					Title:  "Animal Farm",
					Author: "George Orwell",
				},
				{
					Isbn:   60929871,
					Title:  "Brave New World",
					Author: "Aldous Huxley",
				},
				{
					Isbn:   140009728,
					Title:  "Nineteen Eighty-Four",
					Author: "George Orwell",
				},
			},
		},
		{
			"sort on title - descending",
			args{
				datastore.Query{
					Order: []*datastore.PropertyOrder{
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "title"},
							Direction: datastore.PropertyOrder_DESCENDING,
						},
					},
				},
				dbBooks,
			},
			[]*test.Book{
				{
					Isbn:   140009728,
					Title:  "Nineteen Eighty-Four",
					Author: "George Orwell",
				},
				{
					Isbn:   60929871,
					Title:  "Brave New World",
					Author: "Aldous Huxley",
				},
				{
					Isbn:   140008381,
					Title:  "Animal Farm",
					Author: "George Orwell",
				},
				{
					Isbn:   9780140301694,
					Title:  "Alice's Adventures in Wonderland",
					Author: "Lewis Carroll",
				},
			},
		},
		{
			"sort on author then title",
			args{
				datastore.Query{
					Order: []*datastore.PropertyOrder{
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "author"},
							Direction: datastore.PropertyOrder_ASCENDING,
						},
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "title"},
							Direction: datastore.PropertyOrder_ASCENDING,
						},
					},
				},
				dbBooks,
			},
			[]*test.Book{
				{
					Isbn:   60929871,
					Title:  "Brave New World",
					Author: "Aldous Huxley",
				},
				{
					Isbn:   140008381,
					Title:  "Animal Farm",
					Author: "George Orwell",
				},
				{
					Isbn:   140009728,
					Title:  "Nineteen Eighty-Four",
					Author: "George Orwell",
				},
				{
					Isbn:   9780140301694,
					Title:  "Alice's Adventures in Wonderland",
					Author: "Lewis Carroll",
				},
			},
		},
		{
			"sort on author then descending title",
			args{
				datastore.Query{
					Order: []*datastore.PropertyOrder{
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "author"},
							Direction: datastore.PropertyOrder_ASCENDING,
						},
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "title"},
							Direction: datastore.PropertyOrder_DESCENDING,
						},
					},
				},
				dbBooks,
			},
			[]*test.Book{
				{
					Isbn:   60929871,
					Title:  "Brave New World",
					Author: "Aldous Huxley",
				},
				{
					Isbn:   140009728,
					Title:  "Nineteen Eighty-Four",
					Author: "George Orwell",
				},
				{
					Isbn:   140008381,
					Title:  "Animal Farm",
					Author: "George Orwell",
				},
				{
					Isbn:   9780140301694,
					Title:  "Alice's Adventures in Wonderland",
					Author: "Lewis Carroll",
				},
			},
		},
		{
			"filter on author, sort on author then title",
			args{
				datastore.Query{
					Filter: &datastore.Filter{
						FilterType: &datastore.Filter_CompositeFilter{
							CompositeFilter: &datastore.CompositeFilter{
								Filters: []*datastore.Filter{
									&datastore.Filter{
										FilterType: &datastore.Filter_PropertyFilter{
											PropertyFilter: &datastore.PropertyFilter{
												Property: &datastore.PropertyReference{Name: "author"},
												Op:       datastore.PropertyFilter_GREATER_THAN_OR_EQUAL,
												Value: &datastore.Value{
													ValueType: &datastore.Value_StringValue{StringValue: "Geor"},
												},
											},
										},
									},
									&datastore.Filter{
										FilterType: &datastore.Filter_PropertyFilter{
											PropertyFilter: &datastore.PropertyFilter{
												Property: &datastore.PropertyReference{Name: "author"},
												Op:       datastore.PropertyFilter_LESS_THAN,
												Value: &datastore.Value{
													ValueType: &datastore.Value_StringValue{StringValue: "Geor" + "\uFFFF"},
												},
											},
										},
									},
								},
							},
						},
					},
					Order: []*datastore.PropertyOrder{
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "author"},
							Direction: datastore.PropertyOrder_ASCENDING,
						},
						&datastore.PropertyOrder{
							Property:  &datastore.PropertyReference{Name: "title"},
							Direction: datastore.PropertyOrder_ASCENDING,
						},
					},
				},
				dbBooks,
			},
			[]*test.Book{
				{
					Isbn:   140008381,
					Title:  "Animal Farm",
					Author: "George Orwell",
				},
				{
					Isbn:   140009728,
					Title:  "Nineteen Eighty-Four",
					Author: "George Orwell",
				},
			},
		},
		{
			"filter on author, no sort",
			args{
				datastore.Query{
					Filter: &datastore.Filter{
						FilterType: &datastore.Filter_CompositeFilter{
							CompositeFilter: &datastore.CompositeFilter{
								Filters: []*datastore.Filter{
									&datastore.Filter{
										FilterType: &datastore.Filter_PropertyFilter{
											PropertyFilter: &datastore.PropertyFilter{
												Property: &datastore.PropertyReference{Name: "author"},
												Op:       datastore.PropertyFilter_GREATER_THAN_OR_EQUAL,
												Value: &datastore.Value{
													ValueType: &datastore.Value_StringValue{StringValue: "Geor"},
												},
											},
										},
									},
									&datastore.Filter{
										FilterType: &datastore.Filter_PropertyFilter{
											PropertyFilter: &datastore.PropertyFilter{
												Property: &datastore.PropertyReference{Name: "author"},
												Op:       datastore.PropertyFilter_LESS_THAN,
												Value: &datastore.Value{
													ValueType: &datastore.Value_StringValue{StringValue: "Geor" + "\uFFFF"},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				dbBooks,
			},
			[]*test.Book{
				{
					Isbn:   140009728,
					Title:  "Nineteen Eighty-Four",
					Author: "George Orwell",
				},
				{
					Isbn:   140008381,
					Title:  "Animal Farm",
					Author: "George Orwell",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Query(tt.args.query, tt.args.dbBooks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}
