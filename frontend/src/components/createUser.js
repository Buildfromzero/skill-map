export default function CreateUser() {
    function onSubmission(event){

        var firstName = event.target.elements.firstName.value;
        var lastName = event.target.elements.lastName.value;
        console.log(firstName);
        console.log(lastName);

        fetch("https://jsonplaceholder.typicode.com/users", {
            method: "POST",
            body: JSON.stringify({
                firstName,
                lastName,
            }),
            headers: {
                "Content-type": "application/json; charset=UTF-8",
            },
        })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            alert(JSON.stringify(data));
        })
    }
    return (
      <div>
        <h3>Create User</h3>
        <br/>
            <form onSubmit={onSubmission}>
                <div className="mb-3">
                    <label className="form-label">First Name</label>
                    <input type="text" className="form-control" id="firstName"/>
                </div>
                <div className="mb-3">
                    <label className="form-label">Last Name</label>
                    <input type="text" className="form-control" id="lastName"/>
                </div>
                <button type="submit" className="btn btn-primary">Submit</button>
            </form>
        </div>
    );
  }