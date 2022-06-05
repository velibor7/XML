import React from "react";

import JobItem from "./JobItem";
import "./JobList.css";

const JobList = (props) => {
  return (
    <ul className="job-list">
      {props.items.map((job) => (
        <JobItem
          key={job.id}
          id={job.id}
          title={job.title}
          description={job.description}
          skills={job.skills}
          userId={job.userId}
        />
      ))}
    </ul>
  );
};
export default JobList;
