package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const num_monitoramentos = 1
const delay = 5

func main() {

	view_info()

	// loop infinito sem parametros
	for {
		view_menu()
		comando := read_comando()

		switch comando {
		case 1:
			fmt.Println("[1] Monitorando ...")
			init_monitor()
		case 2:
			fmt.Println("[2] Exibindo Logs")
			print_logs()
		case 0:
			fmt.Println("[0] Volte sempre !!!")
			os.Exit(0)
		default:
			fmt.Println("[?] Comando desconhecido ?")
			os.Exit(-1) // sair com err: exit status 255
		}
	}
}

func view_info() {
	nome := "Alex"
	cmd := exec.Command("bash", "-c", "go version")
	version, err := cmd.Output()
	if err != nil {
		fmt.Println((err.Error()))
		return
	}

	fmt.Println("Hola ", nome, "! ", nome, "es variable de tipo ", reflect.TypeOf(nome))
	fmt.Println("Voce esta usando a GO ", string(version))
}

func view_menu() {
	// menu
	fmt.Println("+", strings.Repeat("-", 25), "+")
	fmt.Println("| 1-Iniciar Monitoramento   |")
	fmt.Println("| 2-Exibir Logs             |")
	fmt.Println("| 0-Sair do Programa        |")
	fmt.Println("+", strings.Repeat("-", 25), "+")
}

func read_comando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Voce escolheu a opcao: ", comando, "-> com endereco: ", &comando)
	return comando
}

func init_monitor() {
	fmt.Println("Monitorando...")

	sites := read_file()

	// teste en loop
	for i := 0; i < num_monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Monitorando site: [", i, "] -> ", site)
			ping_url(site)
		}
		// adicionando delay cada 10 seg
		time.Sleep(delay * time.Second)
	}
	fmt.Println("")
}

func ping_url(url string) {
	result, err := http.Get(url)
	if err != nil {
		fmt.Println("[error] (ping_url):", err)
	}

	if result.StatusCode == 200 {
		fmt.Println("Site: ", url, "status: Run")
		register_logs(url, true)
	} else {
		fmt.Println("Site: ", url, "status: erro: ???")
		register_logs(url, false)
	}
}

func read_file() []string {

	var sites []string
	file_names := "sites.txt"

	file3, err := os.Open(file_names)
	if err != nil {
		fmt.Println("[error 3.1] (read_file):", err)
	}

	leitor := bufio.NewReader(file3)
	for {
		linha, err := leitor.ReadString('\n') // ler ate o delimitador da linha com '' simples
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break // quando encontre o final sair do loop
		}
	}
	file3.Close()

	return sites
}

func register_logs(site string, status bool) {
	file_name := "logs.log"
	// using Openfile with flags and premission
	file, err := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	// escrebendo no arquivo verificar documentação para time
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + " -> online: (" + strconv.FormatBool(status) + ")\n")

	file.Close()
}

func print_logs() {

	file, err := ioutil.ReadFile("logs.log")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(file))
}
