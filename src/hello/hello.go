package main

import (
	"encoding/json"
	"fmt"
	"module/src/mypackage"
	"os"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var y = "strings" // Disponivel em todo o package - não da para usar ":="
var x int = 10
var k int

func main() {
	// zeros()
	// pacoteFMT()
	// myType()
	// convert()
	// booleanFunc()
	// numericFunc()
	// stringFunc()
	// numericSystem()
	// constants()
	// iotaType()
	// deslocBit()
	// loopfor()
	// nestedLoop()
	// customWhile()
	// forContinue()
	// conditionIf()
	// conditionSwitch()
	// switchType()
	// operateLogic()
	// firstArray()
	// firstSlices()
	// slicesToSlice()
	// appendSlice()
	// makeSlice()
	// sliceMultDimension()
	// subjacentSlice()
	// mapes()
	// maprange()
	// structtype()
	// structEmbut()
	// anonimoStruct()
	// deferer()
	chamando()

}

func anyThing(x int) {
	// - Usando a variavel [y] declarado fora do escopo
	fmt.Println(y)
	// - Usando a variavel [x] dentro da função main ( z := 20)
	fmt.Println(x)
}

// Variáveis, Valores e TIpos - Valor Zero

var a int
var b float64
var c string
var d bool

func zeros() {
	fmt.Printf("A: %v \n", a)
	fmt.Printf("B: %v \n", b)
	fmt.Printf("C: %v \n", c)
	fmt.Printf("D: %v", d)

}

// Pacote FMT

func pacoteFMT() {
	x := "Oi"
	y := ", Bom dia"
	fmt.Print("Oi", "bom\n")     // Retornar bytes sem espaçamento
	fmt.Println("Oi", "bom dia") // Retorna bytes com espaçamento

	z := fmt.Sprint(x, y)  // Retorna uma String sem espaçamento
	z = fmt.Sprintln(x, y) // Retorna uma String com espaçamento

	fmt.Println(z)
}

// Criar tipagems

type hotdog int

var p hotdog = 10

func myType() {
	numeroComp := 10
	fmt.Printf("%T", p)
	p = hotdog(numeroComp) // p = numeroComp da erro
}

// Converter e nao coerção

func convert() {
	a := int(p)
	fmt.Printf("%T - %v", a, a)
}

// Tipo Booleano

func booleanFunc() {
	var x bool
	fmt.Println(x)
	x = true

	x = (10 < 100)
	y := 10 == 100
	z := 100 > 10
	fmt.Print(x, y, z)
}

func numericFunc() {
	// var num1 uint
	// var num2 int
	// var num3 float32
	// var num4 complex64
	// var num5 byte
	// var num6 rune
	a := "e"
	b := "é"
	c := "色"

	fmt.Printf("%v, %v ,%v \n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v, %v ,%v", d, e, f)

}

func stringFunc() {
	s := `Hello, 
	playground`
	sb := []byte(s)
	fmt.Printf("%v \n%T", sb, sb)

	for _, v := range sb {
		fmt.Printf("%v - %T - %#U - %#X \n", v, v, v, v)
	}
}

func numericSystem() {
	a := 100
	fmt.Printf("%d\t%b\t%#x", a, a, a) //format verb
}

const l = 10 // Ele so se torna algo quando usado
const (
	ha = 10
	hb = 20
	hz = 30
)

var s int
var h float32

func constants() {
	s = l
	h = l
	fmt.Println(s)
	fmt.Println(ha, hb, hz)
}

//Iota

const (
	hx = iota + 1*10
	kz
	ka
)
const (
	ioa = iota
	iob = iota
	ioc = iota
)

// Usado quando não liga para os valores da constantes, so que seja diferente
// Tipo enum

func iotaType() {
	fmt.Println(hx, kz, ka)
}

// Deslocamento de bits

func deslocBit() {
	x := 1
	y := x << 1
	z := x >> 1
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)

}

