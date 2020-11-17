package internal

import (
	"database/sql"
	"errors"

	"github.com/promo/view"
)

type PromoInc struct {
	PromoID     int    `json:"promo_id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

func (u *PromoInc) CreatePromo(db *sql.DB, p *view.Promo) error {
	stmt, _ := db.Prepare("INSERT INTO INCPromo (title,description,start_date,end_date) VALUES (?,?,?,?);")
	stmt.Exec(&p.Title, &p.Description, &p.StartDate, &p.EndDate)
	return nil
}

func (u *PromoInc) GetPromo(id int, db *sql.DB) error {
	row := db.QueryRow("SELECT * FROM INCPromo WHERE promo_id=?;", id)
	row.Scan(&u.PromoID, &u.Title, &u.Description, &u.StartDate, &u.EndDate)
	if u.PromoID <= 0 {
		return errors.New("No promos with this ID")
	}
	return nil
}

func (u *PromoInc) GetAll(db *sql.DB) (*[]view.Promo, error) {
	promos := []view.Promo{}
	rows, _ := db.Query("SELECT * FROM INCPromo ;")
	for rows.Next() {
		u := view.Promo{}
		err := rows.Scan(&u.PromoID, &u.Title, &u.Description, &u.StartDate, &u.EndDate)
		if err != nil {
			return nil, errors.New("No promos to show up")
		}
		promos = append(promos, u)
	}
	return &promos, nil
}

func (u *PromoInc) UpdatePromo(id int, db *sql.DB, p *view.Promo) error {
	stmt, _ := db.Prepare("UPDATE INCPromo SET title=?,description=?,start_date=?,end_date=? WHERE promo_id=?;")
	n, _ := stmt.Exec(p.Title, p.Description, p.StartDate, p.EndDate, id)
	r, _ := n.RowsAffected()
	if r <= 0 {
		return errors.New("No Task with this ID")
	}
	return nil
}

func (u *PromoInc) DeletePromo(id int, db *sql.DB) error {
	r, _ := db.Exec("DELETE FROM INCPromo WHERE promo_id=?;", id)
	n, _ := r.RowsAffected()
	if n <= 0 {
		return errors.New("No task with this ID ")
	}
	return nil
}
