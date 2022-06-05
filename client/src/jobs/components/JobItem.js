import React, { useState, useContext } from "react";

import Card from "../../shared/components/UIElements/Card";

import "./JobItem.css";
import { useHttpClient } from "../../shared/hooks/http-hook";
import { AuthContext } from "../../shared/context/auth-context";

const JobItem = (props) => {
    const { isLoading, error, sendRequest, clearError } = useHttpClient();
    const [showDetailsModal, setShowDetailsModal] = useState(false);
    const auth = useContext(AuthContext);

    const cancelDetailsHandler = () => {
      setShowDetailsModal(false);
    };
  
    const popDetailsModal = () => {
      setShowDetailsModal(true);
    };
  
    return (
      <React.Fragment>
        <Modal
          show={showDetailsModal}
          onCancel={cancelDetailsHandler}
          header={props.title}
        >
          {props.description}
        </Modal>
       
        <li className="job-item" onClick={popDetailsModal}>
          <Card className="job-item__content">
            {isLoading && <Spinner />}

            <div className="job-item__info">
              <h2>{props.title}</h2>
            </div>
          </Card>
        </li>
      </React.Fragment>
    );
  };
  
  export default JobItem;