func loopfor() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

}

func nestedLoop() {
	for horas := 0; horas <= 24; horas++ {
		fmt.Print("Horas:", horas, "\n")
		for min := 0; min < 60; min++ {
			fmt.Print("Minutos: ", min, "\n")
			for segundos := 0; segundos < 60; segundos++ {
				fmt.Print("Segundos: ", segundos, "\n")
			}
		}
	}
}

func customWhile() {
	x := 0
	for x < 10 {
		fmt.Println("Esse while")
		x++
	}
	//Loop infinito
	for {
		fmt.Println("Loop infinito")
		if x < 10 {
			x++
			fmt.Println("Xis")
		} else {
			fmt.Println("Xiss ")
			break
		}
	}
}

func forContinue() {
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)

	}
}

func conditionIf() {
	x := 10
	if x < 100 {
		fmt.Println("Hello")
	}
	if !(x < 100) {
		fmt.Println("Hello is false")
	} else if x < 10 {
		fmt.Println("Hello is true")
	} else {
		fmt.Println("Hello is ...")
	}
}

func conditionSwitch() {
	x := 5
	y := 10
	switch true {
	case x == 5, y > 10:
		fmt.Println("xis igual a cinco")
		fallthrough
	case x == 3, y < 3:
		fmt.Println("XIS IGUAL A TRES")
	case x > 5:
		fmt.Println("xis e maior que cinco")
	default:
		fmt.Println("ta vazio")
	}

}

var ty interface{}

func switchType() {
	ty = true
	switch ty.(type) {
	case int:
		fmt.Println("E int")
	case bool:
		fmt.Println("E bool")
	case string:
		fmt.Println("E string")

	}
}

func operateLogic() {
	xis := 3
	if xis == 2 || xis == 3 {

		fmt.Println(xis)
	}

	if xis%2 == 0 && xis%3 == 2 {
		fmt.Println("É multiplo de dois e treis")
	}
}

func firstArray() {
	var x [5]int
	x[0] = 1
	x[1] = 10
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(len(x))

}

func firstSlices() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	array2 := append(array[:], 6)
	slice2 := append(slice, 6)
	fmt.Println(array2)
	fmt.Println(slice2)
	slice2[2] = 1234
	fmt.Println(slice2)

	sliceStr := []string{"banan", "maca", "jaca"}
	for idx, valor := range sliceStr {
		fmt.Println("Idx", idx, "valor", valor)
	}

	for _, valor := range sliceStr {
		fmt.Println(valor)
	}

}

func slicesToSlice() {
	sabores := []string{"pepperoni", "mozzarela", "marguerita"}

	fatia := sabores[0:2]
	fmt.Println(fatia)

	fmt.Println(sabores[:])

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}
}

func appendSlice() {
	umslice := []int{1, 2, 3, 4}
	doislice := []int{9, 10, 11, 12}

	umslice = append(umslice, 5, 6, 7, 8)
	fmt.Println(doislice, umslice)
	umslice = append(umslice, doislice...)
	fmt.Println(umslice)
}

func makeSlice() {
	makeslice := make([]int, 5, 10)

	makeslice[0], makeslice[1] = 1, 2

	fmt.Println(makeslice, len(makeslice), cap(makeslice))
}

