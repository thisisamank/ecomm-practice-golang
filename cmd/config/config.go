package main

type application struct {
	config config
}

type config struct {
	addr   string
	dbHost string
}
