package servidor

import (
	"crud-basico/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
	}

	usuario := usuario{}
	if err = json.Unmarshal(body, &usuario); err != nil {
		w.Write([]byte("Erro ao converter o corpo da requisição"))
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")
	if err != nil {
		w.Write([]byte("Erro ao criar o usuario"))
	}
	defer stmt.Close()
	insert, err := stmt.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar o statement"))
	}

	idInserido, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o id"))
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! Id: %d", idInserido)))

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
	}
	defer db.Close()

	linhas, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		w.Write([]byte("Erro ao buscar os usuarios"))
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario
		if err = linhas.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuario"))
		}
		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(usuarios); err != nil {
		w.Write([]byte("Erro ao converter o usuario para JSON"))
	}

}

func BuscarUsuarioPorId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o id"))
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
	}
	defer db.Close()

	linha, err := db.Query("SELECT * FROM usuarios WHERE id = ?", usuarioId)
	if err != nil {
		w.Write([]byte("Erro ao buscar o usuario"))
	}
	defer linha.Close()

	var usuario usuario
	if linha.Next() {
		if err = linha.Scan(&usuario.Id, &usuario.Nome, &usuario.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuario"))
		}
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(usuario); err != nil {
			w.Write([]byte("Erro ao converter o usuario para JSON"))
		}
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o id"))
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
	}

	usuario := usuario{}
	if err = json.Unmarshal(body, &usuario); err != nil {
		w.Write([]byte("Erro ao converter o corpo da requisição"))
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")
	if err != nil {
		w.Write([]byte("Erro ao atualizar o usuario"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(usuario.Nome, usuario.Email, usuarioId); err != nil {
		w.Write([]byte("Erro ao executar o statement"))
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o id"))
	}

	db, err := banco.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if err != nil {
		w.Write([]byte("Erro ao deletar o usuario"))
	}
	defer stmt.Close()

	if _, err = stmt.Exec(usuarioId); err != nil {
		w.Write([]byte("Erro ao executar o statement"))
	}

	w.WriteHeader(http.StatusNoContent)
}
