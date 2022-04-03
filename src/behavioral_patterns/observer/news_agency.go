package main

type NewsAgency struct {
	news      string
	observers []Channel
}

func (n *NewsAgency) AddObserver(channel Channel) {
	n.observers = append(n.observers, channel)
}

func (n *NewsAgency) RemoveObserver(channel Channel) {
	n.observers = removeFromSlice(n.observers, channel)
}

func removeFromSlice(observerList []Channel, observerToRemove Channel) []Channel {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observer == observerToRemove {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

func (n *NewsAgency) SetNews(news string) {
	n.news = news
	for _, observer := range n.observers {
		observer.Update(news)
	}
}
