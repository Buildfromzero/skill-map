import React, { useState, useEffect } from 'react';
import User from "./user";

export default function UserList() {

  const [data, setData] = useState([]);

  function reloadUserList(event) {
    // console.log(event);
    fetch("http://localhost:8080/api/users/", {
      method: "GET",
    })
      .then(response => response.json())
      .then(data => {
        console.log(data);
        setData(data);
        // alert(JSON.stringify(data));
      })
  };

  useEffect(() => {
    console.log("effects");
    reloadUserList(null);
  }, [])

  return (
    <div>
      <br />
      <h3>User List</h3>
      <br />
      <button type="simpleQuery" className="btn btn-dark" onClick={reloadUserList}>Reload User List</button>
      <br />
      <table class="table">
        <thead>
          <tr>
            <th scope="col">#ID</th>
            <th scope="col">Full Name</th>
            <th scope="col">Email</th>
            <th scope="col">Skills</th>
          </tr>
        </thead>
        <tbody>
          {
            data.map((item) => (
              <User id={item.id} fullName={item.fullName} email={item.email} />
            ))
          }
        </tbody>
      </table>
    </div>


  );
}