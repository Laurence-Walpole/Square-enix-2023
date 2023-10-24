package main

import (
	"database/sql"
	"strconv"
	"time"
)

type Data struct {
	ID     int64
	InputA float64
	InputB float64
	Output float64
}

type Worker struct {
	ID          int64
	IP          string
	Status      int64
	LastUpdated string
	Created     string
}

type Job struct {
	ID          int64
	Worker      Worker
	_workerid   int64
	Data        Data
	_dataid     int64
	Calculation string
}

var self Worker

func sendQuery(query string, params []string) *sql.Rows {
	rows, err := db.Query(query, params)
	if err != nil {
		e.Logger.Fatal(err.Error())
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			e.Logger.Fatal(err.Error())
		}
	}(rows)
	return rows
}

func createWorkerAndGetAJob(ip string) Job {
	currentTime := getTimeNow()
	self.IP = ip
	self.Created = currentTime
	self.LastUpdated = currentTime
	self.Status = 1

	query := "INSERT INTO `workers` (`ip`, `status`, `last_updated`, `created`) VALUES (?, ?, ?, CURRENT_TIMESTAMP)"
	result, err := db.Exec(query, self.IP, strconv.FormatInt(self.Status, 10), self.LastUpdated)

	if err != nil {
		e.Logger.Fatal(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		e.Logger.Fatal(err.Error())
	}
	self.ID = id

	data := getNextData()
	return makeJob(self.ID, data.ID, makeCalculation(data))
}

func getJob(id int64) Job {
	query := "SELECT * FROM `jobs` WHERE `id` = ?"
	result := sendQuery(query, []string{strconv.FormatInt(id, 10)})
	var job Job
	for result.Next() {
		if err := result.Scan(&job.ID, &job._workerid, &job._dataid, &job.Calculation); err != nil {
			e.Logger.Fatal(err.Error())
		}
		job.Worker = getWorker(job._workerid)
		job.Data = getData(job._dataid)
	}
	return job
}

func getWorker(id int64) Worker {
	query := "SELECT * FROM `workers` WHERE `id` = ?"
	result := sendQuery(query, []string{strconv.FormatInt(id, 10)})
	var worker Worker
	for result.Next() {
		err := result.Scan(&worker.ID, &worker.IP, &worker.Status, &worker.LastUpdated, &worker.Created)
		if err != nil {
			return Worker{}
		}
	}
	return worker
}

func getData(id int64) Data {
	query := "SELECT * FROM `data` WHERE `id` = ?"
	result := sendQuery(query, []string{strconv.FormatInt(id, 10)})
	var data Data
	for result.Next() {
		err := result.Scan(&data.ID, &data.InputA, &data.InputB, &data.Output)
		if err != nil {
			return Data{}
		}
	}
	return data
}

func getNextData() Data {
	var data Data
	row := db.QueryRow("SELECT * FROM `data` WHERE `output` IS NULL AND `id` NOT IN (SELECT data_id FROM jobs) LIMIT 1")
	err := row.Scan(&data.ID, &data.InputA, &data.InputB, &data.Output)
	if err != nil {
		return Data{}
	}
	return data
}

func makeJob(worker_id int64, data_id int64, calculation string) Job {
	var job Job
	query := "INSERT INTO `jobs` (`worker_id`, `data_id`, `calculation`) VALUES (?, ?, ?)"
	result, err := db.Exec(query, worker_id, data_id, calculation)

	if err != nil {
		e.Logger.Fatal(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		e.Logger.Fatal(err.Error())
	}
	job.ID = id
	job._workerid = worker_id
	job._dataid = data_id
	job.Calculation = calculation
	job.Worker = getWorker(worker_id)
	job.Data = getData(data_id)
	return job
}

func makeCalculation(data Data) string {
	calc := strconv.FormatFloat(data.InputA, 'f', -1, 64) + "*" + strconv.FormatFloat(data.InputB, 'f', -1, 64)
	return calc
}

func getTimeNow() string {
	return time.Now().Format("2023-10-24 14:39:35")
}