func sliceMultDimension() {
	ss := [][]int{
		{1, 2, 3, 4},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	sss := [][][]int{
		{
			{1, 2, 3, 4},
			{4, 5, 6, 7},
			{8, 9, 10, 11},
		},
		{
			{1, 2, 3, 4},
			{4, 5, 6, 7},
			{8, 9, 10, 11},
		},
		{
			{1, 2, 3, 4},
			{4, 5, 6, 7},
			{8, 9, 10, 11},
		},
	}
	fmt.Println(ss[1][1])
	fmt.Println(sss[0][1][0])
}

func subjacentSlice() {
	primeiroSlice := []int{1, 2, 3, 4, 5}
	primeiroSlice = append(primeiroSlice[:2], primeiroSlice[4:]...)
	fmt.Println(primeiroSlice)
}

func mapes() {
	amigos := map[string]int{
		"alfredo": 555555,
		"joaona":  131393193,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	amigos["gopher"] = 44444
	fmt.Println(amigos)
	if existe, naoexiste := amigos["romario"]; !naoexiste {
		fmt.Println("NAO EXISTE")
	} else {

		fmt.Println(existe)
	}

}

func maprange() {
	qualquercoisa := map[int]string{
		123:  "muitoLegal",
		98:   "menos legal um pouco",
		9393: "essa e massa",
	}
	fmt.Println(qualquercoisa)
	total := 0
	for key := range qualquercoisa {
		total += key
	}
	fmt.Println(total)
	delete(qualquercoisa, 123)
	fmt.Println(qualquercoisa)
}

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func structtype() {
	client1 := cliente{
		nome:      "Ronald",
		sobrenome: "Luc",
		fumante:   true,
	}
	client2 := cliente{
		"Joana",
		"Darc",
		true,
	}

	fmt.Println(client1, client2)
}

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func structEmbut() {
	pessoa1 := pessoa{
		nome:  "Ronald",
		idade: 30,
	}
	pessoa2 := profissional{
		pessoa:  pessoa{nome: "Maria", idade: 40},
		titulo:  "Maricota",
		salario: 4000,
	}

	fmt.Println(pessoa1, pessoa2)
}

func anonimoStruct() {
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Mario",
		idade: 50,
	}
	fmt.Println(x)
}

func nome() int {
	meuint := 30
	return soma(meuint, 20)
}

func soma(xis, y int) int {
	return xis + y
}

func enumslices() {
	si := []int{10, 10, 1, 2, 3, 4}

	total, _ := somandos(si...)
	fmt.Println(total)
}

func somandos(xises ...int) (int, int) {
	soma := 0
	for _, v := range xises {
		soma += v
	}
	return soma, len(xises)
}

// Executa todo o codigo acima
func deferer() {
	defer fmt.Println("com defer, veio primeiro")
	fmt.Println("sem defer, veio depois")

}

type personal struct {
	nome  string
	idade int
}

func chamando() {
	mauricio := personal{"Mauricio", 30}

	mauricio.bomdia()
}

func (p personal) bomdia() {
	fmt.Println(p.nome, "diz bom dia")
}

type person1 struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	person1
	dentesarrancados int
	salario          float64
}

type arquiteto struct {
	person1
	tipodeconstru    string
	tamanhodaloucura string
}

func (x arquiteto) oibomdiaa() {
	fmt.Println("Meu nome é", x.nome, "e ouve bom dia")
}
func (x dentista) oibomdiaa() {
	fmt.Println("Meu nome é", x.nome, "e ouve bom dia", x.dentesarrancados)
}

type gente interface {
	oibomdiaa()
}

func serhumano(g gente) {
	g.oibomdiaa()
	switch g.(type) {
	case dentista:
		fmt.Println(g.(dentista).salario)
	case arquiteto:
		fmt.Println(g.(arquiteto).sobrenome)
	}
}

func implement() {
	mrdente := dentista{
		person1: person1{
			nome:      "Lau",
			idade:     30,
			sobrenome: "aaka",
		},
		dentesarrancados: 10,
		salario:          3000,
	}
	mrpredio := arquiteto{
		person1: person1{
			nome:      "Vlad",
			sobrenome: "Caus",
			idade:     40,
		},
		tipodeconstru:    "Orizon",
		tamanhodaloucura: "crazy",
	}

	serhumano(mrdente)
	serhumano(mrpredio)
}

func funcaoAnonima() {
	xz := 10

	func(xz int) {
		fmt.Println(xz * 10)
	}(xz)

}

