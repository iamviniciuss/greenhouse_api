package domain

type Greenhouse struct {
	ID        string      `json:"_id" bson:"_id"`
	Name      string      `json:"name" bson:"_name"`
	Sensors   []*Sensor   `json:"sensors"`
	Actuators []*Actuator `json:"actuatorS"`
}

type Sensor struct {
	ID            string        `json:"_id" bson:"_id"`
	Envoironments *Envoironment `json:"envoironment_id"`
	Actuator      *Actuator     `json:"actuator"`
	Name          string        `json:"name" bson:"_name"`
	GreenhouseID  string        `json:"greenhouse_id"`
	IdealValue    []int         `json:"ideal_value" bson:"ideal_value"`
}

type Actuator struct {
	ID            string        `json:"_id" bson:"_id"`
	Name          string        `json:"name" bson:"_name"`
	Envoironments *Envoironment `json:"envoironment_id"`
	Sensor        *Sensor       `json:"sensor"`
	GreenhouseID  string        `json:"greenhouse_id"`
}

type Envoironment struct {
	ID           string `json:"_id" bson:"_id"`
	Name         string `json:"name" bson:"_name"`
	GreenhouseID string `json:"greenhouse_id"`
}
