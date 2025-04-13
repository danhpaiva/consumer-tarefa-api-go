package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "http://localhost:8081"

type Tarefa struct {
	ID        int    `json:"id"`
	Descricao string `json:"descricao"`
}

func main() {
	executarExemplos()
}

func executarExemplos() {
	// Listar tarefas
	fmt.Println("Listando tarefas:")
	listarTarefas()

	// Criar uma nova tarefa
	fmt.Println("Criando uma nova tarefa:")
	novaTarefa := Tarefa{Descricao: "Estudar Go"}
	criarTarefa(novaTarefa)

	// Atualizar uma tarefa existente
	fmt.Println("Atualizando uma tarefa existente:")
	tarefaAtualizada := Tarefa{ID: 1, Descricao: "Estudar Go e criar APIs"}
	atualizarTarefa(tarefaAtualizada)

	// Excluir uma tarefa
	fmt.Println("Excluindo uma tarefa:")
	excluirTarefa(1)
}

func listarTarefas() {
	response, err := http.Get(baseURL + "/tarefas")
	if err != nil {
		fmt.Println("Erro ao listar tarefas:", err)
		return
	}
	defer response.Body.Close()

	var tarefas []Tarefa
	if err := json.NewDecoder(response.Body).Decode(&tarefas); err != nil {
		fmt.Println("Erro ao decodificar resposta:", err)
		return
	}

	for _, tarefa := range tarefas {
		fmt.Println(tarefa)
	}
}

func criarTarefa(tarefa Tarefa) {
	jsonData, err := json.Marshal(tarefa)
	if err != nil {
		fmt.Println("Erro ao serializar tarefa:", err)
		return
	}

	response, err := http.Post(baseURL+"/tarefas", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erro ao criar tarefa:", err)
		return
	}
	defer response.Body.Close()

	var novaTarefa Tarefa
	if err := json.NewDecoder(response.Body).Decode(&novaTarefa); err != nil {
		fmt.Println("Erro ao decodificar resposta:", err)
		return
	}

	fmt.Println("Nova tarefa criada:", novaTarefa)
}

func atualizarTarefa(tarefa Tarefa) {
	jsonData, err := json.Marshal(tarefa)
	if err != nil {
		fmt.Println("Erro ao serializar tarefa:", err)
		return
	}

	request, err := http.NewRequest("PUT", baseURL+"/tarefas/"+fmt.Sprint(tarefa.ID), bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Erro ao atualizar tarefa:", err)
		return
	}
	defer response.Body.Close()

	var tarefaAtualizada Tarefa
	if err := json.NewDecoder(response.Body).Decode(&tarefaAtualizada); err != nil {
		fmt.Println("Erro ao decodificar resposta:", err)
		return
	}

	fmt.Println("Tarefa atualizada:", tarefaAtualizada)
}

func excluirTarefa(id int) {
	request, err := http.NewRequest("DELETE", baseURL+"/tarefas/"+fmt.Sprint(id), nil)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Erro ao excluir tarefa:", err)
		return
	}
	defer response.Body.Close()

	fmt.Println("Tarefa excluída com sucesso")
}
