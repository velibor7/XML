import JobItem from "./JobItem";
import React from "react";

const JobList = (props) => {
  return (
    <div className="jobs">
      <h1>List of Jobs</h1>
      <div className="job-list__container">
        <div className="job__wrapper">
          <div className="job-list__items">
            {props.items?.map((item) => (
              <JobItem item={item} key={item.id} />
            ))}
          </div>
        </div>
      </div>
    </div>
  );
};

export default JobList;