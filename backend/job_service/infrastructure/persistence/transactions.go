package persistence

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/velibor7/XML/job_service/domain"
)

func CreateRelationShips(session *neo4j.Session) error {
	_, err := (*session).WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		result, err := tx.Run(
			"MATCH (job1:Job),(job2:Job) "+
				"WHERE ANY (skill IN job1.skills WHERE skill IN job2.skills) "+
				"MERGE (job1)-[r1:SAME]->(job2)",
			nil)

		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return err
	}
	return nil
}

func DeleteRelationShip(session *neo4j.Session) error {
	_, err := (*session).WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		result, err := tx.Run(
			"MATCH (n:Job)-[r:SAME]->(n) DELETE r",
			nil)

		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return err
	}
	return nil
}

// povezati na osnovu skillova
//MATCH (job1:Job) WHERE job1.title = "Cloud" MATCH (job2:Job)
//WHERE job2.title = "Cloud" CREATE (job1)-[r1:SAME]->(job2)
//MATCH (n: user {uid: "1"}) CREATE (n) -[r: posted]-> (p: post {pid: "42", title: "Good Night",
//msg: "Have a nice and peaceful sleep.", author: n.uid});

func CreateNode(job *domain.Job, session *neo4j.Session) error {
	_, err := (*session).WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {

		result, err := tx.Run(
			"CREATE (job:Job {title:$title, description:$desc, skills:$skills, userId:$userId})",
			map[string]interface{}{
				"title": job.Title, "desc": job.Description, "skills": job.Skills, "userId": job.UserId,
			})

		if err != nil {
			return nil, err
		}
		return result.Consume()
	})
	if err != nil {
		return err
	}
	return nil
}
