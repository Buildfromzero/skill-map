import './App.css';
import UserList from './components/userlist';
import Navbar from './components/navbar';
import CreateUser from './components/createUser';

function App() {
  return (

    <div className="container-fluid">

        <div className="row">
          <div className="col gap-3">
          <Navbar />
          </div>
        </div>

        <br/>

        <div className="row align-items-center">
          <div className="col-4">
          <CreateUser />
          </div>
        </div>

        <br/>

        <div className="row">
          <div className="col">
            <UserList />
          </div>
          </div>

        <br/>

  </div>
  );
}

export default App;
