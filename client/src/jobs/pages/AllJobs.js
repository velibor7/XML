import React, { useState, useEffect, _ } from "react";
import JobList from "../components/JobList";

const AllJobs = () => {
  const [loadedJobs, setLoadedJobs] = useState();


  useEffect(() => {
    const fetchJobs = async () => {

      try {
        fetch(
          `http://localhost:8000/jobs`,
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
            setLoadedJobs(jobs);
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
        <JobList items={loadedJobs}></JobList>
    </>
  );
};

export default AllJobs;