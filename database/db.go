package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Conectar() {
	var err error
	DB, err = sql.Open("sqlite3", "./produtos.db")
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	// TESTE A CONEXÃO:
	if err = DB.Ping(); err != nil {
		log.Fatal("Erro ao fazer ping no banco:", err)
	}

	log.Println("✅ Banco conectado com sucesso!")

	criarTabela()
	adicionarDadosTeste()
}

func criarTabela() {
	query := `
    CREATE TABLE IF NOT EXISTS produtos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        nome TEXT,
        preco REAL,
        quantidade INTEGER,
        descricao TEXT
    );
    `
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}
	log.Println("✅ Tabela criada/verificada com sucesso!")
}

func adicionarDadosTeste() {
	// Verifica se já tem dados
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM produtos").Scan(&count)
	if err != nil {
		log.Fatal("Erro ao contar produtos:", err)
	}

	// Se não tem dados, adiciona alguns de teste
	if count == 0 {
		produtos := []struct {
			nome, descricao string
			preco           float64
			quantidade      int
		}{
			{"Mouse Gamer", "Mouse com iluminação RGB", 199.90, 10},
			{"PC Gamer", "Computador para jogos", 1199.90, 5},
			{"Notebook", "Laptop para trabalho", 4500.99, 3},
			{"Mouse Gamer Pro", "Mouse profissional", 249.90, 7},
		}

		for _, p := range produtos {
			_, err := DB.Exec("INSERT INTO produtos (nome, preco, quantidade, descricao) VALUES (?, ?, ?, ?)",
				p.nome, p.preco, p.quantidade, p.descricao)
			if err != nil {
				log.Fatal("Erro ao inserir produto de teste:", err)
			}
		}
		log.Println("✅ Dados de teste adicionados!")
	} else {
		log.Printf("✅ Banco já tem %d produtos", count)
	}
}
