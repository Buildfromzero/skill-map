import React, { useState, useEffect } from 'react';
import User from "./user";

export default function UserList() {

  const [data, setData] = useState([]);

  function reloadUserList(event){
    console.log(event);
    fetch("https://jsonplaceholder.typicode.com/users", {
            method: "GET",
            headers: {
                "Content-type": "application/json; charset=UTF-8",
            },
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            setData(data);
            alert(JSON.stringify(data));
        })
  };

    return (
        <div>
          <br/>
            <h3>User List</h3>
            <br/>
            <button type="simpleQuery" className="btn btn-dark" onClick={reloadUserList}>Reload User List</button>
            <br/>
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
              data.map((item)=>(
                <User id={item.id} fullName={item.name} email={item.email} />
              ))
            }
          </tbody>
        </table>
        </div>

        
    );
  }