func expression() {
	xis := 10
	yps := func(x int) int {
		return x * 89999
	}

	fmt.Println(xis, "vezes", yps(xis))
}

func retornaFunc() func(int) int {
	return func(ix int) int {
		return ix * 10

	}
}

func chamandoe() {
	xk := retornaFunc()
	xl := xk(3)
	xp := retornaFunc()(3)

	fmt.Println(xl, xp)
}

func closures() {
	a := clo()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

}

func clo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func recursive(x int) int {
	if x == 1 {
		return x
	}

	return x * recursive(x-1)
}

func points() {
	y := &x
	fmt.Println(*y)
}

func ponteiros() {
	x := 11
	estarecebeumponteir(&x)
}

func estarecebeumponteir(x *int) {
	*x++
	fmt.Println("NA FUN", *x)

}

func jsonTratamento() {
	type pessoa struct {
		nome      string
		sobrenome string
	}

	jamesbond := pessoa{
		nome:      "james",
		sobrenome: "bond",
	}

	b, _ := json.Marshal(jamesbond)

	fmt.Println(string(b))
}

func enconder() {
	type pessoa struct {
		nome      string
		sobrenome string
	}
	jamesbond := pessoa{
		nome:      "james",
		sobrenome: "bond",
	}

	enconder := json.NewEncoder(os.Stdout)
	enconder.Encode(jamesbond)
}

func sortStrings() {
	ss := []string{"abobora", "omelete", "cafe", "xuxu", "leite"}
	sort.Strings(ss)
	fmt.Println(ss)
}

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenarPorPotencia []carro

func (a ordenarPorPotencia) Len() int           { return len(a) }
func (a ordenarPorPotencia) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ordenarPorPotencia) Less(i, j int) bool { return a[i].potencia < a[j].potencia }

type ordenarPorConsumo []carro

func (a ordenarPorConsumo) Len() int           { return len(a) }
func (a ordenarPorConsumo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ordenarPorConsumo) Less(i, j int) bool { return a[i].consumo < a[j].consumo }

type ordenarPorLucroDono []carro

func (a ordenarPorLucroDono) Len() int           { return len(a) }
func (a ordenarPorLucroDono) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ordenarPorLucroDono) Less(i, j int) bool { return a[i].nome < a[j].nome }

func mySort() {
	carros := []carro{{"Chevete", 50, 5}, {"Fiat", 50, 4}, {"Fusca", 5, 30}}

	sort.Sort(ordenarPorPotencia(carros))
	fmt.Println(carros)
}

func senha() {
	senha := "3001rodrigo"
	senhaerrada := "3002rodrigo"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sb))

	bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada))

	fmt.Println(sb)

}

var wg = sync.WaitGroup{}

func concorrent() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	wg.Add(2)

	go func1()
	go func2()

	fmt.Println(runtime.NumGoroutine())
	wg.Wait()

}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(20)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, 'n')
		time.Sleep(20)
	}
	wg.Done()
}

func concorrent2() {
	var mu sync.Mutex
	var contador int64
	contador = 0
	totaldegoroutines := 10
	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			mu.Lock()
			v := contador
			fmt.Println("Contador: ", atomic.LoadInt64(&contador))
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(contador)

}

func chamandoOutropkg() {

	mypackage.Printando()
}

// Doc https://rakyll.org/style-packages/

func canal() {
	canal := make(chan int)
	go func() {
		canal <- 42
	}()
	fmt.Println(<-canal)
}

func canal2() {
	canal := make(chan int)

	go send(canal)
	receive(canal)
}

func send(s chan<- int) {
	s <- 42

}
func receive(r <-chan int) {
	fmt.Println("O valor recebido do canal foi:", <-r)
}

func chanelRange() {
	c := make(chan int)
	go meuloop(10, c)

	prints(c)
}
func meuloop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s)
}

func prints(r <-chan int) {
	for v := range r {
		fmt.Println("Recebido no canal: ", v)

	}
}
