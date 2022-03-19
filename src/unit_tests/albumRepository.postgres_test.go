package repository_test

import (
	"database/sql"
	"regexp"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/DATA-DOG/go-sqlmock"

	"acy.com/api/src/entities"
	"acy.com/api/src/repositories"
)


var _ = Describe("Test Repository", func() {
	var repository *repositories.AlbumRepository
	var mock sqlmock.Sqlmock

	BeforeEach(func() {
		var db *sql.DB
		var err error

		db, mock, err = sqlmock.New() // mock sql.DB
		Expect(err).ShouldNot(HaveOccurred())
		
		repository = repositories.NewAlbumRepository(db)
	})
	AfterEach(func() {
		err := mock.ExpectationsWereMet() // make sure all expectations were met
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("FindAll", func() {
		It("found", func(){
			const  sqlSelectAll = `SELECT * FROM "albums"`
			fakeAlbum1 := &entities.Album{ Id: 1, Title: "fake title 1", Artist: "fake artist 1", Price: 100.20, HasRead: true }
			fakeAlbum2 := &entities.Album{ Id: 2, Title: "fake title 2", Artist: "fake artist 1", Price: 20.40, HasRead: false }
			rows := sqlmock.
							NewRows([]string{"id", "title", "artist", "price", "has_read"}).
							AddRow(fakeAlbum1.Id, fakeAlbum1.Title, fakeAlbum1.Artist, fakeAlbum1.Price, fakeAlbum1.HasRead).
							AddRow(fakeAlbum2.Id, fakeAlbum2.Title, fakeAlbum2.Artist, fakeAlbum2.Price, fakeAlbum2.HasRead)

				mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).WillReturnRows(rows)

				albums, err := repository.FindAll()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(albums).Should(Equal([]entities.Album{ *fakeAlbum1, *fakeAlbum2 } ))
			})

		It("not found", func() {
                // ignore sql match
                mock.ExpectQuery(`.+`).WillReturnRows(sqlmock.NewRows(nil))
                albums, err := repository.FindAll()
				Expect(err).ShouldNot(HaveOccurred())
                Expect(albums).Should(Equal([]entities.Album{}))
        })
	})

	Context("FindById", func() {
		It("found", func(){
			const  sqlSelectAll = `SELECT * FROM "albums" WHERE "id" = $1`
			fakeAlbum2 := &entities.Album{ Id: 2, Title: "fake title 2", Artist: "fake artist 1", Price: 20.40, HasRead: false }
			rows := sqlmock.
							NewRows([]string{"id", "title", "artist", "price", "has_read"}).
							AddRow(fakeAlbum2.Id, fakeAlbum2.Title, fakeAlbum2.Artist, fakeAlbum2.Price, fakeAlbum2.HasRead)

				mock.ExpectQuery(regexp.QuoteMeta(sqlSelectAll)).WithArgs(fakeAlbum2.Id).WillReturnRows(rows)

				albums, err := repository.FindById(2)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(albums).Should(Equal(*fakeAlbum2))
			})

		It("not found", func() {
                // ignore sql match
                mock.ExpectQuery(`.+`).WillReturnRows(sqlmock.NewRows(nil))
                albums, err := repository.FindById(2)
				Expect(err).ShouldNot(HaveOccurred())
                Expect(albums).Should(Equal(entities.Album{}))
        })
	})

	Context("Create", func() {
		fakeAlbum1 := &entities.Album{ Id: 1, Title: "fake title 1", Artist: "fake artist 1", Price: 100.20, HasRead: true }
		It("created", func(){
			const sqlInsert = `INSERT INTO "albums" ("title","artist","price","has_read","id") 
                                        VALUES ($1,$2,$3,$4,$5) RETURNING "id"`
			
			rows := sqlmock.
							NewRows([]string{"id", "title", "artist", "price", "has_read"}).
							AddRow(fakeAlbum1.Id, fakeAlbum1.Title, fakeAlbum1.Artist, fakeAlbum1.Price, fakeAlbum1.HasRead)
			mock.ExpectBegin() // begin transaction
			mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).WithArgs(fakeAlbum1.Title, fakeAlbum1.Artist, fakeAlbum1.Price, fakeAlbum1.HasRead, fakeAlbum1.Id).WillReturnRows(rows)
			mock.ExpectCommit() // commit transaction

			albums, err := repository.Create(fakeAlbum1)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(albums).Should(Equal(*fakeAlbum1))
		})
	})
})
