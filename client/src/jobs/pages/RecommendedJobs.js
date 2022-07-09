import React, { useState, useEffect, _ } from "react";
import JobList from "../components/JobList";
import { useParams } from 'react-router-dom' 

const RecommendedJobs = () => {
  const [recommendedJobs, setRecommendedJobs] = useState();
  var id = useParams("id")


  useEffect(() => {
    const fetchJobs = async () => {

      try {
        fetch(
          `http://localhost:8000/jobs/${id}/recommended`,
          { method: "get", dataType: "json"}
        )
          .then((response) => response.json())
          .then((data) => {
            console.log(data)
            var jobs = [];
            data = data['job']
            for (var i = 0; i < data.length; i++) {
              var job = data[i];
              jobs.push(job);
            }
            setRecommendedJobs(jobs);
          });
      } catch (err) {
        console.log("error happend")
        console.log(err);
      };
    };
    fetchJobs();
  }, []);

  return (
    <>
        <JobList items={recommendedJobs}></JobList>
    </>
  );
};

export default RecommendedJobs;