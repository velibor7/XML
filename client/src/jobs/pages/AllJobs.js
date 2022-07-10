import React, { useState, useEffect, useContext} from "react";

import JobList from "../components/JobList";
import Button from "../../shared/components/FormElements/Button";
import { useNavigate } from "react-router-dom";
import { AuthContext } from "../../shared/context/auth-context";


const AllJobs = () => {
  const [loadedJobs, setLoadedJobs] = useState();
  const navigate = useNavigate();
  const auth = useContext(AuthContext);

const CreateNewJob = async () => {
  try {
    navigate(`/newjob`);
  } catch (err) {
    navigate(`/newjob`);
    console.log(err);
  }
}

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
      {auth.isLoggedIn && (
      <div className="job-item__actions">
      <Button info onClick={CreateNewJob}>
        CREATE NEW JOB OFFER
      </Button>
      </div>
      )} 
    </>
  );
};

export default AllJobs;