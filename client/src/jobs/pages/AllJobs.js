import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

import JobList from "../../jobs/components/JobList";
import { useHttpClient } from "../../shared/hooks/http-hook";

/*const DUMMY_DATA = [
  {
    id: 'j1',
    title: 'Software developer',
    description: 'Description of a software developer job',
    skills: ["Node.js", "React", "C#"]
  },
  {
    id: 'j2',
    title: 'Data scientist',
    description: 'Description of a data scientist job',
    skills: ["Node.js", "React", "C#"]
  },
  {
    id: 'j3',
    title: 'HR manager',
    description: 'Description of a HR manager job',
    skills: ["Node.js", "React", "C#"]
  },
];*/

const AllJobs = () => {
  const [loadedJobs, setLoadedJobs] = useState();
  const { isLoading, error, sendRequest, clearError } = useHttpClient();

 /* return (
    <section>
      <h1>All Jobs</h1>
      {DUMMY_DATA.map((job) => {
        return <li key ={job.id}>{job.title}</li>
      })}
    </section>
  );*/

  useEffect(() => {
    const fetchJobs = async () => {
      try {
        const resData = await sendRequest(
          `http://localhost:3000/api_gateway/job_service`
        );
        setLoadedJobs(resData.jobs);
      } catch (err) {}
    };
    fetchJobs();
  }, [sendRequest]);

  return (
    <React.Fragment>
      {!isLoading && loadedJobs && (
        <JobList
          items={loadedJobs}
          //onJobDetails={jobDetailsHandler}
        />
      )}
    </React.Fragment>
  );
};

export default AllJobs;