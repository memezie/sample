package models

type Agent struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (m Agent) GetAgents() (agents []Agent, err error) {
	agents = []Agent{
		{Id: 1, Name: "エージェント1"},
		{Id: 2, Name: "エージェント2"},
	}
	return agents, nil
}

func (m *Agent) GetAgent() {
	m.Name = "エージェント1"
}
