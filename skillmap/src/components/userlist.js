import User from "./user";

export default function UserList() {
    return (

        <div>
            <h3>User List</h3>
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
          <User />
          <User />
          <User />
          <User />
          <User />
        </tbody>
      </table>
        </div>

        
    );
  }