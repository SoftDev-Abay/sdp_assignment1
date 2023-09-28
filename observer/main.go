package main

import "fmt"

// "go/build"

type Article struct {
	title  string
	author string
}

type Subscriber struct {
	name     string
	id       int
	articles []Article
}

func (s *Subscriber) Update(updatedArticles []Article) {

	fmt.Printf("Updated subscriber : %v \n", s.name)
	s.articles = updatedArticles
	fmt.Printf("New articles : %v \n", s.articles)
}

type Publisher struct {
	company_title string
	articles      []Article
	subscribers   []Subscriber
}

func (p *Publisher) AddSubscriber(s Subscriber) {
	p.subscribers = append(p.subscribers, s)
}
func (p *Publisher) removeSubscriber(id int) {
	newSubscribers := []Subscriber{}

	for _, s := range p.subscribers {
		if s.id != id {
			newSubscribers = append(p.subscribers, s)
		}
	}
	p.subscribers = newSubscribers

	p.NotifySubscribers()
}

func (p *Publisher) addArcticle(a Article) {
	p.articles = append(p.articles, a)
	p.NotifySubscribers()
}
func (p *Publisher) removeArcticle(title string) {
	newArticles := []Article{}

	for _, a := range p.articles {
		if a.title != title {
			newArticles = append(p.articles, a)
		}
	}
	p.articles = newArticles

	p.NotifySubscribers()
}

func (p *Publisher) NotifySubscribers() {
	for _, s := range p.subscribers {
		// s.Update(p.articles)
		s.Update(p.articles)
	}
}

func main() {
	subscriber1 := Subscriber{name: "Bob", id: 1}
	subscriber2 := Subscriber{name: "Jerry", id: 2}
	subscriber3 := Subscriber{name: "Tom", id: 3}
	publisher := Publisher{company_title: "New York Times", subscribers: []Subscriber{subscriber1, subscriber2, subscriber3}}
	publisher.addArcticle(Article{"Modern programming", "Bob Smith"})
	publisher.removeArcticle("Modern programming")
}
