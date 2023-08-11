import React, { useState, useEffect } from "react";
import axios from "axios";

const Insert = () => {
  const [mydata, setData] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await axios.get(
          `${process.env.REACT_APP_API}/api/v1/getStudents`
        );

        if (res?.data) {
          setData(res.data.result);
          console.log(res.data.result);
        } else {
          console.log("No data received.");
        }
      } catch (error) {
        console.log("Error fetching data:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  return (
    <div>
      {loading ? (
        <p>Loading...</p>
      ) : mydata.length > 0 ? (
        mydata.map((student) => (
          <div key={student.id}>
            <p>Name: {student.name}</p>
            <p>Age: {student.age}</p>
          </div>
        ))
      ) : (
        <p>No data available</p>
      )}
    </div>
  );
};

export default Insert;
