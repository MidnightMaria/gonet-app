package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type PC struct {
	pcNumber int
	isUsed    bool
	mu        sync.Mutex
}

type User struct {
	username    string
	billingHour int
}

type Billing struct {
	hour  int
	price int
}

var salary = 0
var mu sync.Mutex // Mutex for salary updates
var pcMutex sync.Mutex // Mutex for PC allocation

var users = []User{
	{username: "nesngenes", billingHour: 0},
	{username: "anne_maria_silva", billingHour: 0},
	{username: "keira_knightley", billingHour: 0},
	{username: "akane_akari", billingHour: 0},
	{username: "reina_wolfreign", billingHour: 0},
	{username: "chloe_alberta_elodie", billingHour: 0},
	{username: "hima", billingHour: 0},
	{username: "ceza", billingHour: 0},
	{username: "zize", billingHour: 0},
	{username: "emi", billingHour: 0},
}

var pcs = []PC{
	{pcNumber: 1, isUsed: false, mu: sync.Mutex{}},
	{pcNumber: 2, isUsed: false, mu: sync.Mutex{}},
	{pcNumber: 3, isUsed: false, mu: sync.Mutex{}},
	{pcNumber: 4, isUsed: false, mu: sync.Mutex{}},
	{pcNumber: 5, isUsed: false, mu: sync.Mutex{}},
}

var billing = []Billing{
	{hour: 1, price: 3000},
	{hour: 2, price: 5000},
	{hour: 3, price: 7000},
	{hour: 4, price: 9000},
	{hour: 5, price: 10000},
}

func getRandomPC(pcs []PC) *PC {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	pcMutex.Lock()
	defer pcMutex.Unlock()

	for {
		randPcNumber := rand.Intn(len(pcs))
		selectedPC := &pcs[randPcNumber]

		selectedPC.mu.Lock()
		if !selectedPC.isUsed {
			// Mark the selected PC as used before returning
			selectedPC.isUsed = true
			selectedPC.mu.Unlock()
			return selectedPC
		}
		selectedPC.mu.Unlock()
	}
}

func buyBilling() (billingHour, billingPrice int) {
	randomIndex := rand.Intn(len(billing))
	billingHour = billing[randomIndex].hour
	billingPrice = billing[randomIndex].price

	return billingHour, billingPrice
}

func gonet() int {
	var wg sync.WaitGroup
	wg.Add(len(users))

	for i := 0; i < len(users); i++ {
		selectedPC := getRandomPC(pcs)
		go gonetLogic(&users[i], selectedPC, &wg)
	}

	wg.Wait()
	return salary
}

func gonetLogic(user *User, selectedPC *PC, wg *sync.WaitGroup) {
	defer wg.Done()

	billingHour, billingPrice := buyBilling()
	user.billingHour += billingHour

	fmt.Printf("%s buy %d hour billing for %d rupiah \n", user.username, billingHour, billingPrice)
	fmt.Printf("%s is now using pc %d\n", user.username, selectedPC.pcNumber)

	playingTime := time.Duration(user.billingHour) * time.Second

	mu.Lock()
	time.Sleep(playingTime)
	mu.Unlock()

	fmt.Printf("%s billing is over\n", user.username)

	fmt.Printf("%s left pc %d\n", user.username, selectedPC.pcNumber)

	// Use mutex for concurrent updates to salary
	mu.Lock()
	salary += billingPrice
	mu.Unlock()

	// Mark the PC as unused after the user leaves
	selectedPC.mu.Lock()
	selectedPC.isUsed = false
	selectedPC.mu.Unlock()
}

func main() {
	fmt.Println("......:::::: Welcome to gonet <3 ::::::......")

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	totalSalary := gonet()

	fmt.Printf("Our net is closed. total salary for today is %d\n", totalSalary)
}
