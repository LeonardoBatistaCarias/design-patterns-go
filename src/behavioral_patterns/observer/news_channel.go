package main

type NewsChannel struct {
	news string
}

func (n *NewsChannel) Update(news string) {
	n.SetNews(news)
}

func (n *NewsChannel) SetNews(news string) {
	n.news = news
}

func (n *NewsChannel) GetNews() string {
	return n.news
}
