import User from "./user";

export default function UserList() {

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
            alert(JSON.stringify(data));
        })
  };

    return (


        <div>
          <br/>
            <h3>User List</h3>

            <br/>

            <form onSubmit={reloadUserList}>
                <button type="submit" className="btn btn-dark">Relod Users</button>
            </form>

            <br/>
        <table class="table">
        <thead>
          <tr>
            <th scope="col">#ID</th>
            <th scope="col">First</th>
            <th scope="col">Last</th>
            <th scope="col">Skills</th>
          </tr>
        </thead>
        <tbody>
          <User id="1" firstName="Tom" lastName="Victor" />
          {/* <User />
          <User />
          <User />
          <User /> */}
        </tbody>
      </table>
        </div>

        
    );
  }