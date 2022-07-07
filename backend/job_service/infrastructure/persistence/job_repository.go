package persistence

import (
	"strconv"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/velibor7/XML/job_service/domain"
)

type JobNeo4j struct {
	Driver *neo4j.Driver
}

func NewJobNeo4j(driver *neo4j.Driver) domain.JobInterface {
	return &JobNeo4j{
		Driver: driver,
	}
}

func (db *JobNeo4j) Get(id int) (*domain.Job, error) {
	session := (*db.Driver).NewSession(neo4j.SessionConfig{})

	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)

	job, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			"MATCH (job:Job) WHERE id(job)=$id "+
				"RETURN id(job) AS id, job.title AS title, job.description AS description,"+
				" job.skills AS skills, job.userId AS userId",
			map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		var job *domain.Job
		for result.Next() {
			var lista []string

			id, _ := result.Record().Get("id")
			convId := id.(int64)
			Id := strconv.Itoa(int(convId))
			title, _ := result.Record().Get("title")
			description, _ := result.Record().Get("description")
			userId, _ := result.Record().Get("userId")
			skills, _ := result.Record().Get("skills")
			for _, skill := range skills.([]interface{}) {
				lista = append(lista, skill.(string))
			}
			job = &domain.Job{
				Id:          Id,
				Title:       title.(string),
				Description: description.(string),
				Skills:      lista,
				UserId:      userId.(string),
			}
		}
		return job, err
	})
	if err != nil {
		return nil, err
	}
	return job.(*domain.Job), nil
}

func (db *JobNeo4j) GetAll() ([]*domain.Job, error) {
	session := (*db.Driver).NewSession(neo4j.SessionConfig{})

	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)
	jobs, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			"MATCH (job:Job)"+
				"RETURN id(job) AS id, job.title AS title, job.description AS description, "+
				"job.skills AS skills, job.userId AS userId",
			nil)
		if err != nil {
			return nil, err
		}
		var jobs []*domain.Job
		for result.Next() {
			var lista []string

			id, _ := result.Record().Get("id")
			convId := id.(int64)
			Id := strconv.Itoa(int(convId))
			title, _ := result.Record().Get("title")
			description, _ := result.Record().Get("description")
			userId, _ := result.Record().Get("userId")
			skills, _ := result.Record().Get("skills")
			for _, skill := range skills.([]interface{}) {
				lista = append(lista, skill.(string))
			}

			jobs = append(jobs, &domain.Job{
				Id:          Id,
				Title:       title.(string),
				Description: description.(string),
				Skills:      lista,
				UserId:      userId.(string),
			})
		}
		return jobs, err
	})
	if err != nil {
		return nil, err
	}
	return jobs.([]*domain.Job), nil
}

func (db *JobNeo4j) Create(job *domain.Job) error {
	session := (*db.Driver).NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)
	err := CreateNode(job, &session)

	err = CreateRelationShips(&session)

	err = DeleteRelationShip(&session)

	if err != nil {
		return err
	}

	return nil
}

func (db *JobNeo4j) GetRecommendedJobs(job *domain.Job) ([]*domain.Job, error) {

	session := (*db.Driver).NewSession(neo4j.SessionConfig{})

	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)

	jobs, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			"MATCH (job1:Job {title:$title})-[:SAME]->(job2:Job) "+
				"RETURN id(job2) AS id, job2.title AS title, job2.description AS description, "+
				"job2.skills AS skills, job2.userId AS userId",
			map[string]interface{}{"title": job.Title})
		if err != nil {
			return nil, err
		}
		var jobs []*domain.Job
		for result.Next() {
			var lista []string

			id, _ := result.Record().Get("id")
			convId := id.(int64)
			Id := strconv.Itoa(int(convId))
			title, _ := result.Record().Get("title")
			description, _ := result.Record().Get("description")
			userId, _ := result.Record().Get("userId")
			skills, _ := result.Record().Get("skills")
			for _, skill := range skills.([]interface{}) {
				lista = append(lista, skill.(string))
			}

			jobs = append(jobs, &domain.Job{
				Id:          Id,
				Title:       title.(string),
				Description: description.(string),
				Skills:      lista,
				UserId:      userId.(string),
			})
		}
		return jobs, err
	})
	if err != nil {
		return nil, err
	}
	return jobs.([]*domain.Job), nil
}
