import React, { useState, useEffect, useContext } from "react";
import JobList from "../components/JobList";
import { useParams } from 'react-router-dom' 
import { AuthContext } from "../../shared/context/auth-context";

const RecommendedJobs = () => {
  const [recommendedJobs, setRecommendedJobs] = useState();
  var id = useParams()["id"]
  const auth = useContext(AuthContext);

  useEffect(() => {
    const fetchJobs = async () => {
      try {
        fetch(
          `http://localhost:8000/jobs/${auth.userId}/recommended`,
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
            console.log("jobs");
            console.log(jobs);
          });
          console.log("fetch happend")
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