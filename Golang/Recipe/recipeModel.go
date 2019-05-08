package main

import (
	"database/sql"
	"log"
)

// Recipe definitio model to recipe
type Recipe struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Preparation string `json:"preparation"`
	Time        string `json:"time"`
}

// Count definitio model to count
type Count struct {
	Count int `json:"count"`
}

func countRecips(bd *sql.DB) (Count, error) {
	//return bd.QueryRow("SELECT COUNT(*) FROM recipe").Scan(&c.Count)
	c := Count{}
	rows, err := bd.Query("SELECT COUNT(*) FROM recipe")
	if err != nil {
		return c, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&c.Count); err != nil {
			return c, err
		}
	}
	return c, nil
}

func obtainRecipesByName(bd *sql.DB, name string) ([]Recipe, error) {

	//statement := fmt.Sprintf("SELECT id, name, ingredients, preparation, time FROM recipe WHERE name like '$1'", name)
	//log.Println("SELECT id, name, ingredients, preparation, time FROM recipe WHERE name like $1", name)

	rows, err := bd.Query("SELECT id, name, ingredients, preparation, time FROM recipe WHERE lower(name) LIKE '%' || $1 || '%' ", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	recetas := []Recipe{}
	for rows.Next() {
		var r Recipe
		if err := rows.Scan(&r.ID, &r.Name, &r.Ingredients, &r.Preparation, &r.Time); err != nil {
			return nil, err
		}
		recetas = append(recetas, r)
	}
	return recetas, nil

}

func obtainRecipes(bd *sql.DB, start, cont int) ([]Recipe, error) {
	//statement := fmt.Sprintf("SELECT id, name, ingredients, preparation, time FROM recipe LIMIT $1 OFFSET $2", cont, start)
	rows, err := bd.Query("SELECT id, name, ingredients, preparation, time FROM recipe LIMIT $1 OFFSET $2", cont, start) // paginacion
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	recetas := []Recipe{}
	for rows.Next() {
		var r Recipe
		if err := rows.Scan(&r.ID, &r.Name, &r.Ingredients, &r.Preparation, &r.Time); err != nil {
			return nil, err
		}
		recetas = append(recetas, r)
	}
	return recetas, nil
}

// todo secure query
func (r *Recipe) obtainRecipe(bd *sql.DB) error {
	//statement := fmt.Sprintf("SELECT name, ingredients,preparation,time FROM recipe WHERE id='$1'", r.ID)
	return bd.QueryRow("SELECT name, ingredients,preparation,time FROM recipe WHERE id= $1", r.ID).Scan(&r.Name, &r.Ingredients, &r.Preparation, &r.Time)
}

func (r *Recipe) createRecipe(bd *sql.DB) error {
	//statement := fmt.Sprintf("INSERT INTO recipe(id,name,ingredients,preparation,time) VALUES(DEFAULT,'$1','$2','$3','$4') RETURNING id ", r.Name, r.Ingredients, r.Preparation, r.Time)
	/*
		_, err := bd.Exec(statement)
		if err != nil {
			return err
		}f
	*/
	err := bd.QueryRow("INSERT INTO recipe(id,name,ingredients,preparation,time) VALUES(DEFAULT,$1,$2,$3,$4) RETURNING id ", r.Name, r.Ingredients, r.Preparation, r.Time).Scan(&r.ID)
	if err != nil {
		return err
	}
	return nil // return a object
}

func (r *Recipe) updateRecipeSafe(bd *sql.DB) string {
	//statement := fmt.Sprintf("UPDATE recipe SET  name='%s', ingredients='%s',preparation='%s',time='%s' WHERE id='%d' RETURNING id", r.Name, r.Ingredients, r.Preparation, r.Time, r.ID)
	//_, err := bd.Exec(statement)

	//return err
	rows, err := bd.Query("UPDATE recipe SET  name=$1, ingredients=$2,preparation=$3,time=$4 WHERE id=$5", r.Name, r.Ingredients, r.Preparation, r.Time, r.ID)

	log.Println(rows)
	log.Println(err)

	return "ok"

}

func (r *Recipe) updateRecipe(bd *sql.DB) string {
	//statement := fmt.Sprintf("UPDATE recipe SET  name='%s', ingredients='%s',preparation='%s',time='%s' WHERE id='%d' RETURNING id", r.Name, r.Ingredients, r.Preparation, r.Time, r.ID)
	//_, err := bd.Exec(statement)

	//return err

	err := bd.QueryRow("UPDATE recipe SET  name=$1, ingredients=$2,preparation=$3,time=$4 WHERE id=$5 RETURNING id", r.Name, r.Ingredients, r.Preparation, r.Time, r.ID).Scan(&r.ID)
	if err != nil {
		return "Recipe id Not found"
	}

	return "ok"
}

func (r *Recipe) deleteRecipe(bd *sql.DB) string {
	//statement := fmt.Sprintf("DELETE FROM recipe WHERE id='%d'RETURNING id", r.ID)
	//_, err := bd.Exec(statement)
	//return err

	err := bd.QueryRow("DELETE FROM recipe WHERE id=$1 RETURNING id", r.ID).Scan(&r.ID)
	if err != nil {
		return "Recipe id Not found"
	}

	return "ok"
}
