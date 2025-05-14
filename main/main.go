package main

import (
	"INTRODUCAO_BACKEND/conn"
	"fmt"
)

func main() {
	database, err:=conn.ConnectDatabase()

	if err!=nil{
		panic(err
	}

	query, _:=database.Prepare("INSERT INTO aluno(nome, curso) VALUES ($1, $2)")

	query.Exec("Murilo", "Enegnharia de Software")

	rows,_:=database.Query("SELECT id, nome, ncurso FROM aluno")

	for rows.Next(){
		var (
			nome, curso string
			id int
		)

		rows.Scan(
				&id,
				&nome,
				&curso,
		)

		fmt.Println("---------------------------------------------------------")
		fmt.Printf("ID: %d\nNome: %s\nCurso: %s\n", id, nome, curso)
	}

	updateQuery, _:=database.Prepare("DELETE FROM aluno WHERE id = $1")

	deleteQuery.Exec(2)

	fmt.Println("Aluno com ID 2 removido com sucesso.")

	deleteQuery, _:=database.Prepare("DELETE FROM aluno WHERE id = $1")

	deleteQuery.Exec(2)

	fmt.Println("Aluno com ID 2 removido com sucesso.")
}
