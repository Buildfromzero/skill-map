import Navbar from '../components/navbar';

function Home() {
  return (

    <div className="container-fluid">

      <div className="row">
        <div className="col gap-3">
          <Navbar />
        </div>
      </div>

      <br />

      <nav aria-label="breadcrumb">
        <ol class="breadcrumb">
          <li class="breadcrumb-item active" aria-current="page">Home</li>
        </ol>
      </nav>


      <br />
      <br />

      <div class="row">
        <div class="col-sm-3">
          <div class="card">
            <div class="card-body">
              <h5 class="card-title">50 Total Registered Users</h5>
              <p class="card-text">Go to the User page for managing skills</p>
              <a href="/users" class="btn btn-primary">Manage users</a>
            </div>
          </div>
        </div>
        <div class="col-sm-3">
          <div class="card">
            <div class="card-body">
              <h5 class="card-title">10 Total Skills</h5>
              <p class="card-text">go to the skills page for managing skills</p>
              <a href="/skills" class="btn btn-primary">Manage Skills</a>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Home;
