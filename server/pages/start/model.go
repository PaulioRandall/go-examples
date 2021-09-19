package start

type dataModel struct {
	Title          string
	StartNowButton StartNowButton
}

type StartNowButton struct {
	Text string
	Link string
}

func makeDataModel(j journey) dataModel {
	return dataModel{
		Title: "Start",
		StartNowButton: StartNowButton{
			Text: "Start nows",
			Link: j.nextLink,
		},
	}
}
