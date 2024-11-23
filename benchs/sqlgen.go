package benchs

import (
	"database/sql"

	"github.com/efectn/go-orm-benchmarks/db"
)

var sqldb *sql.DB

func init() {
	st := NewSuite("sqlgen")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, SqlgenInsert)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, SqlgenInsertMulti)
		// st.AddBenchmark("MultiInsert 100 row2", 200*OrmMulti, SqlgenInsertMulti2)
		st.AddBenchmark("Update", 200*OrmMulti, SqlgenUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, SqlgenRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, SqlgenReadSlice)

		db, err := sql.Open(sqlDriver, "root:abcd1234@/go-orm?parseTime=true")
		CheckErr(err)

		sqldb = db
	}
}

func SqlgenInsert(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		_, err := db.InsertOne(ctx, sqldb, m)
		CheckErr(err, b)

		m = &Model{
			Id:      int(m.Id),
			Name:    m.Name,
			Title:   m.Title,
			Fax:     m.Fax,
			Web:     m.Web,
			Age:     int(m.Age),
			Right:   m.Right,
			Counter: m.Counter,
		}
	}
}

func SqlgenInsertMulti(b *B) {
	var ms []Model
	WrapExecute(b, func() {
		InitDB()
		ms = make([]Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, *NewModel())
		}
	})

	for i := 0; i < b.N; i++ {
		for _, m := range ms {
			m.Id = 0
		}
		_, err := db.Insert(ctx, sqldb, ms)
		CheckErr(err, b)
	}
}

// func SqlgenInsertMulti(b *B) {
// 	var ms []Model
// 	WrapExecute(b, func() {
// 		sqlutil.SetDialect("postgres")
// 		InitDB()
// 		ms = make([]Model, 0, 100)
// 		for i := 0; i < 100; i++ {
// 			ms = append(ms, *NewModel())
// 		}
// 	})

// 	for i := 0; i < b.N; i++ {
// 		_, err := sqlutil.InsertInto(ctx, sqldb, ms)
// 		CheckErr(err, b)
// 	}
// }

// func SqlgenInsertMulti2(b *B) {
// 	var ms []Model
// 	WrapExecute(b, func() {
// 		sqlutil.SetDialect("postgres")
// 		InitDB()
// 		ms = make([]Model, 0, 100)
// 		for i := 0; i < 100; i++ {
// 			ms = append(ms, *NewModel())
// 		}
// 	})

// 	for i := 0; i < b.N; i++ {
// 		_, err := sqlutil.InsertInto(ctx, sqldb, ms)
// 		CheckErr(err, b)
// 	}
// }

func SqlgenUpdate(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		err := sqlgenInsert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := db.UpdateByPK(ctx, sqldb, &Model{
			Name:    m.Name,
			Title:   m.Title,
			Fax:     m.Fax,
			Web:     m.Web,
			Age:     int(m.Age),
			Right:   m.Right,
			Counter: m.Counter,
			Id:      m.Id,
		})
		CheckErr(err, b)
	}
}

func SqlgenRead(b *B) {
	var m *Model
	WrapExecute(b, func() {
		InitDB()
		m = NewModel()
		err := sqlgenInsert(m)
		CheckErr(err, b)
	})

	m.Id = 1
	for i := 0; i < b.N; i++ {
		err := db.FindByPK(ctx, sqldb, m)
		CheckErr(err, b)
	}
}

func SqlgenReadSlice(b *B) {
	var m *Model
	WrapExecute(b, func() {
		var err error
		InitDB()
		m = NewModel()
		for i := 0; i < 100; i++ {
			err = sqlgenInsert(m)
			CheckErr(err, b)
		}
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		models := make([]Model, 100)
		var m Model
		rows, err := sqldb.QueryContext(ctx, sqlxSelectMultiSQL)
		CheckErr(err, b)
		for j := 0; rows.Next() && j < len(models); j++ {
			err = rows.Scan(m.Addrs()...)
			CheckErr(err, b)
			models[j] = m
		}
		rows.Close()
	}
}

func sqlgenInsert(m *Model) error {
	_, err := db.InsertOne(ctx, sqldb, m)
	return err
	// return nil
}
