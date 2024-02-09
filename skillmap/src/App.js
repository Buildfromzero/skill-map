import './App.css';
import UserList from './components/userlist';
import Navbar from './components/navbar';
import CreateUser from './components/createUser';

function App() {
  return (

    <div class="container-fluid">

    <div class="row">
    <div class="col gap-3">
    <Navbar />
    </div>
    </div>

    <br/>

    <div class="row align-items-center">
    <div class="col-4">
    <CreateUser />
    </div>
    </div>

    <br/>

    <div class="row">
    <div class="col">
      <UserList />
    </div>
    </div>

    <br/>

  </div>
  );
}

export default App;
