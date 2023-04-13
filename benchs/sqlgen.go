package benchs

import (
	"database/sql"

	sqlutil "github.com/si3nloong/sqlgen/sql"
)

var sqldb *sql.DB

func init() {
	st := NewSuite("sqlgen")
	st.InitF = func() {
		st.AddBenchmark("InsertOne", 200*OrmMulti, SqlgenInsertOne)
		st.AddBenchmark("InsertRaw", 200*OrmMulti, SqlgenInsertRaw)
		st.AddBenchmark("MultiInsert 100 row", 200*OrmMulti, SqlgenInsertMulti)
		st.AddBenchmark("MultiInsert 100 row2", 200*OrmMulti, SqlgenInsertMulti2)
		st.AddBenchmark("Update", 200*OrmMulti, SqlgenUpdate)
		st.AddBenchmark("Read", 200*OrmMulti, SqlgenRead)
		st.AddBenchmark("MultiRead limit 100", 200*OrmMulti, SqlxReadSlice)

		db, err := sql.Open("pgx", OrmSource)
		CheckErr(err)

		sqldb = db
	}
}

func sqlgenInsert(m *Model) error {
	_, err := sqlutil.InsertOne(ctx, sqldb, m)
	CheckErr(err)

	return nil
}

func SqlgenInsert(b *B) {
	var m Model
	WrapExecute(b, func() {
		sqlutil.SetDialect("postgres")
		InitDB()
		m = *NewModel()
	})

	for i := 0; i < b.N; i++ {
		_, err := sqlutil.InsertInto(ctx, sqldb, []Model{m})
		CheckErr(err, b)
	}
}

func SqlgenInsert2(b *B) {
	var m Model
	WrapExecute(b, func() {
		sqlutil.SetDialect("postgres")
		InitDB()
		m = *NewModel()
	})

	for i := 0; i < b.N; i++ {
		_, err := sqlutil.InsertInto(ctx, sqldb, []Model{m})
		CheckErr(err, b)
	}
}

func SqlgenInsertOne(b *B) {
	var m *Model
	WrapExecute(b, func() {
		sqlutil.SetDialect("postgres")
		InitDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		_, err := sqlutil.InsertOne(ctx, sqldb, m)
		CheckErr(err, b)
	}
}

func SqlgenInsertRaw(b *B) {
	var m Model
	WrapExecute(b, func() {
		sqlutil.SetDialect("postgres")
		InitDB()
		m = *NewModel()
	})

	for i := 0; i < b.N; i++ {
		_, err := sqldb.Exec(sqlxInsertSQL, m.Values()...)
		CheckErr(err, b)
	}
}

func SqlgenInsertMulti(b *B) {
	var ms []Model
	WrapExecute(b, func() {
		sqlutil.SetDialect("postgres")
		InitDB()
		ms = make([]Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, *NewModel())
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := sqlutil.InsertInto(ctx, sqldb, ms)
		CheckErr(err, b)
	}
}

func SqlgenInsertMulti2(b *B) {
	var ms []Model
	WrapExecute(b, func() {
		sqlutil.SetDialect("postgres")
		InitDB()
		ms = make([]Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, *NewModel())
		}
	})

	for i := 0; i < b.N; i++ {
		_, err := sqlutil.InsertInto(ctx, sqldb, ms)
		CheckErr(err, b)
	}
}

func SqlgenUpdate(b *B) {
	var m *Model
	WrapExecute(b, func() {
		sqlutil.SetDialect("postgres")
		InitDB()
		m = NewModel()
		err := sqlgenInsert(m)
		CheckErr(err, b)
	})

	for i := 0; i < b.N; i++ {
		_, err := sqlutil.UpdateOne(ctx, sqldb, m)
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

	for i := 0; i < b.N; i++ {
		err := sqlutil.FindOne(ctx, sqldb, m)
		CheckErr(err, b)
	}
}
