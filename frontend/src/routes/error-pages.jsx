import { useRouteError } from "react-router-dom";
import Navbar from '../components/navbar';

export default function ErrorPage() {
  const error = useRouteError();
  console.error(error);

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
          <li class="breadcrumb-item"><a href="/">Home</a></li>
          <li class="breadcrumb-item active" aria-current="users">Users</li>
        </ol>
      </nav>

      <br />


      <h1>Oops!</h1>
      <p>Sorry, an unexpected error has occurred.</p>
      <p>
        <i>{error.statusText || error.message}</i>
      </p>



      <br />
    </div>
  );
}