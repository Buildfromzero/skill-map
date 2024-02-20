import { useParams } from 'react-router-dom';
import React, { useState, useEffect } from 'react';


export default function UserDetail() {
    const [data, setData] = useState([]);
    const { userId } = useParams();
    const usetDetailApi = "http://localhost:8080/api/users/" + userId + "/";

    console.log("UserDetailApi: ", usetDetailApi);

    function updateUser(event) {
        event.preventDefault();

        var fullName = event.target.elements.fullName.value;
        var email = event.target.elements.email.value;
        console.log(fullName);
        console.log(email);

        fetch(usetDetailApi, {
            method: "PATCH",
            body: JSON.stringify({
                "fullName": fullName,
                "email": email,
            })
        })
            .then(response => response.json())
            .then(data => {
                console.log(data);
                alert(JSON.stringify(data));
            })
    }

    function getUserDetails() {
        console.log("get user detail");
        fetch(usetDetailApi, {
            method: "GET",
        })
            .then(response => {
                return response.json();
            })
            .then(data => {
                console.log(data);
                setData(data);
            })
    }

    function deleteUser(event) {
        event.preventDefault();
        console.log("Deleting user...");
        fetch(usetDetailApi, {
            method: "DELETE",
        })
            .then(response => {
                console.log(response);
                return response.json();
            })
            .then(data => {
                console.log(data);
                alert(JSON.stringify(data));
            })
    }


    useEffect(() => {
        getUserDetails();

    }, [])

    return (
        <div>
            <h3>User Details : {data.fullName}</h3>
            <br />
            <form onSubmit={updateUser}>
                <div className="mb-3">
                    <label className="form-label">Full Name</label>
                    <input type="text" className="form-control" id="fullName" defaultValue={data.fullName} />
                </div>
                <div className="mb-3">
                    <label className="form-label">Email</label>
                    <input type="text" className="form-control" id="email" defaultValue={data.email} />
                </div>
                <button type="submit" className="btn btn-primary">Update</button>
                <span style={{ marginLeft: '10px' }}></span>
                <button type="submit" className="btn btn-danger" onClick={deleteUser}>Delete user</button>
            </form>
            <br />

        </div>
    );
}