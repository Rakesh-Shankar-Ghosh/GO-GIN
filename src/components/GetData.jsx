import React from "react";
import { useState, useEffect } from "react";
import axios from "axios";

const Insert = () => {
  const [mydata, setData] = useState();

  const dataHandler = async () => {
    try {
      console.log("I am calling....");
      const res = await axios.get(`${process.env.REACT_APP_API}/getstudents`);

      if (res?.data) {
        setData(res.data);
        console.log(res);
      }
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    dataHandler();
  }, []);

  return (
    <>
      <div>
        {mydata && (
          <div>
            <div>GET DATA WORKING SUCCESSFULLY.......</div>
          </div>
        )}
      </div>
    </>
  );
};

export default Insert;
