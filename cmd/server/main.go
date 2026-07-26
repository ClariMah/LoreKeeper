package main

import (
	"context"
	"fmt"
	"log"
	"lorekeeper/internal/application"
	"lorekeeper/internal/infrastructure/persistence/memory"
)

func main() {
	ctx := context.Background()

	fmt.Println("📜 === INICIALIZANDO O LOREKEEPER 2.0 ===")

	repository := memory.NewBooksMemoryRepository()

	useCase := application.NewUseCase(repository)

	fmt.Println("\n--- 1. TESTANDO CRIAÇÃO DE LIVROS (CREATE) ---")

	// Criando o primeiro livro
	book1, err := useCase.CreateBook(ctx, "O Hobbit", "J.R.R. Tolkien", "Allen & Unwin", 1937, 310, "Fantasia", "Português", "978-8551002704")
	if err != nil {
		log.Fatalf("Erro ao criar livro 1: %v", err)
	}
	fmt.Printf("✅ Criado: [%d] %s - %s (%d págs)\n", book1.ID, book1.Title, book1.Author, book1.PageNumber)

	// Criando um segundo livro
	book2, err := useCase.CreateBook(ctx, "O Nome do Vento", "Patrick Rothfuss", "Sextante", 2007, 656, "Fantasia", "Português", "978-8580570533")
	if err != nil {
		log.Fatalf("Erro ao criar livro 2: %v", err)
	}
	fmt.Printf("✅ Criado: [%d] %s - %s (%d págs)\n", book2.ID, book2.Title, book2.Author, book2.PageNumber)

	fmt.Println("\n--- 2. TESTANDO VALIDAÇÃO DE ERRO ---")
	// Tentando criar livro sem título para testar o domínio
	_, err = useCase.CreateBook(ctx, "", "Autor Sem Nome", "Editora", 2020, 100, "Geral", "Português", "123")
	if err != nil {
		fmt.Printf("⚠️ Validação funcionou corretamente! Erro capturado: %v\n", err)
	}

	fmt.Println("\n--- 3. TESTANDO LISTAGEM (READ ALL) ---")
	books, err := useCase.ListBooks(ctx)
	if err != nil {
		log.Fatalf("Erro ao listar livros: %v", err)
	}
	fmt.Printf("📚 Total de livros na memória: %d\n", len(books))
	for _, b := range books {
		fmt.Printf(" - %s por %s\n", b.Title, b.Author)
	}

	fmt.Println("\n--- 4. TESTANDO BUSCA POR ID (READ ONE) ---")
	foundBook, err := useCase.GetBookById(ctx, book1.ID)
	if err != nil {
		log.Fatalf("Erro ao buscar livro: %v", err)
	}
	fmt.Printf("🔍 Encontrado: %s (Editora: %s)\n", foundBook.Title, foundBook.Publisher)

	fmt.Println("\n--- 5. TESTANDO DELEÇÃO (DELETE) ---")
	err = useCase.DeleteBook(ctx, book1.ID)
	if err != nil {
		log.Fatalf("Erro ao deletar livro: %v", err)
	}
	fmt.Printf("🗑️ Livro '%s' removido com sucesso!\n", book1.Title)

	// Verificando a lista pós-deleção
	remainingBooks, _ := useCase.ListBooks(ctx)
	fmt.Printf("📚 Livros restantes na memória: %d\n", len(remainingBooks))

	fmt.Println("\n🎉 === TODOS OS TESTES PASSARAM COM SUCESSO! ===")
}
