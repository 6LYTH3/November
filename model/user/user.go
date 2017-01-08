package model

import (
	"gopkg.in/mgo.v2"
	"log"
)

type User struct {
	Name   string
	Gender string
	Age    int
	Id     string
}

func connectDB() *mgo.Session {
	session, err := mgo.Dial("192.168.99.100:27017")
	if err != nil {
		log.Fatal(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}

func (u *User) Add() {
	log.Printf("[!] Add")
	s := connectDB()
	defer s.Close()

	c := s.DB("blog").C("user")
	err := c.Insert(u)
	if err != nil {
		log.Fatal(err)
	}
}

func (u *User) FindAll() User {
	s := connectDB()
	defer s.Close()

	c := s.DB("blog").C("user")
	var r []User
	err := c.Find(nil).All(&r)
	if err != nil {
		log.Fatal(err)
	}

	return r[0]
}
